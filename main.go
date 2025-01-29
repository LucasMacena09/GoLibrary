package main

import (
	"github.com/LucasMacena09/GoLibrary/routes"
)

func main() {
	router := routes.SetupRouter()
	router.Run("localhost:8080")
}