package trees

type QuadTreeNode struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *QuadTreeNode
	TopRight    *QuadTreeNode
	BottomLeft  *QuadTreeNode
	BottomRight *QuadTreeNode
}

func construct(grid [][]int) *QuadTreeNode {
	return constructQuadTree(grid, 0, 0, len(grid))
}

func constructQuadTree(grid [][]int, row, col, length int) *QuadTreeNode {
	node := &QuadTreeNode{
		IsLeaf: true,
		Val:    grid[row][col] == 1,
	}

	if !isLeaf(grid, row, col, length) {
		node.IsLeaf = false

		nextGridLength := length / 2

		node.TopLeft = constructQuadTree(grid, row, col, nextGridLength)
		node.TopRight = constructQuadTree(grid, row, col+nextGridLength, nextGridLength)
		node.BottomLeft = constructQuadTree(grid, row+nextGridLength, col, nextGridLength)
		node.BottomRight = constructQuadTree(grid, row+nextGridLength, col+nextGridLength, nextGridLength)
	}

	return node
}

func isLeaf(grid [][]int, row, col, length int) bool {
	firstElement := grid[row][col]

	for i := row; i < row+length; i++ {
		for j := col; j < col+length; j++ {
			if grid[i][j] != firstElement {
				return false
			}
		}
	}
	return true
}
