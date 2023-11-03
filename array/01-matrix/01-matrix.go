// https://leetcode.com/problems/01-matrix/description/
package matrix_01

func updateMatrix(mat [][]int) [][]int {
	result := [][]int{}
	queue := [][]int{}
	for i := range mat {
		i_result := []int{}
		for j := range mat[i] {
			if mat[i][j] == 0 {
				i_result = append(i_result, 0)
			} else {
				i_result = append(i_result, len(mat))
				queue = append(queue, []int{i, j})
			}
		}
		result = append(result, i_result)
	}

	var current []int

	for len(queue) > 0 {
		current, queue = queue[0], queue[1:]
		cur_i, cur_j := current[0], current[1]

		directions := [][2]int{
			{cur_i, cur_j + 1},
			{cur_i, cur_j - 1},
			{cur_i + 1, cur_j},
			{cur_i - 1, cur_j},
		}
		for _, direction := range directions {
			i := direction[0]
			j := direction[1]
			if i >= 0 && j >= 0 && i < len(mat) && j < len(mat[i]) {
				result_in_new_direction := result[i][j] + 1
				current_result := result[cur_i][cur_j]

				if result_in_new_direction < current_result {
					result[cur_i][cur_j] = result_in_new_direction
					queue = append(queue, []int{i, j + 1})
					queue = append(queue, []int{i, j - 1})
					queue = append(queue, []int{i + 1, j})
					queue = append(queue, []int{i - 1, j})
				}
			}

		}
	}

	return result
}
