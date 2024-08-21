package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// vamos a crear una función que devuelve un saludo recibiendo un dato

// Función para un mensaje
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("Nombre vacío!")
	}
	message := fmt.Sprintf(randomFormats(), name)
	return message, nil
}

// Funcion para varios mensajes
func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)
	//usamos for range para iterar los elementos del mapa
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}

		messages[name] = message

	}
	return messages, nil
}

func randomFormats() string {
	formats := []string{
		"!Hola, %v! Bienvenido¡",
		"!Que bueno verte, %v¡",
		"!Saludo, %v¡, !Encantad@ de conocerte¡",
	}
	return formats[rand.Intn(len(formats))]

}
