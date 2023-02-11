package main

import (
	"log"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/gofiber/fiber/v2"
)

var (
	database Values
	queue    Queue
)

func main() {
	// Use an external setup function in order
	// to configure the app in tests as well
	// this step is mandatory as I am using Fiber framework.
	app := Setup()

	// using go cron job scheduler
	s := gocron.NewScheduler(time.UTC)
	// in every five seconds this function will run to
	// update statitics. this will run cuncurrently.
	// by updating statics everytime we get data,
	// we can reduce time complexity O(n) to O(1)
	// space complexity is also O(1). even if user input grows
	// we are using only one struct to save data. that is
	// costant space and constant time operation
	s.Every(5).Seconds().Do(func() {
		front := queue.front
		if front == nil {
			resetDB()
		}
		for front != nil {
			z := time.Now().UTC()
			if z.Sub(front.updated).Seconds() > 60.0 {
				refreshDatabase(front.data)
				queue.Dequeue()
			} else {
				break
			}
			front = front.next
		}
		// fmt.Println("database minimun :", database.min, "database max:", database.max)
	})
	// starting a go routine without
	// blocking the main routine.
	s.StartAsync()

	// start the application on http://localhost:3000
	log.Fatal(app.Listen(":3000"))
}

// Setup Setup a fiber app with all of its routes
func Setup() *fiber.App {
	// Initialize a new app
	app := fiber.New()

	// Register the index route with a simple
	// "OK" response. It should return status
	// code 200
	app.Post("/transactions", Transactions)
	app.Get("/statitics", Statitics)

	// Return the configured app
	return app
}
