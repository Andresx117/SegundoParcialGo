package main

import (
	"fmt"

	gestorarchivos "github.com/Andresx117/SegundoParcialGo/GestorArchivos"
)

func main() {
	fmt.Println()

	CanalNodo := make(chan []gestorarchivos.Nodo)
	CanalNodo <- gestorarchivos.LeerNodos("../NodosSahara.tsp")

}
