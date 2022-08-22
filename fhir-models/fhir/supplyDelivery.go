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

// SupplyDelivery is documented here http://hl7.org/fhir/StructureDefinition/SupplyDelivery
type SupplyDelivery struct {
	Id                 *string                     `bson:"id,omitempty" json:"id,omitempty"`
	Meta               *Meta                       `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules      *string                     `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language           *string                     `bson:"language,omitempty" json:"language,omitempty"`
	Text               *Narrative                  `bson:"text,omitempty" json:"text,omitempty"`
	Extension          []Extension                 `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension  []Extension                 `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier         []Identifier                `bson:"identifier,omitempty" json:"identifier,omitempty"`
	BasedOn            []Reference                 `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	PartOf             []Reference                 `bson:"partOf,omitempty" json:"partOf,omitempty"`
	Status             *SupplyDeliveryStatus       `bson:"status,omitempty" json:"status,omitempty"`
	Patient            *Reference                  `bson:"patient,omitempty" json:"patient,omitempty"`
	Type               *CodeableConcept            `bson:"type,omitempty" json:"type,omitempty"`
	SuppliedItem       *SupplyDeliverySuppliedItem `bson:"suppliedItem,omitempty" json:"suppliedItem,omitempty"`
	OccurrenceDateTime *string                     `bson:"occurrenceDateTime,omitempty" json:"occurrenceDateTime,omitempty"`
	OccurrencePeriod   *Period                     `bson:"occurrencePeriod,omitempty" json:"occurrencePeriod,omitempty"`
	OccurrenceTiming   *Timing                     `bson:"occurrenceTiming,omitempty" json:"occurrenceTiming,omitempty"`
	Supplier           *Reference                  `bson:"supplier,omitempty" json:"supplier,omitempty"`
	Destination        *Reference                  `bson:"destination,omitempty" json:"destination,omitempty"`
	Receiver           []Reference                 `bson:"receiver,omitempty" json:"receiver,omitempty"`
}
type SupplyDeliverySuppliedItem struct {
	Id                  *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension           []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension   []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Quantity            *Quantity        `bson:"quantity,omitempty" json:"quantity,omitempty"`
	ItemCodeableConcept *CodeableConcept `bson:"itemCodeableConcept,omitempty" json:"itemCodeableConcept,omitempty"`
	ItemReference       *Reference       `bson:"itemReference,omitempty" json:"itemReference,omitempty"`
}
type OtherSupplyDelivery SupplyDelivery

// MarshalJSON marshals the given SupplyDelivery as JSON into a byte slice
func (r SupplyDelivery) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherSupplyDelivery
		ResourceType string `json:"resourceType"`
	}{
		OtherSupplyDelivery: OtherSupplyDelivery(r),
		ResourceType:        "SupplyDelivery",
	})
}

// UnmarshalSupplyDelivery unmarshals a SupplyDelivery.
func UnmarshalSupplyDelivery(b []byte) (SupplyDelivery, error) {
	var supplyDelivery SupplyDelivery
	if err := json.Unmarshal(b, &supplyDelivery); err != nil {
		return supplyDelivery, err
	}
	return supplyDelivery, nil
}
