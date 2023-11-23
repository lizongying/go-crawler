package pkg

import (
	"sync"
	"testing"
)

func TestState_RegisterIsReady(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)

	state := NewState("")
	state.RegisterIsReady(func() {
		wg.Done()
	})

	go func() {
		state.BeReady()
	}()

	wg.Wait()
}

func TestState_RegisterIsZero(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)

	state := NewState("")
	state.RegisterIsZero(func() {
		wg.Done()
	})

	go func() {
		state.In()
		state.Out()
	}()

	wg.Wait()
}

func TestState_RegisterIsReadyAndIsZero(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)

	state := NewState("")
	state.RegisterIsReadyAndIsZero(func() {
		wg.Done()
	})

	go func() {
		state.BeReady()
		state.In()
		state.Out()
	}()

	wg.Wait()
}

func TestState_BeReady(t *testing.T) {
	state := NewState("")
	state.BeReady()

	if !state.IsReady() {
		t.Errorf("Expected state to be ready")
	}
}

func TestState_In(t *testing.T) {
	state := NewState("")
	state.In()

	if !state.IsReady() {
		t.Errorf("Expected state to be ready")
	}

	if state.IsZero() {
		t.Errorf("Expected state not to be zero")
	}
}

func TestState_Out(t *testing.T) {
	state := NewState("")
	state.In()
	state.Out()

	if !state.IsReady() {
		t.Errorf("Expected state to be ready")
	}

	if !state.IsZero() {
		t.Errorf("Expected state to be zero")
	}
}

func TestState_IsReady(t *testing.T) {
	state := NewState("")

	if state.IsReady() {
		t.Errorf("Expected state to not be ready")
	}

	state.BeReady()

	if !state.IsReady() {
		t.Errorf("Expected state to be ready")
	}
}

func TestState_IsZero(t *testing.T) {
	state := NewState("")

	if !state.IsZero() {
		t.Errorf("Expected state to be zero")
	}

	state.In()

	if state.IsZero() {
		t.Errorf("Expected state to not be zero")
	}

	state.Out()

	if !state.IsZero() {
		t.Errorf("Expected state to be zero")
	}
}

func TestMultiState_RegisterIsReady(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)

	state1 := NewState("")
	state2 := NewState("")

	multiState := NewMultiState(state1, state2)
	multiState.RegisterIsReady(func() {
		wg.Done()
	})

	go func() {
		state1.BeReady()
		state2.BeReady()
	}()

	wg.Wait()
}

func TestMultiState_RegisterIsZero(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)

	state1 := NewState("")
	state2 := NewState("")

	multiState := NewMultiState(state1, state2)
	multiState.RegisterIsZero(func() {
		wg.Done()
	})

	go func() {
		state1.In()
		state2.In()
		state1.Out()
		state2.Out()
	}()

	wg.Wait()
}

func TestMultiState_RegisterIsReadyAndIsZero(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)

	state1 := NewState("")
	state2 := NewState("")

	multiState := NewMultiState(state1, state2)
	multiState.RegisterIsReadyAndIsZero(func() {
		wg.Done()
	})

	go func() {
		state1.BeReady()
		state1.In()
		state1.Out()
		state2.BeReady()
		state2.In()
		state2.Out()
	}()

	wg.Wait()
}

func TestMultiState_IsReady(t *testing.T) {
	state1 := NewState("")
	state2 := NewState("")

	multiState := NewMultiState(state1, state2)

	if multiState.IsReady() {
		t.Errorf("Expected multiState to not be ready")
	}

	state1.BeReady()

	if multiState.IsReady() {
		t.Errorf("Expected multiState to not be ready")
	}

	state2.BeReady()

	if !multiState.IsReady() {
		t.Errorf("Expected multiState to be ready")
	}
}

func TestMultiState_IsZero(t *testing.T) {
	state1 := NewState("")
	state2 := NewState("")

	multiState := NewMultiState(state1, state2)

	if !multiState.IsZero() {
		t.Errorf("Expected multiState to be zero")
	}

	state1.In()

	if multiState.IsZero() {
		t.Errorf("Expected multiState to not be zero")
	}

	state1.Out()

	if !multiState.IsZero() {
		t.Errorf("Expected multiState to be zero")
	}

	state2.In()

	if multiState.IsZero() {
		t.Errorf("Expected multiState to not be zero")
	}

	state2.Out()

	if !multiState.IsZero() {
		t.Errorf("Expected multiState to be zero")
	}
}

func TestMultiState_IsReadyAndIsZero(t *testing.T) {
	state1 := NewState("")
	state2 := NewState("")

	multiState := NewMultiState(state1, state2)

	if multiState.IsReadyAndIsZero() {
		t.Errorf("Expected multiState to not be ready and zero")
	}

	state1.BeReady()

	if multiState.IsReadyAndIsZero() {
		t.Errorf("Expected multiState to not be ready and zero")
	}

	state2.BeReady()

	if !multiState.IsReadyAndIsZero() {
		t.Errorf("Expected multiState to be ready and zero")
	}

	state1.In()

	if multiState.IsReadyAndIsZero() {
		t.Errorf("Expected multiState to not be ready and zero")
	}

	state1.Out()

	if !multiState.IsReadyAndIsZero() {
		t.Errorf("Expected multiState to be ready and zero")
	}

	state2.In()

	if multiState.IsReadyAndIsZero() {
		t.Errorf("Expected multiState to not be ready and zero")
	}

	state2.Out()

	if !multiState.IsReadyAndIsZero() {
		t.Errorf("Expected multiState to be ready and zero")
	}
}
