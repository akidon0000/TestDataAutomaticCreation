package main

import (
	"fmt"
	"strconv"
	"time"
	"math/rand"

	"github.com/360EntSecGroup-Skylar/excelize"

)

// ===============
const excelPath = "./testData2.xlsx"
const excelSheet = "Sheet1"
const testNum = 500000
const date = "0601"
var quadkeyList[3] string = [3]string {"133002112310210000", "133002022031110000", "132113311032111222"} // 東京，西宮，徳島
// ===============

func main(){
	f, err := excelize.OpenFile(excelPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := 0; i < testNum; i++ {
		if ((i+1)%1000 == 0){
			fmt.Println(i+1)
		}

		rand.Seed(time.Now().UnixNano())
		var a int
		var b int
		var c int
		for i := 0; i < 10; i++ {
			a = int(rand.Intn(8)+1) // 1-8桁の文字列
			b = int(rand.Intn(8)+1)
		}
		random1, _ := MakeRandomStr(a)
		random2, _ := MakeRandomStr(b)
		// c = int(rand.Intn(3))
		c = i % 3

		f.SetCellValue(excelSheet, "A"+strconv.Itoa(i+1), "0")
		f.SetCellValue(excelSheet, "B"+strconv.Itoa(i+1), "test2021-" + date + "-" + strconv.Itoa(i+1))
		f.SetCellValue(excelSheet, "C"+strconv.Itoa(i+1), random1)
		f.SetCellValue(excelSheet, "D"+strconv.Itoa(i+1), random2)
		f.SetCellValue(excelSheet, "E"+strconv.Itoa(i+1), quadkeyList[c])
		f.SetCellValue(excelSheet, "F"+strconv.Itoa(i+1), "NULL")
	}
	f.SaveAs(excelPath)
}

func MakeRandomStr(digit int) (string, error) {
		const letters = "あいうえおかきくけこがぎぐげごさしすせそざじずぜぞたちつてとだぢづでどなにぬねのはひふへほばびぶべぼぱぴぷぺぽやゆよらりるれろわをん"
		rand.Seed(time.Now().UnixNano())

		var a int
		var result string
		result = ""
		for i := 0; i < digit; i++ {
			for i := 0; i < 66; i++ {
				a = rand.Intn(66)
			}
			rs := []rune(letters)
			result += string(rs[a:a+1])
		}
		return result, nil
}
