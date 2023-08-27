package numcn

import "fmt"

func main() {
	chNum := "负十七亿零五十三万七千零一十六"
	num, _ := DecodeToInt64(chNum)
	fmt.Println(num) // -1700537016
	chNumAgain := EncodeFromInt64(num)
	fmt.Println(chNumAgain) // 负十七亿零五十三万七千零一十六

	chFloatNum := "负零点零七三零六"
	fNum, _ := DecodeToFloat64(chFloatNum)
	fmt.Printf("%f\n", fNum) // -0.073060
	chFloatNumAgain := EncodeFromFloat64(fNum)
	fmt.Println(chFloatNumAgain) // 负零点零七三零六
}
