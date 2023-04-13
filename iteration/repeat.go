package iteration

const repeatCount = 5

// Repeat 引数を5回繰り返したものを返す関数.
func Repeat(character string) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return repeated
}
