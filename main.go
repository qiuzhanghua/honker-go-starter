package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/russellromney/honker-go"
)

func main() {
	db, err := honker.Open("app.db", "./honker_ext.dll")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	q := db.Queue("emails", honker.QueueOptions{})

	// Enqueue
	_, err = q.Enqueue(map[string]any{"to": "alice@example.com"}, honker.EnqueueOptions{})
	if err != nil {
		log.Fatal(err)
	}

	// Claim + process + ack
	job, err := q.ClaimOne("worker-1")
	if err != nil {
		log.Fatal(err)
	}

	var body map[string]any
	json.Unmarshal(job.Payload, &body)
	fmt.Printf("sending to %s\n", body["to"])

	if _, err := job.Ack(); err != nil {
		log.Fatal(err)
	}
}
