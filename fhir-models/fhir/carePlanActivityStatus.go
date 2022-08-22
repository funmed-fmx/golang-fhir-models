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

// CarePlanActivityStatus is documented here http://hl7.org/fhir/ValueSet/care-plan-activity-status
type CarePlanActivityStatus int

const (
	CarePlanActivityStatusNotStarted CarePlanActivityStatus = iota
	CarePlanActivityStatusScheduled
	CarePlanActivityStatusInProgress
	CarePlanActivityStatusOnHold
	CarePlanActivityStatusCompleted
	CarePlanActivityStatusCancelled
	CarePlanActivityStatusStopped
	CarePlanActivityStatusUnknown
	CarePlanActivityStatusEnteredInError
)

func (code CarePlanActivityStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *CarePlanActivityStatus) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "not-started":
		*code = CarePlanActivityStatusNotStarted
	case "scheduled":
		*code = CarePlanActivityStatusScheduled
	case "in-progress":
		*code = CarePlanActivityStatusInProgress
	case "on-hold":
		*code = CarePlanActivityStatusOnHold
	case "completed":
		*code = CarePlanActivityStatusCompleted
	case "cancelled":
		*code = CarePlanActivityStatusCancelled
	case "stopped":
		*code = CarePlanActivityStatusStopped
	case "unknown":
		*code = CarePlanActivityStatusUnknown
	case "entered-in-error":
		*code = CarePlanActivityStatusEnteredInError
	default:
		return fmt.Errorf("unknown CarePlanActivityStatus code `%s`", s)
	}
	return nil
}
func (code CarePlanActivityStatus) String() string {
	return code.Code()
}
func (code CarePlanActivityStatus) Code() string {
	switch code {
	case CarePlanActivityStatusNotStarted:
		return "not-started"
	case CarePlanActivityStatusScheduled:
		return "scheduled"
	case CarePlanActivityStatusInProgress:
		return "in-progress"
	case CarePlanActivityStatusOnHold:
		return "on-hold"
	case CarePlanActivityStatusCompleted:
		return "completed"
	case CarePlanActivityStatusCancelled:
		return "cancelled"
	case CarePlanActivityStatusStopped:
		return "stopped"
	case CarePlanActivityStatusUnknown:
		return "unknown"
	case CarePlanActivityStatusEnteredInError:
		return "entered-in-error"
	}
	return "<unknown>"
}
func (code CarePlanActivityStatus) Display() string {
	switch code {
	case CarePlanActivityStatusNotStarted:
		return "Not Started"
	case CarePlanActivityStatusScheduled:
		return "Scheduled"
	case CarePlanActivityStatusInProgress:
		return "In Progress"
	case CarePlanActivityStatusOnHold:
		return "On Hold"
	case CarePlanActivityStatusCompleted:
		return "Completed"
	case CarePlanActivityStatusCancelled:
		return "Cancelled"
	case CarePlanActivityStatusStopped:
		return "Stopped"
	case CarePlanActivityStatusUnknown:
		return "Unknown"
	case CarePlanActivityStatusEnteredInError:
		return "Entered in Error"
	}
	return "<unknown>"
}
func (code CarePlanActivityStatus) Definition() string {
	switch code {
	case CarePlanActivityStatusNotStarted:
		return "Care plan activity is planned but no action has yet been taken."
	case CarePlanActivityStatusScheduled:
		return "Appointment or other booking has occurred but activity has not yet begun."
	case CarePlanActivityStatusInProgress:
		return "Care plan activity has been started but is not yet complete."
	case CarePlanActivityStatusOnHold:
		return "Care plan activity was started but has temporarily ceased with an expectation of resumption at a future time."
	case CarePlanActivityStatusCompleted:
		return "Care plan activity has been completed (more or less) as planned."
	case CarePlanActivityStatusCancelled:
		return "The planned care plan activity has been withdrawn."
	case CarePlanActivityStatusStopped:
		return "The planned care plan activity has been ended prior to completion after the activity was started."
	case CarePlanActivityStatusUnknown:
		return "The current state of the care plan activity is not known.  Note: This concept is not to be used for \"other\" - one of the listed statuses is presumed to apply, but the authoring/source system does not know which one."
	case CarePlanActivityStatusEnteredInError:
		return "Care plan activity was entered in error and voided."
	}
	return "<unknown>"
}
