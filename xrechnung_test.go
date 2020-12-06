//Package xrechnung - Library for processing electronic invoices - german xrechnung **/
package xrechnung

import (
	"reflect"
	"testing"
)

func TestXMLToStructure(t *testing.T) {
	var myXMLData string
	var myXRechnung InvoiceStructure
	type args struct {
		xmlData string
	}
	tests := []struct {
		name           string
		args           args
		wantXstructure InvoiceStructure
		wantErr        bool
	}{
		{
			name: "Test invoice",
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
		xstructure InvoiceStructure
	}
	tests := []struct {
		name        string
		args        args
		wantXMLData string
		wantErr     bool
	}{
		// TODO: Add test cases.
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
