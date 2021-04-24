package main

import (
	"fmt"
	"github.com/baranius/godiator"
	"github.com/baranius/godiator-echo/resources"
	"github.com/baranius/godiator-echo/resources/healthcheck/healthcheck_handlers"
	"github.com/baranius/godiator-echo/utils/pipelines"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize GinGonic
	g := gin.Default()

	//Initialize Mediator
	gdtr := godiator.GetInstance()

	//RegisterPipeline
	gdtr.RegisterPipeline(&pipelines.ValidationPipeline{})

	//Register Handlers
	gdtr.Register(&healthcheck_handlers.HealthCheckRequest{}, healthcheck_handlers.NewHealthCheckHandler)

	// Register Routes
	resources.Register(g)

	// Run GinGonic
	err := g.Run(":8080")
	if err != nil {
		fmt.Println(err.Error())
	}
}
