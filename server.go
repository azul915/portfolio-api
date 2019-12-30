package main

import (
  "github.com/azul915/portfolio-api/src/infrastructure"
)

func main() {
    infrastructure.Router.Run(":1999")
}