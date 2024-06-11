package main

import "todo-app/routes"

func main() {
	r := routes.Router()
	r.Start(":8080")
}
