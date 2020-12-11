//Package xrechnung - Library for processing electronic invoices - german xrechnung 2.0 **/
package xrechnung

import (
	"encoding/xml"
)

//XMLToStructure returns the structure of the xcalculation from XML string
func XMLToStructure(xmlData string) (xstructure interface{}, err error) {
	err = xml.Unmarshal([]byte(xmlData), &xstructure)
	return
}

//StructureToXML returns the XML of the xcalculation from the structure
func StructureToXML(xstructure interface{}) (xmlData string, err error) {
	if myBytesData, err := xml.Marshal(xstructure); err == nil {
		xmlData = string(myBytesData)
	}
	return
}
