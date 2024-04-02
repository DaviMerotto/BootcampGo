package main

import (
	"errors"
	"fmt"
)

func main() {
	for {
		exercicio := "0"
		fmt.Println("Lista de exercicios: 1, 2, 3, 4, 5")
		fmt.Println("Digite o número do exercicio que deseja executar ou 0 para sair:")
		fmt.Scanln(&exercicio)
		if exercicio == "0" {
			break
		}
		switch exercicio {
		case "1":
			exercicio1()
		case "2":
			fmt.Println(exercicio2())
		case "3":
			fmt.Println(exercicio3())
		case "4":
			exercicio4()
		case "5":
			exercicio5()
		default:
			fmt.Println("Exercicio não encontrado")
		}
	}
}

func exercicio1() {
	//Uma empresa de chocolates precisa calcular o imposto de seus funcionários no momento de depositar o salário, para cumprir seu objetivo será necessário criar uma função que retorne o imposto de um salário.
	//Temos a informação que salários até R$50.000 não serão descontados impostos. Salários entre R$50.000 e R$150.000 são descontados 17%. Salários acima de R$150.000 serão descontados 27%.

	salario, imposto := 0.0, 0.0
	fmt.Print("Digite o salário:")
	fmt.Scanln(&salario)
	if salario <= 50000 {
		imposto = 0
	} else if salario > 50000 && salario <= 150000 {
		imposto = salario * 0.17
	} else {
		imposto = salario * 0.27
	}
	fmt.Println("O imposto do salário é:", imposto)
}

func exercicio2() error {
	//Um colégio precisa calcular a média das notas (por aluno). Precisamos criar uma função na qual possamos passar N quantidade de números inteiros e devolva a média.
	//Obs: Caso um dos números digitados seja negativo, a aplicação deve retornar um erro

	notas := []int{}
	media := 0
	qtdNotas := 0
	fmt.Printf("Digite a quantidade de notas do aluno:")
	fmt.Scanln(&qtdNotas)
	fmt.Printf("Digite as notas do aluno:")
	for range qtdNotas {
		nota := 0
		fmt.Scanln(&nota)
		if nota < 0 {
			return errors.New("Erro ao calcular a média devido a nota negativa")
		} else {
			notas = append(notas, nota)
		}
	}

	for _, nota := range notas {
		media += nota
	}

	fmt.Printf("quantidade de notas: %d\n", qtdNotas)
	fmt.Printf("A media do aluno foi de: %d\n", media/qtdNotas)
	return nil
}

func exercicio3() error {
	//Uma empresa marítima precisa calcular o salário de seus funcionários com base no número de horas 	trabalhadas por mês e na categoria do funcionário.
	//Se a categoria for C, seu salário é de R$1.000 por hora
	//Se a categoria for B, seu salário é de $1.500 por hora mais 20% caso tenha passado de 160h mensais
	//Se a categoria for A, seu salário é de $3.000 por hora mais 50% caso tenha passado de 160h mensais
	//Calcule o salário dos funcionários conforme as informações abaixo:
	//Funcionário de categoria C: 162h mensais
	//Funcionário de categoria B: 176h mensais
	//Funcionário de categoria A: 172h mensais
	type Funcionario struct {
		categoria string
		horas     int
	}
	const (
		salarioCategoriaC = 1000.0
		salarioCategoriaB = 1500.0
		salarioCategoriaA = 3000.0
		bonusCategoriaB   = 0.2
		bonusCategoriaA   = 0.5
	)

	funcionarios := []Funcionario{{categoria: "C", horas: 162}, {categoria: "B", horas: 176}, {categoria: "A", horas: 172}}

	for _, funcionario := range funcionarios {
		salario := 0.0
		switch funcionario.categoria {
		case "C":
			salario = salarioCategoriaC * float64(funcionario.horas)
		case "B":
			salario = salarioCategoriaB * float64(funcionario.horas)
			if funcionario.horas > 160 {
				salario += salario * bonusCategoriaB
			}
		case "A":
			salario = salarioCategoriaA * float64(funcionario.horas)
			if funcionario.horas > 160 {
				salario += salario * bonusCategoriaA
			}
		default:
			return errors.New("Categoria de funcionário inválida")
		}
		fmt.Printf("O salário do funcionário de categoria %s é de R$%.2f\n", funcionario.categoria, salario)
	}
	return nil
}

