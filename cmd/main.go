package main

import "backend-kata/cmd/web"

func main() {
	app := web.NewApplication()
	app.Run()
}
