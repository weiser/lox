package scanner

import (
	"testing"

	"github.com/weiser/lox/token"
)

func TestScannerCreationOneCharacterLexemes(t *testing.T) {
	lexToTok := map[string]token.TType{
		"(": token.LEFT_PAREN,
		")": token.RIGHT_PAREN,
		"{": token.LEFT_BRACE,
		"}": token.RIGHT_BRACE,
		",": token.COMMA,
		".": token.DOT,
		"-": token.MINUS,
		"+": token.PLUS,
		";": token.SEMICOLON,
		"*": token.STAR,
	}

	for _, src := range "(){},.-+;*" {
		scanner := MakeScanner(string(src))
		toks := scanner.ScanTokens()
		ans := token.MakeToken(lexToTok[string(src)], string(src), nil, 1)
		if toks[0] != ans {
			t.Errorf("got %v and wanted %v", toks[0], ans)
		}
	}
}

func TestScannerErrors(t *testing.T) {

	for _, src := range "@" {
		scanner := MakeScanner(string(src))
		scanner.ScanTokens()

		if len(scanner.Errors) != 1 {
			t.Errorf("got unexpected number of errors! %v", scanner.Errors)
		}
	}
}
