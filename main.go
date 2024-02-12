package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
)

func main() {
	//OUTPUT NO 1
	input1 := "We Always Mekar"
	input2 := "coding is fun"
	fmt.Println(">>>>>>>>>>>>>> Nomor 1 <<<<<<<<<<<<<")
	fmt.Println("Output 1:", countLetters(input1))
	fmt.Println("Output 2:", countLetters(input2))

	//OUTPUT NO 2
	input4 := []string{"Abc", "bCd"}
	input5 := []string{"Oke", "One"}
	input6 := []string{"Pendanaan", "Terproteksi", "Untuk", "Dampak", "Berarti"}
	fmt.Println(">>>>>>>>>>>>>> Nomor 2 <<<<<<<<<<<<<")
	fmt.Println("Output 1:", sortLetters(input4))
	fmt.Println("Output 2:", sortLetters(input5))
	fmt.Println("Output 3:", sortLetters(input6))
}

// SOAL NOMOR 1
func countLetters(input string) string {
	input = strings.ToLower(input)
	letterCount := make(map[rune]int)

	for _, char := range input {
		if unicode.IsLetter(char) {
			letterCount[char]++
		}
	}

	var output strings.Builder
	for _, char := range input {
		if count, ok := letterCount[char]; ok {
			output.WriteString(fmt.Sprintf("%c=%d, ", char, count))
			delete(letterCount, char)
		}
	}

	result := strings.TrimSuffix(output.String(), ", ")

	return result
}

// SOAL NOMOR 2
func sortLetters(input []string) string {
	// Menggunakan map untuk menghitung jumlah masing-masing huruf
	counts := make(map[rune]int)
	for _, word := range input {
		for _, char := range word {
			// Mengabaikan karakter non-huruf
			if char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z' {
				counts[char]++
			}
		}
	}

	// Mengonversi map menjadi slice untuk pengurutan
	var letterCounts []struct {
		char  rune
		count int
	}
	for char, count := range counts {
		letterCounts = append(letterCounts, struct {
			char  rune
			count int
		}{char, count})
	}

	// Mengurutkan slice berdasarkan jumlah huruf (terbanyak ke terkecil)
	sort.Slice(letterCounts, func(i, j int) bool {
		if letterCounts[i].count == letterCounts[j].count {
			// Jika jumlah sama, urutkan berdasarkan huruf besar kecil
			return letterCounts[i].char < letterCounts[j].char
		}
		return letterCounts[i].count > letterCounts[j].count
	})

	var result strings.Builder
	for _, lc := range letterCounts {
		if lc.count == 1 || (lc.count == 2 && unicode.IsLower(lc.char)) {
			result.WriteRune(lc.char)
		} else if lc.count == 2 && unicode.IsUpper(lc.char) {
			result.WriteRune(lc.char)
		} else if lc.count > 2 || (lc.count == 2 && unicode.IsUpper(lc.char)) {
			result.WriteRune(lc.char)
		}
	}

	return result.String()
}
