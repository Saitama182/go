/*
Условие задачи:
Напишите функцию, которая принимает строку и возвращает самое длинное слово в этой строке и его длину.
Если в строке несколько слов с одинаковой максимальной длиной, верните первое из них.
Слова в строке разделены пробелами.

Примеры:
 "А роза упала на лапу Азора" - самое длинное слово "Азора", длина 5.
 "hello world" - самое длинное слово "hello", длина 5.

Ограничения:
Строка может содержать пробелы и знаки препинания.
*/


package main 

import "fmt"


func findLongestWord(s string) string {

	var count, maxCount int
	var res, maxRes string

	for _, r := range s {
		if r == ' ' || r == ','{
			if count > maxCount{
				maxCount = count
				maxRes = res
			} 
			res = ""
			count = 0
		} else {
			count++
			res += string(r)
		}
	}

	if count > maxCount {
		maxRes = res
	}


	return maxRes
   }






func main() {
	m := findLongestWord("No x in Nixon")

	fmt.Println(m)
}