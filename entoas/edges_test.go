// Copyright 2019-present Facebook
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package entoas

import (
	"os"
	"path/filepath"
	"testing"

	"entgo.io/contrib/entoas/serialization"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"github.com/stretchr/testify/require"
)

func TestWalk_CycleDepth(t *testing.T) {
	var w walk
	require.Equal(t, uint(0), w.cycleDepth())

	w = walk{"a"}
	require.Equal(t, uint(1), w.cycleDepth())

	w = walk{"a", "b"}
	require.Equal(t, uint(1), w.cycleDepth())

	w = walk{"a", "a"}
	require.Equal(t, uint(2), w.cycleDepth())

	w = walk{"a", "b", "b"}
	require.Equal(t, uint(2), w.cycleDepth())

	w = walk{"a", "b", "b", "c"}
	require.Equal(t, uint(1), w.cycleDepth())

	w = walk{"a", "b", "b", "a"}
	require.Equal(t, uint(2), w.cycleDepth())

	w = walk{"a", "a", "b", "a"}
	require.Equal(t, uint(3), w.cycleDepth())
}

func TestWalk_Push(t *testing.T) {
	var w walk
	require.Equal(t, walk(nil), w)

	w.push("a")
	require.Equal(t, walk{"a"}, w)

	w.push("b")
	require.Equal(t, walk{"a", "b"}, w)
}

func TestWalk_Pop(t *testing.T) {
	w := walk{"a", "b", "c"}

	w.pop()
	require.Equal(t, walk{"a", "b"}, w)

	w.pop()
	require.Equal(t, walk{"a"}, w)

	w.pop()
	require.Equal(t, walk{}, w)
}

func TestEdgeTree(t *testing.T) {
	// Load a graph.
	wd, err := os.Getwd()
	require.NoError(t, err)
	g, err := entc.LoadGraph(filepath.Join(wd, "internal", "schema"), &gen.Config{})
	require.NoError(t, err)
	// Extract the Edges for a read operation on the Pet entity.
	var (
		p *gen.Type
		o *gen.Edge
	)
	for _, n := range g.Nodes {
		if n.Name == "Pet" {
			p = n
			for _, e := range n.Edges {
				if e.Name == "owner" {
					o = e
					break
				}
			}
			break
		}
	}
	es, err := EdgeTree(p, serialization.Groups{"test:edge"})
	require.NoError(t, err)
	require.Equal(t, Edges{{Edge: o}}, es)
}