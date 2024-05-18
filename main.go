package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	gestorarchivos "github.com/Andresx117/SegundoParcialGo/GestorArchivos"
	VecinoCercano "github.com/Andresx117/SegundoParcialGo/VecinoCercano"
)

func main() {
	CanalNodo := make(chan []gestorarchivos.Nodo)
	CanalVecino := make(chan gestorarchivos.Resultado)
	CanalInsercion := make(chan gestorarchivos.Resultado)

	var wg sync.WaitGroup

	// Iniciar una goroutine para leer los nodos y enviarlos al canal
	go func() {
		nodos := gestorarchivos.LeerNodos("NodosSahara.tsp")
		CanalNodo <- nodos
	}()

	// Recibir los nodos del canal
	IndiceNodos := <-CanalNodo
	rand.Seed(time.Now().UnixNano())
	IndiceAleatorio := rand.Intn(len(IndiceNodos))
	fmt.Println("Nodo inicio", IndiceNodos[IndiceAleatorio].Nombre)
	//distanciasPrim, distanciasSec := VecinoCercano.Calculo(IndiceNodos)
	//todasDistancias := append(distanciasPrim, distanciasSec...)
	//fmt.Println("Distancias calculadas:", todasDistancias)

	// Calcular la ruta óptima utilizando Vecino Más Cercano
	wg.Add(1)
	go func() {

		rutaVecinoMasCercano, distanciaTotalVecinoMasCercano := VecinoCercano.VecinoMasCercano(IndiceNodos)
		fmt.Println("Ruta utilizando Vecino Más Cercano:", rutaVecinoMasCercano)
		fmt.Println("Distancia total utilizando Vecino Más Cercano:", distanciaTotalVecinoMasCercano)
		Resve := gestorarchivos.CrearResultado(rutaVecinoMasCercano, distanciaTotalVecinoMasCercano)
		CanalVecino <- *Resve
		wg.Done()
	}()

	// Calcular la ruta óptima utilizando Inserción Más Cercana
	wg.Add(1)
	go func() {

		rutaInsercionMasCercana, distanciaTotalInsercionMasCercana := VecinoCercano.InsercionMasCercana(IndiceNodos)
		fmt.Println("Ruta utilizando Inserción Más Cercana:", rutaInsercionMasCercana)
		fmt.Println("Distancia total utilizando Inserción Más Cercana:", distanciaTotalInsercionMasCercana)
		Resin := gestorarchivos.CrearResultado(rutaInsercionMasCercana, distanciaTotalInsercionMasCercana)
		CanalInsercion <- *Resin
		wg.Done()
	}()
	//wg.Wait()
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
