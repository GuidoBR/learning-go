package main

import (
        "testing"
        "strings"
		"os"
		"reflect"
)

const linhaLetraA = `0041;LATIN CAPITAL LETTER A;Lu;0;L;;;;;N;;;;0061;`
const linhas3Da43 = `
003D;EQUALS SIGN;Sm;0;ON;;;;;N;;;;;
003E;GREATER-THAN SIGN;Sm;0;ON;;;;;Y;;;;;
003F;QUESTION MARK;Po;0;ON;;;;;N;;;;;
0040;COMMERCIAL AT;Po;0;ON;;;;;N;;;;;
0041;LATIN CAPITAL LETTER A;Lu;0;L;;;;;N;;;;0061;
0042;LATIN CAPITAL LETTER B;Lu;0;L;;;;;N;;;;0062;
0043;LATIN CAPITAL LETTER C;Lu;0;L;;;;;N;;;;0063;
`

func TestAnalisarLinha(t *testing.T) {
        runa, nome, palavras := AnalisarLinha(linhaLetraA)
        if runa != 'A' {
                t.Errorf("Esperava 'A', veio %q", runa)
        }
        const nomeA = "LATIN CAPITAL LETTER A"
        if nome != nomeA {
            t.Errorf("Esperava %q, veio %q", nomeA, nome)
        }
		palavrasA := []string{"LATIN", "CAPITAL", "LETTER", "A"}
		if !reflect.DeepEqual(palavras, palavrasA) {
			t.Errorf("\n\tEsperado: %q\n\trecebido: %q", palavrasA, palavras)
		}
}

func ExampleListar() {
    texto := strings.NewReader(linhas3Da43)
    Listar(texto, "MARK")
	  // Output: U+003F	?	QUESTION MARK
}

func ExampleListar_doisResultados() {
    texto := strings.NewReader(linhas3Da43)
    Listar(texto, "SIGN")
    // Output:
    // U+003D	=	EQUALS SIGN
    // U+003E	>	GREATER-THAN SIGN
}

func ExampleListar_duasPalavras() {
	texto := strings.NewReader(linhas3Da43)
	Listar(texto, "CAPITAL LATIN")
	// Output:
	// U+0041	A	LATIN CAPITAL LETTER A
	// U+0042	B	LATIN CAPITAL LETTER B
	// U+0043	C	LATIN CAPITAL LETTER C
}

func Example_consultaDuasPalavras() {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{"", "cat", "smiling"}
	main()
	// Output:
	// U+1F638	😸	GRINNING CAT FACE WITH SMILING EYES
	// U+1F63A	😺	SMILING CAT FACE WITH OPEN MOUTH
	// U+1F63B	😻	SMILING CAT FACE WITH HEART-SHAPED EYES
}

func Example() {
        oldArgs := os.Args
        defer func() { os.Args = oldArgs }()
        os.Args = []string{"", "cruzeiro"}
        main()
		// Output:
		// U+20A2	₢	CRUZEIRO SIGN
}

func TestContém(t *testing.T) {
	var casos = []struct {
		fatia		[]string
		procurado	string
		esperado	bool
	} {
		{[]string{"A", "B"}, "B", true},
		{[]string{}, "A", false},
		{[]string{"A", "B"}, "Z", false},
	}
	for _, caso := range casos {
		obtido := contém(caso.fatia, caso.procurado)
		if obtido != caso.esperado {
			t.Errorf("contém(%#v, %#v) esperado: %v; recebido: %v",
				caso.fatia, caso.procurado, caso.esperado, obtido)
		}
	}
}

func TestContémTodos(t *testing.T) {
	casos := []struct {
		fatia      []string
		procurados []string
		esperado   bool
	}{
		{[]string{"A", "B"}, []string{"B"}, true},
		{[]string{}, []string{"A"}, false},
		{[]string{"A"}, []string{}, true},
		{[]string{}, []string{}, true},
		{[]string{"A", "B"}, []string{"Z"}, false},
		{[]string{"A", "B", "C"}, []string{"A", "C"}, true},
		{[]string{"A", "B", "C"}, []string{"A", "Z"}, false},
		{[]string{"A", "B"}, []string{"A", "B", "C"}, false},
	}
	for _, caso := range casos {
		obtido := contémTodos(caso.fatia, caso.procurados)
		if obtido != caso.esperado {
			t.Errorf("contémTodos(%#v, %#v)\nesperado: %v; recebido: %v",
				caso.fatia, caso.procurados, caso.esperado, obtido)
		}
	}
}
