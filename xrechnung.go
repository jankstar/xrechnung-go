//Package xrechnung - Library for processing electronic invoices - german xrechnung **/
package xrechnung

import (
	"encoding/xml"
)

//XMLToStructure liefert aus XML string die Struktur der xrechnung
func XMLToStructure(xmlData string) (xstructure interface{}, err error) {
	err = xml.Unmarshal([]byte(xmlData), &xstructure)
	return
}

//StructureToXML liefert aus Struktur das XML der xrechnung
func StructureToXML(xstructure interface{}) (xmlData string, err error) {
	var myBytesData []byte
	if myBytesData, err = xml.Marshal(xstructure); err == nil {
		xmlData = string(myBytesData)
	}
	return
}
