// Package interpreter implement the Interpreter design pattern.
// It provides a domain-specific language (DSL) for scripting fairytales.
// The Tokenize function performs lexical analysis, generating an abstract syntax token list.
// The Parse function interprets these tokens to produce a human-readable representation of the
// fairytale narrative.
package interpreter

import "fmt"

func Run() {
	fairytale := `
		TITLE Red Hood;
		CHARACTER RedHood;
		RedHood SAY I will visit my grandma!;
		RedHood MOVE Forest;
		RedHood ENCOUNTER Wolf;
		RedHood SAY Hello to you, Mr. Wolf!;
		Wolf EATS Grandma;
		CHARACTER Hunter;
		Hunter SHOOTS Wolf;
		RedHood SAY Thank you, hunter!;
		RedHood ENCOUNTER alive Grandma;
		END;
	`
	fmt.Println(Parse(Tokenize(fairytale)))
}
