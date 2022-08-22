// Copyright 2019 - 2021 The Samply Community
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package fhir

import (
	"encoding/json"
	"fmt"
	"strings"
)

// THIS FILE IS GENERATED BY https://github.com/funmed-ally/golang-fhir-models
// PLEASE DO NOT EDIT BY HAND

// UDIEntryType is documented here http://hl7.org/fhir/ValueSet/udi-entry-type
type UDIEntryType int

const (
	UDIEntryTypeBarcode UDIEntryType = iota
	UDIEntryTypeRfid
	UDIEntryTypeManual
	UDIEntryTypeCard
	UDIEntryTypeSelfReported
	UDIEntryTypeUnknown
)

func (code UDIEntryType) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *UDIEntryType) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "barcode":
		*code = UDIEntryTypeBarcode
	case "rfid":
		*code = UDIEntryTypeRfid
	case "manual":
		*code = UDIEntryTypeManual
	case "card":
		*code = UDIEntryTypeCard
	case "self-reported":
		*code = UDIEntryTypeSelfReported
	case "unknown":
		*code = UDIEntryTypeUnknown
	default:
		return fmt.Errorf("unknown UDIEntryType code `%s`", s)
	}
	return nil
}
func (code UDIEntryType) String() string {
	return code.Code()
}
func (code UDIEntryType) Code() string {
	switch code {
	case UDIEntryTypeBarcode:
		return "barcode"
	case UDIEntryTypeRfid:
		return "rfid"
	case UDIEntryTypeManual:
		return "manual"
	case UDIEntryTypeCard:
		return "card"
	case UDIEntryTypeSelfReported:
		return "self-reported"
	case UDIEntryTypeUnknown:
		return "unknown"
	}
	return "<unknown>"
}
func (code UDIEntryType) Display() string {
	switch code {
	case UDIEntryTypeBarcode:
		return "Barcode"
	case UDIEntryTypeRfid:
		return "RFID"
	case UDIEntryTypeManual:
		return "Manual"
	case UDIEntryTypeCard:
		return "Card"
	case UDIEntryTypeSelfReported:
		return "Self Reported"
	case UDIEntryTypeUnknown:
		return "Unknown"
	}
	return "<unknown>"
}
func (code UDIEntryType) Definition() string {
	switch code {
	case UDIEntryTypeBarcode:
		return "a barcodescanner captured the data from the device label."
	case UDIEntryTypeRfid:
		return "An RFID chip reader captured the data from the device label."
	case UDIEntryTypeManual:
		return "The data was read from the label by a person and manually entered. (e.g.  via a keyboard)."
	case UDIEntryTypeCard:
		return "The data originated from a patient's implant card and was read by an operator."
	case UDIEntryTypeSelfReported:
		return "The data originated from a patient source and was not directly scanned or read from a label or card."
	case UDIEntryTypeUnknown:
		return "The method of data capture has not been determined."
	}
	return "<unknown>"
}
