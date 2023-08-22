package pkg

type ReqSpiderStart struct {
	TaskId  string
	Timeout int //second
	Name    string
	Func    string
	Args    string
}

type ReqSpiderStop struct {
	TaskId  string
	Timeout int //second
}
