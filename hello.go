package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const monitoramentos = 3

const delay = 5

func main() {

	exibeIntroducao()
	lerSitesArquivo()

	for {

		exibeMenu()
		comando := lerComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo Logs...")
		case 0:
			fmt.Println("Saindo do programa")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1)
		}
	}

}

func exibeIntroducao() {
	nome := "Guilherme"
	versao := 1.1
	fmt.Println("Olá, sr", nome)
	fmt.Println("Este programa está na versão", versao)
	fmt.Println("")
}

func lerComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)
	fmt.Println("")

	return comandoLido
}

func exibeMenu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	//sites := [] string {"https://httpbin.org/status/200", "https://alura.com.br", "https://httpbin.org/status/404"}

	sites := lerSitesArquivo()

	for i := 0; i < monitoramentos; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			testaSites(site)
		}
		time.Sleep(delay * time.Second)
	}

	fmt.Println("")

}

func testaSites(site string) {

	resp, err := http.Get(site)

	if err != nil {
		fmt.Printf("Ocorreu um erro", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
		fmt.Println("")
	} else {
		fmt.Println("Site:", site, "esta com problemas. Status Code:", resp.StatusCode)
		fmt.Println("")
	}
}

func lerSitesArquivo() []string {
	var sites []string

	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro", err)
	}

	leitor := bufio.NewReader(arquivo)

	for {

		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		
		sites = append(sites, linha)

		if err == io.EOF {
			break
		}
	}

	arquivo.Close()

	return sites
}
