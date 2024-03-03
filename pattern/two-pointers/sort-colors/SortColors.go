package sortcolors

// SortColors sort an array and then returns it in ascending order
func SortColors(colors []int) []int {

	red, white, blue := 0, 0, len(colors)-1
	for white <= blue {
		if colors[white] == 0 {
			if colors[red] != 0 {
				colors[red], colors[white] = colors[white], colors[red]
			}
			white += 1
			red += 1
		} else if colors[white] == 1 {
			white += 1
		} else if colors[white] == 2 {
			if colors[blue] != 2 {
				colors[blue], colors[white] = colors[white], colors[blue]
			}
			blue -= 1
		}
	}

	return colors
}
