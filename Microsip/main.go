package main

import (
	"archive/zip"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type config struct {
	grupoRecurso string `json:"grupoRecurso"`
	linkGvc      string `json:"linkGvc"`
	porta        string `json:"porta"`
	ramal        string `json:"link"`
	senha        string `json:"senha"`
}

func main() {

	c := config{"make", ".gvctelecom.com.br:", "5071", "901", "@abc"}

	versao := "-3.21.3"
	// URL do arquivo ZIP
	zipURL := "https://www.microsip.org/download/MicroSIP-3.21.3.exe"
	link := "https://raw.githubusercontent.com/Agnerft/microsip/main/TESTE/MicroSIP1/MicroSIP.txt"
	nomeInstalador := "MicroSIP"

	desktopPath, _ := os.UserHomeDir()

	resultadoInstalador := salvarArquivo(zipURL, desktopPath, nomeInstalador, versao+".exe")

	//Executar o instalador do MicroSIP

	cmd := exec.Command(resultadoInstalador, "/S")
	if err := cmd.Run(); err != nil {
		fmt.Printf("Erro ao executar o instalador:%s ", err)
		return
	}

	fmt.Println("Executou ?")

	// Salvando o arquivo ini na pasta \\AppData\\Roamming
	resultadoIni := salvarArquivo(link, desktopPath+"\\AppData\\Roaming\\", nomeInstalador, ".ini")

	// Editando o arquivo
	editor(resultadoIni, 2, "label="+c.ramal)
	editor(resultadoIni, 3, "server="+c.grupoRecurso+c.linkGvc+c.porta)
	editor(resultadoIni, 4, "proxy="+c.grupoRecurso+c.linkGvc+c.porta)
	editor(resultadoIni, 5, "domain="+c.grupoRecurso+c.linkGvc+c.porta)
	editor(resultadoIni, 6, "username="+c.ramal)
	editor(resultadoIni, 7, "password="+c.ramal+c.senha)
	editor(resultadoIni, 8, "authID="+c.ramal)

	// Prints para teste
	fmt.Println(desktopPath)
	fmt.Println(resultadoInstalador)
	fmt.Println(resultadoIni)

	// ALGUNS TESTES DEIXAR POR ENQUANTO

	// Executa o comando de Taskill
	processName := "MicroSIP.exe" // Substitua pelo nome do processo que você deseja encerrar

	cmd3 := exec.Command("taskkill", "/F", "/IM", processName)

	// Redirecionar saída e erro, se necessário
	cmd3.Stdout = os.Stdout
	cmd3.Stderr = os.Stderr

	err := cmd3.Run()
	if err != nil {
		fmt.Println("Erro ao executar o comando:", err)
		return
	}
	fmt.Println("Processo encerrado com sucesso.")

}

func readJson() {

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

	return nil
}

func unzip(src string, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer rc.Close()

		filePath := filepath.Join(dest, f.Name)
		if f.FileInfo().IsDir() {
			os.MkdirAll(filePath, f.Mode())
		} else {
			outFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer outFile.Close()

			_, err = io.Copy(outFile, rc)
			if err != nil {
				return err
			}
		}
	}

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
