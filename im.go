package im_go_sdk

import (
	"bytes"
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"math/big"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
)

type SexType int
type PlatformType string

var (
	OrderTypeAdd int64 = 1
	OrderTypeDel int64 = 2

	SexUnknown SexType = -1
	SexFemale  SexType = 0
	SexMale    SexType = 1

	PlatformH5      PlatformType = "h5"
	PlatformANDROID PlatformType = "android"
	PlatformIOS     PlatformType = "ios"

	waitTime = time.Millisecond * 300
)

func NewIM(secret string) (*im, error) {
	dst, err := base64.RawStdEncoding.DecodeString(secret)
	if err != nil {
		panic("secret error")
	}
	var ss struct {
		AppKey      string `json:"app_key"`
		AppSecret   string `json:"app_secret"`
		IMAppID     string `json:"im_app_id"`
		IMAppKey    string `json:"im_app_key"`
		IMAppSecret string `json:"im_app_secret"`
		IMHost      string `json:"im_host"`
	}
	if err := json.Unmarshal(dst, &ss); err != nil {
		return nil, err
	}
	return &im{
		ss.AppKey,
		ss.AppSecret,
		ss.IMAppID,
		ss.IMAppKey,
		ss.IMAppSecret,
		ss.IMHost,
	}, nil
}

type im struct {
	appKey      string
	appSecret   string
	imAppID     string
	imAppKey    string
	imAppSecret string
	imHost      string
}

func genUniqueIDString(appKey string) string {
	nLen := 9
	container := string([]byte(appKey)[2:]) + "-"
	var str = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	b := bytes.NewBufferString(str)
	length := b.Len()
	bigInt := big.NewInt(int64(length))
	for i := 0; i < nLen; i++ {
		randomInt, _ := rand.Int(rand.Reader, bigInt)
		container += string(str[randomInt.Int64()])
	}
	return container
}

func genRandomString() string {
	nLen := 16
	var container string
	var str = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	b := bytes.NewBufferString(str)
	length := b.Len()
	bigInt := big.NewInt(int64(length))
	for i := 0; i < nLen; i++ {
		randomInt, _ := rand.Int(rand.Reader, bigInt)
		container += string(str[randomInt.Int64()])
		if i == 7 {
			container += strconv.FormatInt(time.Now().Unix(), 10)
		}
	}
	return container
}

func genSign(params url.Values, md5Secret string) string {
	data := encode(params) + "&key=" + md5Secret
	md5Data := md5.Sum([]byte(data))
	return strings.ToLower(hex.EncodeToString(md5Data[:]))
}

func encode(v url.Values) string {
	if v == nil {
		return ""
	}
	var buf strings.Builder
	keys := make([]string, 0, len(v))
	for k := range v {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		vs := v[k]
		if buf.Len() > 0 {
			buf.WriteByte('&')
		}
		buf.WriteString(k)
		buf.WriteByte('=')
		buf.WriteString(vs[0])
	}
	return buf.String()
}

func randomString(nLen int) string {
	var container string
	var str = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	b := bytes.NewBufferString(str)
	length := b.Len()
	bigInt := big.NewInt(int64(length))
	for i := 0; i < nLen; i++ {
		randomInt, _ := rand.Int(rand.Reader, bigInt)
		container += string(str[randomInt.Int64()])
	}
	return container
}

func genGUID() string {
	return randomString(9) + "-" + randomString(4) + "-" + randomString(4) + "-" + randomString(12)
}

func getTimestampS() string {
	t := time.Now()
	return strconv.FormatInt(t.Unix(), 10)
}

func getTimestampMS() string {
	t := time.Now()
	return strconv.FormatInt(t.Unix()*1000+t.UnixNano(), 10)
}
