package trap_set

import "testing"

func __xgo_link_set_trap(trapImpl func(pkgPath string, identityName string, generic bool, pc uintptr, recv interface{}, args []interface{}, results []interface{}) (func(), bool)) {
	panic("failed to link __xgo_link_set_trap")
}

// go run ./cmd/xgo test -v -run TestTrapSet ./test/xgo_test/trap_set
func TestTrapSet(t *testing.T) {
	var haveCalledTrap bool
	__xgo_link_set_trap(func(pkgPath, identityName string, generic bool, pc uintptr, recv interface{}, args, results []interface{}) (func(), bool) {
		haveCalledTrap = true
		return nil, false
	})
	run()

	if !haveCalledTrap {
		t.Fatalf("expect have called trap, actually not called")
	}
}

func run() {
}