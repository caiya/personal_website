package utils

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"strconv"
)

//32位md5加密
func Md5(str string) string {
	ctx := md5.New()
	ctx.Write([]byte(str))
	rs := hex.EncodeToString(ctx.Sum(nil))
	return rs
}

//生成指定位数的随机数字字符串
func GetRandomNum(size int) string {
	s := ""
	for i := 0; i < size; i++ {
		s += strconv.FormatInt(rand.Int63n(10), 10)
	}
	return s
}
