package main

import "dev11/internal/app"

func main() {
	if err := app.Run(); err != nil {
		panic(err)
	}
}
