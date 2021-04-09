package accountapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/google/uuid"
)

func (accountAPI *AccountAPI) Fetch(accountID uuid.UUID) (Account, Links, error) {

	//### http request to fake-api ###
	url := fmt.Sprintf("%s/organisation/accounts/%s", accountAPI.baseurl, accountID.String())
	res, err := http.Get(url)
	if err != nil {
		return Account{}, Links{}, err
	}

	//### read response body ###
	defer res.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(res.Body)

	//### process statuscode ###
	if res.StatusCode == 400 {
		brError := BadRequestError{}
		json.Unmarshal(bodyBytes, &brError)
		return Account{}, Links{}, fmt.Errorf("Error while fetching Account - Bad Request. ErrorMessage: %s", brError.ErrorMessage)
	} else if res.StatusCode != 200 {
		return Account{}, Links{}, fmt.Errorf("Error while fetching Account. Status: %s.", res.Status)
	}

	//### process results ###
	result := struct {
		Data  Account `json:"data"`
		Links Links   `json:"links"`
	}{}
	err = json.Unmarshal(bodyBytes, &result)

	if err != nil {
		return Account{}, Links{}, err
	}

	//### return results ###
	return result.Data, result.Links, nil
}
