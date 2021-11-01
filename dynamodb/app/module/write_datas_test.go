package module

import "testing"

func Test_WriteBySize(t *testing.T) {
	var sizeMB float64 = 1
	WriteBySizeMB(sizeMB)
}

func Test_WriteTestDatas(t *testing.T) {
	WriteDatas()
}
