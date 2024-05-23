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
func VecinoMasCercano(nodos []gestorarchivos.Nodo, indice int) ([]gestorarchivos.Distancia, float64) {
	if len(nodos) == 0 {
		return nil, 0
	}

	visitados := make(map[string]bool)
	var ruta []gestorarchivos.Distancia
	totalDistancia := 0.0

	nodoActual := nodos[indice]
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
		NodoFinal: nodos[indice].Nombre,
		Distancia: distanciaEuclidiana(nodoActual, nodos[indice]),
	})
	totalDistancia += distanciaEuclidiana(nodoActual, nodos[indice])

	return ruta, totalDistancia
}

// InsercionMasCercana calcula la ruta óptima utilizando el algoritmo de inserción más cercana
func InsercionMasCercana(nodos []gestorarchivos.Nodo, indice int) ([]gestorarchivos.Distancia, float64) {
	if len(nodos) == 0 {
		return nil, 0
	}

	// Empezamos con un ciclo que incluye el primer nodo
	var ruta []gestorarchivos.Distancia
	totalDistancia := 0.0

	visitados := make(map[string]bool)
	visitados[nodos[indice].Nombre] = true

	if len(nodos) > 1 {
		visitados[nodos[indice+1].Nombre] = true
		ruta = append(ruta, gestorarchivos.Distancia{
			NodoI:     nodos[indice].Nombre,
			NodoFinal: nodos[indice+1].Nombre,
			Distancia: distanciaEuclidiana(nodos[indice], nodos[indice+1]),
		})
		ruta = append(ruta, gestorarchivos.Distancia{
			NodoI:     nodos[indice+1].Nombre,
			NodoFinal: nodos[indice].Nombre,
			Distancia: distanciaEuclidiana(nodos[indice+1], nodos[indice]),
		})
		totalDistancia = 2 * distanciaEuclidiana(nodos[indice], nodos[indice+1])
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

func DistanciaEuclidianaVecindario(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
}

// Función auxiliar para calcular la distancia total de una ruta
func CalcularDistanciaTotal(ruta []gestorarchivos.Distancia) float64 {
	total := 0.0
	for _, distancia := range ruta {
		total += distancia.Distancia
	}
	return total
}

// Función para encontrar las coordenadas de un nodo en el slice
func EncontrarCoordenadas(nombre string, nodos []gestorarchivos.Nodo) (float64, float64) {
	for _, nodo := range nodos {
		if nodo.Nombre == nombre {
			return nodo.CoorX, nodo.CoorY
		}
	}
	return 0, 0
}

// Función para aplicar el método de vecindario
func AplicarVecindario(resultado gestorarchivos.Resultado, nodos []gestorarchivos.Nodo) float64 {
	ruta := resultado.RutaR
	mejorRuta := ruta
	mejorDistancia := resultado.DistanciaR

	for i := 0; i < len(ruta)-1; i++ {
		for j := i + 1; j < len(ruta); j++ {
			// Crear una nueva ruta intercambiando los nodos
			nuevaRuta := make([]gestorarchivos.Distancia, len(ruta))
			copy(nuevaRuta, ruta)
			nuevaRuta[i], nuevaRuta[j] = nuevaRuta[j], nuevaRuta[i]

			// Recalcular las distancias para los nodos intercambiados
			if i > 0 {
				x1, y1 := EncontrarCoordenadas(nuevaRuta[i-1].NodoI, nodos)
				x2, y2 := EncontrarCoordenadas(nuevaRuta[i-1].NodoFinal, nodos)
				nuevaRuta[i-1].Distancia = DistanciaEuclidianaVecindario(x1, y1, x2, y2)
			}
			x1, y1 := EncontrarCoordenadas(nuevaRuta[i].NodoI, nodos)
			x2, y2 := EncontrarCoordenadas(nuevaRuta[i].NodoFinal, nodos)
			nuevaRuta[i].Distancia = DistanciaEuclidianaVecindario(x1, y1, x2, y2)
			if i < len(ruta)-1 {
				x1, y1 := EncontrarCoordenadas(nuevaRuta[i+1].NodoI, nodos)
				x2, y2 := EncontrarCoordenadas(nuevaRuta[i+1].NodoFinal, nodos)
				nuevaRuta[i+1].Distancia = DistanciaEuclidianaVecindario(x1, y1, x2, y2)
			}
			x1, y1 = EncontrarCoordenadas(nuevaRuta[j].NodoI, nodos)
			x2, y2 = EncontrarCoordenadas(nuevaRuta[j].NodoFinal, nodos)
			nuevaRuta[j].Distancia = DistanciaEuclidianaVecindario(x1, y1, x2, y2)
			if j < len(ruta)-1 {
				x1, y1 := EncontrarCoordenadas(nuevaRuta[j+1].NodoI, nodos)
				x2, y2 := EncontrarCoordenadas(nuevaRuta[j+1].NodoFinal, nodos)
				nuevaRuta[j+1].Distancia = DistanciaEuclidianaVecindario(x1, y1, x2, y2)
			}

			nuevaDistancia := CalcularDistanciaTotal(nuevaRuta)
			if nuevaDistancia < mejorDistancia {
				mejorRuta = nuevaRuta
				mejorDistancia = nuevaDistancia
			}
		}
	}

	resultado.RutaR = mejorRuta
	resultado.DistanciaR = mejorDistancia
	return resultado.DistanciaR
}
