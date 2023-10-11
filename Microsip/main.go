package main

import (
	"Microsip/database"
	"Microsip/models"
	"Microsip/router"
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

var clientConfig []models.ClienteConfig
var docCliente models.Doc

func main() {

	// Iniciar o servidor na porta 8080.
	router.HeandleRequest()
	var doc string

	fmt.Print("Digite o documento da empresa: ")
	_, err := fmt.Scanln(&doc)
	if err != nil {
		fmt.Println("Tente outra vez")
	}

	fmt.Println("Agora vamos verificar se vc está na nossa base de dados, um momento")
	time.Sleep(1)

	// Salvando o arquivo ini na pasta \\AppData\\Roamming

	// Editando o arquivo

	//jsonfile, _ := database.BuscaPorDoc(doc, clientConfig)
	//fmt.Println(jsonfile)

	//if jsonfile == nil {
	//	fmt.Println("Desculpe, não encontrei você . . . ")
	//	return
	//}

	//fmt.Println("Encontrei você, ")
	//fmt.Println("Os ramais que tenho vinculados a sua base são:")
	//if err := json.Unmarshal(jsonfile, &clientConfig); err != nil {
	//	fmt.Println("Erro ao fazer o Unmarshal do JSON:", err)
	//	return
	//}

	//fmt.Println(&clientConfig)

	// Edição e Salvamento do arquivo .ini
	for _, config := range clientConfig {

		for i := range config.QuantRamais {
			fmt.Println(config.QuantRamais[i].Ramal)
		}

		fmt.Println("Porém os que não foram configurados ainda são:")

		for i := range config.QuantRamais {
			if config.QuantRamais[i].InUse == false {
				fmt.Println(config.QuantRamais[i].Ramal)
			}

		}

		var ramal string
		fmt.Print("Por favor informe agora, qual ramal você vai utilizar? ")
		_, err := fmt.Scanln(&ramal)
		if err != nil {
			fmt.Println("Erro ao ler a entrada:", err)
			return
		}
		//database.EditClient(config.ID, config)

		config.Ramal = ramal

		ramalInt, _ := strconv.Atoi(ramal)

		found := false

		for i := range config.QuantRamais {

			if config.QuantRamais[i].InUse != found {
				fmt.Println("Ramal sendo usado.")
				break
			}

			if config.QuantRamais[i].Ramal == ramalInt {
				config.QuantRamais[i].InUse = true
				found = true
				break
			}
		}

		if !found {
			fmt.Println("Ramal não encontrado.")
			return
		}

		fmt.Println(config)
		fmt.Println("aqui?")

		// Codifique a estrutura de dados atualizada para JSON
		updatedJSON, err := json.Marshal(config.QuantRamais)
		if err != nil {
			fmt.Println("Erro ao serializar o JSON:", err)
			return
		}

		fmt.Printf(string(updatedJSON))

		database.EditClient(config.ID, config)
		//
		//
		//
		fmt.Println("Oi")
		//linkCompleto := config.GrupoRecurso + config.LinkGvc + config.Porta
		//utils.Editor(resultadoIni, 2, "label="+config.Ramal)
		//utils.Editor(resultadoIni, 3, "server="+linkCompleto)
		//utils.Editor(resultadoIni, 4, "proxy="+linkCompleto)
		//utils.Editor(resultadoIni, 5, "domain="+linkCompleto)
		//utils.Editor(resultadoIni, 6, "username="+config.Ramal)
		//utils.Editor(resultadoIni, 7, "password="+config.Ramal+config.Senha)
		//utils.Editor(resultadoIni, 8, "authID="+config.Ramal)
	}

}

//unc isValidDocument() {

//docInt, _ := strconv.Atoi(doc)
//var url string
//for i := range url {
//	s := strconv.Itoa(i)
//	url := "https://basesip.makesystem.com.br/clientes/" + s

//	fmt.Println(url)
//	response, err := http.Get(url)
//		if err != nil {
//fmt.Println("Erro ao fazer a solicitação HTTP:", err)
//			return
//		}

//		defer response.Body.Close()

//}

// Verifique o código de status da resposta HTTP
///if response.StatusCode != http.StatusOK {
//fmt.Println("Erro na resposta HTTP:", response.Status)
//return
//}

//decoder := json.NewDecoder(response.Body)
//if err := decoder.Decode(&docCliente); err != nil {
//	fmt.Println("Erro ao decodificar JSON:", err)
//	return
//}

// Itere sobre os clientes e imprima apenas o campo "doc"
//for _, cliente := range docCliente {

//	fmt.Println("Doc do cliente:", cliente.Doc)
//}

// }
