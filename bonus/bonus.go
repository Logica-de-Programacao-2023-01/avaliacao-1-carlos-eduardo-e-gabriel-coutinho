package bonus

import "sort"

//Joãozinho ganhou um kit de construção de torres. O kit é composto por várias barras de madeira, e seus comprimentos são
//conhecidos. As barras podem ser empilhadas umas sobre as outras, desde que seus comprimentos sejam iguais.
//
//Joãozinho quer construir o menor número possível de torres com as barras que tem. Você deve ajudar Joãozinho a usar as
//barras da melhor maneira possível, determinando a altura da torre mais alta e quantas torres podem ser construídas.

func CalculateTowers(barLengths []int) (int, int) {
	sort.Ints(barLengths)
	var tor = 1
	var alt = 1
	var torAtual = 1
	for i := 1; i < len(barLengths); i++ {
		if barLengths[i] != barLengths[i-1] {
			tor++
			torAtual = 1
		} else {
			torAtual++
		}
		if torAtual >= alt {
			alt = torAtual
		}
	}

	return alt, tor

	return 0, 0
}
