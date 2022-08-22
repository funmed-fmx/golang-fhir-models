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

// EnableWhenBehavior is documented here http://hl7.org/fhir/ValueSet/questionnaire-enable-behavior
type EnableWhenBehavior int

const (
	EnableWhenBehaviorAll EnableWhenBehavior = iota
	EnableWhenBehaviorAny
)

func (code EnableWhenBehavior) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *EnableWhenBehavior) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "all":
		*code = EnableWhenBehaviorAll
	case "any":
		*code = EnableWhenBehaviorAny
	default:
		return fmt.Errorf("unknown EnableWhenBehavior code `%s`", s)
	}
	return nil
}
func (code EnableWhenBehavior) String() string {
	return code.Code()
}
func (code EnableWhenBehavior) Code() string {
	switch code {
	case EnableWhenBehaviorAll:
		return "all"
	case EnableWhenBehaviorAny:
		return "any"
	}
	return "<unknown>"
}
func (code EnableWhenBehavior) Display() string {
	switch code {
	case EnableWhenBehaviorAll:
		return "All"
	case EnableWhenBehaviorAny:
		return "Any"
	}
	return "<unknown>"
}
func (code EnableWhenBehavior) Definition() string {
	switch code {
	case EnableWhenBehaviorAll:
		return "Enable the question when all the enableWhen criteria are satisfied."
	case EnableWhenBehaviorAny:
		return "Enable the question when any of the enableWhen criteria are satisfied."
	}
	return "<unknown>"
}
