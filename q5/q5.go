package q5

import "strings"

//Pedro começou a frequentar aulas de programação. Na primeira aula, sua tarefa foi escrever um programa simples. O
//programa deveria fazer o seguinte: na sequência de caracteres fornecida, composta por letras latinas maiúsculas e
//minúsculas, ele:
//
//* deleta todas as vogais;
//* insere um caractere "." antes de cada consoante;
//* substitui todas as consoantes maiúsculas pelas correspondentes em minúsculas.
//
//As vogais são as letras "A", "O", "E", "U", "I", e o restante são consoantes. A entrada do programa é exatamente uma
//sequência de caracteres e a saída deve ser uma única sequência de caracteres, resultante após o processamento do
//programa na sequência de caracteres inicial.
//
//Ajude Pedro a lidar com esta tarefa fácil.

func ProcessString(s string) string {
	strings.ReplaceAll(s, "A", "")
	strings.ReplaceAll(s, "E", "")
	strings.ReplaceAll(s, "I", "")
	strings.ReplaceAll(s, "O", "")
	strings.ReplaceAll(s, "U", "")
	strings.ReplaceAll(s, "a", "")
	strings.ReplaceAll(s, "e", "")
	strings.ReplaceAll(s, "i", "")
	strings.ReplaceAll(s, "o", "")
	strings.ReplaceAll(s, "u", "")

	strings.ToLower(s)

	str := ""

	for i := range s {
		if strings.Contains("BCDFGHJKLMNPQRSTVWXYZbcdfghjklmnpqrstvwxyz", string(s[i])) {
			res := "." + strings.ToLower(string(s[i]))
			str += res
		}
	}

	return str

	return ""
}
