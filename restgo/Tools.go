package restgo

import (
	"math/rand"
	"time"
	"crypto/md5"
	"encoding/hex"
	"github.com/satori/go.uuid"
)

func GetPageOffset(pageSize int ,pageNum int)(offset int) {
	if pageSize!=0 || pageNum!=0 {
		offset = pageSize * (pageNum-1)
	}
	return
}

func GetToken()(token string){
	uuid,_ := uuid.NewV4()
	return MD5(uuid.String())
}

// 生成32位MD5
func MD5(text string) (result string){
	ctx := md5.New()
	ctx.Write([]byte(text))
	result = hex.EncodeToString(ctx.Sum(nil))
	return
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
