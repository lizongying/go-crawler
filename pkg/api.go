package pkg

type ReqStartSpider struct {
	TaskId  string
	Timeout int //second
	Name    string
	Func    string
	Args    string
}

type ReqStopSpider struct {
	TaskId  string
	Timeout int //second
}
