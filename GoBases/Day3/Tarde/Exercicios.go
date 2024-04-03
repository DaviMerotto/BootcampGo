package main

import (
	f "fmt"
	"math/rand"
	"sync"
	"time"
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

type Usuarios struct {
	Nome      string
	Sobrenome string
	Idade     uint
	Email     string
	Senha     string
}

func (u *Usuarios) mudarNome(nome, sobrenome string) {
	u.Nome = nome
	u.Sobrenome = sobrenome
}

func (u *Usuarios) mudarIdade(idade uint) {
	u.Idade = idade
}

func (u *Usuarios) mudarEmail(email string) {
	u.Email = email
}

func (u *Usuarios) mudarSenha(senha string) {
	u.Senha = senha
}

func exercicio1() {
	//Uma empresa de mídia social precisa implementar uma estrutura de usuários com funções que acrescentem informações à estrutura. Para otimizar e economizar memória, eles exigem que a estrutura de usuários ocupe o mesmo lugar na memória para o programa principal e para as funções:
	//A estrutura deve possuir os seguintes campos: Nome, Sobrenome, idade, email e senha
	//E devem implementar as seguintes funcionalidades:
	//mudar o nome: me permite mudar o nome e sobrenome
	//mudar a idade: me permite mudar a idade
	//mudar e-mail: me permite mudar o e-mail
	//mudar a senha: me permite mudar a senha
	u := Usuarios{}
	u.mudarNome("Lucas", "Silva")
	u.mudarIdade(25)
	u.mudarEmail("teste@teste.com")
	u.mudarSenha("123456")
	f.Println(u)
	u.mudarNome("João", "Silva")
	f.Println(u)
}

type produto struct {
	nome       string
	preco      float64
	quantidade int
}

type usuario struct {
	nome      string
	sobrenome string
	email     string
	produtos  []produto
}

func novoProduto(nome string, preco float64) produto {
	return produto{nome: nome, preco: preco}
}

func adicionarProduto(u *usuario, p *produto, quantidade int) {
	p.quantidade = quantidade
	u.produtos = append(u.produtos, *p)
}

func deletarProdutos(u *usuario) {
	u.produtos = nil
}

func exercicio2() {
	//Uma grande empresa de vendas na web precisa adicionar funcionalidades para adicionar produtos aos usuários. Para fazer isso, eles exigem que usuários e produtos tenham o mesmo endereço de memória no main do programa e nas funções.
	//Estruturas necessárias:
	//Usuário: Nome, Sobrenome, E-mail, Produtos (array de produtos).
	//Produto: Nome, preço, quantidade.
	//Algumas funções necessárias:
	//Novo produto: recebe nome e preço, e retorna um produto.
	//Adicionar produto: recebe usuário, produto e quantidade, não retorna nada, adiciona o produto ao usuário.
	//Deletar produtos: recebe um usuário, apaga os produtos do usuário.

	u := usuario{nome: "Lucas", sobrenome: "Silva", email: "teste@teste.com"}

	p1 := novoProduto("Jaqueta", 150.0)
	p2 := novoProduto("Camiseta", 50.0)

	adicionarProduto(&u, &p1, 2)
	adicionarProduto(&u, &p2, 1)

	f.Println(u)
	deletarProdutos(&u)
	f.Println(u)
}

type produtoEx3 struct {
	nome       string
	preco      float64
	quantidade int
}

type servico struct {
	nome            string
	preco           float64
	minutosTrabalho int
}

type manutencao struct {
	nome  string
	preco float64
}

func somarProdutos(produto []produtoEx3, c chan float64) {
	total := 0.0
	for _, p := range produto {
		total += p.preco * float64(p.quantidade)
	}
	c <- total
}

func somarServicos(s []servico, c chan float64) {
	total := 0.0
	for _, s := range s {
		total += s.preco * float64(s.minutosTrabalho/60)
	}
	c <- total
}

func somarManutencao(m []manutencao, c chan float64) {
	total := 0.0
	for _, m := range m {
		total += m.preco
	}
	c <- total
}

func exercicio3() {
	//Uma empresa nacional é responsável pela venda de produtos, serviços e manutenção.
	//Para isso, eles precisam realizar um programa que seja responsável por calcular o preço total dos Produtos, Serviços e Manutenção. Devido à forte demanda e para otimizar a velocidade, eles exigem que o cálculo da soma seja realizado em paralelo por meio de 3 go routines.
	//Precisamos de 3 estruturas:
	//Produtos: nome, preço, quantidade.
	//Serviços: nome, preço, minutos trabalhados.
	//Manutenção: nome, preço.
	//Precisamos de 3 funções:
	//Somar Produtos: recebe um array de produto e devolve o preço total (preço * quantidade).
	//Somar Serviços: recebe uma array de serviço e devolve o preço total (preço * média hora trabalhada, se ele não vier trabalhar por 30 minutos, ele será cobrado como se tivesse trabalhado meia hora).
	//Somar Manutenção: recebe um array de manutenção e devolve o preço total.
	//Os 3 devem ser executados concomitantemente e ao final o valor final deve ser mostrado na tela (somando o total dos 3).

	produtos := []produtoEx3{
		{nome: "Jaqueta", preco: 150.0, quantidade: 2},
		{nome: "Camiseta", preco: 50.0, quantidade: 1},
	}

	servicos := []servico{
		{nome: "Manutenção", preco: 100.0, minutosTrabalho: 120},
		{nome: "Instalação", preco: 50.0, minutosTrabalho: 30},
	}

	manutencao := []manutencao{
		{nome: "Troca de peça", preco: 200.0},
		{nome: "Limpeza", preco: 50.0},
	}

	ch := make(chan float64)

	go somarProdutos(produtos, ch)
	go somarServicos(servicos, ch)
	go somarManutencao(manutencao, ch)

	total := 0.0
	total += <-ch
	total += <-ch
	total += <-ch

	close(ch)

	f.Printf("O preço total via um canal é: %.2f\n", total)
}

func insertionSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := i; j > 0 && arr[j-1] > arr[j]; j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}
	return arr
}

func bubbleSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func selectionSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
	return arr
}

func exercicio4() {
	var wg sync.WaitGroup

	var1 := rand.Perm(100)
	var2 := rand.Perm(1000)
	var3 := rand.Perm(10000)
	var4 := rand.Perm(100000)

	f.Printf("------- 100 elementos -------\n")
	wg.Add(3)
	go func() {
		start := time.Now()
		insertionSort(var1)
		f.Printf("Tempo de execução do insertionSort para 100 elementos: %v\n", time.Since(start))
		wg.Done()
	}()
	go func() {
		start := time.Now()
		bubbleSort(var1)
		f.Printf("Tempo de execução do bubbleSort para 100 elementos: %v\n", time.Since(start))
		wg.Done()
	}()
	go func() {
		start := time.Now()
		selectionSort(var1)
		f.Printf("Tempo de execução do selectionSort para 100 elementos: %v\n", time.Since(start))
		wg.Done()
	}()
	wg.Wait()
	f.Printf("------- 1000 elementos -------\n")

	wg.Add(3)
	go func() {
		defer wg.Done()
		start := time.Now()
		insertionSort(var2)
		f.Printf("Tempo de execução do insertionSort para 1000 elementos: %v\n", time.Since(start))
	}()
	go func() {
		defer wg.Done()
		start := time.Now()
		bubbleSort(var2)
		f.Printf("Tempo de execução do bubbleSort para 1000 elementos: %v\n", time.Since(start))
	}()
	go func() {
		defer wg.Done()
		start := time.Now()
		selectionSort(var2)
		f.Printf("Tempo de execução do selectionSort para 1000 elementos: %v\n", time.Since(start))
	}()
	wg.Wait()
	f.Printf("------- 10000 elementos -------\n")

	wg.Add(3)
	go func() {
		defer wg.Done()
		start := time.Now()
		insertionSort(var3)
		f.Printf("Tempo de execução do insertionSort para 10000 elementos: %v\n", time.Since(start))
	}()
	go func() {
		defer wg.Done()
		start := time.Now()
		bubbleSort(var3)
		f.Printf("Tempo de execução do bubbleSort para 10000 elementos: %v\n", time.Since(start))
	}()
	go func() {
		defer wg.Done()
		start := time.Now()
		selectionSort(var3)
		f.Printf("Tempo de execução do selectionSort para 10000 elementos: %v\n", time.Since(start))
	}()
	wg.Wait()
	f.Printf("------- 100000 elementos -------\n")

	wg.Add(3)
	go func() {
		defer wg.Done()
		start := time.Now()
		insertionSort(var4)
		f.Printf("Tempo de execução do insertionSort para 100000 elementos: %v\n", time.Since(start))
	}()
	go func() {
		defer wg.Done()
		start := time.Now()
		bubbleSort(var4)
		f.Printf("Tempo de execução do bubbleSort para 100000 elementos: %v\n", time.Since(start))
	}()
	go func() {
		defer wg.Done()
		start := time.Now()
		selectionSort(var4)
		f.Printf("Tempo de execução do selectionSort para 100000 elementos: %v\n", time.Since(start))
	}()
	wg.Wait()
}
