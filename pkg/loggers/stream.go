package loggers

type Ch struct {
	Name    string
	Channel chan []byte
}

type Stream struct {
	channels   map[string]chan []byte
	register   chan Ch
	unregister chan string
	channel    chan []byte
}

func NewStream() (s *Stream) {
	s = &Stream{
		channels:   make(map[string]chan []byte),
		register:   make(chan Ch),
		unregister: make(chan string),
		channel:    make(chan []byte, 10),
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
				s.channels[ch.Name] = ch.Channel
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

func (s *Stream) Register(name string, channel chan []byte) {
	s.register <- Ch{
		Name:    name,
		Channel: channel,
	}
}

func (s *Stream) Unregister(name string) {
	s.unregister <- name
}
