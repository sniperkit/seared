// Copyright (C) 2017, Pablo Lalloni <plalloni@gmail.com>.
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// 1. Redistributions of source code must retain the above copyright notice,
//    this list of conditions and the following disclaimer.
//
// 2. Redistributions in binary form must reproduce the above copyright notice,
//    this list of conditions and the following disclaimer in the documentation
//    and/or other materials provided with the distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
// ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE
// LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
// CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF
// SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS
// INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN
// CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
// ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
// POSSIBILITY OF SUCH DAMAGE.

package node

import (
	"fmt"
	"strings"
)

type Kind int

const (
	Terminal Kind = iota
	NonTerminal
)

type Node struct {
	Kind     Kind
	Label    string
	Children []*Node
	Value    string
}

func NewTerminal(value string) *Node {
	return &Node{
		Kind:  Terminal,
		Value: value,
	}
}

func NewNonTerminal(label string, children []*Node) *Node {
	return &Node{
		Kind:     NonTerminal,
		Label:    label,
		Children: children,
	}
}

func (n *Node) Format() string {
	var s string
	switch n.Kind {
	case Terminal:
		s = "\"" + n.Value + "\""
	case NonTerminal:
		ss := []string{}
		for _, child := range n.Children {
			ss = append(ss, child.Format())
		}
		s = "(" + n.Label + " " + strings.Join(ss, " ") + ")"
	default:
		s = fmt.Sprintf("(kind %v node)", n.Kind)
	}
	return s
}
