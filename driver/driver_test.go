package driver

import (
	"testing"
)

func TestConversion1(t *testing.T) {
	angle := -90
	u, err := degressToUs(angle, servoMin, servoMax)
	if u != 575 || err != nil {
		t.Fatalf(`Conversion of angle failed, expected %v but got %v`, 575, u)
	}
}

func TestConversion2(t *testing.T) {
	angle := 0
	u, err := degressToUs(angle, servoMin, servoMax)
	if u != 1449 || err != nil {
		t.Fatalf(`Conversion of angle failed, expected %v but got %v`, 1449, u)
	}
}

func Test3Conversion(t *testing.T) {
	angle := 90
	u, err := degressToUs(angle, servoMin, servoMax)
	if u != 2324 || err != nil {
		t.Fatalf(`Conversion of angle failed, expected %v but got %v`, 2324, u)
	}
}

func Test4Conversion(t *testing.T) {
	angle := -40
	u, err := degressToUs(angle, servoMin, servoMax)
	if u != 1061 || err != nil {
		t.Fatalf(`Conversion of angle failed, expected %v but got %v`, 1061, u)
	}
}

func TestConversion5(t *testing.T) {
	angle := -10
	u, err := degressToUs(angle, servoMin, servoMax)
	if u != 1352 || err != nil {
		t.Fatalf(`Conversion of angle failed, expected %v but got %v`, 1352, u)
	}
}

func TestConversion6(t *testing.T) {
	angle := 30
	u, err := degressToUs(angle, servoMin, servoMax)
	if u != 1741 || err != nil {
		t.Fatalf(`Conversion of angle failed, expected %v but got %v`, 1741, u)
	}
}

func TestInvConversion1(t *testing.T) {
	u, err := usToDegrees(575, servoMin, servoMax)
	if u != -90 || err != nil {
		t.Fatalf(`Conversion of us failed, expected %v but got %v`, -90, u)
	}
}

func TestInvConversion2(t *testing.T) {
	u, err := usToDegrees(1449, servoMin, servoMax)
	if u != 0 || err != nil {
		t.Fatalf(`Conversion of us failed, expected %v but got %v`, 0, u)
	}
}

func TestInvConversion3(t *testing.T) {
	u, err := usToDegrees(2324, servoMin, servoMax)
	if u != 90 || err != nil {
		t.Fatalf(`Conversion of us failed, expected %v but got %v`, 90, u)
	}
}

func TestInvConversion4(t *testing.T) {
	u, err := usToDegrees(1061, servoMin, servoMax)
	if u != -40 || err != nil {
		t.Fatalf(`Conversion of us failed, expected %v but got %v`, -40, u)
	}
}

func TestInvConversion5(t *testing.T) {
	u, err := usToDegrees(1352, servoMin, servoMax)
	if u != -10 || err != nil {
		t.Fatalf(`Conversion of us failed, expected %v but got %v`, -10, u)
	}
}

func TestInvConversion6(t *testing.T) {
	u, err := usToDegrees(1741, servoMin, servoMax)
	if u != 30 || err != nil {
		t.Fatalf(`Conversion of us failed, expected %v but got %v`, 30, u)
	}
}
