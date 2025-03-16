package Clone_Graph

import (
	"reflect"
	"sort"
	"testing"
)

// TestCloneGraph tests the cloneGraph function using several test cases.
func TestCloneGraph(t *testing.T) {
	tests := []struct {
		name    string
		adjList [][]int
	}{
		{
			name:    "example1",
			adjList: [][]int{{2, 4}, {1, 3}, {2, 4}, {1, 3}},
		},
		{
			name:    "example2",
			adjList: [][]int{{}},
		},
		{
			name:    "example3",
			adjList: [][]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Build the original graph from the adjacency list.
			original := buildGraph(tt.adjList)
			// Clone the graph.
			clone := cloneGraph(original)

			// For an empty graph, expect nil.
			if original == nil && clone != nil {
				t.Errorf("Expected nil clone for empty graph, got non-nil")
			}
			if original != nil && clone == nil {
				t.Errorf("Expected non-nil clone for non-empty graph, got nil")
			}

			// If the graph is not empty, verify that the clone is a deep copy.
			if original != nil && clone != nil {
				origMap := getNodeMap(original)
				cloneMap := getNodeMap(clone)
				// The number of nodes in both graphs should match.
				if len(origMap) != len(cloneMap) {
					t.Errorf("Expected same number of nodes, got original %d, clone %d", len(origMap), len(cloneMap))
				}
				// For each node, verify the structure and that it is a deep copy.
				for val, oNode := range origMap {
					cNode, ok := cloneMap[val]
					if !ok {
						t.Errorf("Clone graph missing node with value %d", val)
						continue
					}
					// Check that the pointers are different.
					if oNode == cNode {
						t.Errorf("Node with value %d is not deep copied (same pointer)", val)
					}
					// Compare neighbor values.
					oNeighVals := getNeighborValues(oNode)
					cNeighVals := getNeighborValues(cNode)
					sort.Ints(oNeighVals)
					sort.Ints(cNeighVals)
					if !reflect.DeepEqual(oNeighVals, cNeighVals) {
						t.Errorf("Mismatch in neighbors for node %d: expected %v, got %v", val, oNeighVals, cNeighVals)
					}
				}
			}
		})
	}
}

// getNeighborValues returns the list of neighbor values for a given node.
func getNeighborValues(n *Node) []int {
	values := []int{}
	for _, neigh := range n.Neighbors {
		values = append(values, neigh.Val)
	}
	return values
}

// getNodeMap performs a BFS on the graph starting from root and maps each node value to the node pointer.
func getNodeMap(root *Node) map[int]*Node {
	nodeMap := make(map[int]*Node)
	if root == nil {
		return nodeMap
	}
	queue := []*Node{root}
	visited := make(map[int]bool)
	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]
		if visited[n.Val] {
			continue
		}
		visited[n.Val] = true
		nodeMap[n.Val] = n
		for _, neigh := range n.Neighbors {
			if !visited[neigh.Val] {
				queue = append(queue, neigh)
			}
		}
	}
	return nodeMap
}

// buildGraph converts an adjacency list to a graph and returns the first node (node with val = 1).
func buildGraph(adjList [][]int) *Node {
	if len(adjList) == 0 {
		return nil
	}
	// Create nodes for all indices (1-indexed)
	nodes := make(map[int]*Node)
	for i := 1; i <= len(adjList); i++ {
		nodes[i] = &Node{Val: i}
	}
	// Connect neighbors based on the given adjacency list.
	for i, neighbors := range adjList {
		for _, neighVal := range neighbors {
			nodes[i+1].Neighbors = append(nodes[i+1].Neighbors, nodes[neighVal])
		}
	}

	return nodes[1]
}
