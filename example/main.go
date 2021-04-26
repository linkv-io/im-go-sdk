package main

import (
	"fmt"
	imsdk "github.com/linkv-io/im-go-sdk"
)

func main() {
	appKey := "LM6000141252986692189573"
	appSecret := "a436a3757c8b5da179d72dd37d24f81c"
	im := imsdk.NewIM(appKey, appSecret)
	thirdUID := "test-go-tob"
	aID := "test"
	// 进行帐号绑定
	imToken, imOpenID, err := im.GetTokenByThirdUID(thirdUID, aID, "test-go",
		imsdk.SexUnknown, "http://meet.linkv.sg/app/rank-list/static/img/defaultavatar.cd935fdb.png",
		"", "", "")
	if err != nil {
		panic("im.GetTokenByThirdUID(" + err.Error() + ")")
	}
	fmt.Printf("token:%v\topenID:%v\n", imToken, imOpenID)
}
