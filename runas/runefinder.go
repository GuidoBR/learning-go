package main

import (
    "io"
    "fmt"
    "strconv"
    "strings"
    "bufio"
	"os"
	"log"
)

func AnalisarLinha(ucdLine string) (rune, string, []string) {
    campos := strings.Split(ucdLine, ";")
    código, _ := strconv.ParseInt(campos[0], 16, 32)
	palavras := strings.Fields(campos[1])
    return rune(código), campos[1], palavras
}

// Listar exibe na saída padrão o código, a runa e o nome dos caracteres Unicode
// cujo nome contém o texto da consulta
func Listar(texto io.Reader, consulta string) {
	termos := strings.Fields(consulta)
    varredor := bufio.NewScanner(texto)
    for varredor.Scan() {
        linha := varredor.Text()
        if strings.TrimSpace(linha) == "" {
          continue
        }
        runa, nome, palavrasNome := AnalisarLinha(linha)
        if contémTodos(palavrasNome, termos) {
          fmt.Printf("U+%04X\t%[1]c\t%s\n", runa, nome)
        }
    }
}

func contém(fatia []string, procurado string) bool {
	for _, item := range fatia {
		if item == procurado {
			return true
		}
	}
	return false
}

func contémTodos(fatia []string, procurados []string) bool {
	for _, procurado := range procurados {
		if !contém(fatia, procurado) {
			return false
		}
	}
	return true
}

func main() {
	ucd, err := os.Open("UnicodeData.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer func() { ucd.Close() }()
	consulta := strings.Join(os.Args[1:], " ")
	Listar(ucd, strings.ToUpper(consulta))
}
