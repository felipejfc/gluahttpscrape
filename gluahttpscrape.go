package gluahttpscrape

import (
	"fmt"

	"github.com/yuin/gopher-lua"
)

type httpScrapeModule struct{}

func NewHttpScrapeModule() *httpScrapeModule {
	return &httpScrapeModule{}
}

func (h *httpScrapeModule) Loader(L *lua.LState) int {
	mod := L.SetFuncs(L.NewTable(), map[string]lua.LGFunction{
		"findByTag": h.findByTag,
		//"findByClass": h.findByClass,
		//"findById":    h.findById,
	})
	L.Push(mod)
	return 1
}

func (h *httpScrapeModule) findByTag(L *lua.LState) int {
	body := L.ToString(1)
	tag := L.ToString(2)
	fmt.Printf("bode: %s, tag: %s", body, tag)
	return 0
}
