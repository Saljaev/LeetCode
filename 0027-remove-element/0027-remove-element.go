func removeElement(nums []int, val int) int {
    result := len(nums)
    
    for i := 0; i < len(nums); i++ {
        if nums[i] == val {
            result--
            nums = append(nums[:i],nums[i+1:]...)
            i--
        }
    }
    
    return result 
}