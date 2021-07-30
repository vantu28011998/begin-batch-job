package main

import (
	"Dlgo-pub-batch/cmd/batch"
	"net/http"
)

func main() {
	batch.Task()
	http.ListenAndServe(":8080", nil)
}
