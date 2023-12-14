package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

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

	return nil
}

func SalvarArquivo(link string, destination string, namePath string, extenssao string) string {

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
