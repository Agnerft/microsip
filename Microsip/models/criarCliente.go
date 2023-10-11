package models

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
)

func CriarNovoCliente(cliente *ClienteConfig, url string) ([]byte, error) {
	// Serializar a struct para JSON.
	jsonBytes, err := cliente.ToJSON()
	if err != nil {
		log.Printf("Erro a serializar: %v", err)
	}

	resp, _ := http.Post(url, "application/json", bytes.NewBuffer(jsonBytes))

	defer resp.Body.Close()

	fmt.Println("Passou no CriarNovoCliente?")
	return jsonBytes, nil
}

func (c *ClienteConfig) AdicionarRamais(inicio int, quantidade int) (string, error) {

	for i := inicio; i < inicio+quantidade; i++ {
		novoRamal := Ramal{
			Ramal: i,
			InUse: false,
		}
		c.QuantRamais = append(c.QuantRamais, novoRamal)

	}
	stringRamais := "Ramais criados com sucesso."
	return stringRamais, nil

}
