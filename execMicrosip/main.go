package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {

	resultado := "./main.exe"
	cmd := exec.Command(resultado) // Substitua "./main.exe" pelo caminho real do seu arquivo .exe

	// Defina a saída padrão e os descritores de erro para que você possa ver a saída
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println("Erro ao executar o comando:", err)
	}
}
