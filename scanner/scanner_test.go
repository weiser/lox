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

func TestScannerTokenAndError(t *testing.T) {
	scanner := MakeScanner(")@")
	toks := scanner.ScanTokens()
	errs := scanner.Errors
	if len(errs) != 1 {
		t.Errorf("should have one error. @ is not a valid lexeme")
	}
	if toks[0].TokenType != token.RIGHT_PAREN {
		t.Errorf("token should be ')'. %v", toks)
	}
}

func TestScannerTwoCharLexemes(t *testing.T) {
	scanner := MakeScanner("!=")
	toks := scanner.ScanTokens()
	if toks[0].TokenType != token.BANG_EQUAL {
		t.Errorf("token should be token.BANG_EQUAL, got %v", toks[0])
	}
}

func TestScannerComment(t *testing.T) {
	scanner := MakeScanner(`
	// 123
	=
`)
	toks := scanner.ScanTokens()
	if toks[0].TokenType != token.EQUAL {
		t.Errorf("token should be token.EQUAL, got %v", toks[0])
	}
}

func TestScannerMultiline(t *testing.T) {
	scanner := MakeScanner(`
	// comment
	(( ))
	!
`)
	toks := scanner.ScanTokens()
	if len(toks) != 6 {
		t.Errorf("got wrong number of tokens, got %v", toks)
	}
	if toks[0].TokenType != token.LEFT_PAREN {
		t.Errorf("token should be token.LEFT_PAREN, got %v", toks[0])
	}
}

func TestScannerString(t *testing.T) {
	scanner := MakeScanner(`"hi mom"`)
	toks := scanner.ScanTokens()

	if toks[0].TokenType != token.STRING && toks[0].Literal != "hi mom" {
		t.Errorf("token should be token.STRING, got %v", toks[0])
	}
}
