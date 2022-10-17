package course_schedule

type G struct {
	g         [][]int
	L         []int
	required  []int
	permanent map[int]struct{}
	temporary map[int]bool
	stop      bool
}

func newG(numCourses int, prerequisites [][]int) *G {
	g := G{
		g:         make([][]int, numCourses),
		L:         make([]int, 0, numCourses),
		permanent: make(map[int]struct{}, numCourses),
		temporary: make(map[int]bool, numCourses),
		required:  make([]int, numCourses),
	}
	for _, v := range prerequisites {
		g.g[v[1]] = append(g.g[v[1]], v[0])
		g.required[v[0]]++
	}
	return &g
}

func (g *G) bfs() bool {
	for n, c := range g.required {
		if c > 0 {
			continue
		}
		if _, ex := g.permanent[n]; ex {
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
	if _, ex := g.permanent[n]; ex {
		return
	}
	if g.temporary[n] {
		g.stop = true
		return
	}
	g.temporary[n] = true
	for _, m := range g.g[n] {
		g.visit(m)
		if g.stop {
			return
		}
	}
	g.temporary[n] = false
	g.permanent[n] = struct{}{}
	g.L = append(g.L, n)
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	if numCourses == 1 {
		return true
	}
	if len(prerequisites) == 0 {
		return true
	}
	return newG(numCourses, prerequisites).bfs()
}
