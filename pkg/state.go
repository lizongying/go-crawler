package pkg

import (
	"sync/atomic"
)

// State represents a state with ready and count properties.
type State struct {
	isReady            bool
	count              atomic.Uint32
	fnIsReady          []func()
	fnIsZero           []func()
	fnIsReadyAndIsZero []func()
}

// NewState creates a new State instance.
func NewState() *State {
	return &State{}
}

// RegisterIsReady registers functions to be called when the state is ready.
func (s *State) RegisterIsReady(fns ...func()) {
	s.fnIsReady = append(s.fnIsReady, fns...)
}

// RegisterIsZero registers functions to be called when the state is zero.
func (s *State) RegisterIsZero(fns ...func()) {
	s.fnIsZero = append(s.fnIsZero, fns...)
}

// RegisterIsReadyAndIsZero registers functions to be called when the state is ready and zero.
func (s *State) RegisterIsReadyAndIsZero(fns ...func()) {
	s.fnIsReadyAndIsZero = append(s.fnIsReadyAndIsZero, fns...)
}

// BeReady sets the state as ready and calls the registered functions for ready state.
func (s *State) BeReady() {
	s.isReady = true
	for _, v := range s.fnIsReady {
		v()
	}
}

// In increments the count of the state.
func (s *State) In() {
	if !s.IsReady() {
		s.BeReady()
	}
	s.count.Add(1)
}

// Out decrements the count of the state and calls the registered functions for zero state.
// If the state is ready, it also calls the registered functions for ready and zero state.
func (s *State) Out() {
	s.count.Add(^uint32(0))
	if s.IsZero() {
		for _, v := range s.fnIsZero {
			v()
		}
		if s.IsReady() {
			for _, v := range s.fnIsReadyAndIsZero {
				v()
			}
		}
	}
}

// IsReady checks if the state is ready.
func (s *State) IsReady() bool {
	return s.isReady
}

// IsZero checks if the count of the state is zero.
func (s *State) IsZero() bool {
	return s.count.Load() == 0
}

// isReadyAndIsZero checks if the state is ready and zero.
func (s *State) isReadyAndIsZero() bool {
	return s.IsReady() && s.IsZero()
}

// MultiState represents a collection of states.
type MultiState struct {
	states             []*State
	fnIsReady          []func()
	fnIsZero           []func()
	fnIsReadyAndIsZero []func()
}

// NewMultiState creates a new MultiState instance with the provided states.
func NewMultiState(states ...*State) *MultiState {
	return &MultiState{states: states}
}

// RegisterIsReady registers functions to be called when any state in the MultiState is ready.
func (s *MultiState) RegisterIsReady(fns ...func()) {
	s.fnIsReady = append(s.fnIsReady, fns...)
	for _, v := range s.states {
		fn := func() {
			if s.IsReady() {
				for _, v2 := range s.fnIsReady {
					v2()
				}
			}
		}
		v.RegisterIsReady(fn)
	}
}

// RegisterIsZero registers functions to be called when any state in the MultiState is zero.
func (s *MultiState) RegisterIsZero(fns ...func()) {
	s.fnIsZero = append(s.fnIsZero, fns...)
	for _, v := range s.states {
		fn := func() {
			if s.IsZero() {
				for _, v2 := range s.fnIsZero {
					v2()
				}
			}
		}
		v.RegisterIsZero(fn)
	}
}

// RegisterIsReadyAndIsZero registers functions to be called when any state in the MultiState is ready and zero.
func (s *MultiState) RegisterIsReadyAndIsZero(fns ...func()) {
	s.fnIsReadyAndIsZero = append(s.fnIsReadyAndIsZero, fns...)
	for _, v := range s.states {
		fn := func() {
			if s.isReadyAndIsZero() {
				for _, v2 := range s.fnIsReadyAndIsZero {
					v2()
				}
			}
		}
		v.RegisterIsReadyAndIsZero(fn)
	}
}

// IsReady checks if all states in the MultiState are ready.
func (s *MultiState) IsReady() bool {
	for _, v := range s.states {
		if !v.IsReady() {
			return false
		}
	}
	return len(s.states) > 1
}

// IsZero checks if all states in the MultiState are zero.
func (s *MultiState) IsZero() bool {
	for _, v := range s.states {
		if !v.IsZero() {
			return false
		}
	}
	return true
}

// isReadyAndIsZero checks if all states in the MultiState are ready and zero.
func (s *MultiState) isReadyAndIsZero() bool {
	return s.IsReady() && s.IsZero()
}
