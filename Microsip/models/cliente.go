package models

type ClienteConfig struct {
	ID           int    `json:"id"`
	Doc          int64  `json:"doc"`
	Cliente      string `json:"cliente"`
	QuantRamais  int    `json:"quantRamais"`
	GrupoRecurso string `json:"grupoRecurso"`
	LinkGvc      string `json:"linkGvc"`
	Porta        string `json:"porta"`
	Ramal        string `json:"ramal"`
	Senha        string `json:"senha"`
}
