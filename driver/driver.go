package driver

import (
	"errors"
	"math"
	"time"

	"github.com/corrupt/go-smbus"
)

const address = 0x15
const servoMin = 575
const servoMax = 2325

const REG_CONFIG = 0x00
const REG_SERVO1 = 0x01
const REG_SERVO2 = 0x03

var ErrInvalidAngle = errors.New("invalid angle")
var ErrInvalidReading = errors.New("invalid reading")

type PanAndTiltHat struct {
	idleTimeout   time.Duration
	servo1Timer   *time.Timer
	servo2Timer   *time.Timer
	servoMin      [2]uint16
	servoMax      [2]uint16
	bus           *smbus.SMBus
	servo1Enabled bool
	servo2Enabled bool
}

func (pt *PanAndTiltHat) config() {
	var config byte
	config = 0
	if pt.servo1Enabled {
		config = config | 1
	}
	if pt.servo2Enabled {
		config = config | 1<<1
	}

	pt.bus.Write_byte_data(REG_CONFIG, config)
}

func usToDegrees(us, usmin, usmax uint16) (int, error) {
	if us < usmin || us > usmax {
		return 0, ErrInvalidReading
	}
	servoRange := usmax - usmin
	res := (float64(us-usmin) / float64(servoRange)) * 180.0
	return int(math.Round(res) - 90), nil
}

func degressToUs(angle int, usmin, usmax uint16) (uint16, error) {
	if angle < -90 || angle > 90 {
		return 0, ErrInvalidAngle
	}
	angle += 90
	servoRange := float64(usmax - usmin)
	us := (servoRange / 180) * float64(angle)
	return usmin + uint16(us), nil
}

func (pt *PanAndTiltHat) PanValue() (int, error) {
	us, err := pt.bus.Read_word_data(REG_SERVO1)
	if err != nil {
		return 0, err
	}
	return usToDegrees(us, pt.servoMin[0], pt.servoMax[0])
}

func (pt *PanAndTiltHat) TiltValue() (int, error) {
	us, err := pt.bus.Read_word_data(REG_SERVO2)
	if err != nil {
		return 0, err
	}
	return usToDegrees(us, pt.servoMin[1], pt.servoMax[1])
}

func (pt *PanAndTiltHat) Pan(angle int) error {
	if !pt.servo1Enabled {
		pt.servo1Enabled = true
		pt.config()
	}
	us, err := degressToUs(angle, pt.servoMin[0], pt.servoMax[0])
	if err != nil {
		return err
	}
	err = pt.bus.Write_word_data(REG_SERVO1, us)
	if pt.servo1Timer != nil {
		pt.servo1Timer.Stop()
	}
	if pt.idleTimeout > 0 {
		pt.servo1Timer = time.AfterFunc(pt.idleTimeout, func() {
			pt.servo1Timer = nil
			pt.servo1Enabled = false
			pt.config()
		})
	}
	return err
}

func (pt *PanAndTiltHat) Tilt(angle int) error {
	if !pt.servo2Enabled {
		//enable servo 2
		pt.servo2Enabled = true
		pt.config()
	}
	us, err := degressToUs(angle, pt.servoMin[1], pt.servoMax[1])
	if err != nil {
		return err
	}
	err = pt.bus.Write_word_data(REG_SERVO2, us)
	if pt.servo2Timer != nil {
		pt.servo2Timer.Stop()
	}
	if pt.idleTimeout > 0 {
		pt.servo2Timer = time.AfterFunc(pt.idleTimeout, func() {
			pt.servo1Timer = nil
			pt.servo2Enabled = false
			pt.config()
		})
	}

	return err
}

func (pt *PanAndTiltHat) Close() {
	if pt.servo1Timer != nil {
		pt.servo1Timer.Stop()
	}
	if pt.servo2Timer != nil {
		pt.servo2Timer.Stop()
	}
	pt.servo1Enabled = false
	pt.servo2Enabled = false
	pt.config()
	pt.bus.Bus_close()
}

func Initialize(idleTimeout time.Duration) (*PanAndTiltHat, error) {
	pt := &PanAndTiltHat{
		idleTimeout: idleTimeout,
		servoMin:    [2]uint16{servoMin, servoMin},
		servoMax:    [2]uint16{servoMax, servoMax},
	}
	smb, err := smbus.New(1, address)
	if err != nil {
		return nil, err
	}
	pt.bus = smb
	return pt, nil
}
