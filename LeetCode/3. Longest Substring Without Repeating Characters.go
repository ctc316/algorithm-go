func lengthOfLongestSubstring(s string) int {
    // char: last position of char
    dict := make(map[byte]int)

    leng := len(s)
    max := 0
    count := 0
    start := 0
    for i := 0; i < leng; i++ {
        ch := s[i]
        if pos, ok := dict[ch]; ok && pos >= start {
            count = i - pos
            start = pos
        } else {
            count++
            if count > max {
                max = count
            }
        }
        dict[ch] = i
    }
    return max
}