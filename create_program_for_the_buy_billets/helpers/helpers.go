package helpers

import "strconv"

func Get_format_money(num int) string {
	result := ""
	numStr := ""
	i:=0
	for{
		if num%10==0{
			num/=10
			i++
		}else {
			numStr=strconv.Itoa(num)
			break
		}
		if i==3{
			i=0
			result = "K" + result
		}
		if result=="KK"{
			result="M"
		}
	}
	for j:=0; j<i; j++{
		numStr+="0"
	}
	return numStr+result
}
