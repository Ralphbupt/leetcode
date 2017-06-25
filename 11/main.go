package main

func maxArea0_1(height []int) int {
	max := 0
	for i := 0; i < len(height)-1; i++ {
		for j := i + 1; j < len(height); j++ {
			if height[i] >= height[j] {
				if (j-i)*height[j] > max {
					max = (j - i) * height[j]
				}
			} else {
				if (j-i)*height[i] > max {
					max = (j - i) * height[i]
				}
			}
		}
	}
	return max
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
func maxArea(height []int) int {

	result := 0

	l := 0
	r := len(height) - 1
	for l < r {
		result = max(result, min(height[l], height[r])*(r-l))
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return result
}

func main() {
	maxArea([]int{1, 2, 3})
}
