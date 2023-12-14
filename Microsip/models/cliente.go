package models

import "encoding/json"

type Ramal struct {
	Ramal int  `json:"ramalSelect"`
	InUse bool `json:"inUse"`
}
type ClienteConfig struct {
	ID           int     `json:"id"`
	Doc          int     `json:"doc"`
	Cliente      string  `json:"cliente"`
	GrupoRecurso string  `json:"grupoRecurso"`
	LinkGvc      string  `json:"linkGvc"`
	Porta        string  `json:"porta"`
	Ramal        string  `json:"ramal"`
	Senha        string  `json:"senha"`
	QuantRamais  []Ramal `json:"quantRamaisOpen"`
}

func NewClienteTeste() ClienteConfig {

	return ClienteConfig{
		LinkGvc:     ".gvctelecom.com.br:",
		Senha:       "@abc",
		QuantRamais: []Ramal{},
	}
}

type Cliente struct {
	Clientes []ClienteConfig `json:"clientes"`
}

type Doc []struct {
	Doc string `json:"doc"`
}

func (c *ClienteConfig) ToJSON() ([]byte, error) {
	return json.Marshal(c)
}

// FromJSON desserializa a struct a partir de JSON.
func (c *ClienteConfig) FromJSON(data []byte) error {
	return json.Unmarshal(data, c)
}

type JSONConvertible interface {
	ToJSON() ([]byte, error)
	FromJSON([]byte) error
}
