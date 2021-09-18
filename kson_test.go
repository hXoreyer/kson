package kson

import (
	"testing"
)

func BenchmarkKson(b *testing.B) {
	// TODO: Initialize
	b.ResetTimer()
	obj, _ := KparseByFile("./test.json")
	for i := 0; i < b.N; i++ {
		// TODO: Your Code Here
		_ = obj.GetInt("datas")
		obj.Set("datas", 1234567890)
		obj.Set("arr",
			NewArray().Append(
				NewObject().Set("asd", "bbbccdd"),
				NewObject().Set("bbb", 123)),
		)
	}
	obj.SaveAsFile("./test.json")
}
