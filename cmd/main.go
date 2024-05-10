package main

import "github.com/Lincyaw/PaperGraph-backend/router"

func main() {
	engine := router.New()
	err := engine.Run(":8080")
	if err != nil {
		panic(err)
	}
}
