package main

/*
/*  file   : interpreter.go
/*  author : hariprasad
/*  date   : 3 June 2018
/*  last   : 3 June 2018
/*  this file contains the goalng implementation of Interpreter.
It is writtern based on the example provided. in the site: https://ruslanspivak.com/lsbasi-part1/
*/

import (
	"fmt"
	"strconv"
)

// TokenType defines the Tokens
type TokenType string

const (
	// INTEGER Data type
	INTEGER TokenType = "INTEGER"

	// PLUS operator
	PLUS TokenType = "PLUS"

	// EOF (end-of-file) token is used to indicate that
	//  there is no more input left for lexical analysis
	EOF TokenType = "EOF"
)

// Token is a lexical values
type Token struct {
	Type  TokenType
	Value string
}

//Interpreter does Interpretation for the text.
type Interpreter struct {
	text         string
	position     int
	currentToken *Token
}

// expr -> INTEGER PLUS INTEGER
func (i *Interpreter) expr() string {

	// set current token to the first token taken from the input
	i.currentToken = i.getNextToken()

	// we expect the current token to be a single-digit integer
	left := i.currentToken
	i.eat(INTEGER)

	// we expect the current token to be a '+' token
	op := i.currentToken
	i.eat(PLUS)

	// we expect the current token to be a single-digit integer
	right := i.currentToken
	i.eat(INTEGER)

	// after the above call the self.current_token is set to
	// EOF token
	// at this point INTEGER PLUS INTEGER sequence of tokens
	// has been successfully found and the method can just
	// return the result of adding two integers, thus
	// effectively interpreting client input
	var lValue, rValue, result int
	if left.Type == INTEGER {
		lValue, _ = strconv.Atoi(left.Value)
	}

	if left.Type == INTEGER {
		rValue, _ = strconv.Atoi(right.Value)
	}

	if op.Value == "+" {
		result = lValue + rValue
	}

	return strconv.Itoa(result)
}

// Lexical analyzer (also known as scanner or tokenizer)
// This method is responsible for breaking a sentence
// apart into tokens. One token at a time.
func (i *Interpreter) getNextToken() *Token {
	text := i.text

	// is self.pos index past the end of the self.text ?
	// if so, then return EOF token because there is no more
	// input left to convert into tokens
	if i.position > (len(text) - 1) {
		return &Token{Type: "EOF", Value: ""}
	}

	// get a character at the position self.pos and decide
	// what token to create based on the single character
	currentChar := text[i.position]
	if currentChar == '+' {
		token := &Token{Type: "PLUS", Value: string(currentChar)}
		i.position++
		return token
	}

	// if the character is a digit then convert it to
	// integer, create an INTEGER token, increment self.pos
	// index to point to the next character after the digit,
	// and return the INTEGER token
	if string(currentChar) != "" {
		token := &Token{Type: "INTEGER", Value: string(currentChar)}
		i.position++
		return token
	}

	return nil
}

func (i *Interpreter) eat(tokenType TokenType) {
	// compare the current token type with the passed token
	// type and if they match then "eat" the current token
	// and assign the next token to the self.current_token,
	// otherwise raise an exception.
	if i.currentToken.Type == tokenType {
		i.currentToken = i.getNextToken()
	}
}

func main() {

	text := "3+4"
	fmt.Print("Enter the expression ,eg: 3+4\n")
	fmt.Scan(&text)

	interpreter := &Interpreter{
		text: text,
	}

	result := interpreter.expr()
	fmt.Print("Result: ", result, "\n")
}
