package accountapi

import (
	"testing"

	"github.com/google/uuid"
)

//### Testing Create operations ###
func Test_Create(t *testing.T) {
	result, links, err := Create(testNewAccount)
	if err != nil {
		t.Fatalf(`Create(newAccount)  error: %v`, err)
	}

	//## Checking results ##
	checkAccount(result, t)

	t.Logf("Result: %v", result)
	t.Logf("Links: %v", links)
}

func Test_CreateWithDuplicateID(t *testing.T) {
	_, _, err := Create(testNewAccount)
	if err == nil {
		t.Fatalf(`Create(newAccount) err should not be nil`)
	}
	t.Logf("Error expected: %v", err)
}

//### Testing Fetch operations ###
func Test_Fetch(t *testing.T) {
	u := testNewAccount.ID
	result, links, err := Fetch(u)

	if err != nil {
		t.Fatalf(`Fetch(uuid)  error: %v`, err)
	}

	//## Checking result ##
	checkAccount(result, t)

	t.Logf("Result: %v", result)
	t.Logf("Links: %v", links)
}

//### Testing Delete operations ###
func Test_Delete(t *testing.T) {
	u := testNewAccount.ID
	err := Delete(u, 0)
	if err != nil {
		t.Fatalf(`Delete(uuid,0)  error: %v`, err)
	}
}

func Test_DeleteRecordNotFound(t *testing.T) {
	u := uuid.New() //get ramdom UUID
	err := Delete(u, 0)
	if err == nil {
		t.Fatalf(`Delete(uuid,0) err should not be nil`)
	}
	t.Logf("Error expected: %v", err)
}

//### Function for testing data integrity ###
func checkAccount(account Account, t *testing.T) {

	//### Checking Account ###
	if account.ID != testNewAccount.ID.String() {
		t.Errorf(`Fetch() ID wrong`)
	}
	if account.OrganisationID != testNewAccount.OrganisationID.String() {
		t.Errorf(`Fetch() OrganisationID wrong`)
	}
	if account.CreatedOn.IsZero() {
		t.Errorf(`Fetch() Created on date issue`)
	}
	if account.ModifiedOn.IsZero() {
		t.Errorf(`Fetch() Modified on date issue`)
	}

	//### Checking Attributes ###
	if account.Attributes.CountryCode != testNewAccount.Attributes.CountryCode {
		t.Errorf(`Fetch() Country wrong`)
	}
	if account.Attributes.BaseCurrency != testNewAccount.Attributes.BaseCurrency {
		t.Errorf(`Fetch() Base Currency wrong`)
	}
	if account.Attributes.BankID != testNewAccount.Attributes.BankID {
		t.Errorf(`Fetch() BankID wrong`)
	}
	if account.Attributes.BankIDCode != testNewAccount.Attributes.BankIDCode {
		t.Errorf(`Fetch() BankIDCode wrong`)
	}
	if account.Attributes.AccountNumber != testNewAccount.Attributes.AccountNumber {
		t.Errorf(`Fetch() AccountNumber wrong`)
	}
	if account.Attributes.BIC != testNewAccount.Attributes.BIC {
		t.Errorf(`Fetch() BIC wrong`)
	}
	if account.Attributes.IBAN != testNewAccount.Attributes.IBAN {
		t.Errorf(`Fetch() IBAN wrong`)
	}
	if account.Attributes.CustomerID != testNewAccount.Attributes.CustomerID {
		t.Errorf(`Fetch() CustomerID wrong`)
	}
	if account.Attributes.Name[0] != testNewAccount.Attributes.Name[0] {
		//t.Errorf(`Fetch() Name wrong`)
		t.Errorf(`The documentation say "title", "firstname", etc  was supersided by "names", however the fake-api doesn't seems to reconsign "names"`)
	}
	if account.Attributes.AlternativeNames[0] != testNewAccount.Attributes.AlternativeNames[0] {
		//t.Errorf(`Fetch() AlternativeNames wrong`)
		t.Errorf(`The documentation say "alternative_bank_account_names" was supersided by "alternative_names", however the fake-api still return "alternative_bank_account_names"`)
	}
	if account.Attributes.AccountClassification != testNewAccount.Attributes.AccountClassification {
		t.Errorf(`Fetch() AccountClassification wrong`)
	}
	if account.Attributes.JointAccount != testNewAccount.Attributes.JointAccount {
		t.Errorf(`Fetch() JointAccount wrong`)
	}
	if account.Attributes.AccountMatchingOptOut != testNewAccount.Attributes.AccountMatchingOptOut {
		t.Errorf(`Fetch() AccountMatchingOptOut wrong`)
	}
	if account.Attributes.SecondaryIdentification != testNewAccount.Attributes.SecondaryIdentification {
		t.Errorf(`Fetch() SecondaryIdentification wrong`)
	}
	if account.Attributes.Switched != testNewAccount.Attributes.Switched {
		t.Errorf(`Fetch() Switched wrong`)
	}
	if account.Attributes.Status != testNewAccount.Attributes.Status {
		//t.Errorf(`Fetch() Status wrong`)
		t.Errorf(`Fake-api doesn't seems to recognise "status"`)
	}
}

//### Test data ###
var testNewAccount = NewAccount{
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
