package im_go_sdk

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/linkv-io/im-go-sdk/http"
	"net/url"
	"sort"
	"strings"
)

func (o *im) PushConverseData(fromUID, toUID, objectName, content, pushContent, pushData, deviceID, toAppID, toUserExtSysUserID, isCheckSensitiveWords string) (bool, error) {
	nonce := genGUID()
	timestamp := getTimestampS()

	arr := []string{nonce, timestamp, o.imAppSecret}
	sort.Strings(arr)
	md5Data := md5.Sum([]byte(strings.Join(arr, "")))
	cmimToken := strings.ToLower(hex.EncodeToString(md5Data[:]))
	sha1Data := sha1.Sum([]byte(o.imAppID + "|" + o.imAppKey + "|" + timestamp + "|" + nonce))
	sign := strings.ToUpper(hex.EncodeToString(sha1Data[:]))

	headers := make(map[string]string)
	headers["nonce"] = nonce
	headers["timestamp"] = timestamp
	headers["cmimToken"] = cmimToken
	headers["sign"] = sign
	headers["appkey"] = o.imAppKey
	headers["appId"] = o.imAppID
	headers["appUid"] = fromUID

	params := url.Values{}
	params.Set("fromUserId", fromUID)
	params.Set("toUserId", toUID)
	params.Set("objectName", objectName)
	params.Set("content", content)
	params.Set("appId", o.imAppID)

	if len(pushContent) > 0 {
		params.Set("pushContent", pushContent)
	}

	if len(pushData) > 0 {
		params.Set("pushData", pushData)
	}
	if len(deviceID) > 0 {
		params.Set("deviceId", deviceID)
	}
	if len(toAppID) > 0 {
		params.Set("toUserAppid", toAppID)
	}

	if len(toUserExtSysUserID) > 0 {
		params.Set("toUserExtSysUserId", toUserExtSysUserID)
	}

	if len(isCheckSensitiveWords) > 0 {
		params.Set("isCheckSensitiveWords", isCheckSensitiveWords)
	}
	uri := o.imHost + "/api/rest/message/converse/pushConverseData"

	var errResult error

	for i := 0; i < 3; i++ {

		jsonData, resp, err := http.PostDataWithHeader(uri, params, headers)
		if err != nil {
			return false, err
		}

		if resp.StatusCode != 200 {
			return false, fmt.Errorf("httpStatusCode(%v) != 200", resp.StatusCode)
		}

		var result struct {
			Code int `json:"code"`
		}

		if err := json.Unmarshal(jsonData, &result); err != nil {
			return false, err
		}

		if result.Code == 200 {
			return true, nil
		}

		return false, fmt.Errorf("code not 200(%v)", result.Code)
	}
	return false, errResult
}
