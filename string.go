package item

import (
	"bytes"
	"math/rand"
	"reflect"
	"sync"
	"time"
	"unsafe"
)

var buffPool = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}

// 获取随机字符串
var randStringLowBase = []byte("0123456789abcdefghijklmnopqrstuvwxyz")

func GetRandomString(cnum int) string {
	bs := randStringLowBase
	result := buffPool.Get().(*bytes.Buffer)
	defer func() {
		result.Reset()
		buffPool.Put(result)
	}()
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < cnum; i++ {
		result.WriteByte(bs[r.Intn(len(bs))])
	}

	return result.String()
}

var randStringNumber = []byte("1234567890")

// GetRandomNumString 获取数字字符串
func GetRandomNumString(cnum int) string {
	bs := randStringNumber
	result := buffPool.Get().(*bytes.Buffer)
	defer func() {
		result.Reset()
		buffPool.Put(result)
	}()
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < cnum; i++ {
		result.WriteByte(bs[r.Intn(len(bs))])
	}

	return result.String()
}

// GetRandomNumInt 获取数字字符串
func GetRandomNumInt(cnum int) int {
	bs := randStringNumber
	var result int
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < cnum; i++ {
		result += result*10 + r.Intn(len(bs))
	}

	return result
}

var randStringUpBase = []byte("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYX")

func GetRandomStringUp(cnum int) string {
	bs := randStringUpBase
	result := buffPool.Get().(*bytes.Buffer)
	defer func() {
		result.Reset()
		buffPool.Put(result)
	}()
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < cnum; i++ {
		result.WriteByte(bs[r.Intn(len(bs))])
	}

	return result.String()
}

func CamelToCase(str string) string {
	result := buffPool.Get().(*bytes.Buffer)
	defer func() {
		result.Reset()
		buffPool.Put(result)
	}()
	for i, p := range StringToBytesUnsafe(str) {
		if p > 64 && p < 91 {
			p = p + 32
			if i != 0 {
				result.WriteByte('_')
			}
		}
		result.WriteByte(p)
	}
	return result.String()
}

func CaseToCamel(str string) string {
	result := buffPool.Get().(*bytes.Buffer)
	defer func() {
		result.Reset()
		buffPool.Put(result)
	}()
	ucFlag := true
	for _, p := range StringToBytesUnsafe(str) {
		if p == '_' {
			ucFlag = true
			continue
		}
		if p > 96 && p < 123 {
			if ucFlag {
				p = p - 32
				ucFlag = false
			}
		}
		result.WriteByte(p)
	}
	return result.String()
}

// StringToBytesUnsafe 快速转换 但是不能用于并发场景
func StringToBytesUnsafe(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}

// BytesToString 快速转换 但是不能用于并发场景
func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
