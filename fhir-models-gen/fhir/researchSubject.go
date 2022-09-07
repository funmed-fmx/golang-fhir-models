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

import "encoding/json"

// THIS FILE IS GENERATED BY https://github.com/samply/golang-fhir-models
// PLEASE DO NOT EDIT BY HAND

// ResearchSubject is documented here http://hl7.org/fhir/StructureDefinition/ResearchSubject
type ResearchSubject struct {
	Id                *string               `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                 `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string               `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string               `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative            `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension           `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension           `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        []Identifier          `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status            ResearchSubjectStatus `bson:"status,omitempty" json:"status,omitempty"`
	Period            *Period               `bson:"period,omitempty" json:"period,omitempty"`
	Study             Reference             `bson:"study,omitempty" json:"study,omitempty"`
	Individual        Reference             `bson:"individual,omitempty" json:"individual,omitempty"`
	AssignedArm       *string               `bson:"assignedArm,omitempty" json:"assignedArm,omitempty"`
	ActualArm         *string               `bson:"actualArm,omitempty" json:"actualArm,omitempty"`
	Consent           *Reference            `bson:"consent,omitempty" json:"consent,omitempty"`
}
type OtherResearchSubject ResearchSubject

// MarshalJSON marshals the given ResearchSubject as JSON into a byte slice
func (r ResearchSubject) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherResearchSubject
		ResourceType string `json:"resourceType"`
	}{
		OtherResearchSubject: OtherResearchSubject(r),
		ResourceType:         "ResearchSubject",
	})
}

// UnmarshalResearchSubject unmarshals a ResearchSubject.
func UnmarshalResearchSubject(b []byte) (ResearchSubject, error) {
	var researchSubject ResearchSubject
	if err := json.Unmarshal(b, &researchSubject); err != nil {
		return researchSubject, err
	}
	return researchSubject, nil
}
