package main

import (
	"fmt"
	"log"
	"math"
	"sort"
	"strconv"
)

// HistoricoCPU es un vector de números flotantes
// se usa como memoria, para la función Healthcheck
var HistoricoCPU []float64

func main() {

	// a, b := RemoverCuatros(math.Pow10(101))

	// transformamos los numeros a un string, de manera de poderlo leer
	// de una manera bonita en la pantalla
	// aString := strconv.FormatFloat(a, 'f', -1, 64)
	// bString := strconv.FormatFloat(b, 'f', -1, 64)

	a, b := RemoverCuatros(9463434)
	fmt.Println(a, b)

	fmt.Println(Healthcheck(2, 3))
	fmt.Println(Healthcheck(5, 3))
	fmt.Println(Healthcheck(5, 3))
	fmt.Println(Healthcheck(12, 3))
	fmt.Println(Healthcheck(3, 3))
	fmt.Println(Healthcheck(4, 3))

}

// RemoverCuatros recibe un numero, le quita los 4 y luego devuelve
// dos numeros que sumados dan el numero original, para esta funcion
// lo que hacemos es basicamente restar unos en cada posicion de los 4
// lo hacemos con potencias de 10 elevadas a la posicion de cada numero 4
// para que de esta manera podamos al final, restar y obtener siempre
// un número que no tendrá cuatros en sus cifras
func RemoverCuatros(numero float64) (float64, float64) {
	if numero < 1 || numero > math.Pow10(100) {
		log.Fatal("Input no Validos")
	}
	// Se transforma el numero a un string
	numeroString := strconv.FormatFloat(numero, 'f', -1, 64)

	// Obtenemos b, que simboliza el número que necesito
	// para que, al restar, no queden dígitos 4

	b := ObtenerReductorDeCuatros(numeroString)

	// Creamos la variable "a" que sería igual al numero entrante
	// menos el número que nos quitara los dígitos 4

	a := numero - b

	return a, b
}

// ObtenerReductorDeCuatros devuelve un número que al restar el
// numero original permitira remover todos los digitos 4 del numero
func ObtenerReductorDeCuatros(numero string) float64 {
	// Se obtiene el largo del número menos 1
	largoNumero := len(numero) - 1

	var restaNecesaria float64
	// Se recorre el numero
	for i, r := range numero {
		// convertimos el byte a string
		c := string(r)
		// validamos si es un 4
		if c == "4" {
			// Agregamos lo necesario, para restarle al 4 un digito
			// utilizando una potencia de 10 elevado al largo - posicion (i)
			// tantas veces como 4 aparescan a lo largo del recorrido
			// sumando cada uno de ellos, para formar un numero
			// que, al restarlo, siempre elimine los 4 del número original
			restaNecesaria += math.Pow10(largoNumero - i)
		}
	}
	return restaNecesaria
}

// Healthcheck recibe el % de cpu en un instante y el largo de la ventana
// en la cual se evaluará, basado en la mediana de esta ventana
// si es momento de alarmar o no. el HistoricoCPU es una variable global
// esta se va llenando cada vez que se llama a esta funcion (Healthcheck)
func Healthcheck(cpu float64, largoVentana int) bool {
	if cpu < 0 || cpu > 100 || largoVentana < 0 || float64(largoVentana) > math.Pow10(100) {
		log.Fatal("Input no Validos")
	}

	// Se declara la mediana
	var median float64

	// Se agrega la cpu actual al histórico de cpu
	HistoricoCPU = append(HistoricoCPU, cpu)

	// Esta variable cuando es superior a 0 nos dice cuanto debemos recortar
	recorteVentana := len(HistoricoCPU) - largoVentana
	// Si recorteVentana es mayor a 0, es necesario obtener una ventana
	if recorteVentana >= 0 {
		// Recortamos desde atrás para adelante
		median = ObtenerMediana(HistoricoCPU[recorteVentana:])
	} else {
		// No recortamos, debido a que tenemos menos numeros que la ventana mínima
		median = ObtenerMediana(HistoricoCPU)
	}
	// Se hace el calculo que solicita el requerimiento
	if cpu >= 2*median {
		// Se retorna false, que es un equivalente al 0
		return false // 0
	}
	// Se retorna true, que es un equivalente al 1
	return true // 1
}

// ObtenerMediana nos permite obtener la mediana dado los requerimientos
// la funcion recibe un arreglo de flotantes y retorna un flotante que
// simboliza la mediana
func ObtenerMediana(numeros []float64) float64 {

	// Se declara mediana
	var mediana float64

	// Se hace un ordenamiento ascendente
	sort.Float64s(numeros)

	// Se genera una variable que contiene el largo
	largo := len(numeros)

	// Si el largo es == 1, entonces se retorna instantáneamente
	// el único valor del arreglo, para evitar cálculos
	if largo == 1 {
		mediana = numeros[0]

		// Si el largo es mayor a 1, entonces se calcula
	} else if largo > 1 {

		// Se consulta si el largo es par o impar
		par := ValidarPar(largo)

		// Si es par entonces se suma el promedio de los valores del centro
		if par {
			mediana = (numeros[(largo/2)-1] + numeros[(largo/2)]) / 2

			// Si es impar, entonces solo se obtiene el valor central
		} else {
			mediana = numeros[int((largo / 2))]
		}
	}

	return mediana
}

// ValidarPar entrega dado un valor entero, si este es par o impar
func ValidarPar(valor int) bool {
	if valor%2 == 0 {
		return true
	}
	return false

}
