package main

import (
	"fmt"
	"os/exec"
	"strconv"
)

func main() {
	// Número de qubits``
	n := 3

	// Comando para executar o script Python
	cmd := exec.Command("python3", "deutsch_jozsa.py", strconv.Itoa(n))

	// Capturar a saída do comando
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Erro ao executar o script:", err)
		return
	}

	// Exibir o resultado
	fmt.Println("Resultado:", string(output))
}
