package accountapi

import (
	"testing"

	"github.com/google/uuid"
)

//### Testing Constructor ###
func Test_Constructor(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Fatalf("InitAccountAPI(s) Error: %v", err)
		}
	}()
	accountAPI = InitAccountAPI("http://localhost:8080/v1")
}

func Test_ConstructorSchemaWrong(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Logf("Error expected: %v", err)
		} else {
			t.Fatalf(`InitAccountAPI(s) err should not be nil`)
		}
	}()
	accountAPI = InitAccountAPI("ftp://localhost:8080")
}

func Test_ConstructorQSWrong(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Logf("Error expected: %v", err)
		} else {
			t.Fatalf(`InitAccountAPI(s) err should not be nil`)
		}
	}()
	accountAPI = InitAccountAPI("http://localhost:8080/?hhh")
}

//### Testing Create operations ###
func Test_Create(t *testing.T) {
	//## setup up the mock - newAccount data ##
	newAccount := newTestAccount()

	//## The test ##
	result, links, err := accountAPI.Create(newAccount)
	if err != nil {
		t.Fatalf(`Create(newAccount)  error: %v`, err)
	}

	//## Checking results ##
	checkAccount(newAccount, result, t)

	t.Logf("Result: %v", result)
	t.Logf("Links: %v", links)
}

func Test_CreateWithDuplicateID(t *testing.T) {
	//## setup up the mock - newAccount data & duplication record ##
	newAccount := newTestAccount()
	accountAPI.Create(newAccount) //Create record which duplication ID will be used later

	//## The test ##
	_, _, err := accountAPI.Create(newAccount)
	if err == nil {
		t.Fatalf(`Create(newAccount) err should not be nil`)
	}
	t.Logf("Error expected: %v", err)
}

//### Testing Fetch operations ###
func Test_Fetch(t *testing.T) {
	//## setup up the mock - newAccount data & Create record which will be fetched ##
	newAccount := newTestAccount()
	accountAPI.Create(newAccount)
	u := newAccount.ID

	//## The test ##
	result, links, err := accountAPI.Fetch(u)

	if err != nil {
		t.Fatalf(`Fetch(uuid)  error: %v`, err)
	}

	//## Checking result ##
	checkAccount(newAccount, result, t)

	t.Logf("Result: %v", result)
	t.Logf("Links: %v", links)
}

//### Testing Delete operations ###
func Test_Delete(t *testing.T) {
	//## setup up the mock - newAccount data & create record to be delete ##
	newAccount := newTestAccount()
	accountAPI.Create(newAccount)
	u := newAccount.ID

	//## The test
	err := accountAPI.Delete(u, 0)
	if err != nil {
		t.Fatalf(`Delete(uuid,0)  error: %v`, err)
	}
}

func Test_DeleteRecordNotFound(t *testing.T) {
	//## setup up the mock - get ramdom UUID
	u := uuid.New()

	//## The test
	err := accountAPI.Delete(u, 0)
	if err == nil {
		t.Fatalf(`Delete(uuid,0) err should not be nil`)
	}
	t.Logf("Error expected: %v", err)
}

//### Function for testing data integrity ###
func checkAccount(newAccount NewAccount, account Account, t *testing.T) {

	//### Checking Account ###
	if account.ID != newAccount.ID.String() {
		t.Errorf(`Fetch() ID wrong`)
	}
	if account.OrganisationID != newAccount.OrganisationID.String() {
		t.Errorf(`Fetch() OrganisationID wrong`)
	}
	if account.CreatedOn.IsZero() {
		t.Errorf(`Fetch() Created on date issue`)
	}
	if account.ModifiedOn.IsZero() {
		t.Errorf(`Fetch() Modified on date issue`)
	}

	//### Checking Attributes ###
	if account.Attributes.CountryCode != newAccount.Attributes.CountryCode {
		t.Errorf(`Fetch() Country wrong`)
	}
	if account.Attributes.BaseCurrency != newAccount.Attributes.BaseCurrency {
		t.Errorf(`Fetch() Base Currency wrong`)
	}
	if account.Attributes.BankID != newAccount.Attributes.BankID {
		t.Errorf(`Fetch() BankID wrong`)
	}
	if account.Attributes.BankIDCode != newAccount.Attributes.BankIDCode {
		t.Errorf(`Fetch() BankIDCode wrong`)
	}
	if account.Attributes.AccountNumber != newAccount.Attributes.AccountNumber {
		t.Errorf(`Fetch() AccountNumber wrong`)
	}
	if account.Attributes.BIC != newAccount.Attributes.BIC {
		t.Errorf(`Fetch() BIC wrong`)
	}
	if account.Attributes.IBAN != newAccount.Attributes.IBAN {
		t.Errorf(`Fetch() IBAN wrong`)
	}
	if account.Attributes.CustomerID != newAccount.Attributes.CustomerID {
		t.Errorf(`Fetch() CustomerID wrong`)
	}
	if account.Attributes.Name[0] != newAccount.Attributes.Name[0] {
		//t.Errorf(`Fetch() Name wrong`)
		t.Errorf(`The documentation say "title", "firstname", etc  was supersided by "names", however the fake-api doesn't seems to reconsign "names"`)
	}
	if account.Attributes.AlternativeNames[0] != newAccount.Attributes.AlternativeNames[0] {
		//t.Errorf(`Fetch() AlternativeNames wrong`)
		t.Errorf(`The documentation say "alternative_bank_account_names" was supersided by "alternative_names", however the fake-api still return "alternative_bank_account_names"`)
	}
	if account.Attributes.AccountClassification != newAccount.Attributes.AccountClassification {
		t.Errorf(`Fetch() AccountClassification wrong`)
	}
	if account.Attributes.JointAccount != newAccount.Attributes.JointAccount {
		t.Errorf(`Fetch() JointAccount wrong`)
	}
	if account.Attributes.AccountMatchingOptOut != newAccount.Attributes.AccountMatchingOptOut {
		t.Errorf(`Fetch() AccountMatchingOptOut wrong`)
	}
	if account.Attributes.SecondaryIdentification != newAccount.Attributes.SecondaryIdentification {
		t.Errorf(`Fetch() SecondaryIdentification wrong`)
	}
	if account.Attributes.Switched != newAccount.Attributes.Switched {
		t.Errorf(`Fetch() Switched wrong`)
	}
	if account.Attributes.Status != newAccount.Attributes.Status {
		//t.Errorf(`Fetch() Status wrong`)
		t.Errorf(`Fake-api doesn't seems to recognise "status"`)
	}
}

//### Test data ###
var accountAPI = InitAccountAPI("http://localhost:8080/v1")

func newTestAccount() NewAccount {
	newAccount := NewAccount{
		uuid.New(),
		uuid.MustParse("eb0bd6f5-c3f5-44b2-b677-acd23cdde73c"), //OrgID
		Attributes{
			"GB",                               //Country
			"GBP",                              //Base Currency
			"123456",                           //BankID
			"GBDSC",                            //BankIDCode
			"3456789",                          //AccountNumber
			"NWBKGB22",                         //BIC
			"GB11NWBK40030041426819",           //IBAN
			"ABC123",                           //CustomerID
			[4]string{"Firstname", "Lastname"}, //Name
			[3]string{"Confirmation of Payee"}, //AlternativeNames
			"Personal",                         //AccountClassification
			true,                               //JointAccount
			true,                               //AccountMatchingOptOut
			"ABC123ABC",                        //SecondaryIdentification
			false,                              //Switched
			"confirmed",                        //Status
		},
	}
	return newAccount
}
