//Package xrechnung - Library for processing electronic invoices - german xrechnung **/
package xrechnung

//InvoiceStructure structure of the invoice
type InvoiceStructure struct {
	Invoice struct {
		CustomizationID      string `xml:"CustomizationID"`
		ID                   string `xml:"ID"`
		IssueDate            string `xml:"IssueDate"`
		InvoiceTypeCode      int    `xml:"InvoiceTypeCode"`
		Note                 string `xml:"Note"`
		DocumentCurrencyCode string `xml:"DocumentCurrencyCode"`
		TaxCurrencyCode      string `xml:"TaxCurrencyCode"`
		BuyerReference       string `xml:"BuyerReference"`
		//Supplier de:Lieferant/Kreditor
		AccountingSupplierParty struct {
			Party party `xml:"Party"`
		} `xml:"AccountingSupplierParty"`
		//Customer de:Kunde/Debitor
		AccountingCustomerParty struct {
			Party party `xml:"Party"`
		} `xml:"AccountingCustomerParty"`
		PaymentMeans struct {
			PaymentMeansCode      int `xml:"PaymentMeansCode"`
			PayeeFinancialAccount struct {
				ID string `xml:"ID"`
			} `xml:"PayeeFinancialAccount"`
		} `xml:"PaymentMeans"`
		PaymentTerms struct {
			Note string `xml:"Note"`
		} `xml:"PaymentTerms"`
		TaxTotal struct {
			TaxAmountStr string `xml:"TaxAmount"`
			CurrencyID   string `xml:"currencyID,attr"`
			taxAmount    int64
			TaxSubtotal  struct {
				//todo
			} `xml:"TaxSubtotal"`
		} `xml:"TaxTotal"`
	} `xml:"Invoice"`
}

type party struct {
	PartyName struct {
		Name string `xml:"Name"`
	} `xml:"PartyName"`
	PostalAddress struct {
		StreetName string `xml:"StreetName"`
		CityName   string `xml:"CityName"`
		PostalZone string `xml:"PostalZone"`
		Country    struct {
			IdentificationCode string `xml:"IdentificationCode"`
		} `xml:"Country"`
	} `xml:"PostalAddress"`
	PartyTaxScheme struct {
		CompanyID string `xml:"CompanyID"`
		TaxScheme struct {
			ID string `xml:"ID"`
		} `xml:"TaxScheme"`
	} `xml:"PartyTaxScheme"`
	PartyLegalEntity struct {
		RegistrationName string `xml:"RegistrationName"`
		CompanyID        string `xml:"CompanyID"`
		CompanyLegalForm string `xml:"CompanyLegalForm"`
	} `xml:"PartyLegalEntity"`
	Contact struct {
		Name           string `xml:"Name"`
		Telephone      string `xml:"Telephone"`
		ElectronicMail string `xml:"ElectronicMail"`
	} `xml:"Contact"`
}
