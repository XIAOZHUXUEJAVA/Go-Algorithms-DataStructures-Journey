package others

import (
	"reflect"
	"testing"
)

func TestEncrypt(t *testing.T) {
	plainText := "abC"
	cipherText := encrypt(plainText, 3)

	if !reflect.DeepEqual("deF", cipherText) {
		t.Errorf("expected: %s, output: %s", "deF", cipherText)
	}
}
