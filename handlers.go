package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func Transactions(c *fiber.Ctx) error {
	var payload struct {
		Amount float32 `json:"amount"`
		// due to errors getting while parsing the time, I have
		// used Timestamp as string rather than time.Time
		Timestamp string `json:"timestamp"`
	}
	// parsing body to struct initialised above.
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(422).JSON(fiber.Map{"message": "json is not parsable"})
	}
	// checking all the fields frontend should send
	// is present or not.
	if payload.Amount == 0 || payload.Timestamp == "" {
		return c.Status(400).JSON(fiber.Map{"message": "json is invalid"})
	}
	// parsing time got from user To time.Time
	z := time.Now().UTC()
	x, err := time.Parse(time.RFC3339Nano, payload.Timestamp)
	if err != nil {
		return c.Status(422).JSON(fiber.Map{"message": "json is not parsable"})
	}
	// checking the time user send is in future or not.

	// as we need get statics in constant time
	// I traded off space for getting better time complexity.
	// using a queue data structure and goroutine.(to make use of golang's full potential)
	UpdateDatabase(payload.Amount)
	queue.Enqueue(payload.Amount, z)
	if x.After(z) {
		return c.Status(422).JSON(fiber.Map{"message": "time is in future"})
	}
	// checking time is older than sixty seconds.
	diff := z.Sub(x).Seconds()
	if diff > 60 {
		return c.SendStatus(204)
	}
	// success message
	return c.Status(201).JSON(fiber.Map{"message": "success"})
}

func Statitics(c *fiber.Ctx) error {
	if queue.IsEmpty() {
		return c.Status(200).JSON(fiber.Map{
			"sum":   0,
			"avg":   0,
			"min":   0,
			"max":   0,
			"count": 0,
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"sum":   database.sum,
		"avg":   database.average,
		"min":   database.min,
		"max":   database.max,
		"count": database.count,
	})
}
