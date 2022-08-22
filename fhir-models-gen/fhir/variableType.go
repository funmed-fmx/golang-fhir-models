// Copyright 2019 The Samply Development Community
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

// THIS FILE IS GENERATED BY https://github.com/samply/golang-fhir-models
// PLEASE DO NOT EDIT BY HAND

// VariableType is documented here http://hl7.org/fhir/ValueSet/variable-type
type VariableType int

const (
	VariableTypeDichotomous VariableType = iota
	VariableTypeContinuous
	VariableTypeDescriptive
)

func (code VariableType) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *VariableType) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "dichotomous":
		*code = VariableTypeDichotomous
	case "continuous":
		*code = VariableTypeContinuous
	case "descriptive":
		*code = VariableTypeDescriptive
	default:
		return fmt.Errorf("unknown VariableType code `%s`", s)
	}
	return nil
}
func (code VariableType) String() string {
	return code.Code()
}
func (code VariableType) Code() string {
	switch code {
	case VariableTypeDichotomous:
		return "dichotomous"
	case VariableTypeContinuous:
		return "continuous"
	case VariableTypeDescriptive:
		return "descriptive"
	}
	return "<unknown>"
}
func (code VariableType) Display() string {
	switch code {
	case VariableTypeDichotomous:
		return "Dichotomous"
	case VariableTypeContinuous:
		return "Continuous"
	case VariableTypeDescriptive:
		return "Descriptive"
	}
	return "<unknown>"
}
func (code VariableType) Definition() string {
	switch code {
	case VariableTypeDichotomous:
		return "The variable is dichotomous, such as present or absent."
	case VariableTypeContinuous:
		return "The variable is a continuous result such as a quantity."
	case VariableTypeDescriptive:
		return "The variable is described narratively rather than quantitatively."
	}
	return "<unknown>"
}
