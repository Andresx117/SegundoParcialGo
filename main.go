package main

import (
	"fmt"
	"sync"

	GestorArchivos "github.com/Andresx117/SegundoParcialGo/GestorArchivos"
	VecinoCercano "github.com/Andresx117/SegundoParcialGo/VecinoCercano"
)

func main() {
	CanalNodo := make(chan []GestorArchivos.Nodo)
	CanalVecino := make(chan GestorArchivos.Resultado)
	CanalInsercion := make(chan GestorArchivos.Resultado)

	var wg sync.WaitGroup

	// Iniciar una goroutine para leer los nodos y enviarlos al canal
	go func() {
		nodos := GestorArchivos.LeerNodos("NodosSahara.tsp")
		CanalNodo <- nodos
	}()

	// Recibir los nodos del canal
	IndiceNodos := <-CanalNodo

	// Calcular la ruta óptima utilizando Vecino Más Cercano
	wg.Add(1)
	go func() {
		distanciasPrim, distanciasSec := VecinoCercano.Calculo(IndiceNodos)
		fmt.Println("Distancias calculadas por búsqueda de vecindario:", append(distanciasPrim, distanciasSec...))
		wg.Done()
	}()

	// Calcular la ruta óptima utilizando Vecino Más Cercano
	wg.Add(1)
	go func() {
		rutaVecinoMasCercano, distanciaTotalVecinoMasCercano := VecinoCercano.VecinoMasCercano(IndiceNodos)
		fmt.Println("Ruta utilizando Vecino Más Cercano:", rutaVecinoMasCercano)
		fmt.Println("Distancia total utilizando Vecino Más Cercano:", distanciaTotalVecinoMasCercano)
		Resve := GestorArchivos.CrearResultado(rutaVecinoMasCercano, distanciaTotalVecinoMasCercano)
		CanalVecino <- *Resve
		wg.Done()
	}()

	// Calcular la ruta óptima utilizando Inserción Más Cercana
	wg.Add(1)
	go func() {
		rutaInsercionMasCercana, distanciaTotalInsercionMasCercana := VecinoCercano.InsercionMasCercana(IndiceNodos)
		fmt.Println("Ruta utilizando Inserción Más Cercana:", rutaInsercionMasCercana)
		fmt.Println("Distancia total utilizando Inserción Más Cercana:", distanciaTotalInsercionMasCercana)
		Resin := GestorArchivos.CrearResultado(rutaInsercionMasCercana, distanciaTotalInsercionMasCercana)
		CanalInsercion <- *Resin
		wg.Done()
	}()

	// Esperar a que todas las goroutines terminen
	wg.Wait()

	// Imprimir resultados recibidos de las goroutines
	fmt.Println(<-CanalVecino)
	fmt.Println(<-CanalInsercion)
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
