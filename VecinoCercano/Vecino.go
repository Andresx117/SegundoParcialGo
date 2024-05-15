package vecinocercano

import (
	"math/rand"
	"time"

	gestorarchivos "github.com/Andresx117/SegundoParcialGo/GestorArchivos"
)

func Calculo(CanalNodo <-chan []gestorarchivos.Nodo) {
	rand.Seed(time.Now().UnixNano())
	IndiceNodos := <-CanalNodo
	IndiceAleatorio := rand.Intn(len(IndiceNodos))
	prim := IndiceNodos[:IndiceAleatorio]
	Sec := IndiceNodos[IndiceAleatorio+1:]
}
