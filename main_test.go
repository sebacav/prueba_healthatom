package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEsPar(t *testing.T) {
	par := ValidarPar(88)
	assert.True(t, par, "ValidarPar debería ser par")
	par = ValidarPar(3)
	assert.False(t, par, "ValidarPar debería ser par")
}

func TestObtenerMediana(t *testing.T) {

	numeros := []float64{60}
	mediana := ObtenerMediana(numeros)
	assert.Equal(t, mediana, 60.0, "Al enviar solo un valor a ObtenerMediana, debería entregar el mismo valor")

	numeros = []float64{60, 50, 80, 90}
	mediana = ObtenerMediana(numeros)
	assert.Equal(t, mediana, 70.0, "Al enviar una cantidad par de numero a ObtenerMediana, debería entregar el promedio de los dos del centro")

	numeros = []float64{60, 50, 100, 80, 90}
	mediana = ObtenerMediana(numeros)
	assert.Equal(t, mediana, 80.0, "Al enviar una cantidad impar de numero a ObtenerMediana, debería entregar el promedio de los dos del centro")
}

func TestHealthcheck(t *testing.T) {

	alarma := Healthcheck(2, 3)
	assert.True(t, alarma, "Al enviar el primer valor de CPU a Healthcheck, debería arrojar True")

	alarma = Healthcheck(5, 3)
	assert.True(t, alarma, "Al enviar el segundo valor de CPU a Healthcheck, debería arrojar True, si la CPU no esta sobre la mediana")

	alarma = Healthcheck(5, 3)
	assert.True(t, alarma, "Al enviar el tercer valor de CPU a Healthcheck, debería arrojar True, si la CPU no esta sobre la mediana")

	alarma = Healthcheck(12, 3)
	assert.False(t, alarma, "Al enviar el cuarto valor de CPU a Healthcheck, debería arrojar False, si la CPU esta bajo la mediana")

	alarma = Healthcheck(3, 3)
	assert.True(t, alarma, "Al enviar el quinto valor de CPU a Healthcheck, debería arrojar True, si la CPU no esta sobre la mediana")

	alarma = Healthcheck(4, 3)
	assert.True(t, alarma, "Al enviar el sexto valor de CPU a Healthcheck, debería arrojar True, si la CPU no esta sobre la mediana")
}

func TestObtenerReductorDeCuatros(t *testing.T) {

	resultado := ObtenerReductorDeCuatros("440446544")
	assert.Equal(t, resultado, 110110011.0, "Al enviar un numero como string con varios digitos 4, este debería devolver un numero que al restar el original, este queda sin digitos 4")

	resultado = ObtenerReductorDeCuatros("12356789")
	assert.Equal(t, resultado, 0.0, "Al enviar un numero como string sin digitos 4, este debería devolver un 0 ")

}

func TestRemoverCuatros(t *testing.T) {

	numero := 9463434.0
	parteA, parteB := RemoverCuatros(numero)
	assert.Equal(t, parteA, 9363333.0, "Al enviar un numero una de sus partes (A) debería devolver un numero que no contenga 4")
	assert.Equal(t, parteB, 100101.0, "Al enviar un numero una de sus partes (B) debería devolver un numero que no contenga 4")
	assert.Equal(t, parteA+parteB, numero, "Al enviar un numero una de sus partes (B) debería devolver un numero que no contenga 4")

	numero = 9563738.0
	parteA, parteB = RemoverCuatros(numero)
	assert.Equal(t, parteA, 9563738.0, "Al enviar un numero que no contenga numeros 4 su parte (A), debería devolver el mismo numero enviado")
	assert.Equal(t, parteB, 0.0, "Al enviar un numero que no contenga numeros 4 su parte (B), debería devolver un 0")
	assert.Equal(t, parteA+parteB, numero, "Al enviar un numero, debería devolver 2 numeros que al sumarlos den el mismo numero")

}
