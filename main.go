package main

import (
	"injar/configs"
	"injar/routes"
)

func main() {
	configs.InitDB()
	e := routes.New()
	e.Start(":8000")
}
