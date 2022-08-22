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

// AssertionDirectionType is documented here http://hl7.org/fhir/ValueSet/assert-direction-codes
type AssertionDirectionType int

const (
	AssertionDirectionTypeResponse AssertionDirectionType = iota
	AssertionDirectionTypeRequest
)

func (code AssertionDirectionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *AssertionDirectionType) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "response":
		*code = AssertionDirectionTypeResponse
	case "request":
		*code = AssertionDirectionTypeRequest
	default:
		return fmt.Errorf("unknown AssertionDirectionType code `%s`", s)
	}
	return nil
}
func (code AssertionDirectionType) String() string {
	return code.Code()
}
func (code AssertionDirectionType) Code() string {
	switch code {
	case AssertionDirectionTypeResponse:
		return "response"
	case AssertionDirectionTypeRequest:
		return "request"
	}
	return "<unknown>"
}
func (code AssertionDirectionType) Display() string {
	switch code {
	case AssertionDirectionTypeResponse:
		return "response"
	case AssertionDirectionTypeRequest:
		return "request"
	}
	return "<unknown>"
}
func (code AssertionDirectionType) Definition() string {
	switch code {
	case AssertionDirectionTypeResponse:
		return "The assertion is evaluated on the response. This is the default value."
	case AssertionDirectionTypeRequest:
		return "The assertion is evaluated on the request."
	}
	return "<unknown>"
}
