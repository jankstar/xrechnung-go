//Package xrechnung - Library for processing electronic invoices - german xrechnung 2.0 **/
package xrechnung

import (
	"encoding/xml"
	"errors"
	"fmt"
)

const (
	//Tasks
	CheckAndConvert = "CheckAndConvert"
)

type InvoiceDataElement struct {
	Name        string
	ID          string
	Value       string
	Code        TCode
	DataType    string
	Description string
	Error       error
}

//XMLConvertStructure Converts the XML string data into structural data and checks fields
func (me Invoice2) XMLConvertStructure() (HeadFields []InvoiceDataElement, err error) {
	// check whether CheckAndConvert has already run
	for _, e := range me.task {
		if e == CheckAndConvert {
			err = errors.New("CheckAndConvert already run")
			return
		}
	}

	//-----------------------------------------------------//
	lElement := InvoiceDataElement{
		Name:     "Invoice number",
		ID:       "BT-1",
		Value:    me.ID.Value,
		Code:     TCode{},
		DataType: "Identifier",
		Description: `Eine eindeutige Kennung der Rechnung, die diese im System des Verkäufers identifiziert. \
		Anmerkung: Es ist kein „identification scheme“ zu verwenden.`,
	}
	if lElement.Value == "" {
		lElement.Error = errors.New(fmt.Sprintf("Error field '%s' no value", lElement.Name))
	}
	HeadFields = append(HeadFields, lElement)

	//-----------------------------------------------------//
	lElement = InvoiceDataElement{
		Name:        "Invoice issue date",
		ID:          "BT-2",
		Value:       me.IssueDate.Value,
		Code:        TCode{},
		DataType:    "Date",
		Description: `Das Datum, an dem die Rechnung ausgestellt wurde.`,
	}
	if lElement.Value == "" {
		lElement.Error = errors.New(fmt.Sprintf("Error field '%s' no value", lElement.Name))
	}
	HeadFields = append(HeadFields, lElement)

	//-----------------------------------------------------//
	lElement = InvoiceDataElement{
		Name:     "Invoice type code",
		ID:       "BT-3",
		Value:    me.InvoiceTypeCode.Value,
		Code:     TCode{},
		DataType: "Code",
		Description: `Ein Code, der den Funktionstyp der Rechnung angibt. \
		Anmerkung: Der Rechnungstyp muss gemäß UNTDID 1001 spezifiziert werden. \
		Folgende Codes aus der Codeliste sollen verwendet werden: \
		• 326 (Partial invoice) \
		• 380 (Commercial invoice) \
		• 384 (Corrected invoice) \
		• 389 (Self-billed invoice) \
		• 381 (Credit note) \
		• 875 (Partial construction invoice) \
		• 876 (Partial final construction invoice) \
		• 877 (Final construction invoice)`,
	}
	if lElement.Value == "" {
		lElement.Error = errors.New(fmt.Sprintf("Error field '%s' no value", lElement.Name))
	} else {
		for _, e := range CInvoiceCypeCode {
			if e.ID == lElement.Value {
				lElement.Code = e
				break
			}
		}
		if lElement.Code.Text == "" {
			lElement.Error = errors.New(fmt.Sprintf("Error field '%s' code mapping not valid", lElement.Name))
		}
	}
	HeadFields = append(HeadFields, lElement)

	//-----------------------------------------------------//
	lElement = InvoiceDataElement{
		Name:     "Invoice currency code",
		ID:       "BT-5",
		Value:    me.DocumentCurrencyCode.Value,
		Code:     TCode{},
		DataType: "Code",
		Description: `Die Währung, in der alle Rechnungsbeträge angegeben werden, ausgenommen ist der Umsatzsteuer- \
		Gesamtbetrag, der in der Abrechnungswährung anzugeben ist. \
		Anmerkung: Nur eine Währung ist in der Rechnung zu verwenden, der „Invoice total VAT amount in accounting \
		currency“ (BT-111) ist in der Abrechnungswährung auszuweisen. Die gültigen Währungen sind bei der ISO 4217 \
		„Codes for the representation of currencies and funds“ registriert. Nur die Alpha-3-Darstellung darf verwendet \
		werden.`,
	}
	if lElement.Value == "" {
		lElement.Error = errors.New(fmt.Sprintf("Error field '%s' no value", lElement.Name))
	} else {
		for _, e := range CCurrency {
			if e.CurrencyID == lElement.Value {
				lElement.Code.ID = e.CurrencyID
				lElement.Code.Text = e.CurrencyName
				break
			}
		}
		if lElement.Code.Text == "" {
			lElement.Error = errors.New(fmt.Sprintf("Error field '%s' code mapping not valid", lElement.Name))
		}
	}
	HeadFields = append(HeadFields, lElement)

	//-----------------------------------------------------//
	lElement = InvoiceDataElement{
		Name:     "VAT accounting currency code",
		ID:       "BT-6",
		Value:    me.TaxCurrencyCode.Value,
		Code:     TCode{},
		DataType: "Code",
		Description: `Die für die Umsatzsteuer-Abrechnungs- und -Meldezwecke verwendete Währung, die im Land des Verkäufers gültig \
			ist oder verlangt wird. \
			Anmerkung: Zu Verwenden in Kombination mit „Invoice total VAT amount in accounting currency“ (BT-111), wenn \
			die Umsatzsteuerabrechnungswährung von der Rechnungswährung abweicht. Die gültigen Währungen sind bei \
			der ISO 4217 „Codes for the representation of currencies and funds“ registriert. Nur die Alpha-3-Darstellung darf \
			verwendet werden.`,
	}
	if lElement.Value != "" {
		for _, e := range CCurrency {
			if e.CurrencyID == lElement.Value {
				lElement.Code.ID = e.CurrencyID
				lElement.Code.Text = e.CurrencyName
				break
			}
		}
		if lElement.Code.Text == "" {
			lElement.Error = errors.New(fmt.Sprintf("Error field '%s' code mapping not valid", lElement.Name))
		}
	}
	HeadFields = append(HeadFields, lElement)

	//-----------------------------------------------------//
	lElement = InvoiceDataElement{
		Name:     "Value added tax point date",
		ID:       "BT-7",
		Value:    "",
		Code:     TCode{},
		DataType: "Code",
		Description: `Das Datum, zu dem die Umsatzsteuer für den Verkäufer und für den Erwerber abrechnungsrelevant wird. \
			Die Anwendung von BT-7 und 8 schließen sich gegenseitig aus.`,
	}
	HeadFields = append(HeadFields, lElement)

	//-----------------------------------------------------//
	me.task = append(me.task, CheckAndConvert)

	return
}

//XMLToStructure returns the structure of the xcalculation from XML string
func XMLToStructure(xmlData string) (xInvoice Invoice2, err error) {
	err = xml.Unmarshal([]byte(xmlData), &xInvoice)
	return
}

//StructureToXML returns the XML of the xcalculation from the structure
func StructureToXML(xInvoice Invoice2) (xmlData string, err error) {
	if myBytesData, err := xml.Marshal(xInvoice); err == nil {
		xmlData = string(myBytesData)
	}
	return
}
