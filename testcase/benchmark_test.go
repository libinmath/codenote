package testcase

import (
  "bytes"
  "reflect"
  "testing"
)

var slice1 = []byte{'G','E','E','K','S'}
var slice2 = []byte{'G','E','e','K','S'}


func BenchmarkByteEqual(b *testing.B) {
  for n:= 0; n<b.N; n++{
    bytes.Equal(slice1, slice2)
  }
}

func BenchmarkReflectDeepEqual(b *testing.B) {
  for n:=0;n<b.N;n++{
    reflect.DeepEqual(slice1,slice2)
  }
}
