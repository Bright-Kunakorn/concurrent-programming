func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    // Ensure that nums1 is the smaller array
    if len(nums1) > len(nums2) {
        return findMedianSortedArrays(nums2, nums1)
    }

    m, n := len(nums1), len(nums2)
    left, right := 0, m
    median1, median2 := 0, 0

    for left <= right {
        // Partition the smaller array
        partition1 := (left + right) / 2
        // Calculate the partition for the larger array
        partition2 := (m+n+1)/2 - partition1

        // Handle edge cases for partitions
        maxLeft1, minRight1 := math.MinInt32, math.MaxInt32
        if partition1 != 0 {
            maxLeft1 = nums1[partition1-1]
        }
        if partition1 != m {
            minRight1 = nums1[partition1]
        }

        maxLeft2, minRight2 := math.MinInt32, math.MaxInt32
        if partition2 != 0 {
            maxLeft2 = nums2[partition2-1]
        }
        if partition2 != n {
            minRight2 = nums2[partition2]
        }

        // Check if the partitions are correct
        if maxLeft1 <= minRight2 && maxLeft2 <= minRight1 {
            // Calculate the medians
            if (m+n)%2 == 0 {
                median1 = max(maxLeft1, maxLeft2)
                median2 = min(minRight1, minRight2)
                return float64(median1+median2) / 2.0
            } else {
                median1 = max(maxLeft1, maxLeft2)
                return float64(median1)
            }
        } else if maxLeft1 > minRight2 {
            // Need to move left in nums1
            right = partition1 - 1
        } else {
            // Need to move right in nums1
            left = partition1 + 1
        }
    }

    return 0.0
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
