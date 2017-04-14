package main

import "testing"

const linhaLetraA = `0041;LATIN CAPITAL LETTER A;Lu;0;L;;;;;N;;;;0061;`

func TestAnalisarLinha(t *testing.T) {
        runa, _ := AnalisarLinha(linhaLetraA)
        if runa != 'A' {
                t.Errorf("Esperava A, veio %c", runa)
        }
}
