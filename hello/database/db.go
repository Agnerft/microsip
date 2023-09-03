package database

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func BuscaPorDoc(doc int) (string, error) {
	//doc := 12310400000182
	// ajustar porta
	url := "http://localhost:3000/clientes?doc=" + strconv.Itoa(doc)
	method := "GET"

	//fmt.Println(url)

	payload := strings.NewReader(``)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		//return

	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		//return

	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		//return

	}
	fmt.Println(string(body))

	return string(body), nil
}

func AtualizarINUSE(w http.ResponseWriter, r *http.Request) {
	json, _ := BuscaPorDoc(12310400000182)

	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		//return
	}

	// Obter o ramal desejado dos parâmetros da URL
	ramalParam := r.URL.Query().Get("ramal")
	if ramalParam == "" {
		http.Error(w, "O parâmetro 'ramal' é obrigatório", http.StatusBadRequest)
		//return
	}

	// // Serializar a estrutura de dados de volta em JSON
	// jsonUpdated, err := json.MarshalIndent(clientesData, "", "    ")
	// if err != nil {
	//     http.Error(w, "Erro ao fazer o Marshal", http.StatusInternalServerError)
	//     return
	// }

	// Salvar o JSON atualizado em um arquivo (ou outra fonte)
	err := ioutil.WriteFile("seuarquivo_atualizado.json", json, 0644)
	if err != nil {
		http.Error(w, "Erro ao escrever o arquivo JSON atualizado", http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "JSON atualizado foi salvo com sucesso.")

	//return ramalParam
}
