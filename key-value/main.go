package main

import (
	"context"
	"fmt"
)


func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "token", "senha")
	bookingHotel(ctx, "Victor")

}

func bookingHotel(ctx context.Context, name string){
	token := ctx.Value("token")
	fmt.Println(token)
	fmt.Println(name)
}