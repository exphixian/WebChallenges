func smallerNumbersThanCurrent(nums []int) []int {
    
    numbers := []int{}
    for i := 0; i < len(nums); i++ {
        var total int
        for j := 0; j < len(nums); j++ {
            if nums[i] > nums[j] {
                total++
            }
        }
        numbers = append(numbers, total)
    }
    return numbers
}
