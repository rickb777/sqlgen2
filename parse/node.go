package parse

import (
	"bytes"
	"fmt"
	"io"
)

type Node struct {
	Pkg  string // source code package.
	Name string // identifier name.
	Kind Kind   // source code kind.
	Type string // name of source code type.
	Tags *Tag

	Parent *Node
	Nodes  []*Node
}

func (n *Node) appendNode(node *Node) {
	node.Parent = n
	n.Nodes = append(n.Nodes, node)
}

// Walk traverses the node tree, invoking the callback
// function for each node that is traversed.
func (n *Node) Walk(fn func(*Node)) {
	for _, node := range n.Nodes {
		fn(node)
		node.Walk(fn)
	}
}

// WalkRev traverses the tree in reverse order, invoking
// the callback function for each parent node until
// the root node is reached.
func (n *Node) walkRev(fn func(*Node)) {
	if n.Parent != nil {
		n.Parent.walkRev(fn)
	}
	fn(n) // this was previously inside the if block
}

// Leaves returns a flattened list of all leaf nodes in the Tree.
func (n *Node) Leaves() []*Node {
	var nodes []*Node
	n.Walk(func(node *Node) {
		if len(node.Nodes) == 0 {
			nodes = append(nodes, node)
		}
	})
	return nodes
}

// Path returns the route from the node to the root of the Tree.
func (n *Node) Path() []*Node {
	var nodes []*Node
	n.walkRev(func(node *Node) {
		nodes = append(nodes, node)
	})
	return nodes
}

func (n *Node) String() string {
	buf := &bytes.Buffer{}
	n.indented(buf, "")
	return buf.String()
}

func (n *Node) indented(w io.Writer, indent string) {
	if len(n.Pkg) > 0 {
		fmt.Fprintf(w, "%s%s %s.%s %d\n", indent, n.Name, n.Pkg, n.Type, n.Kind)
	} else {
		fmt.Fprintf(w, "%s%s %s %d\n", indent, n.Name, n.Type, n.Kind)
	}
	deeper := indent + "  "
	for _, c := range n.Nodes {
		c.indented(w, deeper)
	}
}
