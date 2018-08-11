package controllers

import (
	"firefly/app/models"
	"fmt"
	"github.com/revel/revel"
	"github.com/revel/modules/jobs/app/jobs"
)

// Periodically count the bookings in the database.
type BookInfoFetch struct {
	Qbs
}

func (c BookInfoFetch) Run() {
	c.Dial()
	count := models.BookCount(c.q)
	fmt.Printf("There are %d books.\n", count)
	fmt.Printf("DouBan keys: %s.\n", doubanKey)
	c.Close()
}

func init() {
	revel.OnAppStart(func() {
		jobs.Schedule("@every 1m", BookInfoFetch{})
	})
}
