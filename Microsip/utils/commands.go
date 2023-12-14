package utils

import (
	"fmt"
	"os"
	"os/exec"
)

var desktopPath, _ = os.UserHomeDir()

func Comandos(resultado string) {

	resultadoDel := desktopPath + "AppData\\Local\\MicroSIP\\Uninstall"

	cmd := exec.Command(resultadoDel, "/S")
	if err := cmd.Run(); err != nil {
		fmt.Printf("Erro ao executar o desinstalador: %s ", err)

	}

	fmt.Println("Removido")

	cmd = exec.Command(resultado, "/S")
	if err := cmd.Run(); err != nil {
		fmt.Printf("Erro ao executar o instalador:%s ", err)
		return
	}

	// Executa o comando de Taskill
	processName := "MicroSIP.exe" // Substitua pelo nome do processo que você deseja encerrar

	cmd3 := exec.Command("taskkill", "/F", "/IM", processName)

	// Redirecionar saída e erro, se necessário
	cmd3.Stdout = os.Stdout
	cmd3.Stderr = os.Stderr

	err := cmd3.Run()
	if err != nil {
		fmt.Println("Erro ao executar o comando:", err)
		return
	}
	fmt.Println("Processo encerrado com sucesso.")

}

func Remover() {

}
