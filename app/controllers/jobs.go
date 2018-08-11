package controllers

import (
	"fmt"

	"github.com/lyhopq/firefly/app/models"
	"github.com/revel/modules/jobs/app/jobs"
	"github.com/revel/revel"
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
