package main

import (
	"fmt"

	gestorarchivos "github.com/Andresx117/SegundoParcialGo/GestorArchivos"
	VecinoCercano "github.com/Andresx117/SegundoParcialGo/VecinoCercano"
)

func main() {
	CanalNodo := make(chan []gestorarchivos.Nodo)

	// Iniciar una goroutine para leer los nodos y enviarlos al canal
	go func() {
		nodos := gestorarchivos.LeerNodos("NodosSahara.tsp")
		CanalNodo <- nodos
	}()

	// Recibir los nodos del canal
	IndiceNodos := <-CanalNodo
	distanciasPrim, distanciasSec := VecinoCercano.Calculo(IndiceNodos)
	todasDistancias := append(distanciasPrim, distanciasSec...)
	fmt.Println("Distancias calculadas:", todasDistancias)

	// Calcular la ruta óptima utilizando Vecino Más Cercano
	rutaVecinoMasCercano, distanciaTotalVecinoMasCercano := VecinoCercano.VecinoMasCercano(IndiceNodos)
	fmt.Println("Ruta utilizando Vecino Más Cercano:", rutaVecinoMasCercano)
	fmt.Println("Distancia total utilizando Vecino Más Cercano:", distanciaTotalVecinoMasCercano)

	// Calcular la ruta óptima utilizando Inserción Más Cercana
	rutaInsercionMasCercana, distanciaTotalInsercionMasCercana := VecinoCercano.InsercionMasCercana(IndiceNodos)
	fmt.Println("Ruta utilizando Inserción Más Cercana:", rutaInsercionMasCercana)
	fmt.Println("Distancia total utilizando Inserción Más Cercana:", distanciaTotalInsercionMasCercana)
}

//todasDistancias es un arreglo con las distancias entre distintos pares de nodos, así: [nodo Inicial, nodo Final, distancia]

/*for _, distancia := range todasDistancias {
		fmt.Printf("Desde %s hasta %s: %.2f\n", distancia.NodoI, distancia.NodoFinal, distancia.Distancia)
	}

	// Calcular la ruta óptima utilizando el algoritmo del vecino más cercano
	rutaOptima, distanciaTotal := VecinoCercano.VecinoMasCercano(IndiceNodos)

	// Imprimir la ruta óptima
	fmt.Println("\nRuta óptima (Vecino más cercano):")
	for _, distancia := range rutaOptima {
		fmt.Printf("Desde %s hasta %s: %.2f\n", distancia.NodoI, distancia.NodoFinal, distancia.Distancia)
	}
	fmt.Printf("Distancia total de la ruta óptima: %.2f\n", distanciaTotal)
}

/*func main() {
	fmt.Println()

	CanalNodo := make(chan []gestorarchivos.Nodo)
	CanalNodo <- gestorarchivos.LeerNodos("NodosSahara.tsp")
	//IndiceNodos := <-CanalNodo
	//vecinocercano.Calculo(IndiceNodos)
}
*/
