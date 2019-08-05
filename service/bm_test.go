package service

import (
	"fmt"
	"regexp"
	"strings"
	"testing"
)

func BenchmarkRex(b *testing.B) {
	b.SetBytes(1)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		F1()
	}
}

func BenchmarkRex1(b *testing.B) {
	b.SetBytes(1)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		F2()
	}
}

func TestF1(t *testing.T) {
	fmt.Println(F1())
}

func F1() string {
	store := []string{"paranormal", "paranormal", "graal", "none"}
	res := make([]string, 0)

	for _, v := range store {
		re := regexp.MustCompile(`a.`)
		s := re.FindAllString(v, -1)
		res = append(res, s...)
	}
	return strings.Join(res, ",")
}

func F2() string {
	store := []string{"paranormal", "paranormal", "graal", "none"}
	res := make([]string, 0)

	re := regexp.MustCompile(`a.`)

	for _, v := range store {
		s := re.FindAllString(v, -1)
		res = append(res, s...)
	}
	return strings.Join(res, ",")
}
