package VecinoCercano

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"

	"github.com/Andresx117/SegundoParcialGo/GestorArchivos"
)

// Calculo realiza el cálculo de las distancias utilizando el algoritmo de búsqueda de vecindario
func Calculo(IndiceNodos []GestorArchivos.Nodo) ([]GestorArchivos.Distancia, []GestorArchivos.Distancia) {
	rand.Seed(time.Now().UnixNano())
	IndiceAleatorio := rand.Intn(len(IndiceNodos))
	prim := IndiceNodos[:IndiceAleatorio]
	Sec := IndiceNodos[IndiceAleatorio-1:]

	var distanciasPrim []GestorArchivos.Distancia
	var distanciasSec []GestorArchivos.Distancia
	var wg sync.WaitGroup

	wg.Add(2)
	fmt.Println(IndiceNodos[IndiceAleatorio])

	go calcularDistancias(prim, &distanciasPrim, &wg)
	go calcularDistancias(Sec, &distanciasSec, &wg)

	wg.Wait()

	return distanciasPrim, distanciasSec
}

// calcularDistancias calcula las distancias entre los nodos en la lista dada
func calcularDistancias(nodos []GestorArchivos.Nodo, distancias *[]GestorArchivos.Distancia, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < len(nodos)-1; i++ {
		for j := i + 1; j < len(nodos); j++ {
			distancia := distanciaEuclidiana(nodos[i], nodos[j])
			*distancias = append(*distancias, GestorArchivos.Distancia{
				NodoI:     nodos[i].Nombre,
				NodoFinal: nodos[j].Nombre,
				Distancia: distancia,
			})
		}
	}
}

// distanciaEuclidiana calcula la distancia euclidiana entre dos nodos
func distanciaEuclidiana(nodo1, nodo2 GestorArchivos.Nodo) float64 {
	deltaX := nodo1.CoorX - nodo2.CoorX
	deltaY := nodo1.CoorY - nodo2.CoorY
	return math.Sqrt(deltaX*deltaX + deltaY*deltaY)
}

/*func busquedaVecindario(nodos []gestorarchivos.Nodo) ([]gestorarchivos.Distancia, float64) {
	// Inicializar la ruta como una permutación de los nodos
	ruta := make([]gestorarchivos.Nodo, len(nodos))
	copy(ruta, nodos)

	// Calcular la distancia total inicial
	distanciaTotal := calcularDistanciaTotal(ruta)

	// Variable para indicar si se realizó un intercambio en la iteración anterior
	intercambio := true

	// Realizar iteraciones hasta que no se realice ningún intercambio en una iteración completa
	for intercambio {
		intercambio = false
		for i := 0; i < len(ruta)-1; i++ {
			for j := i + 1; j < len(ruta); j++ {
				// Aplicar el intercambio
				ruta[i], ruta[j] = ruta[j], ruta[i]

				// Calcular la nueva distancia total
				nuevaDistancia := calcularDistanciaTotal(ruta)

				// Si la nueva distancia es menor, se acepta el intercambio
				if nuevaDistancia < distanciaTotal {
					distanciaTotal = nuevaDistancia
					intercambio = true
				} else {
					// Si no es menor, se deshace el intercambio
					ruta[i], ruta[j] = ruta[j], ruta[i]
				}
			}
		}
	}

	// Construir la lista de distancias basada en la ruta final
	var distancias []gestorarchivos.Distancia
	for i := 0; i < len(ruta)-1; i++ {
		distancia := gestorarchivos.Distancia{
			NodoI:     ruta[i].Nombre,
			NodoFinal: ruta[i+1].Nombre,
			Distancia: distanciaEuclidiana(ruta[i], ruta[i+1]),
		}
		distancias = append(distancias, distancia)
	}

	// Agregar la distancia desde el último nodo hasta el primero para cerrar el ciclo
	distanciaFinal := gestorarchivos.Distancia{
		NodoI:     ruta[len(ruta)-1].Nombre,
		NodoFinal: ruta[0].Nombre,
		Distancia: distanciaEuclidiana(ruta[len(ruta)-1], ruta[0]),
	}
	distancias = append(distancias, distanciaFinal)

	return distancias, distanciaTotal
}*/

/*func calcularDistanciaTotal(ruta []gestorarchivos.Nodo) float64 {
	distanciaTotal := 0.0
	for i := 0; i < len(ruta)-1; i++ {
		distanciaTotal += distanciaEuclidiana(ruta[i], ruta[i+1])
	}
	// Agregar la distancia desde el último nodo hasta el primero para cerrar el ciclo
	distanciaTotal += distanciaEuclidiana(ruta[len(ruta)-1], ruta[0])
	return distanciaTotal
}*/
