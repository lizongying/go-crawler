package main

type ExtraOk struct {
	Count int
}

type DataOk struct {
	Id    string `column:"id"  bson:"_id" json:"id"`
	Count int    `column:"Count"  bson:"count" json:"count"`
	A     uint   `column:"a"`
	B     uint32 `column:"b"`
	C     string `column:"c"`
	D     string `column:"d"`
}
