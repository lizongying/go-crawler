package pkg

import (
	"sync/atomic"
)

type State struct {
	set          bool
	count        atomic.Int32
	setFn        []func()
	zeroFn       []func()
	setAndZeroFn []func()
}

func NewState() *State {
	return &State{}
}
func (s *State) RegisterSetFn(fns ...func()) {
	s.setFn = append(s.setFn, fns...)
}
func (s *State) RegisterZeroFn(fns ...func()) {
	s.zeroFn = append(s.zeroFn, fns...)
}
func (s *State) RegisterSetAndZeroFn(fns ...func()) {
	s.setAndZeroFn = append(s.setAndZeroFn, fns...)
}
func (s *State) Set() {
	s.set = true
	for _, v := range s.setFn {
		v()
	}
}
func (s *State) IsSet() bool {
	return s.set
}
func (s *State) In() {
	s.count.Add(1)
}
func (s *State) Out() {
	s.count.Add(-1)
	for _, v := range s.zeroFn {
		v()
	}
	for _, v := range s.setAndZeroFn {
		v()
	}
}
func (s *State) Zero() bool {
	return s.count.Load() <= 0
}
func (s *State) SetAndZero() bool {
	return s.IsSet() && s.Zero()
}

type MultiState struct {
	states       []*State
	setFn        []func()
	zeroFn       []func()
	setAndZeroFn []func()
}

func NewMultiState(states ...*State) *MultiState {
	return &MultiState{states: states}
}
func (s *MultiState) RegisterSetFn(fns ...func()) {
	s.setFn = append(s.setFn, fns...)
	for _, v := range s.states {
		fn := func() {
			if s.IsSet() {
				for _, v2 := range s.setFn {
					v2()
				}
			}
		}
		v.RegisterSetFn(fn)
	}
}
func (s *MultiState) RegisterZeroFn(fns ...func()) {
	s.zeroFn = append(s.zeroFn, fns...)
	for _, v := range s.states {
		fn := func() {
			if s.Zero() {
				for _, v2 := range s.zeroFn {
					v2()
				}
			}
		}
		v.RegisterZeroFn(fn)
	}
}
func (s *MultiState) RegisterSetAndZeroFn(fns ...func()) {
	s.setAndZeroFn = append(s.setAndZeroFn, fns...)
	for _, v := range s.states {
		fn := func() {
			if s.SetAndZero() {
				for _, v2 := range s.setAndZeroFn {
					v2()
				}
			}
		}
		v.RegisterSetAndZeroFn(fn)
	}
}
func (s *MultiState) IsSet() bool {
	for _, v := range s.states {
		if !v.IsSet() {
			return false
		}
	}
	return true
}
func (s *MultiState) Zero() bool {
	for _, v := range s.states {
		if !v.Zero() {
			return false
		}
	}
	return true
}
func (s *MultiState) SetAndZero() bool {
	return s.IsSet() && s.Zero()
}
