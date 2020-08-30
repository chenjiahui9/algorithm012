func findCircleNum(M [][]int) int {
	n := len(M)
	p := make([]int,n)
	for i := 0; i < n; i++ {
		p[i] = i
	}

	var f func(int) int; f = func(x int) int {
		if p[x] != x {
			p[x] = f(p[x])
		}
		return p[x]
	}

	ans := n 
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if M[j][i] == 1 {
				pi,pj := f(i),f(j)
				if pi != pj {
					p[pj] = pi
					ans--
				}
			}
		}
	}
	return ans
}
//并查集
type UnitFind struct {
	count int
	parent []int
}
func NewUnitFind(n int) *UnitFind {
	parent := make([]int,n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	return &UnitFind{
		count : n,
		parent : parent,
	}
}
//路径压缩后 查询为O(1)
func (u *UnitFind) Find(i int) int {
	
	root := i
	for u.parent[root] != root {
		root = u.parent[root]
	}
	for u.parent[i] == i {
		i,u.parent[i] = u.parent[i],root
	}
	return root
	
}

func (u *UnitFind) Union(i,j int){
	
	pi := u.Find(i)
	pj := u.Find(j)
	if pi != pj {
		u.parent[i] = pj
		u.count--
	}
}
