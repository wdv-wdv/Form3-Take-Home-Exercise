package accountapi

import (
	"fmt"
	"net/url"
)

type AccountAPI struct {
	baseurl string // = "http://localhost:8080"
}

func InitAccountAPI(s string) AccountAPI {
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	if !(u.Scheme == "http" || u.Scheme == "https") {
		panic(`Only "http" & "https" Scheme allowed`)
	}
	if u.RawQuery != "" {
		panic(`No query string allowed`)
	}

	accountAPI := AccountAPI{baseurl: fmt.Sprintf("%s://%s%s", u.Scheme, u.Host, u.Path)}
	return accountAPI
}

// func (acountAPI *AcountAPI) isInit(){
// 	if acountAPI.baseurl == ""
// }
