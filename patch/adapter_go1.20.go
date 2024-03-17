//go:build go1.20 && !go1.21
// +build go1.20,!go1.21

package patch

import (
	"cmd/compile/internal/base"
	"cmd/compile/internal/ir"
	"cmd/compile/internal/reflectdata"
	"cmd/compile/internal/types"
)

const genericTrapNeedsWorkaround = false

// different with go1.20
func NewSignature(pkg *types.Pkg, recv *types.Field, tparams, params, results []*types.Field) *types.Type {
	return types.NewSignature(pkg, recv, tparams, params, results)
}

func SetConvTypeWordPtr(conv *ir.ConvExpr, t *types.Type) {
	conv.TypeWord = reflectdata.TypePtrAt(base.Pos, types.NewPtr(t))
}

func getFuncResultsType(funcType *types.Type) *types.Type {
	panic("getFuncResultsType should not be called above go1.19")
}
