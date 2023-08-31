package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	url := "https://www.microsip.org/downloads" // Substitua pelo URL que você deseja consultar
	tagName := "a"                              // Substitua pela tag HTML que você deseja encontrar

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Erro ao fazer a solicitação HTTP:", err)
		return
	}
	defer resp.Body.Close()

	tokenizer := html.NewTokenizer(resp.Body)
	found := false

	for {
		tokenType := tokenizer.Next()

		switch tokenType {
		case html.ErrorToken:
			if tokenizer.Err() == io.EOF {
				// Chegou ao final do documento HTML
				break
			}
			fmt.Println("Erro ao analisar o token HTML:", tokenizer.Err())
			return

		case html.StartTagToken, html.SelfClosingTagToken:
			token := tokenizer.Token()

			if token.Data == tagName {
				found = true
				fmt.Println("Tag encontrada:", token.String())
				break
			}
		case html.TextToken:
			if found {
				content := strings.TrimSpace(tokenizer.Token().Data)
				if content != "" {
					fmt.Println("Conteúdo da tag:", content)
					return
				}
			}
		}

		if found {
			break
		}
	}
}

func downloadFile(url string, destPath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	outFile, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, resp.Body)
	if err != nil {
		return err
	}

	fmt.Println("Arquivo baixado com sucesso.")
	return nil
}

func salvarArquivo(link string, destination string, namePath string, extenssao string) string {

	// pegando o path C:\\%userprofile%
	//desktopPath, _ := os.UserHomeDir()

	// criando a pasta passando o path e o nome da pasta
	destinationFolder := filepath.Join(destination, namePath)
	if err := os.MkdirAll(destinationFolder, os.ModePerm); err != nil {
		fmt.Println("Erro ao criar a pasta na área de trabalho:", err)

	} // If da criação da pasta

	// salvando o arquivo
	arquivoTmp := filepath.Join(destinationFolder, namePath)
	if err := downloadFile(link, arquivoTmp+extenssao); err != nil {
		fmt.Println("Erro ao baixar o arquivo:", err)

	}
	fmt.Printf("Arquivo %s salvo com sucesso\n", namePath+extenssao)
	return arquivoTmp + extenssao

}
