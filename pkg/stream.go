package pkg

type Stream interface {
	Register(name string, channel chan []byte)
	Unregister(name string)
}
