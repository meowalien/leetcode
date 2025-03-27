package Evaluate_Division

/*
You are given an array of variable pairs equations and an array of real numbers values, where equations[i] = [Ai, Bi] and values[i] represent the equation Ai / Bi = values[i]. Each Ai or Bi is a string that represents a single variable.

You are also given some queries, where queries[j] = [Cj, Dj] represents the jth query where you must find the answer for Cj / Dj = ?.

Return the answers to all queries. If a single answer cannot be determined, return -1.0.

Note: The input is always valid. You may assume that evaluating the queries will not result in division by zero and that there is no contradiction.

Note: The variables that do not occur in the list of equations are undefined, so the answer cannot be determined for them.

Example 1:

Input: equations = [["a","b"],["b","c"]], values = [2.0,3.0], queries = [["a","c"],["b","a"],["a","e"],["a","a"],["x","x"]]
Output: [6.00000,0.50000,-1.00000,1.00000,-1.00000]
Explanation:
Given: a / b = 2.0, b / c = 3.0
queries are: a / c = ?, b / a = ?, a / e = ?, a / a = ?, x / x = ?
return: [6.0, 0.5, -1.0, 1.0, -1.0 ]
note: x is undefined => -1.0
Example 2:

Input: equations = [["a","b"],["b","c"],["bc","cd"]], values = [1.5,2.5,5.0], queries = [["a","c"],["c","b"],["bc","cd"],["cd","bc"]]
Output: [3.75000,0.40000,5.00000,0.20000]
Example 3:

Input: equations = [["a","b"]], values = [0.5], queries = [["a","b"],["b","a"],["a","c"],["x","y"]]
Output: [0.50000,2.00000,-1.00000,-1.00000]

Constraints:

1 <= equations.length <= 20
equations[i].length == 2
1 <= Ai.length, Bi.length <= 5
values.length == equations.length
0.0 < values[i] <= 20.0
1 <= queries.length <= 20
queries[i].length == 2
1 <= Cj.length, Dj.length <= 5
Ai, Bi, Cj, Dj consist of lower case English letters and digits.
*/
func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	// 1. 建圖
	graph := buildGraph(equations, values)

	// 2. 逐一處理查詢
	ans := make([]float64, len(queries))
	for i, q := range queries {
		start, end := q[0], q[1]
		ans[i] = queryAnswer(graph, start, end)
	}
	return ans
}

// 以 map[string]map[string]float64 的結構來建圖
func buildGraph(equations [][]string, values []float64) map[string]map[string]float64 {
	graph := make(map[string]map[string]float64)
	for i, eq := range equations {
		a, b := eq[0], eq[1]
		ratio := values[i]

		if graph[a] == nil {
			graph[a] = make(map[string]float64)
		}
		if graph[b] == nil {
			graph[b] = make(map[string]float64)
		}

		// a -> b
		graph[a][b] = ratio
		// b -> a
		graph[b][a] = 1.0 / ratio
	}
	return graph
}

func queryAnswer(graph map[string]map[string]float64, start, end string) float64 {
	// 若圖裡沒有 start 或 end，直接回傳 -1
	if graph[start] == nil || graph[end] == nil {
		return -1.0
	}
	// 若兩者相同
	if start == end {
		return 1.0
	}
	// 用 DFS (或 BFS) 來找路徑
	visited := make(map[string]bool)
	return dfs(graph, start, end, 1.0, visited)
}

// 以 DFS 方式尋找路徑並計算累積比值
func dfs(graph map[string]map[string]float64, curr, target string, product float64, visited map[string]bool) float64 {
	visited[curr] = true
	// 若當前節點的鄰居包含 target，直接算出結果
	for neighbor, ratio := range graph[curr] {
		if visited[neighbor] {
			continue
		}
		newProduct := product * ratio
		if neighbor == target {
			return newProduct
		}
		// 否則繼續往下搜
		res := dfs(graph, neighbor, target, newProduct, visited)
		if res != -1.0 {
			return res
		}
	}
	// 找不到就回傳 -1
	return -1.0
}
