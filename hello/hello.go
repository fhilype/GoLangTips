package main

import (
	"fmt"
	"net/http"
	"os"
	"reflect"
	"strconv"
)

func main() {
	exibeIntroducao()
	/*
		if comando == 1 {
			fmt.Println("Monitorando...")
		} else if comando == 2 {
			fmt.Println("Exibindo LOGs...")
		} else if comando == 0 {
			fmt.Println("Saindo do programa...")
		} else {
			fmt.Println("Comando inesperado.")
		}
	*/
	for {
		exibeMenu()
		switch lerComando() {
		case 1:
			iniciarMonitoramento()
			break
		case 2:
			fmt.Println("Exibindo LOGs...")
			break
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0) // Informa o Sistema Operacional que o programa foi encerrado com sucesso
			break
		default:
			fmt.Println("Comando inesperado.")
			os.Exit(-1)
			break
		}
	}
}

func exibeIntroducao() {
	var nome string = "Fhilype"
	sobrenome := "Medeiros"
	var idade int = 31
	var versao float32 = 1.1
	fmt.Println("Olá", nome, sobrenome, "seja bem vindo")

	impressao := "Sua idade é "
	impressao += strconv.Itoa(idade)
	impressao += " anos"
	fmt.Println(impressao)

	fmt.Println("Este programa esta na versão: ", versao)

	fmt.Println("Tipo da variável idade:", reflect.TypeOf(idade))
	fmt.Println("Tipo da variável versao:", reflect.TypeOf(versao))

	nomeDevolvido, sobrenomeDevolvido := devolveNomeCompleto(nome, sobrenome)
	fmt.Println("nomeDevolvido", nomeDevolvido)
	fmt.Println("sobrenomeDevolvido", sobrenomeDevolvido)
}

func devolveNomeCompleto(nome string, sobrenome string) (string, string) {
	return nome, sobrenome
}

func exibeMenu() {
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir LOGs")
	fmt.Println("0 - Sair")
}

func lerComando() int {
	var comando int
	/*
		fmt.Scanf
		%d é o modificador onde a leitura espera um número int como entrada.
		& indica o endereço. Logo &comando, o endereço da variável comando.
		fmt.Scanf("%d", &comando)
		fmt.Println("O dado de entrada foi", comando, "e seu endereço é", &comando)
	*/
	/*
		fmt.Scan (preferível)
		não recebe um modificador como parametro.
		& indica o endereço. Logo &comando, o endereço da variável comando.
	*/
	fmt.Scan(&comando)
	fmt.Println("O dado de entrada foi", comando, "e seu endereço é", &comando)
	return comando
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	var url string = "https://random-status-code.herokuapp.com/"
	/*
		o '_', para funções que retornam múltiplos valores, ignora o valor na posição onde o '_' é colocado
	*/
	response, _ := http.Get(url)
	if response.StatusCode == 200 {
		fmt.Println("Site:", url, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", url, "está com problemas. Status Code:", response.StatusCode)
	}
}
