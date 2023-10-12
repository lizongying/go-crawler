package main

type ExtraDetail struct {
	Id int
}

type DataWord struct {
	Id      int    `bson:"_id" json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
