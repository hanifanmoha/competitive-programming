package minimumgeneticmutation

const GENE_LENGTH = 8

type Node struct {
	gene      string
	flag      bool
	mutations []*Node
}

func NewNode(gene string) *Node {
	return &Node{
		gene:      gene,
		mutations: make([]*Node, 0),
	}
}

func (n *Node) AddMutation(m *Node) {
	n.mutations = append(n.mutations, m)
}

type QueueItem struct {
	node *Node
	step int
}

func isNeighbor(geneA, geneB string) bool {
	count := 0
	for i := range GENE_LENGTH {
		if geneA[i] != geneB[i] {
			count++
		}
	}
	return count == 1
}

func minMutation(startGene string, endGene string, bank []string) int {

	N := len(bank)
	validGenes := make(map[string]*Node)
	for i := range N {
		gene := bank[i]
		validGenes[gene] = NewNode(gene)
	}

	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			geneA, geneB := bank[i], bank[j]
			vga, vgb := validGenes[geneA], validGenes[geneB]
			if isNeighbor(geneA, geneB) {
				vga.AddMutation(vgb)
				vgb.AddMutation(vga)
			}
		}
	}

	// calculate step

	queue := make([]*QueueItem, 0)
	idx := 0

	// if start gene valid, add it as initial queue

	if vg, exist := validGenes[startGene]; exist {
		vg.flag = true
		queue = append(queue, &QueueItem{node: vg, step: 0})
	} else {
		// if start gene invalid, add its mutation as initial queue
		for _, vg := range validGenes {
			if isNeighbor(vg.gene, startGene) {
				vg.flag = true
				queue = append(queue, &QueueItem{node: vg, step: 1})
			}
		}
	}

	for idx < len(queue) {

		qItem := queue[idx]
		qNode := qItem.node
		qStep := qItem.step

		if qNode.gene == endGene {
			return qStep
		}

		// add neighbor to queue if not yet added
		for _, m := range qNode.mutations {
			if !m.flag {
				m.flag = true
				queue = append(queue, &QueueItem{
					node: m,
					step: qStep + 1,
				})
			}
		}

		idx++
	}

	return -1
}
