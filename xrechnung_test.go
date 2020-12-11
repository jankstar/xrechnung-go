//Package xrechnung - Library for processing electronic invoices - german xrechnung **/
package xrechnung

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func TestXMLToStructure(t *testing.T) {
	var myXMLData string
	var myFileData []byte
	var myXRechnung CInvoice2__ubl

	xmlFile, err := os.Open("01.01a-INVOICE_ubl.xml")
	defer xmlFile.Close()
	myFileData, _ = ioutil.ReadAll(xmlFile)
	if err != nil {
		fmt.Print(err)
		return
	}
	myXMLData = string(myFileData)

	type args struct {
		xmlData string
	}
	tests := []struct {
		name           string
		args           args
		wantXstructure CInvoice2__ubl
		wantErr        bool
	}{
		{
			name: "Test XML to Structure invoice",
			args: args{
				xmlData: myXMLData,
			},
			wantXstructure: myXRechnung,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotXstructure, err := XMLToStructure(tt.args.xmlData)
			if (err != nil) != tt.wantErr {
				t.Errorf("XMLToStructure() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotXstructure, tt.wantXstructure) {
				t.Errorf("XMLToStructure() = %v, want %v", gotXstructure, tt.wantXstructure)
			}
		})
	}
}

func TestStructureToXML(t *testing.T) {
	type args struct {
		xstructure CInvoice2__ubl
	}
	var myInvoice CInvoice2__ubl

	tests := []struct {
		name        string
		args        args
		wantXMLData string
		wantErr     bool
	}{
		{
			"Test Structure to XML invoice",
			args{myInvoice},
			"",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotXMLData, err := StructureToXML(tt.args.xstructure)
			if (err != nil) != tt.wantErr {
				t.Errorf("StructureToXML() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotXMLData != tt.wantXMLData {
				t.Errorf("StructureToXML() = %v, want %v", gotXMLData, tt.wantXMLData)
			}
		})
	}
}
