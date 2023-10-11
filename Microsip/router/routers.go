package router

import (
	"Microsip/controle"
	"Microsip/mid"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func HeandleRequest() {
	r := mux.NewRouter()
	r.Use(mid.ContentType)

	r.HandleFunc("/", controle.Home)
	r.HandleFunc("/execute", controle.SaveExeEIni)

	r.HandleFunc("/documento/{doc}", controle.FindUser)
	r.HandleFunc("/remover", controle.Remover)
	r.HandleFunc("/criando", controle.CriandoConta)

	fmt.Println("Servidor está ouvindo na porta 8080...")
	err := http.ListenAndServe("0.0.0.0:8080", r)
	if err != nil {
		fmt.Println("Erro ao iniciar o servidor:", err)
	}
}
