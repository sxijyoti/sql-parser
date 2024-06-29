package main

import (
	"fmt"
	"strings"
)

type keyword struct {
	selectKeyword string
	fromKeyword   string
}

const (
	selectKeywordVal = "SELECT"
	fromKeywordVal   = "FROM"
)

type symbol struct {
	asteriskSymbol  rune
	semicolonSymbol rune
	commaSymbol     rune
}

const (
	asteriskSymbolVal  = '*'
	semicolonSymbolVal = ';'
	commaSymbolVal     = ','
)

func LexSQL(query string, keywords keyword, symbols symbol) (tokens []string, columns []string, table []string) {
	query = strings.ToUpper(query)
	selectIndex := strings.Index(query, keywords.selectKeyword)
	fromIndex := strings.Index(query, keywords.fromKeyword)
	semicolonIndex := strings.Index(query, string(symbols.semicolonSymbol))

	if selectIndex == -1 || fromIndex == -1 || semicolonIndex == -1 {
		return nil, nil, nil
	}
	tokens = append(tokens, keywords.selectKeyword)
	tokens = append(tokens, keywords.fromKeyword)
	// tokens = append(token,string(symbol.semicolonSymbol))
	columnstr := strings.TrimSpace(query[selectIndex+len(keywords.selectKeyword) : fromIndex])
	tablestr := strings.TrimSpace(query[fromIndex+len(keywords.fromKeyword) : semicolonIndex])

	if columnstr == string(symbols.asteriskSymbol) {
		columns = append(columns, string(symbols.asteriskSymbol))
		tokens = append(tokens, string(symbols.asteriskSymbol))
	} else {
		columns = strings.Split(columnstr, string(symbols.commaSymbol))
		for i := range columns {
			columns[i] = strings.TrimSpace(columns[i])
			tokens = append(tokens, columns[i])
		}
	}
	table = append(table, strings.TrimSpace(tablestr))
	tokens = append(tokens, strings.TrimSpace(tablestr))
	tokens = append(tokens, string(symbols.semicolonSymbol))

	return tokens, columns, table

}

func main() {

	keywords := keyword{
		selectKeyword: selectKeywordVal,
		fromKeyword:   fromKeywordVal,
	}

	symbols := symbol{
		asteriskSymbol:  asteriskSymbolVal,
		semicolonSymbol: semicolonSymbolVal,
		commaSymbol:     commaSymbolVal,
	}

	query := "SELECT * FROM Table_name ;"
	tokens, columns, table := LexSQL(query, keywords, symbols)

	fmt.Println(" ")
	fmt.Println("Tokens: ")
	fmt.Println(" ")
	for _, c := range tokens {
		fmt.Println(c)
	}
	fmt.Println(" ")
	fmt.Println("Columns: ")
	for _, c := range columns {
		fmt.Println(c)
	}
	fmt.Println(" ")
	fmt.Println("Table: ")
	for _, c := range table {
		fmt.Println(c)
	}
	fmt.Println(" ")
}
