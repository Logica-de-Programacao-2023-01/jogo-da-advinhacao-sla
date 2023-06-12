package main

import (
	"fmt"
	"math/rand"
)

func main() {
	result := rand.Intn(100) + 1
	var inp int
	fmt.Println("adivinhe um numero inteiro entre 1 e 100:")
	var s_n string
	X_Tentativas := []int{}
	tentativa := 1

	for {
		fmt.Scanln(&inp)
		if inp < result {
			fmt.Println("Esse resultado é menor, tente novamente:")
			tentativa++
		} else if inp > result {
			fmt.Println("Esse resultado é maior, tente novamente:")
			tentativa++
		} else {
			if inp == result {
				X_Tentativas = append(X_Tentativas, tentativa)
				fmt.Println("Parabens você ganhou!!")
				fmt.Printf("Você usou %d tentativas\n", tentativa)
				fmt.Println("Deseja jogar novamente? (s/n)")
				fmt.Scanln(&s_n)
				if s_n == "n" {
					for i, X_Tenta := range X_Tentativas {
						fmt.Printf("%d tentativas no jogo %d\n", X_Tenta, i+1)
					}
					break
				} else {
					tentativa = 1
					result = rand.Intn(100) + 1
					fmt.Println("Adivinhe um nÚmero inteiro entre 1 e 100:")
				}

			}
		}
	}
}
