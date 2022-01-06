package main

import (
	"builder_service/internal/model"
	"context"
	"fmt"
	"time"
)

func main() {

}

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
