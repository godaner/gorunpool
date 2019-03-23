package gorunpool

import (
	"testing"
	"fmt"
	"time"
)

func TestGoRunPool_Run(t *testing.T) {
	pool:=NewPool(InitConfig{
		Size:10,
	})
	for i:=0;i<1000;i++{
		pool.Run(Task{
			Input:map[string]interface{}{
				"a":1,
				"b":"123",
			},
			Process: func(input Params) (output Params, err error) {
				fmt.Println("Process:")
				fmt.Println("Input:",input)
				return Params{
					"d":12331,
				},nil
			},
			Callback: func(output Params, err error) {
				fmt.Println("Callback:")
				fmt.Println(output,err)

			},

		})
	}
	time.Sleep(10000000)
}