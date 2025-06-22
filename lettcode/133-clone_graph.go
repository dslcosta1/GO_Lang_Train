/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

var visited map[int]*Node

func cloneNode(node *Node) *Node {
    newNode := &Node{
        Val: node.Val,
        Neighbors: []*Node{},
    }

    visited[node.Val] = newNode

    for _, neighbor := range node.Neighbors {
        if newNeighbor, ok := visited[neighbor.Val]; ok {
            newNode.Neighbors = append(newNode.Neighbors, newNeighbor)
        } else {
            newNode.Neighbors = append(newNode.Neighbors, cloneNode(neighbor))
        }
    }

    return newNode
}

func cloneGraph(node *Node) *Node {
    if (node == nil) {
        return nil
    }
    visited = nil
    visited = make(map[int]*Node)

    return cloneNode(node)
}
