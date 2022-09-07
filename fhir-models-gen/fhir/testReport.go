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
	"time"
)

// THIS FILE IS GENERATED BY https://github.com/samply/golang-fhir-models
// PLEASE DO NOT EDIT BY HAND

// TestReport is documented here http://hl7.org/fhir/StructureDefinition/TestReport
type TestReport struct {
	Id                *string                 `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                   `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                 `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                 `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative              `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension             `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension             `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        *Identifier             `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Name              *string                 `bson:"name,omitempty" json:"name,omitempty"`
	Status            TestReportStatus        `bson:"status,omitempty" json:"status,omitempty"`
	TestScript        Reference               `bson:"testScript,omitempty" json:"testScript,omitempty"`
	Result            TestReportResult        `bson:"result,omitempty" json:"result,omitempty"`
	Score             *float64                `bson:"score,omitempty" json:"score,omitempty"`
	Tester            *string                 `bson:"tester,omitempty" json:"tester,omitempty"`
	Issued            *time.Time              `bson:"issued,omitempty" json:"issued,omitempty"`
	Participant       []TestReportParticipant `bson:"participant,omitempty" json:"participant,omitempty"`
	Setup             *TestReportSetup        `bson:"setup,omitempty" json:"setup,omitempty"`
	Test              []TestReportTest        `bson:"test,omitempty" json:"test,omitempty"`
	Teardown          *TestReportTeardown     `bson:"teardown,omitempty" json:"teardown,omitempty"`
}
type TestReportParticipant struct {
	Id                *string                   `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension               `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension               `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              TestReportParticipantType `bson:"type,omitempty" json:"type,omitempty"`
	Uri               string                    `bson:"uri,omitempty" json:"uri,omitempty"`
	Display           *string                   `bson:"display,omitempty" json:"display,omitempty"`
}
type TestReportSetup struct {
	Id                *string                 `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension             `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension             `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Action            []TestReportSetupAction `bson:"action,omitempty" json:"action,omitempty"`
}
type TestReportSetupAction struct {
	Id                *string                         `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Operation         *TestReportSetupActionOperation `bson:"operation,omitempty" json:"operation,omitempty"`
	Assert            *TestReportSetupActionAssert    `bson:"assert,omitempty" json:"assert,omitempty"`
}
type TestReportSetupActionOperation struct {
	Id                *string                `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension            `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension            `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Result            TestReportActionResult `bson:"result,omitempty" json:"result,omitempty"`
	Message           *string                `bson:"message,omitempty" json:"message,omitempty"`
	Detail            *string                `bson:"detail,omitempty" json:"detail,omitempty"`
}
type TestReportSetupActionAssert struct {
	Id                *string                `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension            `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension            `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Result            TestReportActionResult `bson:"result,omitempty" json:"result,omitempty"`
	Message           *string                `bson:"message,omitempty" json:"message,omitempty"`
	Detail            *string                `bson:"detail,omitempty" json:"detail,omitempty"`
}
type TestReportTest struct {
	Id                *string                `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension            `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension            `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Name              *string                `bson:"name,omitempty" json:"name,omitempty"`
	Description       *string                `bson:"description,omitempty" json:"description,omitempty"`
	Action            []TestReportTestAction `bson:"action,omitempty" json:"action,omitempty"`
}
type TestReportTestAction struct {
	Id                *string                         `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Operation         *TestReportSetupActionOperation `bson:"operation,omitempty" json:"operation,omitempty"`
	Assert            *TestReportSetupActionAssert    `bson:"assert,omitempty" json:"assert,omitempty"`
}
type TestReportTeardown struct {
	Id                *string                    `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Action            []TestReportTeardownAction `bson:"action,omitempty" json:"action,omitempty"`
}
type TestReportTeardownAction struct {
	Id                *string                        `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                    `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                    `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Operation         TestReportSetupActionOperation `bson:"operation,omitempty" json:"operation,omitempty"`
}
type OtherTestReport TestReport

// MarshalJSON marshals the given TestReport as JSON into a byte slice
func (r TestReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherTestReport
		ResourceType string `json:"resourceType"`
	}{
		OtherTestReport: OtherTestReport(r),
		ResourceType:    "TestReport",
	})
}

// UnmarshalTestReport unmarshals a TestReport.
func UnmarshalTestReport(b []byte) (TestReport, error) {
	var testReport TestReport
	if err := json.Unmarshal(b, &testReport); err != nil {
		return testReport, err
	}
	return testReport, nil
}
