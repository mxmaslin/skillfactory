package helpers


func GetMaxInt() int {
	maxUint := ^uint(0) 
	maxInt := int(maxUint >> 1) 
	return maxInt
}


func GetMinInt() int {
	maxInt := GetMaxInt()
	minInt := -maxInt - 1
	return minInt
}