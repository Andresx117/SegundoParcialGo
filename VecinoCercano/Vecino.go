package VecinoCercano

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"

	gestorarchivos "github.com/Andresx117/SegundoParcialGo/GestorArchivos"
)

// Calcula la distancia euclidiana entre dos nodos
func distanciaEuclidiana(nodo1, nodo2 gestorarchivos.Nodo) float64 {
	deltaX := nodo1.CoorX - nodo2.CoorX
	deltaY := nodo1.CoorY - nodo2.CoorY
	return math.Sqrt(deltaX*deltaX + deltaY*deltaY)
}

// VecinoMásCercano calcula la ruta óptima utilizando el algoritmo del vecino más cercano
func VecinoMasCercano(nodos []gestorarchivos.Nodo) ([]gestorarchivos.Distancia, float64) {
	if len(nodos) == 0 {
		return nil, 0
	}

	visitados := make(map[string]bool)
	var ruta []gestorarchivos.Distancia
	totalDistancia := 0.0

	nodoActual := nodos[0]
	visitados[nodoActual.Nombre] = true

	for len(visitados) < len(nodos) {
		nodoMasCercano := gestorarchivos.Nodo{}
		minDistancia := math.MaxFloat64

		for _, nodo := range nodos {
			if !visitados[nodo.Nombre] {
				dist := distanciaEuclidiana(nodoActual, nodo)
				if dist < minDistancia {
					nodoMasCercano = nodo
					minDistancia = dist
				}
			}
		}

		if (nodoMasCercano != gestorarchivos.Nodo{}) {
			ruta = append(ruta, gestorarchivos.Distancia{
				NodoI:     nodoActual.Nombre,
				NodoFinal: nodoMasCercano.Nombre,
				Distancia: minDistancia,
			})
			totalDistancia += minDistancia
			nodoActual = nodoMasCercano
			visitados[nodoActual.Nombre] = true
		}
	}

	// Regresar al nodo inicial para completar el ciclo
	ruta = append(ruta, gestorarchivos.Distancia{
		NodoI:     nodoActual.Nombre,
		NodoFinal: nodos[0].Nombre,
		Distancia: distanciaEuclidiana(nodoActual, nodos[0]),
	})
	totalDistancia += distanciaEuclidiana(nodoActual, nodos[0])

	return ruta, totalDistancia
}

// Calcula las distancias entre los nodos en la lista dada
func calcularDistancias(nodos []gestorarchivos.Nodo, distancias *[]gestorarchivos.Distancia, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < len(nodos)-1; i++ {
		for j := i + 1; j < len(nodos); j++ {
			distancia := distanciaEuclidiana(nodos[i], nodos[j])
			*distancias = append(*distancias, gestorarchivos.Distancia{
				NodoI:     nodos[i].Nombre,
				NodoFinal: nodos[j].Nombre,
				Distancia: distancia,
			})
		}
	}
}

func Calculo(IndiceNodos []gestorarchivos.Nodo) ([]gestorarchivos.Distancia, []gestorarchivos.Distancia) {
	rand.Seed(time.Now().UnixNano())
	IndiceAleatorio := rand.Intn(len(IndiceNodos))
	prim := IndiceNodos[:IndiceAleatorio]
	Sec := IndiceNodos[IndiceAleatorio-1:]

	var distanciasPrim []gestorarchivos.Distancia
	var distanciasSec []gestorarchivos.Distancia
	var wg sync.WaitGroup

	wg.Add(2)
	fmt.Println(IndiceNodos[IndiceAleatorio])

	go calcularDistancias(prim, &distanciasPrim, &wg)
	go calcularDistancias(Sec, &distanciasSec, &wg)

	wg.Wait()

	return distanciasPrim, distanciasSec
}
