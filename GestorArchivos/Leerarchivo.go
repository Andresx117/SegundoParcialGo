package gestorarchivos

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Nodo struct {
	Nombre   string
	CoorX    float64
	CoorY    float64
	Visitado bool
}

func CrearNodos(nombre string, X, Y float64) *Nodo {
	Cn := &Nodo{
		Nombre:   nombre,
		CoorX:    X,
		CoorY:    Y,
		Visitado: false,
	}

	return Cn

}

func LeerNodos(NombreArchivo string) []Nodo {

	var ColeccionNodos []Nodo

	file, err := os.Open(NombreArchivo)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Lee cada línea del archivo
	for scanner.Scan() {
		linea := scanner.Text()
		campos := strings.Split(linea, " ")

		// Convierte los valores a los tipos adecuados
		Z := campos[0]
		X, _ := strconv.ParseFloat(campos[1], 32)
		Y, _ := strconv.ParseFloat(campos[2], 32)

		ColeccionNodos = append(ColeccionNodos, *CrearNodos(Z, X, Y))

	}
	return ColeccionNodos
}
