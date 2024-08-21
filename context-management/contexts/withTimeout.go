package contexts

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"
)

func WithTimeoutContext() {
	// generamos un contexto con un timeout de 3 seg
	ctx, cancelFunc := context.WithTimeout(context.Background(), 3*time.Second)
	// usamos defer para que cuando el main se termine de ejecutar, el cierre del timeout
	// se haga de forma automatica
	defer cancelFunc()

	// Hacemos una busqueda con el contexto
	res, err := Search(ctx, "hola")
	if err != nil {
		log.Println(err)
	}

	log.Printf("El resultado es: %s\n", res)
}

// hace una busqueda en un contexto de timeout, si en un cierto periodo de tiempo
// no responde, se cancela.
func Search(ctx context.Context, query string) (string, error) {
	resp := make(chan string)

	go func() {
		resp <- RandomSleepAndReturnAPI(query)
		close(resp)
	}()

	// esperamos por cualquier respuesta que nos llegue, sea el de la api o del contexto
	for {
		select {
		case respondio := <-resp:
			return respondio, nil
		case <-ctx.Done():
			return "", ctx.Err()
		}
	}
}

// esta funcion simula la llamada a una api y el tiempo de espera de respuesta
func RandomSleepAndReturnAPI(query string) string {
	// tenemos que generar un numero random entre 1 y 5 seg
	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	// generamos una duracion random
	randomDuration := time.Duration(random.Int63n(int64(5 * time.Second)))

	// aqui se realiza la pausa con la duracion dada
	time.Sleep(randomDuration)

	// respondemos con un mensaje despues del sleep
	return fmt.Sprintf("It took us %v...", randomDuration)
}
