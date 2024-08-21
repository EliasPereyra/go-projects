package contexts

import (
	"context"
	"fmt"
)

func WithValueContext() {
	// en un contexto de value podemos pasar datos que estan en el scope de request
	// en toda la app de forma segura
	ctx := context.WithValue(context.Background(), "userId", 14)

	// le pasamos el contexto para que pueda acceder al valor
	ProcessRequest(ctx)
}

func ProcessRequest(ctx context.Context) {
	// Extraemos el valor del contexto
	// usamos .(int) como assertion para verificar si el valor es un integer
	userId, ok := ctx.Value("userId").(int)
	if !ok {
		fmt.Println("No se encontro el userId en el contexto")
		return
	}

	// mostramos el userId
	fmt.Printf("El userId es: %d\n", userId)

	// volvemos a pasar el contexto, para ver como el contexto se puede pasar
	// a traves de la cadena de llamadas de manera segura
	FurtherProcessing(ctx)
}

// aqui se haria trabajo extra, o cualquier logica que ayuda a un mejor procesamiento
func FurtherProcessing(ctx context.Context) {
	userId, ok := ctx.Value("userId").(int)
	if !ok {
		fmt.Println("No se encontro el userId en el contexto")
		return
	}

	// mostramos el userId
	fmt.Printf("Proceso extra para user Id: %d\n", userId)
}
