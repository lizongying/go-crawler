package main

type ExtraDetail struct {
	Keyword string
	ItemId  string
}

type DataWord struct {
	Id      string `bson:"_id" json:"id"`
	Keyword string `json:"keyword"`
	Content string `json:"content"`
}
