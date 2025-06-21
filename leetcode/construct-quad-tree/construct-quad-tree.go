package constructquadtree

type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}

func construct2(grid [][]int, size, r, c int) *Node {
	if size == 1 {
		return &Node{
			Val:    grid[r][c] == 1,
			IsLeaf: true,
		}
	}

	isSame := true
	val := grid[r][c]
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if grid[r+i][c+j] != val {
				isSame = false
				break
			}
		}
		if !isSame {
			break
		}
	}

	if isSame {
		return &Node{
			Val:    val == 1,
			IsLeaf: true,
		}
	}

	return &Node{
		IsLeaf:      false,
		TopLeft:     construct2(grid, size/2, r, c),
		TopRight:    construct2(grid, size/2, r, c+size/2),
		BottomLeft:  construct2(grid, size/2, r+size/2, c),
		BottomRight: construct2(grid, size/2, r+size/2, c+size/2),
	}
}

func construct(grid [][]int) *Node {
	size := len(grid)
	return construct2(grid, size, 0, 0)
}
