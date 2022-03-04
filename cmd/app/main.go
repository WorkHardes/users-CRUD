package main

import (
	app "github.com/users-CRUD/internal/app"
)

const configsDir = "configs"

func main() {
	app.Run(configsDir)
}
