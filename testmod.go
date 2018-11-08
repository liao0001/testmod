package testmod

import (
	"github.com/mozillazg/go-pinyin"
	"strings"
)

func GetPinyin(str string,split string) string {
	ss:=pinyin.LazyConvert(str,nil)
	return strings.Join(ss,split)
}
