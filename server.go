package main

import (
	"portfolio-api/api/infrastructure"
)

func main() {
	infrastructure.Router.Run(":1999")
}
