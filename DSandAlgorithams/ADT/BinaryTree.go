/* 
Binary Trees are the Trees which can have a maximum of two leaves for each node. 
example:
              1
        2  			3
		4		5	6		7     


Properties: 
1. maximum posible elements for n internalNodes is (2^n -1) .
2. minimum possible elements for n internalNodes is Log base 2 (n) + 1. => say as l2n plus 1
3. minimum possible level for n leafNodes is Log base 2 (N + 1)  => say as l2  
4. number of leaf nodes is always greater than internalNodes.  
    leafNodes = internalNodes + 1

Types: 
1. Full Binary Tree - each node must have either 0 or 2 elements. 
2. Complete Binary Tree - last internal node is completely filled with no elements left for further operations. 
3. Perfect Binary Tree - It satisfies Property[1]. All nodes have exactly two elements. 
4. Balanced Binary Tree - Trees having n levels statisfies O(Log n). They also provide O(Log n) for search operations.  
   performance wise, Balanced Binary Tree is prefered. 
5. Pathological Tree - All the Internal node has exactly only one element. It looks and acts similiar to linked list.  


Calcualtions: 
  if total Leaves = 10 then
   minimum possible levels : Log base (11) = 3.45  = ~ 3 nodes. 

  if 3 internalNodes then 
  maximum possible leaves nodes : 2^3 -1 = 8 -1 = 7 elements 
  minimum possible leaves nodes : Log base 2 (n) + 1 = 2.5  = ~3 elements. 

                        		1                                  level #1
						2				3						   level #2
				4			5		6			7           	   level #3	
			8		9	10							
											max possible
											leaf node


 */
package main

import (
	"fmt"
	"os"
	"io"
)

type BinaryNode struct {
	left  *BinaryNode
	right *BinaryNode
	data  int64
}

type BinaryTree struct {
	root *BinaryNode
}

func (t *BinaryTree) insert(data int64) *BinaryTree {
	if t.root == nil {
		t.root = &BinaryNode{data: data, left: nil, right: nil}
	} else {
		t.root.insert(data)
	}
	return t
}

func (n *BinaryNode) insert(data int64) {
	if n == nil {
		return
	} else if data <= n.data {
		if n.left == nil {
			n.left = &BinaryNode{data: data, left: nil, right: nil}
		} else {
			n.left.insert(data)
		}
	} else {
		if n.right == nil {
			n.right = &BinaryNode{data: data, left: nil, right: nil}
		} else {
			n.right.insert(data)
		}
	}
}

func print(w io.Writer, node *BinaryNode, ns int, ch rune) {
	if node == nil {
		return
	}

	for i := 0; i < ns; i++ {
		fmt.Fprint(w, " ")
	}
	fmt.Fprintf(w, "%c:%v\n", ch, node.data)
	print(w, node.left, ns+2, 'L')
	print(w, node.right, ns+2, 'R')
}

func main() {
	tree := &BinaryTree{}
	tree.insert(100).
		insert(-20).
		insert(-50).
		insert(-15).
		insert(-60).
		insert(50).
		insert(60).
		insert(55).
		insert(85).
		insert(15).
		insert(5).
		insert(-10)
	print(os.Stdout, tree.root, 0, 'M')
}

