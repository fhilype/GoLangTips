package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
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

	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir LOGs")
	fmt.Println("0 - Sair")

	var comando int
	/*
		fmt.Scanf
		%d é o modificador onde a leitura espera um número int como entrada.
		& indica o endereço. Logo &comando, o endereço da variável comando.
	*/
	fmt.Scanf("%d", &comando)
	fmt.Println("O dado de entrada foi", comando, "e seu endereço é", &comando)
	/*
		fmt.Scan (preferível)
		não recebe um modificador como parametro.
		& indica o endereço. Logo &comando, o endereço da variável comando.
	*/
	fmt.Scan(&comando)
	fmt.Println("O dado de entrada foi", comando, "e seu endereço é", &comando)
}
