package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type HealthData struct {
	Month	int		`json:"month"`
	Day		int		`json:"day"`
	Weifgt	float64	`json:"weight"`
}

var health_data = []HealthData{{
	Month:	1,
	Day:	1,
	Weifgt:	55.5,
}, {
	Month:	1,
	Day:	2,
	Weifgt:	66.6,
}, {
	Month:	1,
	Day:	3,
	Weifgt:	77.7,
}}

func main() {
	handler1 := func(w http.ResponseWriter, r *http.Request) {
		var buf bytes.Buffer
		enc := json.NewEncoder(&buf)
		if err := enc.Encode(&health_data); err != nil {
			log.Fatal(err)
		}
		fmt.Println(buf.String())

		w.Header().Set("Access-Control-Allow-Origin","http://localhost:3000")

		_, err := fmt.Fprint(w, buf.String())
		if err != nil {
			return
		}
	}

	// GET /tasks
	http.HandleFunc("/health_data", handler1)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