func exercicio4() {
	//Os professores de uma universidade na Colômbia precisam calcular algumas estatísticas de notas dos alunos de um curso, sendo necessário calcular os valores mínimo, máximo e médio de suas notas.
	//Será necessário criar uma função que indique que tipo de cálculo deve ser realizado (mínimo, máximo ou média) e que retorna outra função (e uma mensagem caso o cálculo não esteja definido) que pode ser passado uma quantidade N de inteiros e retorne o cálculo que foi indicado na função anterior

	avgFunc, _ := getFunc(average)
	maxFunc, _ := getFunc(maximum)
	minFunc, _ := getFunc(minimum)

	min, _ := minFunc(6, 7, 8, 3, 2, 5, 6, 7, 10, 9)
	fmt.Printf("minimum: %.2f\n", min)
	avg, _ := avgFunc(6, 7, 8, 3, 2, 5, 6, 7, 10, 9)
	fmt.Printf("average: %.2f\n", avg)
	max, _ := maxFunc(6, 7, 8, 3, 2, 5, 6, 7, 10, 9)
	fmt.Printf("maximum: %.2f\n", max)
}

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func getFunc(t string) (func(values ...float64) (float64, error), error) {
	switch t {
	case minimum:
		return min, nil
	case average:
		return avg, nil
	case maximum:
		return max, nil
	}
	return nil, errors.New("invalid function type")
}

func min(values ...float64) (float64, error) {
	if len(values) == 0 {
		return 0, errors.New("empty list")
	}
	min := values[0]
	for _, v := range values {
		if v < min {
			min = v
		}
	}
	return min, nil
}

func avg(values ...float64) (float64, error) {
	if len(values) == 0 {
		return 0, errors.New("empty list")
	}
	sum := 0.0
	for _, v := range values {
		sum += v
	}
	return sum / float64(len(values)), nil
}

func max(values ...float64) (float64, error) {
	if len(values) == 0 {
		return 0, errors.New("empty list")
	}
	max := values[0]
	for _, v := range values {
		if v > max {
			max = v
		}
	}
	return max, nil
}

func exercicio5() {
	//Um abrigo de animais precisa descobrir quanta comida comprar para os animais de estimação. No momento eles só têm tarântulas, hamsters, cachorros e gatos, mas a previsão é que haja muito mais animais para abrigar.
	//Cães precisam de 10 kg de alimento
	//Gatos 5 kg
	//Hamster 250 gramas.
	//Tarântula 150 gramas.
	//Precisamos:
	//Implementar uma função Animal que receba como parâmetro um valor do tipo texto com o animal especificado e que retorne uma função com uma mensagem (caso não exista o animal)
	//Uma função para cada animal que calcule a quantidade de alimento com base na quantidade necessária do animal digitado.
	animal1, _ := Animal(dog)
	animal2, _ := Animal(cat)
	animal3, _ := Animal(hamster)
	animal4, _ := Animal(tarantula)
	_, erro := Animal("elefante")

	fmt.Printf("A quantidade de alimento para o(os) cachorro(os) em gramas é: %dg\n", animal1(2))
	fmt.Printf("A quantidade de alimento para o(os) gato(os) em gramas é: %dg\n", animal2(5))
	fmt.Printf("A quantidade de alimento para o(os) hamster(ers) em gramas é: %dg\n", animal3(9))
	fmt.Printf("A quantidade de alimento para a(as) tarantula(as) em gramas é: %dg\n", animal4(6))

	if erro != nil {
		fmt.Println(erro)
	}
}

const (
	dog           = "dog"
	cat           = "cat"
	hamster       = "hamster"
	tarantula     = "tarantula"
	dogfood       = 10000
	catfood       = 5000
	hamsterfood   = 250
	tarantulafood = 150
)

func dogFunc(quantity int) int {
	return dogfood * quantity
}

func catFunc(quantity int) int {
	return catfood * quantity
}

func hamsterFunc(quantity int) int {
	return hamsterfood * quantity
}

func tarantulaFunc(quantity int) int {
	return tarantulafood * quantity
}

func Animal(animal string) (func(quantity int) int, error) {
	switch animal {
	case dog:
		return dogFunc, nil
	case cat:
		return catFunc, nil
	case hamster:
		return hamsterFunc, nil
	case tarantula:
		return tarantulaFunc, nil
	default:
		return nil, errors.New("Animal não encontrado")
	}
}
