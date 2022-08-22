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

// SlotStatus is documented here http://hl7.org/fhir/ValueSet/slotstatus
type SlotStatus int

const (
	SlotStatusBusy SlotStatus = iota
	SlotStatusFree
	SlotStatusBusyUnavailable
	SlotStatusBusyTentative
	SlotStatusEnteredInError
)

func (code SlotStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *SlotStatus) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "busy":
		*code = SlotStatusBusy
	case "free":
		*code = SlotStatusFree
	case "busy-unavailable":
		*code = SlotStatusBusyUnavailable
	case "busy-tentative":
		*code = SlotStatusBusyTentative
	case "entered-in-error":
		*code = SlotStatusEnteredInError
	default:
		return fmt.Errorf("unknown SlotStatus code `%s`", s)
	}
	return nil
}
func (code SlotStatus) String() string {
	return code.Code()
}
func (code SlotStatus) Code() string {
	switch code {
	case SlotStatusBusy:
		return "busy"
	case SlotStatusFree:
		return "free"
	case SlotStatusBusyUnavailable:
		return "busy-unavailable"
	case SlotStatusBusyTentative:
		return "busy-tentative"
	case SlotStatusEnteredInError:
		return "entered-in-error"
	}
	return "<unknown>"
}
func (code SlotStatus) Display() string {
	switch code {
	case SlotStatusBusy:
		return "Busy"
	case SlotStatusFree:
		return "Free"
	case SlotStatusBusyUnavailable:
		return "Busy (Unavailable)"
	case SlotStatusBusyTentative:
		return "Busy (Tentative)"
	case SlotStatusEnteredInError:
		return "Entered in error"
	}
	return "<unknown>"
}
func (code SlotStatus) Definition() string {
	switch code {
	case SlotStatusBusy:
		return "Indicates that the time interval is busy because one  or more events have been scheduled for that interval."
	case SlotStatusFree:
		return "Indicates that the time interval is free for scheduling."
	case SlotStatusBusyUnavailable:
		return "Indicates that the time interval is busy and that the interval cannot be scheduled."
	case SlotStatusBusyTentative:
		return "Indicates that the time interval is busy because one or more events have been tentatively scheduled for that interval."
	case SlotStatusEnteredInError:
		return "This instance should not have been part of this patient's medical record."
	}
	return "<unknown>"
}
