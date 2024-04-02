package main

import "fmt"

func main() {
	for {
		exercicio := "0"
		fmt.Println("Lista de exercicios: 1, 2, 3, 4")
		fmt.Println("Digite o número do exercicio que deseja executar ou 0 para sair:")
		fmt.Scanln(&exercicio)
		if exercicio == "0" {
			break
		}
		switch exercicio {
		case "1":
			exercicio1()
		case "2":
			exercicio2()
		case "3":
			exercicio3()
		case "4":
			exercicio4()
		default:
			fmt.Println("Exercicio não encontrado")
		}
	}
}

func exercicio1() {
	var idade int = 25
	var name string = "Davi"
	fmt.Printf("Nome: %s Idade: %d", name, idade)
}

func exercicio2() {
	var temperatura int = 27
	var umidade uint = 36
	var pressao float32 = 1018.25
	fmt.Printf("temperatura %dC, umidade %d%%, pressao %.2f hPa", temperatura, umidade, pressao)
}

func exercicio3() {
	//var 1nome string
	//var sobrenome string
	//var int idade
	//1sobrenome := 6
	//var licenca_para_dirigir = true
	//var estatura da pessoa int
	//quantidadeDeFilhos := 2

	fmt.Println("As variaveis que estão declaradas corretas são sobrenome, licenca_para_dirigir e quantidadeDeFilhos")
	fmt.Println("As variaveis que não estão declaradas corretamente são 1nome, 1sobrenome, idade e estatura da pessoa")
	fmt.Println("Variaveis não podem começar com numeros, portanto 1nome deveria ser somente nome")
	fmt.Println("1sobrenome deveria ser somente sobrenome, e do tipo string, fora que não deveria uma nova declaração via inferencia, pois a variavel já estava declarada")
	fmt.Println("idade esta declarado de forma errada, com o tipo primeiro depois o nome, e deveria ser o contrario")
	fmt.Println("estatura da pessoa deveria estar sem espaços, estaturaDaPessoa por exemplo")

	fmt.Printf("segue o codigo corrigido \n\n")

	fmt.Println("var nome string")
	fmt.Println("var sobrenome string")
	fmt.Println("var idade int")
	fmt.Println("sobrenome = \"teste\"")
	fmt.Println("var licenca_para_dirigir = true")
	fmt.Println("var estaturaDaPessoa int")
	fmt.Println("quantidadeDeFilhos := 2")
}

func exercicio4() {
	//var sobrenome string = "Silva"
	//var idade int = "25"
	//boolean := "false";
	//var salario string = 4585.90
	//var nome string = "Fellipe"

	fmt.Println("As variaveis que estão declaradas corretas são sobrenome, nome")
	fmt.Println("As variaveis que não estão declaradas corretamente são idade, boolean e salario")
	fmt.Println("idade esta declarado de forma errada, com o tipo int e o valor entre aspas, deveria ser somente o valor exemplo 25")
	fmt.Println("boolean esta declarado de forma errada, com o tipo string e o valor entre aspas, deveria ser somente o valor exemplo false que atribuiria o tipo bool, remover também o ;")
	fmt.Println("salario esta declarado de forma errada, com o tipo string e o valor float, deveria ser somente o valor exemplo 4585.90 e estar declarado como float")
	fmt.Printf("Segue o codigo corrigido\n\n")

	fmt.Println("var sobrenome string = \"Silva\"")
	fmt.Println("var idade int = 25")
	fmt.Println("boolean := false")
	fmt.Println("var salario float = 4585.20")
	fmt.Println("var nome string = \"Fellipe\"")
}
