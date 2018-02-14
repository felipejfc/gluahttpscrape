package gluahttpscrape

import (
	"strings"

	luar "github.com/layeh/gopher-luar"
	"github.com/yhat/scrape"
	"github.com/yuin/gopher-lua"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

type httpScrapeModule struct{}

func NewHttpScrapeModule() *httpScrapeModule {
	return &httpScrapeModule{}
}

func (h *httpScrapeModule) Loader(L *lua.LState) int {
	mod := L.SetFuncs(L.NewTable(), map[string]lua.LGFunction{
		"find_attr_by_class":  h.findAttrByClass,
		"find_attr_by_id":     h.findAttrById,
		"find_attr_by_tag":    h.findAttrByTag,
		"find_attrs_by_class": h.findAttrsByClass,
		"find_attrs_by_id":    h.findAttrsById,
		"find_attrs_by_tag":   h.findAttrsByTag,
		"find_text_by_id":     h.findTextById,
		"find_text_by_class":  h.findTextByClass,
		"find_text_by_tag":    h.findTextByTag,
	})
	L.Push(mod)
	return 1
}

func getMatcher(selector, value string) scrape.Matcher {
	var matcher func(*html.Node) bool
	if selector == "class" {
		matcher = scrape.ByClass(value)
	} else if selector == "id" {
		matcher = scrape.ById(value)
	} else if selector == "tag" {
		matcher = scrape.ByTag(atom.Lookup([]byte(value)))
	}
	return matcher
}

func (h *httpScrapeModule) findAttr(selector string, L *lua.LState) int {
	body := L.ToString(1)
	attr := L.ToString(2)
	query := L.ToString(3)
	L.Pop(3)
	root, err := html.Parse(strings.NewReader(body))
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}
	results := scrape.FindAll(root, getMatcher(selector, query))
	attrResults := []string{}
	for _, result := range results {
		attrResults = append(attrResults, scrape.Attr(result, attr))
	}
	L.Push(luar.New(L, attrResults))
	L.Push(lua.LNil)
	return 2
}

func (h *httpScrapeModule) findAttrs(selector string, L *lua.LState) int {
	body := L.ToString(1)
	attrsCount := L.ToInt(2)
	attrs := []string{}
	for i := 1; i <= attrsCount; i++ {
		attrNow := L.ToString(2 + i)
		attrs = append(attrs, attrNow)
	}
	query := L.ToString(2 + attrsCount + 1)
	L.Pop(2 + attrsCount + 1)
	root, err := html.Parse(strings.NewReader(body))
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}
	results := scrape.FindAll(root, getMatcher(selector, query))
	attrResults := []map[string]string{}
	for _, result := range results {
		attrResults = append(attrResults, make(map[string]string))
		idx := len(attrResults) - 1
		for _, attr := range attrs {
			attrResults[idx][attr] = scrape.Attr(result, attr)
		}
	}
	L.Push(luar.New(L, attrResults))
	L.Push(lua.LNil)
	return 2
}

func (h *httpScrapeModule) findText(selector string, L *lua.LState) int {
	body := L.ToString(1)
	query := L.ToString(2)
	L.Pop(2)
	root, err := html.Parse(strings.NewReader(body))
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}
	results := scrape.FindAll(root, getMatcher(selector, query))
	attrResults := []string{}
	for _, result := range results {
		attrResults = append(attrResults, scrape.Text(result))
	}
	L.Push(luar.New(L, attrResults))
	L.Push(lua.LNil)
	return 2
}

func (h *httpScrapeModule) findAttrByClass(L *lua.LState) int {
	return h.findAttr("class", L)
}

func (h *httpScrapeModule) findAttrById(L *lua.LState) int {
	return h.findAttr("id", L)
}

func (h *httpScrapeModule) findAttrByTag(L *lua.LState) int {
	return h.findAttr("tag", L)
}

func (h *httpScrapeModule) findAttrsByClass(L *lua.LState) int {
	return h.findAttrs("class", L)
}

func (h *httpScrapeModule) findAttrsById(L *lua.LState) int {
	return h.findAttrs("id", L)
}

func (h *httpScrapeModule) findAttrsByTag(L *lua.LState) int {
	return h.findAttrs("tag", L)
}

func (h *httpScrapeModule) findTextByClass(L *lua.LState) int {
	return h.findText("class", L)
}

func (h *httpScrapeModule) findTextById(L *lua.LState) int {
	return h.findText("id", L)
}

func (h *httpScrapeModule) findTextByTag(L *lua.LState) int {
	return h.findText("tag", L)
}
