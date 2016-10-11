package basisSupporting

// AccessToken holds parameters to get access-token.
type AccessToken struct {
	GrantType string `json:"grant_type"`
	AppID     string `json:"appid"`
	Secret    string `json:"secret"`
}

// AccessTokenReturnValue holds return-value.
type AccessTokenReturnValue struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

// ServerIP holds parameters to get server-IP.
type ServerIP struct {
	AccessToken string `json:"access_token"`
}

// ServerIPReturnValue holds return-value.
type ServerIPReturnValue struct {
	IPList []string `json:"ip_list"`
}

// ErrMsg holds error-message.
type ErrMsg struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}
