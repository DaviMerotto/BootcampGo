package main

import (
	"fmt"
	f "fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	for {
		exercicio := "0"
		f.Println("Lista de exercicios: 1, 2")
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
		default:
			f.Println("Exercicio não encontrado")
		}
	}
}

func exercicio1() {
	//Um escritório de contabilidade precisa acessar os dados de seus funcionários para poder realizar diferentes liquidações. Para isso, eles têm todos os detalhes necessários em um arquivo .txt.
	//É necessário desenvolver a funcionalidade para poder ler o arquivo .txt indicado pelo cliente, porém, eles não passaram o arquivo para ser lido pelo nosso programa.
	//Desenvolva o código necessário para ler os dados do arquivo chamado “customers.txt” (lembre-se do que você viu sobre o pacote “os”).
	//Como não temos o arquivo necessário, será obtido um erro e, neste caso, o programa terá que exibir um panic ao tentar ler um arquivo que não existe, exibindo a mensagem “o arquivo indicado não foi encontrado ou está danificado ”.

	file, e := os.OpenFile("customers.txt", os.O_RDONLY, 0666)
	if e != nil {
		panic("o arquivo indicado não foi encontrado ou está danificado")
	}
	defer file.Close()
}

type customer struct {
	arquivo int
	name    string
	rg      string
	phone   string
	address string
}

func genId(c *customer) error {
	c.arquivo = rand.Intn(1000)
	if c.arquivo == 0 {
		return f.Errorf("O valor %d é inavlido para o ID", c.arquivo)
	}
	return nil
}

func readFile(path string) []customer {
	defer func() {
		err := recover()
		if err != nil {
			f.Println(err)
		}
	}()
	file, e := os.OpenFile(path, os.O_RDONLY, 0666)
	if e != nil {
		panic("o arquivo indicado não foi encontrado ou está danificado")
	}
	data, e := os.ReadFile(path)
	customers := []customer{}
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		fields := strings.Split(line, ";")
		idStr := fields[0]
		nameStr := fields[1]
		rgStr := fields[2]
		phoneStr := fields[3]
		addressStr := fields[4]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Println("Erro ao converter ID para inteiro")
			return nil
		}
		c := customer{
			arquivo: id,
			name:    nameStr,
			rg:      rgStr,
			phone:   phoneStr,
			address: addressStr,
		}
		customers = append(customers, c)
	}
	defer file.Close()
	return customers
}

func exercicio2() {
	c := customer{}
	err := genId(&c)
	if err != nil {
		panic(err)
	}

	customers := readFile("customers.txt")
	exists := false
	for _, cust := range customers {
		if cust.arquivo == c.arquivo {
			exists = true
		}
	}
	if exists == true {
		f.Printf("O cliente de arquivo %d já existe!\n", c.arquivo)
	} else {
		f.Printf("O cliente de arquivo %d foi registrado!\n", c.arquivo)
	}
}
