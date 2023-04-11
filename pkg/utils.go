package goparsy

import (
	"golang.org/x/net/html"
	"strconv"
)

func GetMultiFields(results []*html.Node, returnType string) []any {
	var result []any
	for _, r := range results {
		result = append(result, ConvertToType(r.Data, returnType))
	}
	return result
}

func ConvertToType(data string, returnType string) any {
	var result any
	switch returnType {
	case "STRING":
		result = data
	case "INTEGER":
		result, _ = strconv.ParseInt(data, 10, 32)
	case "FLOAT":
		result, _ = strconv.ParseFloat(data, 32)
	}
	return result
}
