package keyword

import (
	"testing"
)

func TestGoroutineWithPanic(t *testing.T) {
	GoroutineWithPanic()
}

func TestRecoverNotInDefer(t *testing.T) {
	RecoverNotInDefer()
}

func TestPanicNested(t *testing.T) {
	PanicNested()
}

func TestRecoverExperiment(t *testing.T) {
	RecoverExperiment()
}
