package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type HealthData struct {
	UserID      int       `json:"user_id"`
	Weifgt   float64   `json:"weight"`
	Date time.Time `json:"date"`
}

var health_data = []HealthData{{
	UserID:      1,
	Weifgt:   55.5,
	Date: time.Now(),
}, {
	UserID:      2,
	Weifgt:   66.6,
	Date: time.Now(),
}, {
	UserID:      3,
	Weifgt:   77.7,
	Date: time.Now(),
}}

func main() {
	handler1 := func(w http.ResponseWriter, r *http.Request) {
		var buf bytes.Buffer
		enc := json.NewEncoder(&buf)
		if err := enc.Encode(&health_data); err != nil {
			log.Fatal(err)
		}
		fmt.Println(buf.String())

		_, err := fmt.Fprint(w, buf.String())
		if err != nil {
			return
		}
	}

	// GET /tasks
	http.HandleFunc("/health_data", handler1)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
