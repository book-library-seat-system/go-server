package util

// CheckErr panic错误
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
