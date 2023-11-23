package pkg

type Stream interface {
	Register(id uint32, channel chan []byte)
	Unregister(id uint32)
}
