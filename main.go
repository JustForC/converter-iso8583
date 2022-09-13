package main

import (
	"converter-iso8583/routes"
)

func main() {
	e := routes.Routes()
	e.Logger.Fatal(e.Start(":8000"))
}
