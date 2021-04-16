package main

import "fmt"


func archive(word string) string {
	archived := ""
	word_len := len(word)
	counter := 1
	for i := 0; i < word_len; i++ {
		if word_len - i - 1 > 0 {
			if word[i] == word[i + 1] {
				counter++
			} else {
				if counter > 1 {
					archived += fmt.Sprintf("%v%v", string(word[i]), counter)
				} else {
					archived += fmt.Sprintf("%v", string(word[i]))
				}
				counter = 1
			}
		} else {
			if counter > 1 {
				archived += fmt.Sprintf("%v%v", string(word[i]), counter)
			} else {
				archived += fmt.Sprintf("%v", string(word[i]))
			}
		}

	}
	return archived
}

func main() {
	fmt.Println("aaaa", archive("aaaa"))
	fmt.Println("abaaa", archive("abaaa"))
	fmt.Println("aaab", archive("aaab"))
	fmt.Println("aaaba", archive("aaaba"))
}