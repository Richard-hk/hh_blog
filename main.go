package main

import (
	"hh_blog/config"
	"hh_blog/routes"
)

func main() {
	config.ConfInit()
	routes.RouterInit()
}
