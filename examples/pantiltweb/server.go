package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/eisenwinter/go-panandtilt/driver"
	"github.com/go-chi/chi/v5"
)

func main() {
	pt, err := driver.Initialize(2 * time.Second)
	if err != nil {
		log.Panic(err)
	}
	defer pt.Close()
	tmpl, err := template.ParseFiles("gui.html")
	if err != nil {
		log.Panic(err)
	}

	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		if err := tmpl.Execute(w, nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	r.Get("/api/{direction}/{angle}", func(w http.ResponseWriter, r *http.Request) {
		if direction := chi.URLParam(r, "direction"); direction != "" {
			if angle := chi.URLParam(r, "angle"); angle != "" {
				a, e := strconv.Atoi(angle)
				if e != nil {
					w.WriteHeader(400)
					return
				}
				a = a - 90
				if direction == "pan" {
					pt.Pan(a)
					w.Write([]byte(fmt.Sprintf("{pan: %d}", a)))
					return
				}
				if direction == "tilt" {
					pt.Tilt(a)
					w.Write([]byte(fmt.Sprintf("{tilt: %d}", a)))
					return
				}
			}
		}

		w.WriteHeader(500)
	})
	fmt.Printf("Listening on port 9595")
	err = http.ListenAndServe(":9595", r)
	if err != nil {
		panic(err)
	}
}
