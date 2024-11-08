package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

const monitoramentos = 1

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
			imprimeLogs()
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
	/*
		var urls []string
		urls = append(urls, "https://random-status-code.herokuapp.com/")
		urls = append(urls, "https://alura.com.br/")
		urls = append(urls, "https://caelum.com.br/")
		fmt.Println(urls)
	*/

	urls := lerURLsDoArquivo()

	for j := 1; j <= monitoramentos; j++ {
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
		if j != monitoramentos {
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
		registraLog(url, true)
	} else {
		fmt.Println("Site:", url, "está com problemas. Status Code:", response.StatusCode)
		registraLog(url, false)
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

func lerURLsDoArquivo() []string {
	var urls []string
	/*
		Tratando caso de erro

		le e armazena o arquivo inteiro em bytes na variável
		arquivo, err := ioutil.ReadFile("urls.txt")
	*/
	arquivo, err := os.Open("urls.txt")
	if err == nil {
		reader := bufio.NewReader(arquivo)
		for {
			row, err := reader.ReadString('\n') // aspas simples pois \n no caso não se trata de uma string, mas um delimitador
			row = strings.TrimSpace(row)
			fmt.Println(row)
			urls = append(urls, row)
			if err == io.EOF {
				break
			}
		}
	} else {
		fmt.Println("arquivo:", err)
	}
	arquivo.Close() // fechar o arquivo é uma boa prática pois o mesmo pode precisar ser lido novamente
	return urls
}

func registraLog(url string, status bool) {
	/*
		O_RDWR para ler e escrever no arquivo
		O_CREATE para criar o arquivo se ele não existir
		O_APPEND para não sobrescrever linhas
		0666 permissão padrão do sistema
	*/
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err == nil {
		/*
			https://go.dev/src/time/format.go
			Documentação para formatação de data e hora
		*/
		arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - URL: " + url + " - online: " + strconv.FormatBool(status) + "\n")
	} else {
		fmt.Println("registraLog:", err)
	}
	arquivo.Close()
}

func imprimeLogs() {
	fmt.Println("Exibindo LOGs...")
	arquivo, err := ioutil.ReadFile("log.txt")
	if err == nil {
		fmt.Println(string(arquivo))
	} else {
		fmt.Println("imprimeLogs:", err)
	}
}
