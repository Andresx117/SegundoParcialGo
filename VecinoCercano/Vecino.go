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

// InsercionMasCercana calcula la ruta óptima utilizando el algoritmo de inserción más cercana
func InsercionMasCercana(nodos []gestorarchivos.Nodo) ([]gestorarchivos.Distancia, float64) {
	if len(nodos) == 0 {
		return nil, 0
	}

	// Empezamos con un ciclo que incluye el primer nodo
	var ruta []gestorarchivos.Distancia
	totalDistancia := 0.0

	visitados := make(map[string]bool)
	visitados[nodos[0].Nombre] = true

	if len(nodos) > 1 {
		visitados[nodos[1].Nombre] = true
		ruta = append(ruta, gestorarchivos.Distancia{
			NodoI:     nodos[0].Nombre,
			NodoFinal: nodos[1].Nombre,
			Distancia: distanciaEuclidiana(nodos[0], nodos[1]),
		})
		ruta = append(ruta, gestorarchivos.Distancia{
			NodoI:     nodos[1].Nombre,
			NodoFinal: nodos[0].Nombre,
			Distancia: distanciaEuclidiana(nodos[1], nodos[0]),
		})
		totalDistancia = 2 * distanciaEuclidiana(nodos[0], nodos[1])
	}

	// Inserción más cercana
	for len(visitados) < len(nodos) {
		nodoMasCercano := gestorarchivos.Nodo{}
		minIncremento := math.MaxFloat64
		posicion := 0

		// Encontrar el nodo no visitado más cercano y la mejor posición para insertarlo
		for _, nodo := range nodos {
			if !visitados[nodo.Nombre] {
				for i := 0; i < len(ruta); i++ {
					nodoI := encontrarNodo(ruta[i].NodoI, nodos)
					nodoF := encontrarNodo(ruta[i].NodoFinal, nodos)
					incremento := distanciaEuclidiana(nodoI, nodo) + distanciaEuclidiana(nodo, nodoF) - ruta[i].Distancia
					if incremento < minIncremento {
						nodoMasCercano = nodo
						minIncremento = incremento
						posicion = i
					}
				}
			}
		}

		// Insertar el nodo en la posición encontrada
		if (nodoMasCercano != gestorarchivos.Nodo{}) {
			nodoI := encontrarNodo(ruta[posicion].NodoI, nodos)
			nodoF := encontrarNodo(ruta[posicion].NodoFinal, nodos)
			ruta = append(ruta[:posicion+1], ruta[posicion:]...) // Hacer espacio para la nueva distancia
			ruta[posicion] = gestorarchivos.Distancia{
				NodoI:     nodoI.Nombre,
				NodoFinal: nodoMasCercano.Nombre,
				Distancia: distanciaEuclidiana(nodoI, nodoMasCercano),
			}
			ruta[posicion+1] = gestorarchivos.Distancia{
				NodoI:     nodoMasCercano.Nombre,
				NodoFinal: nodoF.Nombre,
				Distancia: distanciaEuclidiana(nodoMasCercano, nodoF),
			}
			totalDistancia += minIncremento
			visitados[nodoMasCercano.Nombre] = true
		}
	}

	return ruta, totalDistancia
}

// Encuentra un nodo por su nombre
func encontrarNodo(nombre string, nodos []gestorarchivos.Nodo) gestorarchivos.Nodo {
	for _, nodo := range nodos {
		if nodo.Nombre == nombre {
			return nodo
		}
	}
	return gestorarchivos.Nodo{}
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
