package testcase

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

var a = []string{"a", "b", "c"}

func TestStringJoin1plus(t *testing.T) {
	//方式1：+
	ret := a[0] + a[1] + a[2]
	fmt.Printf("%s\n", ret)
}

func TestStringJoin2Sprintf(t *testing.T) {
	//方式2：fmt.Sprintf
	ret := fmt.Sprintf("%s%s%s", a[0], a[1], a[2])
	fmt.Printf("%s\n", ret)

}
func TestStringJoin3Builder(t *testing.T) {
	//方式3：strings.Builder
	var sb strings.Builder
	sb.WriteString(a[0])
	sb.WriteString(a[1])
	sb.WriteString(a[2])
	ret := sb.String()
	fmt.Printf("%s\n", ret)
}
func TestStringJoin4Buffer(t *testing.T) {
	//方式4：bytes.Buffer
	buf := new(bytes.Buffer)
	buf.Write([]byte(a[0]))
	buf.Write([]byte(a[1]))
	buf.Write([]byte(a[2]))
	ret := buf.String()
	fmt.Printf("%s\n", ret)
}

func TestStringJoin5(t *testing.T) {
	//方式5：strings.Join
	ret := strings.Join(a, "")
	fmt.Printf("%s\n", ret)
}

// strings.Join ≈ strings.Builder > bytes.Buffer > "+" > fmt.Sprintf
func BenchmarkStringJoin1plus(b *testing.B) {
	//方式1：+
	for n := 0; n < b.N; n++ {
		_ = a[0] + a[1] + a[2]
	}
}

func BenchmarkStringJoin2Sprintf(b *testing.B) {
	//方式2：fmt.Sprintf
	for n := 0; n < b.N; n++ {
		_ = fmt.Sprintf("%s%s%s", a[0], a[1], a[2])
	}

}
func BenchmarkStringJoin3Builder(b *testing.B) {
	//方式3：strings.Builder
	for n := 0; n < b.N; n++ {
		var sb strings.Builder
		sb.WriteString(a[0])
		sb.WriteString(a[1])
		sb.WriteString(a[2])
		_ = sb.String()
	}
}
func BenchmarkStringJoin4Buffer(b *testing.B) {
	//方式4：bytes.Buffer
	for n := 0; n < b.N; n++ {
		buf := new(bytes.Buffer)
		buf.Write([]byte(a[0]))
		buf.Write([]byte(a[1]))
		buf.Write([]byte(a[2]))
		_ = buf.String()
	}
}

func BenchmarkString5Join(b *testing.B) {
	//方式5：strings.Join
	for n := 0; n < b.N; n++ {
		_ = strings.Join(a, "")
	}
}
