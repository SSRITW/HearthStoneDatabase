package restgo

import (
	"math/rand"
	"time"
	"crypto/md5"
	"encoding/hex"
)

func GetPageOffset(pageSize int ,pageNum int)(offset int) {
	if pageSize!=0 || pageNum!=0 {
		offset = pageSize * (pageNum-1)
	}
	return
}

// 生成32位MD5
func MD5(text string) string{
	ctx := md5.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
}

//获取八位随机数做盐值
func GetRandomSalt() string {
	return GetRandomString(8)
}

//生成随机字符串
func GetRandomString(length int) string{
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
