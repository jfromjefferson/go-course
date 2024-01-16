package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 8*time.Second)
	defer cancel()

	bookHotel(ctx)
}

func bookHotel(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("Hotel booking cancelled due to timeout")
		return
	case <-time.After(5 * time.Second):
		fmt.Println("Hotel booked")
	}
}
