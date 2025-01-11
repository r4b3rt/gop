/*
 * Copyright (c) 2025 The GoPlus Authors (goplus.org). All rights reserved.
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

package formatutil

import (
	"log"
	"os"
	"path"
	"strings"
	"testing"

	"github.com/goplus/gop/scanner"
	"github.com/goplus/gop/token"
)

func doSplitStmts(src []byte) (ret []string) {
	fset := token.NewFileSet()
	base := fset.Base()
	f := fset.AddFile("", base, len(src))

	var s scanner.Scanner
	s.Init(f, src, nil, scanner.ScanComments)
	stmts := splitStmts(&s)

	ret = make([]string, len(stmts))
	for i, stmt := range stmts {
		ret[i] = stmtKind(stmt)
	}
	return
}

func stmtKind(s aStmt) string {
	tok, at := s.tok()
	if tok == token.FUNC {
		if isFuncDecl(s.words[at+1:]) {
			return "FUNC"
		}
		return "FNCALL"
	}
	return tok.String()
}

func testFrom(t *testing.T, pkgDir, sel string) {
	if sel != "" && !strings.Contains(pkgDir, sel) {
		return
	}
	t.Helper()
	log.Println("Parsing", pkgDir)
	in, err := os.ReadFile(pkgDir + "/in.data")
	if err != nil {
		t.Fatal("Parsing", pkgDir, "-", err)
	}
	out, err := os.ReadFile(pkgDir + "/out.expect")
	if err != nil {
		t.Fatal("Parsing", pkgDir, "-", err)
	}
	ret := strings.Join(doSplitStmts(in), "\n") + "\n"
	if ret != string(out) {
		t.Fatal("Parsing", pkgDir, "- failed:\n"+ret)
	}
}

func testFromDir(t *testing.T, sel, relDir string) {
	dir, err := os.Getwd()
	if err != nil {
		t.Fatal("Getwd failed:", err)
	}
	dir = path.Join(dir, relDir)
	fis, err := os.ReadDir(dir)
	if err != nil {
		t.Fatal("ReadDir failed:", err)
	}
	for _, fi := range fis {
		name := fi.Name()
		if strings.HasPrefix(name, "_") {
			continue
		}
		t.Run(name, func(t *testing.T) {
			testFrom(t, dir+"/"+name, sel)
		})
	}
}

func TestSplitStmts(t *testing.T) {
	testFromDir(t, "", "./_testdata")
}
