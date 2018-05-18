// Version 1: Sorting & two pointer, Time: O(n), Space: O(n)
type Pair struct {
    index int
    val int
}

type Pairs []Pair

//sort.Interface - Len, Swap, Less
func (p Pairs) Len() int {
    return len(p)
}
func (p Pairs) Swap(i, j int) {
    p[i], p[j] = p[j], p[i]
}
func (p Pairs) Less(i, j int) bool {
    return p[i].val < p[j].val
}

func twoSum(nums []int, target int) []int {
    // record index for every number
    pairs := make(Pairs, len(nums))
    for i := 0; i < len(nums); i++ {
        pairs[i] = Pair{i, nums[i]}
    }

    // sort
    sort.Sort(pairs)

    // find solution
    left := 0
    right := len(nums) - 1
    for left < right {
        diff := target - pairs[right].val - pairs[left].val
        if (diff < 0) {
            right -= 1
        } else if (diff > 0) {
            left += 1
        } else {
            return []int{pairs[left].index, pairs[right].index}
        }
    }

    return []int{}
}


// Version 2: Use dictionary (hash map) to record previous index of number, Time: O(n), Space: O(n)
func twoSum(nums []int, target int) []int {
    // val:index
    dict := make(map[int]int)

    for key, val := range nums {
        // check if the diff scanned
        if theOtheridx, ok := dict[target - val]; ok {
            return []int{theOtheridx, key}
        }
        // store value-index pair
        if _, ok := dict[val]; !ok {
            dict[val] = key
        }
    }

    return []int{}
}