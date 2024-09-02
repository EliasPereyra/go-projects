package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Ref: https://refactoring.guru/es/design-patterns/decorator/go/example
// Componente base
type ISandwich interface {
	getDescription() string
	getPrice() int
}

// Componente concreto
type Simple struct {
}

func (m *Simple) getPrice() int {
	return 20
}

func (m *Simple) getDescription() string {
	return "Un sandwich con tomate, lechuga y mayonesa"
}

// Decorador concreto
type Completa struct {
	sandwich ISandwich
}

func (s *Completa) getPrice() int {
	return s.sandwich.getPrice() + 15
}

func (s *Completa) getDescription() string {
	return s.sandwich.getDescription() + " + queso y huevo"
}

// Otro decorador concreto
type Doble struct {
	sandwich ISandwich
}

func (s *Doble) getPrice() int {
	return s.sandwich.getPrice() + 25
}

func (s *Doble) getDescription() string {
	return s.sandwich.getDescription() + " con doble milanesas + queso y huevo"
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	sandwich := &Simple{}

	println("Que tipo de milanesa le gustaria ordenar?")
	// leer la entrada
	choice, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("There was an error", err)
		return
	}
	// basado en la entrada devuelvo el tipo
	switch strings.ToLower(strings.TrimSpace(choice)) {
	case "simple":
		// pido uno sencilla
		simple := &Simple{}
		fmt.Printf("Uno sencilla noma sale %d\n", simple.getPrice())
		fmt.Println("Quiere mas información?")
		reader := bufio.NewReader(os.Stdin)
		choice, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal("There was an error", err)
			return
		}
		if strings.ToLower(strings.TrimSpace(choice)) == "si" {
			fmt.Println(simple.getDescription())
		}
		return
	case "completa":
		// pido uno Completa
		completa := &Completa{
			sandwich: sandwich,
		}
		fmt.Printf("Uno completa sale %d\n", completa.getPrice())
		fmt.Println("Quiere mas información?")
		reader := bufio.NewReader(os.Stdin)
		choice, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal("There was an error", err)
			return
		}
		if strings.ToLower(strings.TrimSpace(choice)) == "Si" {
			fmt.Println(completa.getDescription())
		}
		return
	case "doble":
		// pido uno Doble
		doble := &Doble{
			sandwich: sandwich,
		}
		fmt.Printf("Uno doble sale %d\n", doble.getPrice())
		fmt.Println("Quiere mas información?")
		reader := bufio.NewReader(os.Stdin)
		choice, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal("There was an error", err)
			return
		}
		if strings.ToLower(strings.TrimSpace(choice)) == "Si" {
			fmt.Println(doble.getDescription())
		}
		return
	default:
		fmt.Println("Debes ingresar la opcion correcta.")
	}
}
