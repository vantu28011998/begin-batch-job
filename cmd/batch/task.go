package batch

import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
)

func Task() {
	s := gocron.NewScheduler(time.UTC)
	s.Every(5).Seconds().Do(func() { fmt.Println("Hello") })

	s.StartAsync()
}
