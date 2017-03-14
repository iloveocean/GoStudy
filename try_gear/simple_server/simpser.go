package main

import (
	"fmt"

	"github.com/teambition/gear"
	"github.com/teambition/gear/logging"
)

func main() {
	app := gear.New()

	//Add logging middleware
	app.UseHandler(logging.Default())

	// Add router middleware
	router := gear.NewRouter()
	router.Use(func(ctx *gear.Context) error {
		// do some things
		fmt.Println("Router middleware...", ctx.Path)
		return nil
	})
	router.Get("/", func(ctx *gear.Context) error {
		return ctx.HTML(200, "<h1>Hello, Gear!</h1>")
	})
	app.UseHandler(router)
	app.Error(app.Listen(":3000"))
}
