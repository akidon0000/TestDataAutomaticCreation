package main

import (
	"fmt"
	"strconv"
	"time"
	"math/rand"
	"github.com/360EntSecGroup-Skylar/excelize"
)

// ===============
const testNum = 1000000
const forNum = 5

const date = "0711"
const quadkeyShort = "13"
const excelSheet = "Sheet1"
// ===============

func main(){
	// ============タイム測定=============
	part_processing_start := time.Now()
	// =================================
	for i := 0; i < forNum; i++{
		excelPath := "./users" + strconv.Itoa(i+1) + ".xlsx"
		excelWrite(excelPath,i+1)
		// ============タイム測定=============
		part_processing_end := time.Now()
		time_mysql := int((part_processing_end.Sub(part_processing_start)).Seconds())
		// =================================
		hours := int(time_mysql / 3600)
		minutes := int(time_mysql % 3600/60)
		seconds := int(time_mysql % 60)

		fmt.Println("作成時間 : ", hours, "時間", minutes, "分",seconds, "秒")
	}
}

func excelWrite(excelPath string,iNum int){
	f, err := excelize.OpenFile(excelPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := 0; i < testNum; i++ {
		if ((i+1)%1000 == 0){
			str := strconv.Itoa(iNum) + "-" + strconv.Itoa(i+1)
			fmt.Println(str)
		}

		rand.Seed(time.Now().UnixNano())
		var a int
		var b int
		var c int
		for i := 0; i < 10; i++ {
			a = int(rand.Intn(8)+1) // 1-8桁の文字列
			b = int(rand.Intn(8)+1)
		}
		c = 16
		random1, _ := MakeRandomStr(a)
		random2, _ := MakeRandomStr(b)
		random3, _ := MakeRandomQuadkey(c)

		quadkey := quadkeyShort + random3

		f.SetCellValue(excelSheet, "A"+strconv.Itoa(i+1), "0")
		f.SetCellValue(excelSheet, "B"+strconv.Itoa(i+1), "test2021-" + date + "-" + strconv.Itoa(i+1))
		f.SetCellValue(excelSheet, "C"+strconv.Itoa(i+1), random1)
		f.SetCellValue(excelSheet, "D"+strconv.Itoa(i+1), random2)
		f.SetCellValue(excelSheet, "E"+strconv.Itoa(i+1), quadkey)
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
func MakeRandomQuadkey(digit int) (string, error) {
	const letters = "0123"
	rand.Seed(time.Now().UnixNano())

	var a int
	var result string
	result = ""
	for i := 0; i < digit; i++ {
		for i := 0; i < 4; i++ {
			a = rand.Intn(4)
		}
		rs := []rune(letters)
		result += string(rs[a:a+1])
	}
	return result, nil
}
