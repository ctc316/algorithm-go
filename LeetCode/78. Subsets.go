func recur(values []int, counts []int, start int) [][]int {
    var set []int
    for i:=0; i< len(counts); i++ {
        if counts[i] > 0 {
            set = append(set, values[i])
        }
    }
    ans := [][]int{set}
    for i:= start; i<len(counts); i++ {
        if(counts[i] != 0) {
            counts[i] -= 1
            ans = append(ans, recur(values, counts, i)...)
            counts[i] += 1
        }
    }
    return ans
}

func subsets(nums []int) [][]int {
    // build values and counts
    var counts, values []int 
    for i:=0; i<len(nums); i++ {
        val := nums[i]
        found := false
        for j:=0; j<len(values); j++ {
            if val == values[j] {
                counts[j] += 1
                found = true
                break
            }
        }
        if(!found) {
            values = append(values, val)
            counts = append(counts, 1)
        }
    }
    return recur(values, counts, 0)
}