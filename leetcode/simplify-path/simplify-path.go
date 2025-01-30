package simplify_path

type Stack struct {
	index int
	data  []string
}

func (s *Stack) Init() {
	s.data = make([]string, 0)
	s.index = -1
}

// 0 1 2 _ _

func (s *Stack) Push(str string) {
	s.index++
	if len(s.data) >= s.index+1 {
		s.data[s.index] = str
	} else {
		s.data = append(s.data, str)
	}
}

func (s *Stack) Pop() {
	s.index--
	if s.index < -1 {
		s.index = -1
	}
}

func (s *Stack) Get() []string {
	return s.data[:s.index+1]
}

func splitPath(path string) []string {

	path = path + "/"
	start := 0
	res := make([]string, 0)

	for i := range path {
		if path[i] == '/' {
			if i > start {
				res = append(res, path[start:i])
			}
			start = i + 1
		}
	}

	return res
}

func joinPath(dirs []string) string {
	if len(dirs) == 0 {
		return "/"
	}
	str := ""
	for _, dir := range dirs {
		str += "/" + dir
	}
	return str
}

func simplifyPath(path string) string {
	// fmt.Println("Path :", path)

	dirs := splitPath(path)
	// fmt.Println("Split :", dirs)

	stack := Stack{}
	stack.Init()

	for _, dir := range dirs {
		if dir == ".." {
			stack.Pop()
			continue
		} else if dir == "." || dir == "" {
			continue
		}
		stack.Push(dir)
	}

	// fmt.Println("Simplified : ", stack.Get())

	return joinPath(stack.Get())
}
