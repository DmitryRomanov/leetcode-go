// https://leetcode.com/problems/domino-and-tromino-tiling/submissions/
package domino_and_tromino_tiling

func numTilings(n int) int {
	var dfs func(int, int) int
	// - 0 - обе клетки в колонке свободны
	// - 1 — верхняя клетка в колонке занята, нижняя свободна
	// - 2 — нижняя клетка в колонке занята, верхняя свободна
	cache := make([][3]int, 0, n+2)
	for i := 0; i < n+2; i++ {
		cache = append(cache, [3]int{-1, -1, -1})
	}
	dfs = func(col int, state int) int {
		if cache[col][state] != -1 {
			return cache[col][state]
		}
		// Описываем 4 угла очередной секции:
		// {l,r}               {t,b}
		//  ^ левый или правый  ^ верхний или нижний
		// lt, lb могут быть заняты в следствии того
		// как мы положили предыдущие фигуры
		lt := !(state == 1)
		lb := !(state == 2)

		rt := col+1 <= n
		rb := col+1 <= n

		// _ x
		// _ x <- край доски
		// ^
		// последняя колонка
		if lt && lb && !rt && !rb {
			return 1
		}

		result := 0

		// _ _
		// _ _
		// две колонки свободные
		if lt && lb && rt && rb {
			result += dfs(col+1, 0)
			result += dfs(col+1, 1)
			result += dfs(col+1, 2)
			result += dfs(col+2, 0)
		}

		// x _
		// _ _
		// левый верхний угол занят
		if !lt && lb && rt && rb {
			result += dfs(col+1, 2)
			result += dfs(col+2, 0)
		}

		// _ _
		// x _
		// левый нижний угол занят
		if lt && !lb && rt && rb {
			result += dfs(col+1, 1)
			result += dfs(col+2, 0)
		}

		cache[col][state] = result % (1e9 + 7)
		return cache[col][state]
	}

	return dfs(1, 0)
}
