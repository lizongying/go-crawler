package test_parse_spider

type DataParse struct {
	Data struct {
		B uint8 `_json:"b"`
		C int   `_json:"c.1"`
	} `_re:"a = ([^<]+)"`
}
