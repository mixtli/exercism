package pascal

const testVersion = 1

func Triangle(n int) [][]int {
	if n == 0 {
		return [][]int{{}}
	}
	if n == 1 {
		return [][]int{{1}}
	}
	prev := Triangle(n - 1)
	last_row := prev[len(prev)-1]
	new_row := []int{1}
	for i := 0; i < len(last_row)-1; i++ {
		new_row = append(new_row, last_row[i]+last_row[i+1])
	}
	new_row = append(new_row, 1)
	triangle := append(prev, new_row)
	return triangle
}
