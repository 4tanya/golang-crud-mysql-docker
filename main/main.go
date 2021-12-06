package main

import (
	"github.com/router"
)

func main() {
    router := router.GetRouter()

    router.Run("localhost:8080")
}


