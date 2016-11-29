package shortUrl

import (
	//"fmt"
	"math"
	"strings"
)

var base = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

var base_reverse = map[string]int{"0": 0, "1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9, "A": 10, "B": 11, "C": 12, "D": 13, "E": 14, "F": 15, "G": 16, "H": 17, "I": 18, "J": 19, "K": 20, "L": 21, "M": 22, "N": 23, "O": 24, "P": 25, "Q": 26, "R": 27, "S": 28, "T": 29, "U": 30, "V": 31, "W": 32, "X": 33, "Y": 34, "Z": 35, "a": 36, "b": 37, "c": 38, "d": 39, "e": 40, "f": 41, "g": 42, "h": 43, "i": 44, "j": 45, "k": 46, "l": 47, "m": 48, "n": 49, "o": 50, "p": 51, "q": 52, "r": 53, "s": 54, "t": 55, "u": 56, "v": 57, "w": 58, "x": 59, "y": 60, "z": 61}

var baseLen = len(base)

func Encode(s string) int64 {
	sList := strings.Split(s, "")
	var sum float64 = 0
	for index, value := range sList {
		sum += float64(base_reverse[value]) * math.Pow(float64(baseLen), float64(index))
	}
	return int64(sum)
}

func Decode(i int64) string {
	var stringDecode string = ""

	for i > 0 {
		remainder := i % int64(baseLen)
		i = i / int64(baseLen)
		stringDecode += base[remainder]
	}
	return stringDecode
}
