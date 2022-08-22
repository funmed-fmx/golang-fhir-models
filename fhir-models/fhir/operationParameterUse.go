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

// OperationParameterUse is documented here http://hl7.org/fhir/ValueSet/operation-parameter-use
type OperationParameterUse int

const (
	OperationParameterUseIn OperationParameterUse = iota
	OperationParameterUseOut
)

func (code OperationParameterUse) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *OperationParameterUse) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "in":
		*code = OperationParameterUseIn
	case "out":
		*code = OperationParameterUseOut
	default:
		return fmt.Errorf("unknown OperationParameterUse code `%s`", s)
	}
	return nil
}
func (code OperationParameterUse) String() string {
	return code.Code()
}
func (code OperationParameterUse) Code() string {
	switch code {
	case OperationParameterUseIn:
		return "in"
	case OperationParameterUseOut:
		return "out"
	}
	return "<unknown>"
}
func (code OperationParameterUse) Display() string {
	switch code {
	case OperationParameterUseIn:
		return "In"
	case OperationParameterUseOut:
		return "Out"
	}
	return "<unknown>"
}
func (code OperationParameterUse) Definition() string {
	switch code {
	case OperationParameterUseIn:
		return "This is an input parameter."
	case OperationParameterUseOut:
		return "This is an output parameter."
	}
	return "<unknown>"
}
