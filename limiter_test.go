package limiter

import (
	"testing"
)

func BenchmarkRunSmall(b *testing.B) {
	l := New(10)
	for i := 0; i < b.N; i++ {
		for j := 0; j < 1_000; j++ {
			l.Run(func() {
				// code to be executed with the limiter
			})
		}
		l.Wait()
	}

}

func BenchmarkRunWaitMedium(b *testing.B) {
	l := New(100)
	for i := 0; i < b.N; i++ {
		for j := 0; j < 1_000; j++ {
			l.Run(func() {
				// code to be executed with the limiter
			})
		}
		l.Wait()
	}

}

func BenchmarkWaitRunLarge(b *testing.B) {
	l := New(500)
	for i := 0; i < b.N; i++ {
		for j := 0; j < 1_000; j++ {
			l.Run(func() {
				// code to be executed with the limiter
			})
		}
		l.Wait()
	}
}
