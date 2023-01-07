package main

import (
	"context"
	"fmt"
	"time"
)

func enrichContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, "request-id", 200)
}

func doSomethingCool(ctx context.Context) {
	rId := ctx.Value("request-id")
	fmt.Println(rId)

	for {
		select {
		case <-ctx.Done():
			fmt.Println("time out")
			return
		default:
			fmt.Println("doing something cool")
		}
		time.Sleep(time.Millisecond * 500)
	}

}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*6)
	defer cancel()

	ctxx := enrichContext(ctx)

	// melihat eror
	fmt.Println(ctxx.Err())

	go doSomethingCool(ctxx)

	select {
	case <-ctxx.Done():
		fmt.Println("oh no, i've been exceeded the deadline")

		// melihat eror
		fmt.Println(ctxx.Err())
	}

	time.Sleep(time.Second * 2)
}
