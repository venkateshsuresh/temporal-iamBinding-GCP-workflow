package main

import (
	"github.com/gin-gonic/gin"
	"personalproject/temporal/worker"
)

func main() {
	r := gin.Default()

	channel1 := make(chan interface{})
	defer func() {
		channel1 <- struct{}{}
	}()
	go iamWorkFlowInitialize(channel1)

	r.GET("/iambinding", worker.IamWorkFlow)
	r.Run()
}

func iamWorkFlowInitialize(channel <-chan interface{}) {
	err := worker.IamWorker.Run(channel)
	if err != nil {
		panic(err)
	}
}
