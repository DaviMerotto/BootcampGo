package main

import (
	"fmt"
	"time"
)

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
	palavra := ""
	fmt.Print("Digite uma palavra:")
	fmt.Scanln(&palavra)
	fmt.Println("Tamanho da palavra digitada é:", len(palavra))
	fmt.Println("Soletrando a palavra:", palavra)
	for i := 0; i < len(palavra); i++ {
		fmt.Printf("%s ", string(palavra[i]))
	}
	fmt.Println()
}

func exercicio2() {
	type Pessoa struct {
		nome        string
		sobrenome   string
		idade       int
		empregado   bool
		contratacao string
		salario     float64
	}

	pessoa1 := Pessoa{
		nome:        "Lucas",
		sobrenome:   "Silva",
		idade:       25,
		empregado:   false,
		contratacao: "02/01/2006",
		salario:     2500.00,
	}
	pessoa2 := Pessoa{
		nome:        "Matheus",
		sobrenome:   "Silva",
		idade:       21,
		empregado:   true,
		contratacao: "02/09/2023",
		salario:     9000.00,
	}
	pessoa3 := Pessoa{
		nome:        "João",
		sobrenome:   "Silva",
		idade:       35,
		empregado:   true,
		contratacao: "02/09/2020",
		salario:     100900.00,
	}

	numPessoa := 0
	fmt.Println("Qual das pessoas deseja consultar financiamento? 1, 2 ou 3")
	fmt.Scanln(&numPessoa)
	switch numPessoa {
	case 1:
		contratacao, _ := time.Parse("02/01/2006", pessoa1.contratacao)
		oneYearAgo := time.Now().AddDate(-1, 0, 0)
		if pessoa1.idade > 22 && contratacao.Before(oneYearAgo) && pessoa1.empregado {
			if pessoa1.salario > 100000.00 {
				fmt.Println("Financiamento aprovado e sem juros para pessoa 1")
			} else {
				fmt.Println("Financiamento aprovado para pessoa 1")
			}
		} else {
			fmt.Println("Financiamento negado para pessoa 1")
		}
	case 2:
		contratacao, _ := time.Parse("02/01/2006", pessoa2.contratacao)
		oneYearAgo := time.Now().AddDate(-1, 0, 0)
		if pessoa2.idade > 22 && contratacao.Before(oneYearAgo) && pessoa2.empregado {
			if pessoa2.salario > 100000.00 {
				fmt.Println("Financiamento aprovado e sem juros para pessoa 2")
			} else {
				fmt.Println("Financiamento aprovado para pessoa 2")
			}
		} else {
			fmt.Println("Financiamento negado para pessoa 2")
		}
	case 3:
		contratacao, _ := time.Parse("02/01/2006", pessoa3.contratacao)
		oneYearAgo := time.Now().AddDate(-1, 0, 0)
		if pessoa3.idade > 22 && contratacao.Before(oneYearAgo) && pessoa3.empregado {
			if pessoa3.salario > 100000.00 {
				fmt.Println("Financiamento aprovado e sem juros para pessoa 3")
			} else {
				fmt.Println("Financiamento aprovado para pessoa 3")
			}
		} else {
			fmt.Println("Financiamento negado para pessoa 3")
		}
	}
}

func exercicio3() {
	data := ""
	fmt.Println("Digite uma data no formato dd/mm/yyyy:")
	fmt.Scanln(&data)
	date, _ := time.Parse("02/01/2006", data)
	fmt.Printf("Day: %d Month: %s \n", date.Day(), date.Month())

	// Ao meu ver essa é a maneira menos complexa de trazer o nome do mês (em inglês)
	// Se precisar do nome do mês em português, teria que fazer um map com os nomes dos meses
}

func exercicio4() {
	var employees = map[string]int{"Benjamin": 20, "Manuel": 26, "Brenda": 19, "Dario": 44, "Pedro": 30}
	fmt.Printf("A idade de Benjamin é %d\n", employees["Benjamin"])
	plusTwentyOneYears := make(map[string]int)
	for name, age := range employees {
		if age > 21 {
			plusTwentyOneYears[name] = age
		}
	}
	fmt.Printf("Á exatos %d funcionarios com mais de 21 anos\n", len(plusTwentyOneYears))
	fmt.Printf("Adicionando Federico ao map de funcionarios\n")
	employees["Federico"] = 25
	fmt.Printf("Agora temos %d funcionarios\n", len(employees))
	fmt.Printf("Removendo Pedro do map de funcionarios\n")
	delete(employees, "Pedro")
	fmt.Printf("Agora temos %d funcionarios\n", len(employees))
	fmt.Printf("Listando todos os funcionarios e suas idades\n")
	for name, age := range employees {
		fmt.Printf("Nome: %s Idade: %d\n", name, age)
	}
}
