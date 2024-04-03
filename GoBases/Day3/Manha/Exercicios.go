package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
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

type produto struct {
	id         int
	preco      float64
	quantidade int
}

func (p produto) gerarLinha() string {
	return fmt.Sprintf("%d;%f;%d\n", p.id, p.preco, p.quantidade)
}

func geraCsv(path string, products []produto) error {

	if len(products) == 0 {
		return errors.New("Nenhum produto para salvar")
	}

	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0600)

	if err != nil {
		return fmt.Errorf("Erro ao criar/abrir o arquivo %s: %w", path, err)
	}

	defer file.Close()
	_, err = file.WriteString("ID;Preço;Quantidade\n")

	if err != nil {
		return fmt.Errorf("Erro ao escrever o cabeçalho no arquivo %s: %w", file.Name(), err)
	}

	for _, p := range products {
		_, err := file.WriteString(p.gerarLinha())
		if err != nil {
			return fmt.Errorf("Erro ao escrever a linha no arquivo %s: %w", file.Name(), err)
		}
	}
	fmt.Printf("Arquivo %s criado com sucesso\n", file.Name())
	return nil
}

func exercicio1() {
	//Uma empresa que vende produtos de limpeza precisa de:
	//Implementar uma funcionalidade  para guardar um arquivo de texto, com a informação de produtos comprados, separados por ponto e vírgula (csv).
	//Deve possuir o ID do produto, preço e a quantidade.
	//Estes valores podem ser hardcodeados ou escritos manualmente em uma variável.

	produtos := []produto{
		{1, 10.8, 5},
		{2, 20.5, 10},
		{3, 30.6, 15},
	}

	fmt.Printf("%e", geraCsv("produtos.csv", produtos))
}

func exercicio2() {
	//A mesma empresa precisa ler o arquivo armazenado, para isso exige que:
	//Seja impresso na tela os valores tabelados, com título ( à esquerda para o ID e à direita para o Preço e Quantidade), o preço, a quantidade e abaixo do preço o total deve ser exibido (somando preço por quantidade)

	produtos := []produto{}

	file, err := os.Open("produtos.csv")
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo")
		return
	}
	defer file.Close()

	data, err := os.ReadFile("produtos.csv")
	if err != nil {
		fmt.Println("Erro ao ler o arquivo")
		return
	}

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		if line == "ID;Preço;Quantidade" {
			continue
		}
		if line == "" {
			continue
		}
		fields := strings.Split(line, ";")
		idStr := fields[0]
		precoStr := fields[1]
		quantidadeStr := fields[2]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Println("Erro ao converter ID para inteiro")
			return
		}
		preco, err := strconv.ParseFloat(precoStr, 64)
		if err != nil {
			fmt.Println("Erro ao converter preço para float64")
			return
		}
		quantidade, err := strconv.Atoi(quantidadeStr)
		if err != nil {
			fmt.Println("Erro ao converter quantidade para inteiro")
			return
		}
		p := produto{
			id:         id,
			preco:      preco,
			quantidade: quantidade,
		}
		produtos = append(produtos, p)
	}

	total := 0.0
	for _, p := range produtos {
		total += float64(p.quantidade) * p.preco
	}

	fmt.Println("ID\tPreço\tQuantidade")
	for _, p := range produtos {
		fmt.Printf("%d\t%.2f\t%d\n", p.id, p.preco, p.quantidade)
	}
	fmt.Printf("\t%.2f\t\n", total)
}
