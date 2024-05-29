package Parcial

import (
	"math"
	"sync"

	gestorarchivos "github.com/Andresx117/SegundoParcialGo/GestorArchivos"
	VecinoCercano "github.com/Andresx117/SegundoParcialGo/VecinoCercano"
)

func Parcial() {
	CanalNodo := make(chan []gestorarchivos.Nodo)
	CanalVecino := make(chan gestorarchivos.Resultado)
	CanalInsercion := make(chan gestorarchivos.Resultado)
	//CanalVecindario := make(chan gestorarchivos.Resultado)

	var wg sync.WaitGroup

	// Iniciar una goroutine para leer los nodos y enviarlos al canal
	go func() {
		nodos := gestorarchivos.LeerNodos("NodosSahara.tsp")
		CanalNodo <- nodos
	}()

	// Recibir los nodos del canal
	IndiceNodos := <-CanalNodo

	// Calcular la ruta óptima utilizando Vecino Más Cercano
	wg.Add(1)
	go func() {
		var resultado []gestorarchivos.Resultado
		var x int
		var Resultadofinal gestorarchivos.Resultado
		Distanciafinal := math.Inf(x)
		for i := 0; i < len(IndiceNodos); i++ {
			rutaVecinoMasCercano, distanciaTotalVecinoMasCercano := VecinoCercano.VecinoMasCercano(IndiceNodos, i)
			//fmt.Println("Ruta utilizando Vecino Más Cercano:", rutaVecinoMasCercano)
			//fmt.Println("Distancia total utilizando Vecino Más Cercano:", distanciaTotalVecinoMasCercano)
			Resve := gestorarchivos.CrearResultado(rutaVecinoMasCercano, distanciaTotalVecinoMasCercano)
			resultado = append(resultado, *Resve)
		}

		for _, nodo := range resultado {

			if float64(Distanciafinal) >= nodo.DistanciaR {
				Resultadofinal = nodo
				Distanciafinal = nodo.DistanciaR

			}

		}
		println("Distancia final Vecino", Resultadofinal.DistanciaR)
		CanalVecino <- Resultadofinal
		wg.Done()
	}()

	// Calcular la ruta óptima utilizando Inserción Más Cercana
	wg.Add(1)
	go func() {
		var resultado []gestorarchivos.Resultado
		var x int
		var Resultadofinal gestorarchivos.Resultado
		Distanciafinal := math.Inf(x)
		for i := 0; i < len(IndiceNodos)-1; i++ {

			rutaInsercionMasCercana, distanciaTotalInsercionMasCercana := VecinoCercano.InsercionMasCercana(IndiceNodos, i)
			//fmt.Println("Ruta utilizando Inserción Más Cercana:", rutaInsercionMasCercana)
			//fmt.Println("\nDistancia total utilizando Inserción Más Cercana: ", distanciaTotalInsercionMasCercana, " ")
			Resin := gestorarchivos.CrearResultado(rutaInsercionMasCercana, distanciaTotalInsercionMasCercana)
			resultado = append(resultado, *Resin)
		}

		for _, nodo := range resultado {

			if float64(Distanciafinal) >= nodo.DistanciaR {
				Resultadofinal = nodo
				Distanciafinal = nodo.DistanciaR

			}

		}
		println("\nDistancia final Insercion", Resultadofinal.DistanciaR, " ")
		CanalInsercion <- Resultadofinal
		wg.Done()
	}()

	wg.Add(1)
	go func() {

		print("Vecinadrio mas cercano ", VecinoCercano.AplicarVecindario(<-CanalVecino, IndiceNodos), " ")
		wg.Done()

	}()
	wg.Add(1)
	go func() {

		print("Vecinadrio Insercion mas cercano ", VecinoCercano.AplicarVecindario(<-CanalInsercion, IndiceNodos), " ")
		wg.Done()

	}()
	wg.Wait()

}
