package sjis

import (
	"io/ioutil"
	"strings"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
	"golang.org/x/text/width"
)

var defaultSjisEncoder = japanese.ShiftJIS.NewEncoder()

// ConvertUtf8ToSjis converts utf-8 string to Shift_JIS string.
func ConvertUtf8ToSjis(utf8 string) (string, error) {
	if utf8 == "" {
		return "", nil
	}

	strReader := strings.NewReader(utf8)
	sjisReader := transform.NewReader(strReader, defaultSjisEncoder)
	byt, err := ioutil.ReadAll(sjisReader)
	if err != nil {
		return "", err
	}
	return string(byt), nil
}

// ConvertToFullWidth converts half-width character to full-width.
func ConvertToFullWidth(str string) string {
	return width.Widen.String(str)
}
