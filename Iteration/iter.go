package iteration

func Repeat(character string, numberRepeat int) string {
	var repeated string
	for i := 0; i < numberRepeat; i++ {
		repeated += character
	}
	return repeated
}
