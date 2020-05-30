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

	//// Setup Broker Provider (Google Pub/Sub)
	//z.SetupBroker(gcp.NewBroker(
	//	gcp.ProjectID("slints-usa"),
	//	gcp.TopicPrefix("dev.movies."),
	//))

	z.Start()
}
