package course_schedule

// https://en.wikipedia.org/wiki/Topological_sorting#Algorithms Depth-first search
// https://ru.algorithmica.org/cs/graph-traversals/cycle/

type G struct {
	g         [][]int
	L         []int
	required  []int
	permanent []bool
	temporary []bool
	stop      bool
}

func newG(numCourses int, prerequisites [][]int) *G {
	g := G{
		g:         make([][]int, numCourses),
		L:         make([]int, 0, numCourses),
		permanent: make([]bool, numCourses),
		temporary: make([]bool, numCourses),
		required:  make([]int, numCourses),
	}
	for _, v := range prerequisites {
		g.g[v[1]] = append(g.g[v[1]], v[0])
		g.required[v[0]]++
	}
	return &g
}

func (g *G) dfs() bool {
	for n, c := range g.required {
		if c > 0 {
			continue
		}
		if g.permanent[n] {
			continue
		}
		g.visit(n)
		if g.stop {
			return false
		}
	}
	if len(g.L) == len(g.g) {
		return true
	}
	return false
}

func (g *G) visit(n int) {
	if g.permanent[n] {
		return
	}
	if g.temporary[n] {
		g.stop = true
		return
	}
	g.temporary[n] = true
	for _, m := range g.g[n] {
		if g.temporary[m] {
			g.stop = true
			return
		}
		g.visit(m)
		if g.stop {
			return
		}
	}
	g.temporary[n] = false
	g.permanent[n] = true
	g.L = append(g.L, n)
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	if numCourses == 1 {
		return true
	}
	if len(prerequisites) == 0 {
		return true
	}
	return newG(numCourses, prerequisites).dfs()
}
