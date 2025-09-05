package pkg

import (
	"sync/atomic"
)

// State represents a state with ready and count properties.
type State struct {
	name               string
	isReady            atomic.Bool
	count              atomic.Uint32
	fnIsReady          []func()
	fnIsZero           []func()
	fnIsReadyAndIsZero []func()
}

// NewState creates a new State instance.
func NewState(name string) *State {
	return &State{name: name}
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
	if s.isReady.CompareAndSwap(false, true) {
		for _, v := range s.fnIsReady {
			v()
		}
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
	return s.isReady.Load()
}

// IsZero checks if the count of the state is zero.
func (s *State) IsZero() bool {
	return s.count.Load() == 0
}

// IsReadyAndIsZero checks if the state is ready and zero.
func (s *State) IsReadyAndIsZero() bool {
	return s.IsReady() && s.IsZero()
}

func (s *State) Clear() {
	s.isReady.Store(false)
	s.count.Store(0)
}

func (s *State) Count() uint32 {
	return s.count.Load()
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

// IsReady checks if all states in the MultiState are ready.
func (m *MultiState) IsReady() bool {
	for _, v := range m.states {
		if !v.IsReady() {
			return false
		}
	}
	return len(m.states) > 1
}

// IsZero checks if all states in the MultiState are zero.
func (m *MultiState) IsZero() bool {
	for _, v := range m.states {
		if !v.IsZero() {
			return false
		}
	}
	return true
}

// RegisterIsReady registers functions to be called when any state in the MultiState is ready.
func (m *MultiState) RegisterIsReady(fns ...func()) {
	m.fnIsReady = append(m.fnIsReady, fns...)
	for _, v := range m.states {
		fn := func() {
			if m.IsReady() {
				for _, v2 := range m.fnIsReady {
					v2()
				}
			}
		}
		v.RegisterIsReady(fn)
	}
}

// RegisterIsZero registers functions to be called when any state in the MultiState is zero.
func (m *MultiState) RegisterIsZero(fns ...func()) {
	m.fnIsZero = append(m.fnIsZero, fns...)
	for _, v := range m.states {
		fn := func() {
			if m.IsZero() {
				for _, v2 := range m.fnIsZero {
					v2()
				}
			}
		}
		v.RegisterIsZero(fn)
	}
}

// RegisterIsReadyAndIsZero registers functions to be called when any state in the MultiState is ready and zero.
func (m *MultiState) RegisterIsReadyAndIsZero(fns ...func()) {
	m.fnIsReadyAndIsZero = append(m.fnIsReadyAndIsZero, fns...)
	for _, v := range m.states {
		fn := func() {
			if m.IsReadyAndIsZero() {
				for _, v2 := range m.fnIsReadyAndIsZero {
					v2()
				}
			}
		}
		v.RegisterIsReadyAndIsZero(fn)
	}
}

// IsReadyAndIsZero checks if all states in the MultiState are ready and zero.
func (m *MultiState) IsReadyAndIsZero() bool {
	return m.IsReady() && m.IsZero()
}

func (m *MultiState) Clear() {
	for _, v := range m.states {
		v.Clear()
	}
	return
}
