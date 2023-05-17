package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()   // uma vez que ele iniciar ficar rodando na thread principal
	ctx, cancel := context.WithTimeout(ctx, time.Second * 3) // timeout de 3 segundos e cancelado
	defer cancel()
	bookingHotel(ctx)

}

func bookingHotel(ctx context.Context){
	//
	select {
		case <-ctx.Done():
            fmt.Println("Hotel booking canceled. TImeout reached.")
			return
		case <-time.After(1 * time.Second):
			fmt.Println("Hotel booking started.")
			return
	}
	
}