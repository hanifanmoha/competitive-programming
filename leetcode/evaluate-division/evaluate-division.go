package evaluatedivision

type Node struct {
	symbol string
	edges  []*Edge
	flag   bool
}

func NewNode(symbol string) *Node {
	return &Node{
		symbol: symbol,
	}
}

func (n *Node) AddEdge(edge *Edge) {
	n.edges = append(n.edges, edge)
}

func (n *Node) Find(symbol string) float64 {

	if n.flag {
		return -1
	}

	// fmt.Println(n.symbol, " : Find", symbol, ", flag", n.flag)

	if symbol == n.symbol {
		return 1
	}

	for _, edge := range n.edges {

		// fmt.Println("=> move to ", edge.node.symbol)
		n.flag = true
		d := edge.node.Find(symbol)
		n.flag = false
		if d != -1 {
			dist := d * edge.val
			return dist
		}
	}

	return -1
}

type Edge struct {
	node *Node
	val  float64
	// flag bool
}

func NewEdge(node *Node, val float64) *Edge {
	return &Edge{
		node: node,
		val:  val,
	}
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {

	nodes := make(map[string]*Node)
	n := len(equations)
	for i := range n {
		equation := equations[i]
		value := values[i]
		symbol1, symbol2 := equation[0], equation[1]

		node1, ok1 := nodes[symbol1]
		if !ok1 {
			node1 = NewNode(symbol1)
			nodes[symbol1] = node1
		}

		node2, ok2 := nodes[symbol2]
		if !ok2 {
			node2 = NewNode(symbol2)
			nodes[symbol2] = node2
		}

		edge := NewEdge(node2, value)
		node1.AddEdge(edge)

		reverseEdge := NewEdge(node1, 1/value)
		node2.AddEdge(reverseEdge)
	}

	nQuery := len(queries)

	res := make([]float64, nQuery)

	for i, query := range queries {
		s1, s2 := query[0], query[1]
		node1, ok1 := nodes[s1]
		_, ok2 := nodes[s2]
		// fmt.Println(i, s1, s2, ok1, ok2)
		if !ok1 || !ok2 {
			res[i] = -1
			continue
		}
		res[i] = node1.Find(s2)
	}

	return res
}
