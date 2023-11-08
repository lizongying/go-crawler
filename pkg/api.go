package pkg

type ReqJobStart struct {
	Name    string  `json:"name"`
	Func    string  `json:"func"`
	Args    string  `json:"args,omitempty"`
	Mode    JobMode `json:"mode,omitempty"`
	Spec    string  `json:"spec,omitempty"`
	Timeout uint32  `json:"timeout,omitempty"` // second
}
