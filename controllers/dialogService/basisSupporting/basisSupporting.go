package basisSupporting

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/astaxie/beego"

	"weChat/controllers"
	basisModels "weChat/models/dialogService/basisSupporting"
)

const (
	// URLAccessToken holds url to get access-token.
	URLAccessToken = "https://api.weixin.qq.com/cgi-bin/token"
	// URLServerIP holds url to get server-IP.
	URLServerIP = "https://api.weixin.qq.com/cgi-bin/getcallbackip"
)

// BasisController implements basis-supporting.
type BasisController struct {
	beego.Controller
}

// GetAccessToken gets access-token.
func (basis *BasisController) GetAccessToken() {
	accessToken := &basisModels.AccessToken{}

	err := json.Unmarshal(basis.Ctx.Input.RequestBody, accessToken)
	if err != nil {
		fmt.Println("json.Unmarshal():", err)
		basis.Data["json"] = err.Error()
		basis.ServeJSON()
		return
	}

	// request method: http.Get
	resp, err := http.Get(URLAccessToken + "?grant_type=" + accessToken.GrantType +
		"&appid=" + accessToken.AppID + "&secret=" + accessToken.Secret)
	if err != nil {
		basis.Data["json"] = err.Error()
		basis.ServeJSON()
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		basis.Data["json"] = err.Error()
		basis.ServeJSON()
		return
	}

	returnValue := &basisModels.AccessTokenReturnValue{}
	err = json.Unmarshal(body, returnValue)
	if err == nil {
		controllers.AccessToken = returnValue.AccessToken
		basis.Data["json"] = returnValue
		basis.ServeJSON()
		return
	}

	errMsg := &basisModels.ErrMsg{}
	err = json.Unmarshal(body, errMsg)
	if err == nil {
		basis.Data["json"] = errMsg
		basis.ServeJSON()
		return
	}

	basis.Data["json"] = err.Error()
	basis.ServeJSON()
	return
}

// GetServerIP gets server-IP.
func (basis *BasisController) GetServerIP() {
	serverIP := &basisModels.ServerIP{}

	err := json.Unmarshal(basis.Ctx.Input.RequestBody, serverIP)
	if err != nil {
		basis.Data["json"] = err.Error()
		basis.ServeJSON()
		return
	}

	// request method: http.Get
	resp, err := http.Get(URLServerIP + "?access_token=" + serverIP.AccessToken)
	if err != nil {
		basis.Data["json"] = err.Error()
		basis.ServeJSON()
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		basis.Data["json"] = err.Error()
		basis.ServeJSON()
		return
	}

	returnValue := &basisModels.ServerIPReturnValue{}
	err = json.Unmarshal(body, returnValue)
	// fmt.Println("returnValue:", returnValue)
	fmt.Println("num:", len(returnValue.IPList))
	if err == nil {
		basis.Data["json"] = returnValue
		basis.ServeJSON()
		return
	}

	errMsg := &basisModels.ErrMsg{}
	err = json.Unmarshal(body, errMsg)
	fmt.Println("errMsg:", errMsg)
	if err == nil {
		basis.Data["json"] = errMsg
		basis.ServeJSON()
		return
	}

	basis.Data["json"] = err.Error()
	basis.ServeJSON()
}
