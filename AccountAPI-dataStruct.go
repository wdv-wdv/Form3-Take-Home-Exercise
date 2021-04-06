package accountapi

import (
	"time"

	"github.com/google/uuid"
)

type NewAccount struct {
	ID             uuid.UUID
	OrganisationID uuid.UUID
	Attributes     Attributes
}

type newAccount struct {
	Type           string     `json:"type"`
	ID             string     `json:"id"`
	OrganisationID string     `json:"organisation_id"`
	Attributes     Attributes `json:"attributes"`
}

type Account struct {
	CreatedOn      time.Time  `json:"created_on"`
	ID             string     `json:"id"`
	ModifiedOn     time.Time  `json:"modified_on"`
	OrganisationID string     `json:"organisation_id"`
	Version        int        `json:"version"`
	Attributes     Attributes `json:"attributes"`
}

type Attributes struct {
	CountryCode             string    `json:"country"`
	BaseCurrency            string    `json:"base_currency,omitempty"`
	BankID                  string    `json:"bank_id,omitempty"`
	BankIDCode              string    `json:"bank_id_code,omitempty"`
	AccountNumber           string    `json:"account_number"`
	BIC                     string    `json:"bic,omitempty"`
	IBAN                    string    `json:"iban"`
	CustomerID              string    `json:"customer_id"`
	Name                    [4]string `json:"name"`
	AlternativeNames        [3]string `json:"alternative_names"`
	AccountClassification   string    `json:"account_classification"`
	JointAccount            bool      `json:"joint_account"`
	AccountMatchingOptOut   bool      `json:"account_matching_opt_out"`
	SecondaryIdentification string    `json:"secondary_identification"`
	Switched                bool      `json:"switched"`
	Status                  string    `json:"status"`
}

type Links struct {
	Self  string `json:"self"`
	First string `json:"first"`
	Last  string `json:"last"`
	Next  string `json:"next"`
	Prev  string `json:"prev"`
}

type BadRequestError struct {
	ErrorMessage string `json:"error_message"`
	ErrorCode    string `json:"error_code"`
}
