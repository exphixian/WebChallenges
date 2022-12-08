func numberOfSteps(num int) int {
    for i := 0; num >= 0; i++ {
        if num == 0 {
            return i
        }else if num %2 == 0 {
            num = num/2
        } else {
            num = num -1
        }
    }
    return 0
}
