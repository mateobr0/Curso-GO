package numeros

import (
	"fmt"
	"sort"
)

// Desarrolla un programa que tome una lista de palabras y las ordene alfabéticamente.
func Ordenar(x []string) []string {
	sort.Strings(x)
	return x
}

// Escribe una función que tome un arreglo de enteros como parámetro
// y devuelva la suma de los números pares y la suma de los números impares por separado.
func SumaParesEImpares(x []int) (int, int) {
	var (
		pares   int
		impares int
	)
	for _, v := range x {
		if v%2 == 0 {
			pares = pares + v
		} else {
			impares = impares + v
		}
	}
	return pares, impares
}

// Implementa una función que tome un número entero positivo como parámetro
// y devuelva true si es un número perfecto y false en caso contrario.
// Un número perfecto es aquel cuya suma de sus divisores propios
// (excluyendo al propio número) es igual al número.
func EsPerfecto(x uint) bool {
	z := int(x)
	var divisoresPropios int
	for i := 1; i < z; i++ {
		if z%i == 0 {
			divisoresPropios += i
		}
	}

	return divisoresPropios == z
}

// Crea una función que tome dos slices de enteros como parámetros
// y devuelva un nuevo slice que contenga los elementos comunes entre ambos slices.
func ElementosComunes(x, z []int) []int {
	var (
		menor       []int = x
		mayor       []int = z
		comunes     []int
		iteraciones int
	)

	if len(x) > len(z) {
		menor = z
		mayor = x
	}
	//al poner el slice menor en el ciclo exterior las iteraciones disminuyen
	for _, vx := range menor {
		iteraciones++
		for _, vz := range mayor {
			iteraciones++
			if vx == vz {
				comunes = append(comunes, vz)
			}
		}
	}
	fmt.Println("Iteraciones en la func ElementosComunes", iteraciones)
	return comunes
}

// ahora usando map
func ElementosComunesMap(x, z []int) []int {
	var (
		menor       []int = x
		mayor       []int = z
		comunes     []int
		iteraciones int
	)

	if len(x) > len(z) {
		menor = z
		mayor = x
	}

	elementos := make(map[int]bool)
	for _, v := range menor {
		iteraciones++
		elementos[v] = true
	}

	for _, v := range mayor {
		iteraciones++
		if elementos[v] {
			comunes = append(comunes, v)
		}
	}
	fmt.Println("Iteraciones en la func ElementosComunesMap ", iteraciones)
	return comunes
}