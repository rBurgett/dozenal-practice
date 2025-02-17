package dozenal

import (
	"math/rand"
	"strconv"
	"strings"
)

func GetRandom(minI int, maxI int) int {
	return minI + rand.Intn(maxI-minI)
}

func ToBase12(num int) string {
	numStr := strings.ToUpper(strconv.FormatInt(int64(num), 12))
	numStr = strings.ReplaceAll(numStr, "A", "X")
	numStr = strings.ReplaceAll(numStr, "B", "E")

	return numStr
}

func FromBase12(numStr string) (int, error) {
	numStr = strings.ToUpper(numStr)
	numStr = strings.ReplaceAll(numStr, "X", "A")
	numStr = strings.ReplaceAll(numStr, "E", "B")
	num, err := strconv.ParseInt(numStr, 12, 64)
	if err != nil {
		return 0, err
	}

	return int(num), nil
}
