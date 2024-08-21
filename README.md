# Saludos en go 

Este paquete proporciona una forma simple de obtener saludos personalizados en go 

# Instalación

Ejecuta el siguiente comando para instalar este paquete:

´´´´ bash
go get -u github.com/AndreaBur/greetings

# Uso 

Aquí tienes un ejemplo de cómo utilizar el paquete en tu código :

´´´´´go

package main

import (
	"fmt"
	"log"

	"github.com/AndreaBur/greetings"
)

func main() {

	log.SetPrefix("greetings")
	log.SetFlags(0)

	//para probar la función Hellos tenemos
	names := []string{"Alex", "Roel", "Juan"}
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
	/*
		//ahora vamos a capturar este error
		message, err := greetings.Hello("Andrea")
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(message)
	*/

}

--------------------------------------------------
