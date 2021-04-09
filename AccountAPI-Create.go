package accountapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (accountAPI *AccountAPI) Create(r NewAccount) (Account, Links, error) {

	//### Generate request body payload ###
	j, _ := json.Marshal(struct {
		Data newAccount `json:"data"`
	}{

		Data: newAccount{
			"accounts",
			r.ID.String(),
			r.OrganisationID.String(),
			r.Attributes,
		},
	})

	requestBody := bytes.NewBuffer(j)

	//### http request to fake-api ###
	res, err := http.Post(
		fmt.Sprintf("%s/organisation/accounts", accountAPI.baseurl), //"http://localhost:8080",
		"Content-Type: application/vnd.api+json",
		requestBody,
	)

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
		return Account{}, Links{}, fmt.Errorf("Error while creating new Account - Bad Request. ErrorMessage: %s", brError.ErrorMessage)
	} else if res.StatusCode != 201 {
		return Account{}, Links{}, fmt.Errorf("Error while creating new Account. Status: %s.", res.Status)
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
