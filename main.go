package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	m := map[string]func(){}
	ctx, cancel := context.WithCancel(context.Background())

	m["sad"] = cancel

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done(): // if cancel() execute
				fmt.Println("done")
				return
			default:

				time.Sleep(3 * time.Second)
			}
		}
	}(ctx)

	go func(m map[string]func()) {
		time.Sleep(3 * time.Second)
		m["sad"]()
	}(m)

	for {
		time.Sleep(1 * time.Second)
		fmt.Println("finish")
	}
}

/*
func process(building *model.Building) {
	forever := make(chan struct{})
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Duration(building.TimeNeed)*time.Second))

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done(): // if cancel() execute
				//	sendResponce
				return
			default:
				time.Sleep(1 * time.Second)
			}
		}
	}(ctx)

	go func() {
		time.Sleep(1 * time.Second)
		cancel()
	}()

	<-forever
	fmt.Println("finish")
}
*/
