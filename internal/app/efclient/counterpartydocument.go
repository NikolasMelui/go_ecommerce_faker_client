package efclient

// CounterpartyDocument ...
type CounterpartyDocument struct {
	ID           int          `json:"id"`
	Counterparty Counterparty `json:"counterparty"`
	Name         string       `json:"name"`
	Description  string       `json:"description"`
	URL          string       `json:"url"`
}

// CounterpartyDocumentData ...
type CounterpartyDocumentData struct {
	Counterparty Counterparty `json:"counterparty"`
	Name         string       `json:"name"`
	Description  string       `json:"description"`
	URL          string       `json:"url"`
}
