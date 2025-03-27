package Evaluate_Division

type UnionFind struct {
	parent map[string]string
	weight map[string]float64 // weight[x] = x / parent[x]
}

func newUnionFind() *UnionFind {
	return &UnionFind{
		parent: make(map[string]string),
		weight: make(map[string]float64),
	}
}

func (uf *UnionFind) find(x string) (string, float64) {
	// 若 x 的父節點就是 x 自己，表示 x 是根
	if uf.parent[x] == x {
		return x, 1.0
	}
	// 否則遞迴往上找根
	p := uf.parent[x]
	root, rootRatio := uf.find(p)
	// 將 x 直接連到 root
	// 並更新 x 與 root 的比值 => (x / p) * (p / root) = x / root
	uf.parent[x] = root
	uf.weight[x] *= rootRatio
	return uf.parent[x], uf.weight[x]
}

func (uf *UnionFind) union(x, y string, ratio float64) {
	// ratio = x / y
	// 先找 x, y 的根與各自相對根的比值
	rootX, wX := uf.find(x) // wX = x / rootX
	rootY, wY := uf.find(y) // wY = y / rootY

	// 若已經在同一個連通分量就不用再 union
	if rootX == rootY {
		return
	}
	// 把 rootX 當作 rootY 的子節點
	uf.parent[rootX] = rootY
	// weight[rootX] = (x / y) * (y / rootY) / (x / rootX)
	//               = ratio * wY / wX
	uf.weight[rootX] = ratio * wY / wX
}

func (uf *UnionFind) add(x string) {
	if _, ok := uf.parent[x]; !ok {
		uf.parent[x] = x
		uf.weight[x] = 1.0
	}
}

func calcEquation1(equations [][]string, values []float64, queries [][]string) []float64 {
	uf := newUnionFind()
	// 1. 建立並查集
	for i, eq := range equations {
		a, b := eq[0], eq[1]
		ratio := values[i]
		uf.add(a)
		uf.add(b)
		uf.union(a, b, ratio)
	}
	// 2. 查詢
	ans := make([]float64, len(queries))
	for i, q := range queries {
		x, y := q[0], q[1]
		if _, ok := uf.parent[x]; !ok {
			ans[i] = -1.0
			continue
		}
		if _, ok := uf.parent[y]; !ok {
			ans[i] = -1.0
			continue
		}
		rootX, wX := uf.find(x) // wX = x / rootX
		rootY, wY := uf.find(y) // wY = y / rootY
		if rootX != rootY {
			ans[i] = -1.0
		} else {
			// x / y = (x / rootX) / (y / rootY) = wX / wY
			ans[i] = wX / wY
		}
	}
	return ans
}
