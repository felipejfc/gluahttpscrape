package gluahttpscrape

import (
	"strings"

	"github.com/layeh/gopher-luar"
	"github.com/yhat/scrape"
	"github.com/yuin/gopher-lua"
	"golang.org/x/net/html"
)

type httpScrapeModule struct{}

func NewHttpScrapeModule() *httpScrapeModule {
	return &httpScrapeModule{}
}

func (h *httpScrapeModule) Loader(L *lua.LState) int {
	mod := L.SetFuncs(L.NewTable(), map[string]lua.LGFunction{
		"find_attr_by_class": h.findAttrByClass,
		"find_text_by_class": h.findTextByClass,
		//"findById":    h.findById,
	})
	L.Push(mod)
	return 1
}

func getMatcher(selector, value string) scrape.Matcher {
	matcher := func(n *html.Node) bool {
		return scrape.Attr(n, selector) == value
	}
	return matcher
}

func (h *httpScrapeModule) findAttrByClass(L *lua.LState) int {
	body := L.ToString(1)
	attr := L.ToString(2)
	class := L.ToString(3)
	root, err := html.Parse(strings.NewReader(body))
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}
	results := scrape.FindAll(root, getMatcher("class", class))
	attrResults := []string{}
	for _, result := range results {
		attrResults = append(attrResults, scrape.Attr(result, attr))
	}
	L.Push(luar.New(L, attrResults))
	L.Push(lua.LNil)
	return 2
}

func (h *httpScrapeModule) findTextByClass(L *lua.LState) int {
	body := L.ToString(1)
	class := L.ToString(2)
	root, err := html.Parse(strings.NewReader(body))
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}
	results := scrape.FindAll(root, getMatcher("class", class))
	attrResults := []string{}
	for _, result := range results {
		attrResults = append(attrResults, scrape.Text(result))
	}
	L.Push(luar.New(L, attrResults))
	L.Push(lua.LNil)
	return 2
}
