func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    if (nums1 == nil || len(nums1) == 0) && 
        (nums2 == nil || len(nums2) == 0) {
        return 0
    }

    total := len(nums1) + len(nums2)

    // Even
    if total % 2 == 0 {
        return (float64(findKthNum(nums1, nums2, total / 2)) + float64(findKthNum(nums1, nums2, total / 2 + 1))) / 2  
    }

    // Odd
    return float64(findKthNum(nums1, nums2, total / 2 + 1))
}

func findKthNum(nums1 []int, nums2 []int, k int) int {
    // an array is empty
    if len(nums1) == 0 {
        return nums2[k - 1]
    }
    if len(nums2) == 0 {
        return nums1[k - 1]
    }

    // last one
    if k == 1 {
        if nums1[0] < nums2[0] {
            return nums1[0]
        }
        return nums2[0]
    }

    mid := k / 2
    remain := k - mid
    if (len(nums1) < mid) {
        return findKthNum(nums1, nums2[mid:], remain)
    }
    if (len(nums2) < mid) {
        return findKthNum(nums1[mid:], nums2, remain)
    }

    if (nums1[mid - 1] < nums2[mid - 1]) {
        return findKthNum(nums1[mid:], nums2, remain)
    } else {
        return findKthNum(nums1, nums2[mid:], remain)
    }
}