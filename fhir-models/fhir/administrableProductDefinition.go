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

// AdministrableProductDefinition is documented here http://hl7.org/fhir/StructureDefinition/AdministrableProductDefinition
type AdministrableProductDefinition struct {
	Id                    *string                                               `bson:"id,omitempty" json:"id,omitempty"`
	Meta                  *Meta                                                 `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules         *string                                               `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language              *string                                               `bson:"language,omitempty" json:"language,omitempty"`
	Text                  *Narrative                                            `bson:"text,omitempty" json:"text,omitempty"`
	Extension             []Extension                                           `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension     []Extension                                           `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier            []Identifier                                          `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status                PublicationStatus                                     `bson:"status" json:"status,omitempty"`
	FormOf                []Reference                                           `bson:"formOf,omitempty" json:"formOf,omitempty"`
	AdministrableDoseForm *CodeableConcept                                      `bson:"administrableDoseForm,omitempty" json:"administrableDoseForm,omitempty"`
	UnitOfPresentation    *CodeableConcept                                      `bson:"unitOfPresentation,omitempty" json:"unitOfPresentation,omitempty"`
	ProducedFrom          []Reference                                           `bson:"producedFrom,omitempty" json:"producedFrom,omitempty"`
	Ingredient            []CodeableConcept                                     `bson:"ingredient,omitempty" json:"ingredient,omitempty"`
	Device                *Reference                                            `bson:"device,omitempty" json:"device,omitempty"`
	Property              []AdministrableProductDefinitionProperty              `bson:"property,omitempty" json:"property,omitempty"`
	RouteOfAdministration []AdministrableProductDefinitionRouteOfAdministration `bson:"routeOfAdministration" json:"routeOfAdministration,omitempty"`
}
type AdministrableProductDefinitionProperty struct {
	Id                   *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension            []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension    []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type                 CodeableConcept  `bson:"type" json:"type,omitempty"`
	ValueCodeableConcept *CodeableConcept `bson:"valueCodeableConcept,omitempty" json:"valueCodeableConcept,omitempty"`
	ValueQuantity        *Quantity        `bson:"valueQuantity,omitempty" json:"valueQuantity,omitempty"`
	ValueDate            *time.Time       `bson:"valueDate,omitempty" json:"valueDate,omitempty"`
	ValueBoolean         *bool            `bson:"valueBoolean,omitempty" json:"valueBoolean,omitempty"`
	ValueAttachment      *Attachment      `bson:"valueAttachment,omitempty" json:"valueAttachment,omitempty"`
	Status               *CodeableConcept `bson:"status,omitempty" json:"status,omitempty"`
}
type AdministrableProductDefinitionRouteOfAdministration struct {
	Id                        *string                                                            `bson:"id,omitempty" json:"id,omitempty"`
	Extension                 []Extension                                                        `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension         []Extension                                                        `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Code                      CodeableConcept                                                    `bson:"code" json:"code,omitempty"`
	FirstDose                 *Quantity                                                          `bson:"firstDose,omitempty" json:"firstDose,omitempty"`
	MaxSingleDose             *Quantity                                                          `bson:"maxSingleDose,omitempty" json:"maxSingleDose,omitempty"`
	MaxDosePerDay             *Quantity                                                          `bson:"maxDosePerDay,omitempty" json:"maxDosePerDay,omitempty"`
	MaxDosePerTreatmentPeriod *Ratio                                                             `bson:"maxDosePerTreatmentPeriod,omitempty" json:"maxDosePerTreatmentPeriod,omitempty"`
	MaxTreatmentPeriod        *Duration                                                          `bson:"maxTreatmentPeriod,omitempty" json:"maxTreatmentPeriod,omitempty"`
	TargetSpecies             []AdministrableProductDefinitionRouteOfAdministrationTargetSpecies `bson:"targetSpecies,omitempty" json:"targetSpecies,omitempty"`
}
type AdministrableProductDefinitionRouteOfAdministrationTargetSpecies struct {
	Id                *string                                                                            `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                                                                        `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                                                                        `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Code              CodeableConcept                                                                    `bson:"code" json:"code,omitempty"`
	WithdrawalPeriod  []AdministrableProductDefinitionRouteOfAdministrationTargetSpeciesWithdrawalPeriod `bson:"withdrawalPeriod,omitempty" json:"withdrawalPeriod,omitempty"`
}
type AdministrableProductDefinitionRouteOfAdministrationTargetSpeciesWithdrawalPeriod struct {
	Id                    *string         `bson:"id,omitempty" json:"id,omitempty"`
	Extension             []Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension     []Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Tissue                CodeableConcept `bson:"tissue" json:"tissue,omitempty"`
	Value                 Quantity        `bson:"value" json:"value,omitempty"`
	SupportingInformation *string         `bson:"supportingInformation,omitempty" json:"supportingInformation,omitempty"`
}
type OtherAdministrableProductDefinition AdministrableProductDefinition

// MarshalJSON marshals the given AdministrableProductDefinition as JSON into a byte slice
func (r AdministrableProductDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherAdministrableProductDefinition
		ResourceType string `json:"resourceType"`
	}{
		OtherAdministrableProductDefinition: OtherAdministrableProductDefinition(r),
		ResourceType:                        "AdministrableProductDefinition",
	})
}

// UnmarshalAdministrableProductDefinition unmarshals a AdministrableProductDefinition.
func UnmarshalAdministrableProductDefinition(b []byte) (AdministrableProductDefinition, error) {
	var administrableProductDefinition AdministrableProductDefinition
	if err := json.Unmarshal(b, &administrableProductDefinition); err != nil {
		return administrableProductDefinition, err
	}
	return administrableProductDefinition, nil
}
