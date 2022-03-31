func subtractProductAndSum(n int) int {
    s := strconv.Itoa(n)

	pro := 1
	sum := 0

	for _, r := range s {
		val, _ := strconv.Atoi(string(r))

		pro *= val
		sum += val
	}
    
    return pro - sum
}