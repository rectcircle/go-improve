package concurrent

import (
	"testing"
)

func TestSyncOnce(t *testing.T) {
	SyncOnce()
}

func TestSyncWaitGroup(t *testing.T) {
	SyncWaitGroup()
}

func TestSyncPool(t *testing.T) {
	SyncPool()
}
