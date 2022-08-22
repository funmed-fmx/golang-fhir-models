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

// ActionPrecheckBehavior is documented here http://hl7.org/fhir/ValueSet/action-precheck-behavior
type ActionPrecheckBehavior int

const (
	ActionPrecheckBehaviorYes ActionPrecheckBehavior = iota
	ActionPrecheckBehaviorNo
)

func (code ActionPrecheckBehavior) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *ActionPrecheckBehavior) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "yes":
		*code = ActionPrecheckBehaviorYes
	case "no":
		*code = ActionPrecheckBehaviorNo
	default:
		return fmt.Errorf("unknown ActionPrecheckBehavior code `%s`", s)
	}
	return nil
}
func (code ActionPrecheckBehavior) String() string {
	return code.Code()
}
func (code ActionPrecheckBehavior) Code() string {
	switch code {
	case ActionPrecheckBehaviorYes:
		return "yes"
	case ActionPrecheckBehaviorNo:
		return "no"
	}
	return "<unknown>"
}
func (code ActionPrecheckBehavior) Display() string {
	switch code {
	case ActionPrecheckBehaviorYes:
		return "Yes"
	case ActionPrecheckBehaviorNo:
		return "No"
	}
	return "<unknown>"
}
func (code ActionPrecheckBehavior) Definition() string {
	switch code {
	case ActionPrecheckBehaviorYes:
		return "An action with this behavior is one of the most frequent action that is, or should be, included by an end user, for the particular context in which the action occurs. The system displaying the action to the end user should consider \"pre-checking\" such an action as a convenience for the user."
	case ActionPrecheckBehaviorNo:
		return "An action with this behavior is one of the less frequent actions included by the end user, for the particular context in which the action occurs. The system displaying the actions to the end user would typically not \"pre-check\" such an action."
	}
	return "<unknown>"
}
