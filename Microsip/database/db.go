package database

import (
	"Microsip/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

var (
	login string
	senha string
	c     models.ClienteConfig
)

func BuscaPorDoc(doc string, c []models.ClienteConfig) ([]byte, error) {
	//doc := 12310400000182
	login := "root"
	senha := "agner102030"
	url := "https://" + login + ":" + senha + "@" + "basesip.makesystem.com.br/clientes?doc=" + doc
	method := "GET"

	//fmt.Println(url)

	payload := strings.NewReader(``)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Printf("Erro 1", err)
		return nil, err

	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("Erro 2", err)
		return nil, err

	}

	defer res.Body.Close()

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic cm9vdDphZ25lcjEwMjAzMA==")

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Erro 3", err)
		return nil, err

	}

	respBody := string(body)
	bodySemColchetes := strings.ReplaceAll(respBody, "[", "")
	bodySemColchetes = strings.ReplaceAll(bodySemColchetes, "]", "")

	fmt.Println(len(bodySemColchetes))

	if len(bodySemColchetes) == 0 {

		return nil, err
	}

	contagem := strings.Count(respBody, "doc")
	if contagem > 1 {
		fmt.Println("Existe clientes repetidos.")

		return nil, err
	}

	err = json.Unmarshal(body, &c)
	if err != nil {
		fmt.Printf("Aconteceu algo de errado com o json ou o cliente. \n")

		return nil, err
	}

	//fmt.Println(c[0].Cliente)

	return body, nil
}

func EditClient(id int, c models.ClienteConfig) {

	jsonByte := Instancie(c)

	idInt := strconv.Itoa(id)

	url := "https://basesip.makesystem.com.br/clientes/" + idInt
	method := "PUT"

	payload := strings.NewReader(string(jsonByte))

	client := &http.Client{}

	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		fmt.Printf("Erro ao criar a Request: %s", err)
	}

	req.Header.Add("Content-Type", "application/json")

	req.Header.Add("Authorization", "Basic cm9vdDphZ25lcjEwMjAzMA==")

	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("Erro ao conectar: %s", err)

	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Erro ao Ler o body: %s", err)
	}

	fmt.Println(string(body))

}

func Instancie(c models.ClienteConfig) []byte {
	fmt.Println(c)

	client := &models.ClienteConfig{
		ID:           c.ID,
		Doc:          c.Doc,
		Cliente:      c.Cliente,
		GrupoRecurso: c.GrupoRecurso,
		LinkGvc:      c.LinkGvc,
		Porta:        c.Porta,
		Ramal:        c.Ramal,
		Senha:        c.Senha,
		QuantRamais:  c.QuantRamais,
	}

	fmt.Println(&client)

	jsonBytes, err := json.Marshal(&client)
	if err != nil {
		fmt.Printf("Erro ao converter Struct em JSON: %s", err)
	}
	return jsonBytes

}
