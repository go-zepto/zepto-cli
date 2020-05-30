package main

import (
	"github.com/go-zepto/templates/default/controllers"
	"github.com/go-zepto/zepto"
)

func main() {
	// Create Zepto
	z := zepto.NewZepto(
		zepto.Name("web"),
		zepto.Version("latest"),
	)

	// Create web app
	a := z.NewWeb()

	// Movies
	a.GET("/", controllers.HelloIndex)

	// Setup HTTP Server
	z.SetupHTTP("0.0.0.0:8000", a)

	z.Start()
}
