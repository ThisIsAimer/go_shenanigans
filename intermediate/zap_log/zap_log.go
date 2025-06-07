package main

import (
	"fmt"

	"go.uber.org/zap"
)

func main(){
	fmt.Println("zap logger!")
	
	logger, err := zap.NewProduction()
	if err != nil {
		fmt.Println("error is:", err)
		return
	}


	//flush zap buffer
	defer logger.Sync()

	logger.Info("this is a zap log message")

}