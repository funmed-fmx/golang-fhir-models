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

import "encoding/json"

// THIS FILE IS GENERATED BY https://github.com/funmed-ally/golang-fhir-models
// PLEASE DO NOT EDIT BY HAND

// MedicationAdministration is documented here http://hl7.org/fhir/StructureDefinition/MedicationAdministration
type MedicationAdministration struct {
	Id                    *string                             `bson:"id,omitempty" json:"id,omitempty"`
	Meta                  *Meta                               `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules         *string                             `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language              *string                             `bson:"language,omitempty" json:"language,omitempty"`
	Text                  *Narrative                          `bson:"text,omitempty" json:"text,omitempty"`
	Extension             []Extension                         `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension     []Extension                         `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier            []Identifier                        `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Instantiates          []string                            `bson:"instantiates,omitempty" json:"instantiates,omitempty"`
	PartOf                []Reference                         `bson:"partOf,omitempty" json:"partOf,omitempty"`
	Status                string                              `bson:"status" json:"status"`
	StatusReason          []CodeableConcept                   `bson:"statusReason,omitempty" json:"statusReason,omitempty"`
	Category              *CodeableConcept                    `bson:"category,omitempty" json:"category,omitempty"`
	Subject               Reference                           `bson:"subject" json:"subject"`
	Context               *Reference                          `bson:"context,omitempty" json:"context,omitempty"`
	SupportingInformation []Reference                         `bson:"supportingInformation,omitempty" json:"supportingInformation,omitempty"`
	Performer             []MedicationAdministrationPerformer `bson:"performer,omitempty" json:"performer,omitempty"`
	ReasonCode            []CodeableConcept                   `bson:"reasonCode,omitempty" json:"reasonCode,omitempty"`
	ReasonReference       []Reference                         `bson:"reasonReference,omitempty" json:"reasonReference,omitempty"`
	Request               *Reference                          `bson:"request,omitempty" json:"request,omitempty"`
	Device                []Reference                         `bson:"device,omitempty" json:"device,omitempty"`
	Note                  []Annotation                        `bson:"note,omitempty" json:"note,omitempty"`
	Dosage                *MedicationAdministrationDosage     `bson:"dosage,omitempty" json:"dosage,omitempty"`
	EventHistory          []Reference                         `bson:"eventHistory,omitempty" json:"eventHistory,omitempty"`
}
type MedicationAdministrationPerformer struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Function          *CodeableConcept `bson:"function,omitempty" json:"function,omitempty"`
	Actor             Reference        `bson:"actor" json:"actor"`
}
type MedicationAdministrationDosage struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Text              *string          `bson:"text,omitempty" json:"text,omitempty"`
	Site              *CodeableConcept `bson:"site,omitempty" json:"site,omitempty"`
	Route             *CodeableConcept `bson:"route,omitempty" json:"route,omitempty"`
	Method            *CodeableConcept `bson:"method,omitempty" json:"method,omitempty"`
	Dose              *Quantity        `bson:"dose,omitempty" json:"dose,omitempty"`
}
type OtherMedicationAdministration MedicationAdministration

// MarshalJSON marshals the given MedicationAdministration as JSON into a byte slice
func (r MedicationAdministration) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherMedicationAdministration
		ResourceType string `json:"resourceType"`
	}{
		OtherMedicationAdministration: OtherMedicationAdministration(r),
		ResourceType:                  "MedicationAdministration",
	})
}

// UnmarshalMedicationAdministration unmarshals a MedicationAdministration.
func UnmarshalMedicationAdministration(b []byte) (MedicationAdministration, error) {
	var medicationAdministration MedicationAdministration
	if err := json.Unmarshal(b, &medicationAdministration); err != nil {
		return medicationAdministration, err
	}
	return medicationAdministration, nil
}
