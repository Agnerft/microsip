package models

type QuantRamaisOpen struct {
	Ramal int  `json:"ramal"`
	INUSE bool `json:"INUSE"`
}

type ClienteConfig struct {
	ID              int               `json:"id"`
	Doc             int64             `json:"doc"`
	Cliente         string            `json:"cliente"`
	QuantRamaisOpen []QuantRamaisOpen `json:"quantRamaisOpen"`
	GrupoRecurso    string            `json:"grupoRecurso"`
	LinkGvc         string            `json:"linkGvc"`
	Porta           string            `json:"porta"`
	Ramal           string            `json:"ramal"`
	Senha           string            `json:"senha"`
}

type Cliente struct {
	Clientes []ClienteConfig `json:"clientes"`
}
