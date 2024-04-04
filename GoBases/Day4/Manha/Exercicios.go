package main

import (
	"errors"
	f "fmt"
)

func main() {
	for {
		exercicio := "0"
		f.Println("Lista de exercicios: 1, 2, 3, 4")
		f.Println("Digite o número do exercicio que deseja executar ou 0 para sair:")
		f.Scanln(&exercicio)
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
			f.Println("Exercicio não encontrado")
		}
	}
}

type erroSalario struct {
	msg string
}

func (e *erroSalario) Error() string {
	return "Salario inferior ao minimo: 15000"
}

func exercicio1() {
	//Em sua função “main”, defina uma variável chamada “salario”  e atribua um valor do tipo “int”.
	//Crie um erro personalizado com uma struct que implemente “Error()” com a mensagem “erro: O salário digitado não está dentro do valor mínimo" em que seja disparado quando “salário” for menor que  15.000. Caso contrário, imprima no console a mensagem “Necessário pagamento de imposto”.

	salario := 0
	f.Printf("Digite um salario: ")
	f.Scanln(&salario)
	if salario < 15000 {
		err := erroSalario{}
		f.Println(err.Error())
	} else {
		f.Println("Necessário pagamento de imposto")
	}

}
func exercicio2() {
	//Faça a mesma coisa que no exercício anterior, mas reformule o código para que, em vez de “Error()”, seja implementado  “errors.New()”.

	salario := 0
	f.Printf("Digite um salario: ")
	f.Scanln(&salario)
	if salario < 15000 {
		f.Println(errors.New("Salario inferior ao minimo: 15000"))
	} else {
		f.Println("Necessário pagamento de imposto")
	}
}
func exercicio3() {
	//Repita o processo anterior, mas agora implementando "fmt.Errorf()", para que a mensagem de erro receba como parâmetro o valor de "salario", indicando que não atinge o mínimo tributável (a mensagem exibida pelo console deve dizer : "erro: o mínimo tributável é 15.000 e o salário informado é: [salario]”, onde [salario] é o valor do tipo int passado pelo parâmetro).

	salario := 0
	f.Printf("Digite um salario: ")
	f.Scanln(&salario)
	if salario < 15000 {
		f.Println(f.Errorf("erro: o mínimo tributável é 15.000 e o salário informado é: %d", salario))
	} else {
		f.Println("Necessário pagamento de imposto")
	}
}

func calcSalario(horas int, valorHora float64) (float64, error) {
	if horas < 80 {
		e := errors.New("erro: o trabalhador não pode ter trabalhado menos de 80 horas por mês")
		return 0, f.Errorf("%w. O valor de horas digitado foi: %d", e, horas)
	}
	salario := float64(horas) * valorHora
	if salario >= 20000 {
		salario = salario * 0.9
	}
	return salario, nil

}
func exercicio4() {
	//Vamos fazer com que nosso programa seja um pouco mais complexo e útil.
	//Desenvolva as funções necessárias para permitir que a empresa calcule:
	//Salário mensal de um funcionário segundo a quantidade de horas trabalhadas.
	//A função receberá as horas trabalhadas no mês e o valor da hora como parâmetro.
	//Esta função deve retornar mais de um valor (salário calculado e erro).
	//No caso de o salário mensal ser igual ou superior a R$ 20.000, 10% devem ser descontados como imposto.
	//Se o número de horas mensais inseridas for menor que 80 ou um número negativo, a função deverá retornar um erro. Deve indicar "erro: o trabalhador não pode ter trabalhado menos de 80 horas por mês".

	//Desenvolva o código necessário para cumprir as funcionalidades solicitadas, usando “errors.New()”, “fmt.Errorf()” e “errors.Unwrap()”. Não esqueça de realizar as validações dos retornos de erro em sua função “main()”.

	salario, horas, valorHora := 0.0, 0, 0.0
	f.Printf("Digite a quatidade de horas trabalhadas: ")
	f.Scanln(&horas)
	f.Printf("Digite o valor da hora: ")
	f.Scanln(&valorHora)
	salario, err := calcSalario(horas, valorHora)
	if err != nil {
		f.Println(err)
		f.Printf("O erro especifico foi: %s\n", errors.Unwrap(err))
	} else {
		f.Println(salario)
	}
}
