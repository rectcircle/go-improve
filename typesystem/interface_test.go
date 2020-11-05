package typesystem

import (
	"testing"
)

func TestNonPtrOrPtrReceiver(t *testing.T) {
	NonPtrOrPtrReceiver()
}

func TestInterfaceNil(t *testing.T) {
	InterfaceNil()
}

func TestInterfaceNested(t *testing.T) {
	InterfaceNested()
}

func TestStructNested(t *testing.T) {
	StructNested()
}
