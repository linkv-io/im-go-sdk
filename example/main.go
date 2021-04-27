package main

import (
	"fmt"
	imsdk "github.com/linkv-io/im-go-sdk"
)

func main() {
	secret := ""
	im, err := imsdk.NewIM(secret)
	if err != nil {
		panic(fmt.Errorf("imsdk.NewIM(secret) error(%v)", err))
	}

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

	toUID := "1100"
	objectName := "RC:textMsg"
	content := "测试单聊"
	fmt.Println(im.PushConverseData(thirdUID, toUID, objectName, content, "", "", "", "", "", ""))

	content = "测试 事件"
	fmt.Println(im.PushEventData(thirdUID, toUID, objectName, content, "", "", "", ""))
}
