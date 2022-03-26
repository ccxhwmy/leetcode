package sort

func bubbleSort(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		for j := 1; j < len(nums) - i; j++ {
			if nums[j - 1] < nums[j] {
				nums[j - 1], nums[j] = nums[j], nums[j - 1]
			}
		}
	}
	return nums
}

func selectSort(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	return nums
}

func insertSort(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i; j > 0; j-- {
			if nums[j - 1] > nums[j] {
				nums[j - 1], nums[j] = nums[j], nums[j - 1]
			}
		}
	}
	return nums
}

func quickSort(nums []int) []int {
	var quick func(nums []int, left ,right int) []int
	quick = func(nums []int, left, right int) []int {
		if left < right {
			return nil
		}
		i, j, pivot := left, right, nums[left]
		for i < j {
			for i < j && nums[j] >= pivot {
				j--
			}
			for i < j && nums[i] <= pivot {
				i++
			}
			nums[i], nums[j] = nums[j], nums[i]
		}
		nums[i], nums[left] = nums[left], nums[i]
		quick(nums, left, i - 1)
		quick(nums, i + 1, right)
		return nums
	}
	return quick(nums, 0, len(nums) - 1)
}


