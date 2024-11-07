package main

import (
	"fmt"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"time"
)

const monitoramentos = 3

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

	exibeMarcas()
}

func devolveNomeCompleto(nome string, sobrenome string) (string, string) {
	fmt.Println("")
	return nome, sobrenome
}

func exibeMenu() {
	fmt.Println("")
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir LOGs")
	fmt.Println("0 - Sair")
	fmt.Println("")
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
	/*
		Todo Array em Go, tem tamanho fixo [9]

		var urls [4]string
		urls[0] = "https://random-status-code.herokuapp.com/"
		urls[1] = "https://alura.com.br/"
		urls[2] = "https://caelum.com.br/"
		fmt.Println(urls)
		fmt.Println("array length", len(urls))
		fmt.Println("array capacity:", cap(urls))
		fmt.Println(reflect.TypeOf(urls))
	*/
	var urls []string
	urls = append(urls, "https://random-status-code.herokuapp.com/")
	urls = append(urls, "https://alura.com.br/")
	urls = append(urls, "https://caelum.com.br/")
	fmt.Println(urls)

	for j := 0; j < monitoramentos; j++ {
		/*
			Laço for padrão
			for i := 0; i < len(urls); i++ {
				testUrl(urls[i])
			}
		*/

		/*
			Laço for com range
			Onde i, é o índice do slice
		*/
		for i, url := range urls {
			fmt.Println("")
			fmt.Println("Testando a url", i)
			testUrl(url)
			fmt.Println("")
		}
		if j != (monitoramentos - 1) {
			time.Sleep(3 * time.Second)
		}
	}
}

func testUrl(url string) {
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

func exibeMarcas() {
	fmt.Println("")
	/*
		Todo Slice é um Array com capacidade infinita
	*/
	marcas := []string{"Mitsubishi", "Toyota", "Honda", "BYD"}
	marcas = append(marcas, "Fiat")
	fmt.Println(marcas)
	fmt.Println("slice length:", len(marcas))
	/*
		Devido o append realizado no slice e um novo item foi adicionado,
		sua capacidade é dobrada de acordo com a capacidade anterior.
	*/
	fmt.Println("slice capacity:", cap(marcas))
	fmt.Println(reflect.TypeOf(marcas))
}
