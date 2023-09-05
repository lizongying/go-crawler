package response

import (
	"github.com/lizongying/go-css/css"
	"github.com/lizongying/go-xpath/xpath"
	"github.com/tidwall/gjson"
	"reflect"
)

type ParserType uint8

const (
	ParserUnknown ParserType = iota
	ParserJson
	ParserXpath
	ParserCss
	ParserRe
)

type Parser struct {
	rootType  ParserType
	rootPath  string
	rootJson  gjson.Result
	rootXpath []*xpath.Selector
	rootCss   []*css.Selector
	rootRaw   string
	leafType  ParserType
}

func (p *Parser) ParsingRoot(tag reflect.StructTag) {
	path := ""
	path = tag.Get("_json")
	if path != "" {
		p.rootType = ParserJson
		p.rootPath = path
		return
	}
	path = tag.Get("_xpath")
	if path != "" {
		p.rootType = ParserXpath
		p.rootPath = path
		return
	}
	path = tag.Get("_css")
	if path != "" {
		p.rootType = ParserCss
		p.rootPath = path
		return
	}
	path = tag.Get("_re")
	if path != "" {
		p.rootType = ParserRe
		p.rootPath = path
		return
	}
}

func (p *Parser) ParsingLeaf(tag reflect.StructTag) {
	var leafType ParserType
	path := ""
	path = tag.Get("_json")
	if path != "" {
		leafType = ParserJson
	}
	path = tag.Get("_xpath")
	if path != "" {
		leafType = ParserXpath
	}
	path = tag.Get("_css")
	if path != "" {
		leafType = ParserCss
	}
	path = tag.Get("_re")
	if path != "" {
		leafType = ParserRe
	}
	if leafType == ParserUnknown {
		return
	}

	switch p.rootType {
	case ParserJson:
	}
}
