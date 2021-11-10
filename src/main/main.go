package main

import (
	"Bulldog0.1/src/router"
	//"os"
)

func main() {
	//os.MkdirAll("/var/www/vartow/uploads", os.ModePerm);
	e := router.New()
	e.Start(":8080")

}