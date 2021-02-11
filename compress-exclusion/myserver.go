package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/event", func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Content-Disposition", "Attachment")
		log.Printf("Request: %v\n", r)
		log.Printf("Response: %v\n", w)

		content := fillContent(1600)
		http.ServeContent(w, r, "data.txt", time.Now(), content)
	})

	http.HandleFunc("/yolo", func(w http.ResponseWriter, r *http.Request){
		w.Header().Add("Content-Type", "text/yolo")
		w.Header().Set("Content-Disposition", "Attachment")
		log.Printf("Request: %v\n", r)
		log.Printf("Response: %v\n", w)

		content := fillContent(1600)
		http.ServeContent(w, r, "data.txt", time.Now(), content)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Disposition", "Attachment")
		log.Printf("Request: %v\n", r)
		log.Printf("Response: %v\n", w)

		content := fillContent(1600)
		http.ServeContent(w, r, "data.txt", time.Now(), content)
	})


	fmt.Printf("Starting server at port 9000\n")
	if err := http.ListenAndServe(":9000", nil); err != nil {
		log.Fatal(err)
	}
}


func fillContent(length int64) io.ReadSeeker {
	charset := "-ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, length)

	for i := range b {
		b[i] = charset[i%len(charset)]
	}

	if length > 0 {
		b[0] = '|'
		b[length-1] = '|'
	}

	return bytes.NewReader(b)
}