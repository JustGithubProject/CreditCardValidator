package main

func getSliceByCardNumber(cardNumber int) []int {
	var resultSlice []int

	for cardNumber > 0 {
		num := cardNumber % 10
		resultSlice = append([]int{num}, resultSlice...)
		cardNumber /= 10
	}

	return resultSlice
}


func doubleEverySecondInSlice(nums []int) {
	for i := len(nums) - 2; i >= 0; i -= 2 {
		nums[i] *= 2
	}
}


func replaceEverySecondByCondition(nums []int) {
	for i := len(nums) - 2; i >= 0; i -= 2 {
		if nums[i] > 9 {
			// Getting first and second digits of nums[i]
			copyNum := nums[i]
			firstNum := copyNum % 10
			copyNum /= 10
			secondNum := copyNum % 10
			
			nums[i] = firstNum + secondNum
		}
	}
}


func getSumOfSlice(nums []int) int {
	sumRes := 0

	for i := 0; i < len(nums); i++ {
		sumRes += nums[i]
	}

	return sumRes
}


func IsCardValid(cardNumber int) bool {
	nums := getSliceByCardNumber(cardNumber)
	doubleEverySecondInSlice(nums)
	replaceEverySecondByCondition(nums)
	sumRes := getSumOfSlice(nums)

	if sumRes % 10 == 0 {
		return true
	} 

	return false

}