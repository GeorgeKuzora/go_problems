package main

func moveZeroes(nums []int)  {
	idx := 0

    for i := range nums {
        if nums[i] != 0 {
            if i != idx {
                nums[i], nums[idx] = nums[idx], nums[i]
            }
            idx++
        }
    }
}
