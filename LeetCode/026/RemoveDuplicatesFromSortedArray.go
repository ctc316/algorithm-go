// [] -> []
// [1,1,2,3,4,4,4,5] -> [1,2,3,4,5]

// find same value
// move the different value over it

func removeDuplicates(nums []int) int {
    size := len(nums)
    if size == 0 {return 0}
    fixedNum := 1
    for i:=1; i<size; i++ {
        if nums[i] != nums[i-1] {
            nums[fixedNum] = nums[i]
            fixedNum++
        }
    }
    return fixedNum
}

