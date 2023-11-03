package pkg

type ReqJobStart struct {
	Name    string       `json:"name"`
	Func    string       `json:"func"`
	Args    string       `json:"args,omitempty"`
	Mode    ScheduleMode `json:"mode,omitempty"`
	Spec    string       `json:"spec,omitempty"`
	Timeout uint32       `json:"timeout,omitempty"` // second
}

type ReqJobStop struct {
	Id string `json:"id"` // uuid
}
