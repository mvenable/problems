package reverse_or_rotate

func ReverseOrRotate(str string, chunkSize int) string  {
	// While the directions specifically call out checking str being empty that case
	// is already covered by the two checks present.
	if chunkSize <= 0 ||  chunkSize > len(str) {
		return ""
	}


}