package goparsy

import (
	"encoding/json"
	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"strings"
)

type Parsy struct {
	Selectors map[string]Field
	Html      string
}

type Field struct {
	Selector     string
	SelectorType string `yaml:"selectorType"`
	Multiple     bool
	ReturnType   string `yaml:"returnType"`
	Children     map[string]Field
}

func NewParsyFromFilepath(definitionPath string, html string) (Parsy, error) {
	fileData, err := os.ReadFile(definitionPath)
	if err != nil {
		log.Fatal(err)
	}
	data := make(map[string]Field)
	err2 := yaml.Unmarshal(fileData, &data)
	if err2 != nil {
		log.Fatal(err2)
	}
	return Parsy{Selectors: data, Html: html}, nil

}

func NewParsyFromString(definitionString string, html string) (Parsy, error) {
	data := make(map[string]Field)
	err2 := yaml.Unmarshal([]byte(definitionString), &data)
	if err2 != nil {
		log.Fatal(err2)
	}
	return Parsy{Selectors: data, Html: html}, nil
}

func (p Parsy) ExtractField(node *html.Node, field Field) any {
	if field.SelectorType == "CSS" {
		field.Selector = Convert(field.Selector, GLOBAL)
		field.Selector += "/text()"
	}
	if field.Multiple {
		found, err := htmlquery.QueryAll(node, field.Selector)
		if err != nil {
			return nil
		}
		return GetMultiFields(found, field.ReturnType)
	}
	item := htmlquery.FindOne(node, field.Selector)
	if field.Children != nil {
		childrenResult := map[string]any{}
		for key, childField := range field.Children {
			childrenResult[key] = p.ExtractField(item, childField)
		}
		return childrenResult
	}
	if item != nil {
		return ConvertToType(item.Data, field.ReturnType)
	}
	return nil
}

func (p Parsy) GetMap() map[string]any {
	result := map[string]any{}
	doc, err := htmlquery.Parse(strings.NewReader(p.Html))
	if err != nil {
		log.Fatal(err)
	}
	for key, field := range p.Selectors {
		result[key] = p.ExtractField(doc, field)
	}
	return result
}

func (p Parsy) GetJSON() ([]byte, error) {
	return json.Marshal(p.GetMap())
}
