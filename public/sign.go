package public

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

const base64Table = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

var coder = base64.NewEncoding(base64Table)

/**
 * 参数生成
 */
func MakeParams(uid, rid, timestamp string) (params_str, sign_str string) {
	var s, p string
	b := bytes.Buffer{}
	appkey := getAppkey()
	b.WriteString("uid=")
	b.WriteString(uid)
	b.WriteString("&rid=")
	b.WriteString(rid)
	b.WriteString("&timestamp=")
	b.WriteString(timestamp)
	p = b.String()
	b.WriteString("&appkey=")
	b.WriteString(appkey)
	s = b.String()
	return p, s
}

func MakeSign(sign_str string) (str string) {
	h := md5.New()
	h.Write([]byte(sign_str)) // 需要加密的字符串为 uid=uid&rid=rid&timestamp=timestamp&appkey=appkey
	md5Byte := h.Sum(nil)     //byte 类型
	fmt.Println(md5Byte)
	md5Str := fmt.Sprintf("%s", hex.EncodeToString(md5Byte)) // 输出加密结果
	return md5Str
}

/**
 * appkey 声明
 */
func getAppkey() string {
	var appkey string = "FD24013D68D6DD310154E3E311B4A6A4"
	return appkey
}

func base64Encode(encode_byte []byte) []byte {
	return []byte(coder.EncodeToString(encode_byte))
}

func base64Decode(decode_byte []byte) ([]byte, error) {
	return coder.DecodeString(string(decode_byte))
}
