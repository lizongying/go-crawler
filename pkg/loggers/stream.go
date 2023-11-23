package loggers

type Ch struct {
	Id      uint32
	Channel chan []byte
}

type Stream struct {
	channels   map[uint32]chan []byte
	register   chan Ch
	unregister chan uint32
	channel    chan []byte
}

func NewStream() (s *Stream) {
	s = &Stream{
		channels:   make(map[uint32]chan []byte),
		register:   make(chan Ch),
		unregister: make(chan uint32),
		channel:    make(chan []byte, 8),
	}
	go func() {
		for msg := range s.channel {
			for _, ch := range s.channels {
				select {
				case ch <- msg:
				default:
					<-ch
					ch <- msg
				}
			}
		}
	}()
	go func() {
		for {
			select {
			case ch := <-s.register:
				s.channels[ch.Id] = ch.Channel
			case chName := <-s.unregister:
				delete(s.channels, chName)
			}
		}
	}()

	return
}

func (s *Stream) Write(p []byte) (n int, err error) {
	select {
	case s.channel <- p:
	default:
		<-s.channel
		s.channel <- p
	}
	return
}

func (s *Stream) Register(id uint32, channel chan []byte) {
	s.register <- Ch{
		Id:      id,
		Channel: channel,
	}
}

func (s *Stream) Unregister(id uint32) {
	s.unregister <- id
}
