package gluahttpscrape

import (
	"strings"
	"testing"

	"github.com/yuin/gopher-lua"
)

var httpBody = `<!DOCTYPE html><html><body><h1 id=\"testid\" class=\"testclass\" href=\"testhref\">My First Heading</h1><h1 id=\"testid\" class=\"testclass\" href=\"testhref2\">My First Heading</h1><p>My first paragraph.</p></body></html>`

func TestFindAttrByClass(t *testing.T) {
	if err := evalLua(t, `
		local scrape = require("scrape")
		response, error = scrape.find_attr_by_class("`+httpBody+`", "href", "testclass")
		assert_equal("testhref", response[1])
		assert_equal("testhref2", response[2])
	`); err != nil {
		t.Errorf("Failed to evaluate script: %s", err)
	}
}

func TestFindTextByClass(t *testing.T) {
	if err := evalLua(t, `
		local scrape = require("scrape")
		response, error = scrape.find_text_by_class("`+httpBody+`", "testclass")
		assert_equal("My First Heading", response[1])
		assert_equal("My First Heading", response[2])
	`); err != nil {
		t.Errorf("Failed to evaluate script: %s", err)
	}
}

func TestFindAttrById(t *testing.T) {
	if err := evalLua(t, `
		local scrape = require("scrape")
		response, error = scrape.find_attr_by_id("`+httpBody+`", "href", "testid")
		assert_equal("testhref", response[1])
		assert_equal("testhref2", response[2])
	`); err != nil {
		t.Errorf("Failed to evaluate script: %s", err)
	}
}

func TestFindTextById(t *testing.T) {
	if err := evalLua(t, `
		local scrape = require("scrape")
		response, error = scrape.find_text_by_id("`+httpBody+`", "testid")
		assert_equal("My First Heading", response[1])
		assert_equal("My First Heading", response[2])
	`); err != nil {
		t.Errorf("Failed to evaluate script: %s", err)
	}
}

func TestFindAttrByTag(t *testing.T) {
	if err := evalLua(t, `
		local scrape = require("scrape")
		response, error = scrape.find_attr_by_tag("`+httpBody+`", "href", "h1")
		assert_equal("testhref", response[1])
		assert_equal("testhref2", response[2])
	`); err != nil {
		t.Errorf("Failed to evaluate script: %s", err)
	}
}

func TestFindTextByTag(t *testing.T) {
	if err := evalLua(t, `
		local scrape = require("scrape")
		response, error = scrape.find_text_by_tag("`+httpBody+`", "h1")
		assert_equal("My First Heading", response[1])
		assert_equal("My First Heading", response[2])
	`); err != nil {
		t.Errorf("Failed to evaluate script: %s", err)
	}
}

func evalLua(t *testing.T, script string) error {
	L := lua.NewState()
	defer L.Close()

	L.PreloadModule("scrape", NewHttpScrapeModule().Loader)

	L.SetGlobal("assert_equal", L.NewFunction(func(L *lua.LState) int {
		expected := L.Get(1)
		actual := L.Get(2)

		if expected.Type() != actual.Type() || expected.String() != actual.String() {
			t.Errorf("Expected %s %q, got %s %q", expected.Type(), expected, actual.Type(), actual)
		}

		return 0
	}))

	L.SetGlobal("assert_contains", L.NewFunction(func(L *lua.LState) int {
		contains := L.Get(1)
		actual := L.Get(2)

		if !strings.Contains(actual.String(), contains.String()) {
			t.Errorf("Expected %s %q contains %s %q", actual.Type(), actual, contains.Type(), contains)
		}

		return 0
	}))

	return L.DoString(script)
}
