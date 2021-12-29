package token

import "testing"

func TestTokenCreation(t *testing.T) {
	tok := MakeToken(LEFT_PAREN, "{", "", 1)
	ans := Token{TokenType: LEFT_PAREN, Lexeme: "{", Literal: "", Line: 1}
	if tok != ans {
		t.Errorf("got %v and wanted %v", tok, ans)
	}
}
