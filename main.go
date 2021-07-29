package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-co-op/gocron"
)

func main() {

	s := gocron.NewScheduler(time.UTC)
	s.Every(5).Seconds().Do(func() { fmt.Println("Hello World") })
	s.Every(10).Seconds().Do(func() { fmt.Println("Word Hello") })

	s.StartAsync()
	http.ListenAndServe(":8080", nil)

}
