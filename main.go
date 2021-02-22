package main

import (
	"carrinho-api-gorm/api"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	api := api.Server{}
	api.StartServer()
}
