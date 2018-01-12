func coinChange(coins []int, amount int) int {
    sort.Ints(coins)
    var dp = make([]int, amount+1)
    dpLen := len(dp)
    for i := 1; i < dpLen; i++ { dp[i] = math.MaxInt32 }
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
            val1 := dp[i]
            val2 := dp[i-coin] + 1
            if val1 < val2 {
                dp[i] = val1
            } else {
                dp[i] = val2
            }
        }
    }
    if dp[amount] == math.MaxInt32 {
        return -1
    }
    return dp[amount]
}