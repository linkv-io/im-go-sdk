package main

import (
	"fmt"
	imsdk "github.com/linkv-io/im-go-sdk"
)

func main() {
	appKey := "xxxxx"
	appSecret := "xxxx"
	im := imsdk.NewIM(appKey, appSecret)
	thirdUID := "test-go-tob"
	aID := "test"
	// 进行帐号绑定
	imToken, imOpenID, err := im.GetTokenByThirdUID(thirdUID, aID, "test-go",
		imsdk.SexUnknown, "http://xxxxx/app/rank-list/static/img/defaultavatar.cd935fdb.png",
		"", "", "")
	if err != nil {
		panic("im.GetTokenByThirdUID(" + err.Error() + ")")
	}
	fmt.Printf("token:%v\topenID:%v\n", imToken, imOpenID)
}
