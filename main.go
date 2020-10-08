package main

import "ffly-plus/router"

func main() {
	server := router.InitRouter()
	server.GinEngine.Run(":8000")
}
