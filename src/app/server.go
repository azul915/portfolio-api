package main

import (
  "github.com/azul915/portfolio-api/src/app/infrastructure"
)

func main() {
  infrastructure.Router.Run(":1999")
}