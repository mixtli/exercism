package binarysearchtree

const testVersion = 1

type SearchTreeData struct {
	left  *SearchTreeData
	right *SearchTreeData
	data  int
}

// Bst creates a new binary search tree
func Bst(data int) *SearchTreeData {
	return &SearchTreeData{data: data}
}

// Insert inserts an integer into a binary search tree
func (bst *SearchTreeData) Insert(data int) {
	curr := bst
	for {
		if data > curr.data {
			if curr.right != nil {
				curr = curr.right
			} else {
				curr.right = Bst(data)
				return
			}
		} else {
			if curr.left != nil {
				curr = curr.left
			} else {
				curr.left = Bst(data)
				return
			}
		}
	}
}

// MapString returns an array of strings by applying f to each sorted element
func (bst *SearchTreeData) MapString(f func(int) string) []string {
	var str []string
	for _, i := range bst.MapInt(func(j int) int { return j }) {
		str = append(str, f(i))
	}
	return str

}

// MapInt returns a sorted set of integers for a search tree
func (bst *SearchTreeData) MapInt(f func(int) int) []int {
	if bst == nil {
		return []int{}
	}
	right := append([]int{f(bst.data)}, bst.right.MapInt(f)...)
	return append(bst.left.MapInt(f), right...)
}
