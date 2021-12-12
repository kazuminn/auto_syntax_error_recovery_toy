package main

import (
	"fmt"
)

type token struct {
	type_token string
	value      string
}

func main() {
	input := "if ( true { "
	token_list := []token{}

	o := ""	
	for _, v := range input {
		if string(v) != " " {
			o += string(v)
			continue
		}

		if o == "if" {
			token_list = append(token_list, token{"if", "if"})
			o = ""
		} else if o == "(" {
			token_list = append(token_list, token{"brace_left", "("})
			o = ""
		} else if o == ")" {
			token_list = append(token_list, token{"brace_right", ")"})
			o = ""
		} else if o == "true" {
			token_list = append(token_list, token{"bool", "true"})
			o = ""
		} else if o == "{" {
			token_list = append(token_list, token{"tiny_brace_left", "{"})
			o = ""
		}
	}

	//構文解析 and 構文エラー修正
	new_token_list := []token{}
	for i, t := range token_list {
		if t.type_token == "bool" && token_list[2].type_token != "brace_right" {
			new_token_list = append(new_token_list, token_list[:i + 1]...)
			new_token_list = append(new_token_list, token{"brace_right", ")"})
			new_token_list = append(new_token_list, token_list[i+1:]...)
		}
	}

	for _, v := range new_token_list {
		fmt.Print(string(v.value) + " ") //if ( true ) {
	}
}
