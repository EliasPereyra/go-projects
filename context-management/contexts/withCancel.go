package contexts

import (
	"context"
	"fmt"
	"time"
)

func WithCancel() {
	// creo un contexto para manejar operaciones concurrentes
	ctx, cancel := context.WithCancel(context.Background())

	// lanzo una go routine que va a estar escuchando al contexto cancel
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Go routine 1 se cancelo:", ctx.Err())
				return
			default:
				fmt.Println("Go routine 1 se esta ejecutando")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Go routine 2 se cancelo:", ctx.Err())
				return
			default:
				fmt.Println("Go routine 2 se esta ejecutando")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	// Simulamos un trabajo que se ejecute en el main
	fmt.Println("Trabajando en el main...")
	time.Sleep(10 * time.Second)

	// cancelamos el contexto, el cual dara una señal a todas las go routines para que paren
	fmt.Println("Cancelando el contexto...")
	cancel()

	// se le da tiempo a las go routines para que terminen
	time.Sleep(1 * time.Second)
	fmt.Println("Finalizando el programa...")
}
