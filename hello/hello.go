package main

import (
	"encoding/json"
	"fmt"
	"hello/database"
	"hello/models"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	desktopPath, _ := os.UserHomeDir()
	nomeInstalador := "MicroSIP"

	link := "https://raw.githubusercontent.com/Agnerft/microsip/main/TESTE/MicroSIP1/MicroSIP.txt"
	resultadoIni := salvarArquivo(link, desktopPath+"\\AppData\\Roaming\\", nomeInstalador, ".ini")

	jsonfile, _ := database.BuscaPorDoc(12310400000182)

	var clienteConfig []models.ClienteConfig
	if err := json.Unmarshal([]byte(jsonfile), &clienteConfig); err != nil {
		fmt.Println("Erro ao fazer o Unmarshal do JSON:", err)
		return
	}

	fmt.Println(clienteConfig)

	for _, config := range clienteConfig {
		linkCompleto := config.GrupoRecurso + config.LinkGvc + config.Porta
		editor(resultadoIni, 2, "label="+config.Ramal)
		editor(resultadoIni, 3, "server="+linkCompleto)
		editor(resultadoIni, 4, "proxy="+linkCompleto)
		editor(resultadoIni, 5, "domain="+linkCompleto)
		editor(resultadoIni, 6, "username="+config.Ramal)
		editor(resultadoIni, 7, "password="+config.Ramal+config.Senha)
		editor(resultadoIni, 8, "authID="+config.Ramal)

	}

}

func editor(resultado string, numeroLinha int, novoValor string) {

	//resultado := "C:\\Users\\USER\\Microsip\\Microsip.txt"

	file, err := os.OpenFile(resultado, os.O_RDWR, 0644)
	if err != nil {
		log.Fatalf("Erro ao abrir o arquivo: %v", err)
	}
	defer file.Close()

	// Lê o conteúdo do arquivo
	conteudo, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("Erro ao ler o conteúdo do arquivo: %v", err)
	}

	//numeroLinha := 2
	//novoValor := "label=7848"

	// Converte o conteúdo para uma string
	conteudoArquivo := string(conteudo)

	linhas := strings.Split(conteudoArquivo, "\n")

	if numeroLinha > 0 && numeroLinha < len(linhas) {
		linhas[numeroLinha-1] = novoValor
	}

	novoConteudoArquivo := strings.Join(linhas, "\n")

	_, err = file.Seek(0, 0)
	if err != nil {
		log.Fatal("Erro ao mover o ponteiro: %v", err)

	}

	_, err = file.WriteString(novoConteudoArquivo)
	if err != nil {
		log.Fatal("Erro ao salvar novo conteudo: %v", err)
	}

	err = file.Truncate(int64(len(novoConteudoArquivo)))
	if err != nil {
		log.Fatal("Erro ao truncar: %v", err)
	}

	//fmt.Println(novoConteudoArquivo)
	fmt.Println("Alterações salvas com sucesso.")

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
