func strStr(haystack string, needle string) int {
    /** KMP **/
    //[a b c d a b c a b a a]
    //[a b c a b c]
    //[0 0 0 1 2 3]

    //[a a b a a a b a a a c]
    //[a a b a a a c]
    //[0 1 0 1 2 1 0]

    needleLen := len(needle)
    if(needleLen == 0) {return 0}

    // prefix
    prefix := make([]int, needleLen) 
    for i, j := 1, 0; i < needleLen;  {
        if needle[i] == needle[j] {
            prefix[i] = j+1
            i++
            j++
        } else {
            if j == 0 {
                prefix[i] = 0
                i++
            } else { j = prefix[j-1] }
        }
    }

    // run check
    targetLen := len(haystack)
    for i, j := 0, 0; i < targetLen && j < needleLen; {
        if haystack[i] == needle[j]{
            i++
            j++
            if j==needleLen {return i-j}
        } else {
            if j == 0 {
                i++
            } else { j = prefix[j-1] }
        }
    }
    return -1

    /***   Scan Method     ***/
    // needleLen := len(needle)
    // if(needleLen == 0) {return 0}
    // stackLen := len(haystack)
    // invalidPos := stackLen-needleLen + 1
    // for i:=0; i<invalidPos; i++ {
    //     match := true
    //     for j:=0 ; j < needleLen; j++ {
    //         if haystack[i+j] != needle[j] {
    //             match = false
    //             break
    //         }
    //     }
    //     if match {return i}
    // }
    // return -1
}