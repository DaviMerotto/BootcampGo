package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		exercicio := "0"
		fmt.Println("Lista de exercicios: 1, 2")
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
		default:
			fmt.Println("Exercicio não encontrado")
		}
	}
}

type Alunos struct {
	Nome      string
	Sobrenome string
	RG        string
	Data      time.Time
}

func (a Alunos) detalhes() {
	fmt.Println("Nome: ", a.Nome)
	fmt.Println("Sobrenome: ", a.Sobrenome)
	fmt.Println("RG: ", a.RG)
	fmt.Println("Data: ", a.Data)
}

func exercicio1() {
	//Uma universidade precisa cadastrar os alunos e gerar uma funcionalidade para imprimir os detalhes dos dados de cada um deles, conforme o exemplo abaixo:
	//Nome: [Nome do aluno]
	//Sobrenome: [Sobrenome do aluno]
	//RG: [RG do aluno]
	//Data de admissão: [Data de admissão do aluno]
	//Os valores que estão entre parênteses devem ser substituídos pelos dados fornecidos pelos alunos.
	//Para isso é necessário gerar uma estrutura Alunos com as variáveis ​​Nome, Sobrenome, RG, Data e que tenha um método de detalhamento

	aluno1 := Alunos{Nome: "Lucas", Sobrenome: "Silva", RG: "123456789", Data: time.Date(2010, 1, 2, 15, 4, 5, 0, time.UTC)}
	aluno2 := Alunos{Nome: "João", Sobrenome: "Silva", RG: "987654321", Data: time.Now()}
	aluno1.detalhes()
	aluno2.detalhes()
}

type loja struct {
	produtos []produto
}
type produto struct {
	tipo  string
	nome  string
	preco float64
}
type iProduto interface {
	CalcularCusto()
}
type iEcommerce interface {
	Total() float64
	Adicionar(p produto)
}

const (
	p             = "Pequeno"
	m             = "Médio"
	g             = "Grande"
	aditioncostm  = 0.03
	aditioncostg  = 0.06
	shippingcostg = 2500
)

func (l *loja) Adicionar(p produto) {
	l.produtos = append(l.produtos, p)
}
func (l loja) Total() float64 {
	total := 0.0
	for _, prod := range l.produtos {
		total += prod.CalcularCusto()
	}
	return total
}
func (prod produto) CalcularCusto() float64 {
	total := 0.0
	if prod.tipo == p {
		total += prod.preco
	} else if prod.tipo == m {
		total += prod.preco + (prod.preco * aditioncostm)
	} else if prod.tipo == g {
		total += prod.preco + (prod.preco * aditioncostg) + shippingcostg
	}
	return total
}
func novoProduto(tipo, nome string, preco float64) produto {
	return produto{tipo, nome, preco}
}
func novaLoja() iEcommerce {
	return &loja{}
}
func exercicio2() {
	loja := novaLoja()

	loja.Adicionar(novoProduto(p, "copo", 12.9))
	loja.Adicionar(novoProduto(m, "prato", 23.4))
	loja.Adicionar(novoProduto(g, "panela", 50.0))

	fmt.Println(loja.Total())

	// Diversas lojas de e-commerce precisam realizar funcionalidades no Go para gerenciar produtos e devolver o valor do preço total.
	// as empresas têm 3 tipos de produtos:
	// Pequeno, Médio e Grande.
	// Existem custos adicionais para manter o produto no armazém da loja e custos de envio.
	// Lista de custos adicionais:
	// Pequeno: O custo do produto (sem custo adicional)
	// Médio: O custo do produto + 3% pela disponibilidade no estoque
	// Grande: O custo do produto + 6%  pela disponibilidade no estoque + um custo adicional pelo envio de 2500.
	// Requisitos:
	// Criar uma estrutura “loja” que guarde uma lista de produtos.
	// Criar uma estrutura “produto” que guarde o tipo de produto, nome e preço
	// Criar uma interface “Produto” que possua o método “CalcularCusto”
	// Criar uma interface “Ecommerce” que possua os métodos “Total” e “Adicionar”.
	// Será necessário uma função “novoProduto” que receba o tipo de produto, seu nome e preço, e devolva um produto.
	// Será necessário uma função “novaLoja” que retorne um Ecommerce.
	// Interface Produto:
	// Deve possuir o método “CalcularCusto”, onde o mesmo deverá calcular o custo adicional segundo o tipo de produto.
	// Interface Ecommerce:
	// Deve possuir o método “Total”, onde o mesmo deverá retornar o preço total com base no custo total dos produtos + o adicional citado anteriormente (caso a categoria tenha)
	// Deve possuir o método “Adicionar”, onde o mesmo deve receber um novo produto e adicioná-lo a lista da loja
}
