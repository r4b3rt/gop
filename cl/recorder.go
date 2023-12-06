/*
 * Copyright (c) 2021 The GoPlus Authors (goplus.org). All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package cl

import (
	"go/types"

	"github.com/goplus/gop/ast"
	"github.com/goplus/gop/ast/fromgo"
	"github.com/goplus/gop/cl/internal/typesutil"
	"github.com/goplus/gox"
)

type goxRecorder struct {
	rec Recorder
}

// Member maps identifiers to the objects they denote.
func (p *goxRecorder) Member(id ast.Node, obj types.Object) {
	tv := typesutil.NewTypeAndValueForObject(obj)
	switch v := id.(type) {
	case *ast.SelectorExpr:
		sel := v.Sel
		// TODO: record event for a Go ident
		if _, ok := fromgo.CheckIdent(sel); !ok {
			p.rec.Use(sel, obj)
			p.rec.Type(v, tv)
		}
	case *ast.Ident: // it's in a classfile and impossible converted from Go
		p.rec.Use(v, obj)
		p.rec.Type(v, tv)
	}
}

type typesRecorder struct {
	Recorder
	types map[ast.Expr]types.TypeAndValue
}

func (rec *typesRecorder) Type(expr ast.Expr, tv types.TypeAndValue) {
	rec.types[expr] = tv
	rec.Recorder.Type(expr, tv)
}

func newTypeRecord(rec Recorder) *typesRecorder {
	return &typesRecorder{rec, make(map[ast.Expr]types.TypeAndValue)}
}

func (rec *typesRecorder) recordTypeValue(ctx *blockCtx, expr ast.Expr) {
	e := ctx.cb.Get(-1)
	rec.Type(expr, typesutil.NewTypeAndValueForValue(e.Type, e.CVal))
}

func (rec *typesRecorder) recordTypeVariable(ctx *blockCtx, expr ast.Expr) {
	e := ctx.cb.Get(-1)
	t, _ := gox.DerefType(e.Type)
	rec.Type(expr, typesutil.NewTypeAndValueForVariable(t))
}

func (rec *typesRecorder) recordTypeMapIndex(ctx *blockCtx, expr ast.Expr) {
	e := ctx.cb.Get(-1)
	t, _ := gox.DerefType(e.Type)
	rec.Type(expr, typesutil.NewTypeAndValueForMapIndex(t))
}

func (rec *typesRecorder) indexExpr(ctx *blockCtx, expr *ast.IndexExpr) {
	if tv, ok := rec.types[expr.X]; ok {
		switch tv.Type.(type) {
		case *types.Map:
			rec.recordTypeMapIndex(ctx, expr)
			return
		case *types.Slice:
			rec.recordTypeVariable(ctx, expr)
			return
		}
	}
	switch expr.X.(type) {
	case *ast.CompositeLit:
		rec.recordTypeValue(ctx, expr)
	default:
		rec.recordTypeVariable(ctx, expr)
	}
}

func (rec *typesRecorder) recordCallExpr(ctx *blockCtx, v *ast.CallExpr, fnt types.Type) {
	e := ctx.cb.Get(-1)
	rec.Type(v.Fun, typesutil.NewTypeAndValueForValue(fnt, nil))
	rec.Type(v, typesutil.NewTypeAndValueForCallResult(e.Type, e.CVal))
}

func (rec *typesRecorder) recordCompositeLit(ctx *blockCtx, v *ast.CompositeLit, typ types.Type) {
	rec.Type(v.Type, typesutil.NewTypeAndValueForType(typ))
	rec.Type(v, typesutil.NewTypeAndValueForValue(typ, nil))
}

func (rec *typesRecorder) recordType(ctx *blockCtx, typ ast.Expr, t types.Type) {
	rec.Type(typ, typesutil.NewTypeAndValueForType(t))
}

func (rec *typesRecorder) recordIdent(ctx *blockCtx, ident *ast.Ident, obj types.Object) {
	rec.Use(ident, obj)
	rec.Type(ident, typesutil.NewTypeAndValueForObject(obj))
}

func (rec *typesRecorder) recordExpr(ctx *blockCtx, expr ast.Expr, rhs bool) {
	switch v := expr.(type) {
	case *ast.Ident:
	case *ast.BasicLit:
		rec.recordTypeValue(ctx, v)
	case *ast.CallExpr:
	case *ast.SelectorExpr:
		rec.recordTypeVariable(ctx, v)
	case *ast.BinaryExpr:
		rec.recordTypeValue(ctx, v)
	case *ast.UnaryExpr:
		rec.recordTypeValue(ctx, v)
	case *ast.FuncLit:
	case *ast.CompositeLit:
	case *ast.SliceLit:
	case *ast.RangeExpr:
	case *ast.IndexExpr:
		rec.indexExpr(ctx, v)
	case *ast.IndexListExpr:
	case *ast.SliceExpr:
		rec.recordTypeValue(ctx, v)
	case *ast.StarExpr:
		rec.recordTypeVariable(ctx, v)
	case *ast.ArrayType:
	case *ast.MapType:
	case *ast.StructType:
	case *ast.ChanType:
	case *ast.InterfaceType:
	case *ast.ComprehensionExpr:
	case *ast.TypeAssertExpr:
	case *ast.ParenExpr:
		rec.recordTypeValue(ctx, v)
	case *ast.ErrWrapExpr:
	case *ast.FuncType:
	case *ast.Ellipsis:
	case *ast.KeyValueExpr:
	default:
	}
}