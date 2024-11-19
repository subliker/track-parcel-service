package session

type userSession struct {
	state interface{}
}

func (s *userSession) State() interface{} {
	return s.state
}

func (s *userSession) SetState(state interface{}) {
	s.state = state
}

func (s *userSession) ClearState() {
	s.state = nil
}
