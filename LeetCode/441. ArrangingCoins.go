func arrangeCoins(n int) int {
    return int((-1 + math.Sqrt(float64(1+8*n)))/2)
}

// 1
// 3
// 6
// 10
// 15 = 1+2+3+4+5

// (1+x)*x/2 < n
// x^2 + x < 2n
// x^2 + x - 2n < 0

// x < (-1 + Sqrt(1-4*(-2n)))/2
// x > (-1 - Sqrt(1-4*(-2n)))/2

// x < (-1 + Sqrt(1+8n)))/2