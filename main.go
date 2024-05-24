package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func execCmd(cmd string) (string, error) {
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}

func main() {
	myToken := "Personal access tokens (classic)"
	myUser := "GitHub user"

	// Obtém o diretório de trabalho atual
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Erro ao obter o diretório atual:", err)
		return
	}

	// Obtém o nome do diretório atual
	_, dirName := filepath.Split(dir)

	// Verifica se o diretório contém espaços
	if strings.Contains(dir, " ") {
		fmt.Println("Não é possivel criar um repositório chamado:", dirName)
		fmt.Println(" -> Use ' _ ' ou ' - ' para separar palavras no lugar de espaços")
		return
	}

	fmt.Println("O nome do seu repositório será:", dirName)

	var repoName string = dirName
	var commitMessage string

	fmt.Print("Mensagem de commit: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		commitMessage = scanner.Text()
	}
	commitMessage = strings.TrimSpace(commitMessage)

	fmt.Println(execCmd("git init"))
	fmt.Println(execCmd("git branch -m master main"))
	fmt.Println(execCmd("git add ."))
	fmt.Println(execCmd(fmt.Sprintf("git commit -m \"%s\"", commitMessage)))

	fmt.Println(execCmd(fmt.Sprintf("curl -u %s:%s "+
		"https://api.github.com/user/repos -d '{\"name\":\"%s\"}'", myUser, myToken, repoName)))
	fmt.Println(execCmd(fmt.Sprintf("git remote add origin git@github.com:%s/"+
		"%s.git", myUser, repoName)))
	fmt.Println(execCmd("git push -u origin main"))
}
