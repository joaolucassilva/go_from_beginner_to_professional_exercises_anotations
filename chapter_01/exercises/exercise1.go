package exercises

import (
	"fmt"
	"math/rand"
	"strings"
)

func RandNumber() {
	r := rand.Intn(5) + 1
	stars := strings.Repeat("*", r)
	fmt.Println(stars)
}
