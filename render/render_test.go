package render

import (
	"strings"
	"testing"

	"github.com/XPLassal/simple-snake-on-go/structs"
)

func BenchmarkRender(b *testing.B) {
	n := 20
	snake := structs.NewSnake(n)
	apple := structs.NewApples(n, snake)

	var sb strings.Builder
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		sb.Reset()
		RenderField(n, apple, snake, &sb)
		_ = sb.String()
	}
}
