package controle

import (
	"Microsip/database"
	"Microsip/models"
	"Microsip/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"

	"github.com/gorilla/mux"

	"net/http"
)

var (
	versao = "-3.21.3"
	// URL do arquivo ZIP
	zipURL         = "https://www.microsip.org/download/MicroSIP-3.21.3.exe"
	link           = "https://raw.githubusercontent.com/Agnerft/microsip/main/TESTE/MicroSIP1/MicroSIP.txt"
	nomeInstalador = "MicroSIP"
	clientConfig   []models.ClienteConfig
	desktopPath, _ = os.UserHomeDir()
)

func Home(w http.ResponseWriter, r *http.Request) {

	htmlResponse := "template/hello.html"

	w.Header().Set("Content-Disposition", "attachment; filename=hello.html")
	w.Header().Set("Content-Type", "application/octet-stream")

	//fmt.Println(htmlResponse)

	http.ServeFile(w, r, htmlResponse)
}

func SaveExeEIni(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)

		http.ServeFile(w, r, "template/4xx.html")
		return
	}

	resultadoInstalador := utils.SalvarArquivo(zipURL, desktopPath, nomeInstalador, versao+".exe")
	resultadoIni := utils.SalvarArquivo(link, desktopPath+"\\AppData\\Roaming\\", nomeInstalador, ".ini")

	fmt.Fprintln(w, "Arquivo %s salvo com sucesso", resultadoIni)

	//Executar o instalador do MicroSIP

	utils.Comandos(resultadoInstalador)
}

func SaveIni(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
	}

	//vars := mux.Vars(r)

	resultadoIni := utils.SalvarArquivo(link, desktopPath+"\\AppData\\Roaming\\", nomeInstalador, ".ini")

	for _, config := range clientConfig {
		linkCompleto := config.GrupoRecurso + config.LinkGvc + config.Porta
		utils.Editor(resultadoIni, 2, "label="+config.Ramal)
		utils.Editor(resultadoIni, 3, "server="+linkCompleto)
		utils.Editor(resultadoIni, 4, "proxy="+linkCompleto)
		utils.Editor(resultadoIni, 5, "domain="+linkCompleto)
		utils.Editor(resultadoIni, 6, "username="+config.Ramal)
		utils.Editor(resultadoIni, 7, "password="+config.Ramal+config.Senha)
		utils.Editor(resultadoIni, 8, "authID="+config.Ramal)
	}

	fmt.Println("Cliente salvo e ajustado")

}

func CriandoConta(w http.ResponseWriter, r *http.Request) {
	// Verificar se a solicitação é do tipo POST.
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	// Decodificar os dados JSON do corpo da solicitação.
	//var configTest models.ClienteTeste
	configTest := models.NewClienteTeste()
	err := json.NewDecoder(r.Body).Decode(&configTest)
	if err != nil {
		http.Error(w, "Erro ao decodificar os dados JSON", http.StatusBadRequest)
		return
	}

	fmt.Printf("Novo usuário criado: %+v\n", configTest)

	// Responder com uma mensagem de sucesso.
	// w.WriteHeader(http.StatusCreated)
	// fmt.Fprintf(w, "Usuário criado com sucesso!\n")

	// clienteCriado, err := models.CriarNovoCliente(&configTest, "https://root:agner102030@basesip.makesystem.com.br/clientes")
	// if err != nil {
	// 	fmt.Println("Erro")
	// }

	// fmt.Println(len(clienteCriado))
	// fmt.Println("Deu certo?")

	// _, err = r.Body.Read(clienteCriado)
	// if err != nil {
	// 	http.Error(w, "Não deu", http.StatusBadRequest)
	// }

	//http.ServeFile(w, r, "hello.html")

	//client := &http.Client{}

	jsonBytes, err := configTest.ToJSON()
	if err != nil {
		log.Printf("Erro a serializar: %v", err)
	}

	req, err := http.Post("https://root:agner102030@basesip.makesystem.com.br/clientes", "application/json", bytes.NewBuffer(jsonBytes))
	if err != nil {
		fmt.Println(err)
	}

	//fmt.Println(req.Body)

	//fmt.Println(jsonBytes)

	defer req.Body.Close()

}

func FindUser(w http.ResponseWriter, r *http.Request) {
	// Verificar se a solicitação é do tipo GET.

	vars := mux.Vars(r)
	fmt.Println(vars)

	match, _ := regexp.MatchString("^[0-9]+$", vars["doc"])
	fmt.Println(match)

	if r.Method != http.MethodGet {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return

	}

	jsonfile, _ := database.BuscaPorDoc(vars["doc"], clientConfig)
	fmt.Println(string(jsonfile))

	if jsonfile == nil {
		fmt.Println("Desculpe, não encontrei você . . . ")
		return
	}

	if err := json.Unmarshal(jsonfile, &clientConfig); err != nil {
		fmt.Println("Erro ao fazer o Unmarshal do JSON:", err)
		return
	}

	resultadoIni := utils.SalvarArquivo(link, desktopPath+"\\AppData\\Roaming\\", nomeInstalador, ".ini")

	for _, config := range clientConfig {
		fmt.Println("oi")
		linkCompleto := config.GrupoRecurso + config.LinkGvc + config.Porta
		utils.Editor(resultadoIni, 2, "label="+config.Ramal)
		utils.Editor(resultadoIni, 3, "server="+linkCompleto)
		utils.Editor(resultadoIni, 4, "proxy="+linkCompleto)
		utils.Editor(resultadoIni, 5, "domain="+linkCompleto)
		utils.Editor(resultadoIni, 6, "username="+config.Ramal)
		utils.Editor(resultadoIni, 7, "password="+config.Ramal+config.Senha)
		utils.Editor(resultadoIni, 8, "authID="+config.Ramal)

		fmt.Println("O Cliente " + config.Cliente + " foi salvo no arquivo ini")
	}

}

func Remover(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Oi")
	if r.Method != http.MethodGet {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	resultado := desktopPath + "\\AppData\\Local\\MicroSIP\\Uninstall.exe"

	cmd := exec.Command(resultado, "/S")
	if err := cmd.Run(); err != nil {
		fmt.Printf("Erro ao executar o desinstalador: %s ", err)
	}

	fmt.Println("Removido")

}
