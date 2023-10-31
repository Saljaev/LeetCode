func pascalRow(rowIndex int) []int {
    result := make([]int, rowIndex + 1)
    c := 1

    result[0] = c

    for i := 1; i <= rowIndex; i++ {
        c = c*(rowIndex + 1 - i)/i
        result[i] = c
    }

    return result
}

func getRow(rowIndex int) []int {
    return pascalRow(rowIndex)
}