package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

const shortDuration = 1 * time.Microsecond

func main() {
	d := time.Now().Add(shortDuration)

	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}

	http.ListenAndServe()
}
