package accountapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/google/uuid"
)

func (accountAPI *AccountAPI) Delete(accountID uuid.UUID, version int) error {

	//### http request to fake-api ###
	url := fmt.Sprintf("%s/organisation/accounts/%s?version=%v", accountAPI.baseurl, accountID.String(), version)
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		return err
	}

	//### read response body ###
	defer res.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(res.Body)

	//### process statuscode ###
	if res.StatusCode == 400 {
		brError := BadRequestError{}
		json.Unmarshal(bodyBytes, &brError)
		return fmt.Errorf("Error while deleting Account - Bad Request. ErrorMessage: %s", brError.ErrorMessage)
	} else if res.StatusCode != 204 {
		return fmt.Errorf("Error while deleting Account. Status: %s.", res.Status)
	}

	//### return nil on success ###
	return nil
}
