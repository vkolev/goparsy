package goparsy

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
)

func TestNewParsyFromString(t *testing.T) {
	var tests = []struct {
		description, html string
		want              Parsy
	}{
		{
			`title:
  selector: //h1/text()
  selectorType: XPATH
  multiple: false
  returnType: STRING`,
			`<html><head><title>Test</title></head><body><h1>Hello</h1></body></html>`,
			Parsy{
				Selectors: map[string]Field{
					"title": Field{
						Selector:     "//h1/text()",
						SelectorType: "XPATH",
						Multiple:     false,
						ReturnType:   "STRING",
						Children:     nil,
					},
				},
				Html: `<html><head><title>Test</title></head><body><h1>Hello</h1></body></html>`,
			},
		},
		{
			`title:
  selector: //h1/text()
  selectorType: XPATH
  multiple: false
  returnType: STRING
  children:
    child:
      selector: //h2/text()
      selectorType: XPATH
      multiple: true
      returnType: STRING
    second_child:
      selector: //h3/text()
      selectorType: XPATH
      multiple: false
      returnType: STRING`,
			`<html><head><title>Test</title></head><body><h1>Hello</h1></body></html>`,
			Parsy{
				Selectors: map[string]Field{
					"title": Field{
						Selector:     "//h1/text()",
						SelectorType: "XPATH",
						Multiple:     false,
						ReturnType:   "STRING",
						Children: map[string]Field{
							"child": Field{
								Selector:     "//h2/text()",
								SelectorType: "XPATH",
								Multiple:     true,
								ReturnType:   "STRING",
								Children:     nil,
							},
							"second_child": Field{
								Selector:     "//h3/text()",
								SelectorType: "XPATH",
								Multiple:     false,
								ReturnType:   "STRING",
								Children:     nil,
							},
						},
					},
				},
				Html: `<html><head><title>Test</title></head><body><h1>Hello</h1></body></html>`,
			},
		},
	}
	for _, tt := range tests {
		parsy, err := NewParsyFromString(tt.description, tt.html)
		if err != nil {
			t.Errorf("Failed with %s", err)
		}
		assert.Equal(t, parsy, tt.want)
	}
}

func TestNewParsyFromFilepath(t *testing.T) {
	path, err := os.Getwd()
	if err != nil {
		t.Errorf("Test file can not be located")
	}
	testPathBase := strings.Replace(path, "/pkg", "/test_assets/test_base.yaml", 1)
	testPathChildren := strings.Replace(path, "/pkg", "/test_assets/test_children.yaml", 1)

	var tests = []struct {
		description, html string
		want              Parsy
	}{
		{
			testPathBase,
			`<html><head><title>Test</title></head><body><h1>Hello</h1></body></html>`,
			Parsy{
				Selectors: map[string]Field{
					"title": Field{
						Selector:     "//h1/text()",
						SelectorType: "XPATH",
						Multiple:     false,
						ReturnType:   "STRING",
						Children:     nil,
					},
				},
				Html: `<html><head><title>Test</title></head><body><h1>Hello</h1></body></html>`,
			},
		},
		{
			testPathChildren,
			`<html><head><title>Test</title></head><body><h1>Hello</h1></body></html>`,
			Parsy{
				Selectors: map[string]Field{
					"title": Field{
						Selector:     "//h1/text()",
						SelectorType: "XPATH",
						Multiple:     false,
						ReturnType:   "STRING",
						Children: map[string]Field{
							"child": Field{
								Selector:     "//h2/text()",
								SelectorType: "XPATH",
								Multiple:     true,
								ReturnType:   "STRING",
								Children:     nil,
							},
							"second_child": Field{
								Selector:     "//h3/text()",
								SelectorType: "XPATH",
								Multiple:     false,
								ReturnType:   "STRING",
								Children:     nil,
							},
						},
					},
				},
				Html: `<html><head><title>Test</title></head><body><h1>Hello</h1></body></html>`,
			},
		},
	}
	for _, tt := range tests {
		parsy, err := NewParsyFromFilepath(tt.description, tt.html)
		if err != nil {
			t.Errorf("Failed with error: %s", err)
		}
		assert.Equal(t, parsy.Selectors, tt.want.Selectors)
	}
}

func TestParsy_Execute(t *testing.T) {
	var tests = []struct {
		description, html string
		want              map[string]interface{}
	}{
		{
			`title:
  selector: //h1/text()
  selectorType: XPATH
  multiple: false
  returnType: STRING`,
			`<html><head><title>Test</title></head><body><h1>Hello</h1></body></html>`,
			map[string]interface{}{"title": "Hello"},
		},
		{
			`title:
  selector: h1
  selectorType: CSS
  multiple: false
  returnType: STRING`,
			`<html><head><title>Test</title></head><body><h1>Hello</h1></body></html>`,
			map[string]interface{}{"title": "Hello"},
		},
	}
	for _, tt := range tests {
		parsy, err := NewParsyFromString(tt.description, tt.html)
		if err != nil {
			t.Errorf("Failed with %s", err)
		}
		result := parsy.GetMap()
		resultJson, _ := parsy.GetJSON()
		wantJson, _ := json.Marshal(tt.want)
		assert.Equal(t, result, tt.want)
		assert.Equal(t, result["test"], tt.want["test"])
		assert.Equal(t, resultJson, wantJson)
	}
}
