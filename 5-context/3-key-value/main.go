package main

import "context"

func main() {
	ctx := context.WithValue(context.Background(), "name", "John")

	bookHotel(ctx)
}

func bookHotel(ctx context.Context) {
	name := ctx.Value("name").(string)

	println(name)
}
