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

	appData = os.Getenv("APPDATA")
)

func Home(w http.ResponseWriter, r *http.Request) {

	//htmlResponse := "template/hello.html"

	w.Write([]byte(appData))

	//w.Header().Set("Content-Disposition", "attachment; filename=hello.html")
	//w.Header().Set("Content-Type", "application/octet-stream")

	//fmt.Println(htmlResponse)

	http.ServeFile(w, r, desktopPath)
}

func SaveExeEIni(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)

		http.ServeFile(w, r, "template/4xx.html")
		return
	}

	resultadoInstalador := utils.SalvarArquivo(zipURL, desktopPath, nomeInstalador, versao+".exe")
	resultadoIni := utils.SalvarArquivo(link, appData+"\\", nomeInstalador, ".ini")

	fmt.Fprintln(w, "Arquivo %s salvo com sucesso", resultadoIni)

	//Executar o instalador do MicroSIP

	utils.Comandos(resultadoInstalador)
}

func SaveIni(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
	}

	//vars := mux.Vars(r)

	resultadoIni := utils.SalvarArquivo(link, appData+"\\", nomeInstalador, ".ini")

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

	configTest := models.NewClienteTeste()
	err := json.NewDecoder(r.Body).Decode(&configTest)
	if err != nil {
		http.Error(w, "Erro ao decodificar os dados JSON", http.StatusBadRequest)
		return
	}

	fmt.Printf("Novo usuário criado: %+v\n", configTest)

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
		//fmt.Println("oi")
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
	microsipINI := "MicroSIP.ini"
	contacts := "Contacts.xml"
	resultado := desktopPath + "\\AppData\\Local\\MicroSIP\\Uninstall.exe"
	apagarArquivos := desktopPath + "\\AppData\\Roaming\\MicroSIP\\"
	cmd1 := exec.Command(resultado, "/S")
	if err := cmd1.Run(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		erro := []byte("Desinstalador não existe.")
		w.Write(erro)

		//fmt.Printf("Erro ao executar o desinstalador: %s ", err)
	}

	cmd2 := exec.Command("cmd", "/C", "del", apagarArquivos+microsipINI)
	if err := cmd2.Run(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		erro := []byte("Arquivo não Existe")
		w.Write(erro)

		fmt.Printf("Erro ao executar o desinstalador: %s ", err)
		//w.Write()
	}

	cmd3 := exec.Command("cmd", "/C", "del", apagarArquivos+contacts)
	if err := cmd3.Run(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		erro := []byte("Arquivo não Existe")
		w.Write(erro)
		fmt.Printf("Erro ao executar o desinstalador: %s ", err)
		//fmt.Println(apagarArquivos)
	}

	w.WriteHeader(http.StatusOK)
	fmt.Println("Removidos")

}

// func CreateRamal(w http.ResponseWriter, r *http.Request) {

// 	vars := mux.Vars(r)
// 	fmt.Println(vars)

// 	match, _ := regexp.MatchString("^[0-9]+$", vars["doc"])
// 	fmt.Println(match)

// 	jsonfile, _ := database.BuscaPorDoc(vars["doc"], clientConfig)

// 	if err := json.Unmarshal(jsonfile, &clientConfig); err != nil {
// 		fmt.Println("Erro ao fazer o Unmarshal do JSON:", err)
// 		return
// 	}

// 	for _, config := range clientConfig {

// 		for i := range config.QuantRamais {
// 			fmt.Println(config.QuantRamais[i].Ramal)
// 		}

// 		fmt.Println("Porém os que não foram configurados ainda são:")

// 		for i := range config.QuantRamais {
// 			if config.QuantRamais[i].InUse == false {
// 				fmt.Println(config.QuantRamais[i].Ramal)
// 			}

// 		}
// 		var ramal string

// 		//database.EditClient(config.ID, config)

// 		config.Ramal = ramal

// 		ramalInt, _ := strconv.Atoi(ramal)

// 		found := false

// 		for i := range config.QuantRamais {

// 			if config.QuantRamais[i].InUse != found {
// 				fmt.Println("Ramal sendo usado.")
// 				break
// 			}

// 			if config.QuantRamais[i].Ramal == ramalInt {
// 				config.QuantRamais[i].InUse = true
// 				found = true
// 				break
// 			}
// 		}

// 		if !found {
// 			fmt.Println("Ramal não encontrado.")
// 			return
// 		}

// 		fmt.Println(config)
// 		fmt.Println("aqui?")

// 		// Codifique a estrutura de dados atualizada para JSON
// 		updatedJSON, err := json.Marshal(config.QuantRamais)
// 		if err != nil {
// 			fmt.Println("Erro ao serializar o JSON:", err)
// 			return
// 		}

// 		fmt.Printf(string(updatedJSON))

// 		database.EditClient(config.ID, config)
// 		//
// 		//
// 		//
// 		fmt.Println("Oi")
// 		//linkCompleto := config.GrupoRecurso + config.LinkGvc + config.Porta
// 		//utils.Editor(resultadoIni, 2, "label="+config.Ramal)
// 		//utils.Editor(resultadoIni, 3, "server="+linkCompleto)
// 		//utils.Editor(resultadoIni, 4, "proxy="+linkCompleto)
// 		//utils.Editor(resultadoIni, 5, "domain="+linkCompleto)
// 		//utils.Editor(resultadoIni, 6, "username="+config.Ramal)
// 		//utils.Editor(resultadoIni, 7, "password="+config.Ramal+config.Senha)
// 		//utils.Editor(resultadoIni, 8, "authID="+config.Ramal)
// 	}
// }
