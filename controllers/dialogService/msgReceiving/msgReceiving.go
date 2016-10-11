package msgReceiving

import (
	"encoding/json"
	// "fmt"

	"github.com/astaxie/beego"

	// "weChat/controllers"
	msgModels "weChat/models/dialogService/msgReceiving"
)

// const ()

// MsgReceivingController implement.
type MsgReceivingController struct {
	beego.Controller
}

// VerifyMsg -
func (msgReceiving *MsgReceivingController) VerifyMsg() {
	ver := &msgModels.Verification{}

	err := json.Unmarshal(msgReceiving.Ctx.Input.RequestBody, ver)
	if err != nil {
		msgReceiving.Data["json"] = err.Error()
		msgReceiving.ServeJSON()
		return
	}

	// str := ver.Nonce + ver.TimeStamp + controllers.AccessToken

	// s := "sha1 this string"
	// //产生一个散列值得方式是 sha1.New()，sha1.Write(bytes)，然后 sha1.Sum([]byte{})。这里我们从一个新的散列开始。
	// h := sha1.New()
	// //写入要处理的字节。如果是一个字符串，需要使用[]byte(s) 来强制转换成字节数组。
	// h.Write([]byte(s))
	// //这个用来得到最终的散列值的字符切片。Sum 的参数可以用来都现有的字符切片追加额外的字节切片：一般不需要要。
	// bs := h.Sum(nil)
	// //SHA1 值经常以 16 进制输出，例如在 git commit 中。使用%x 来将散列结果格式化为 16 进制字符串。
	// fmt.Println(s)
	// fmt.Printf("%x\n", bs)

}

// RecMsg -
func (msgReceiving *MsgReceivingController) RecMsg() {}

// RecEvent -
func (msgReceiving *MsgReceivingController) RecEvent() {}

// RecSRResult -
func (msgReceiving *MsgReceivingController) RecSRResult() {}
