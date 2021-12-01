package benchmarks

import "testing"

func BenchmarkCalculate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Calculate(2)
	}
}

/*func BenchmarkSample(b *testing.B) {
	var vb = []byte("12345")
	for i := 0; i < b.N; i++ {
		if x := fmt.Sprintf("%s", vb); x != "12345" {
			b.Fatalf("Unexpected string: %s", x)
		}
	}
}*/
