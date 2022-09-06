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

// SupplyRequest is documented here http://hl7.org/fhir/StructureDefinition/SupplyRequest
type SupplyRequest struct {
	Id                  *string                  `bson:"id,omitempty" json:"id,omitempty"`
	Meta                *Meta                    `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules       *string                  `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language            *string                  `bson:"language,omitempty" json:"language,omitempty"`
	Text                *Narrative               `bson:"text,omitempty" json:"text,omitempty"`
	Extension           []Extension              `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension   []Extension              `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier          []Identifier             `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status              *SupplyRequestStatus     `bson:"status,omitempty" json:"status,omitempty"`
	Category            *CodeableConcept         `bson:"category,omitempty" json:"category,omitempty"`
	Priority            *RequestPriority         `bson:"priority,omitempty" json:"priority,omitempty"`
	ItemCodeableConcept CodeableConcept          `bson:"itemCodeableConcept" json:"itemCodeableConcept"`
	ItemReference       Reference                `bson:"itemReference" json:"itemReference"`
	Quantity            Quantity                 `bson:"quantity" json:"quantity"`
	Parameter           []SupplyRequestParameter `bson:"parameter,omitempty" json:"parameter,omitempty"`
	OccurrenceDateTime  *time.Time               `bson:"occurrenceDateTime,omitempty" json:"occurrenceDateTime,omitempty"`
	OccurrencePeriod    *Period                  `bson:"occurrencePeriod,omitempty" json:"occurrencePeriod,omitempty"`
	OccurrenceTiming    *Timing                  `bson:"occurrenceTiming,omitempty" json:"occurrenceTiming,omitempty"`
	AuthoredOn          *time.Time               `bson:"authoredOn,omitempty" json:"authoredOn,omitempty"`
	Requester           *Reference               `bson:"requester,omitempty" json:"requester,omitempty"`
	Supplier            []Reference              `bson:"supplier,omitempty" json:"supplier,omitempty"`
	ReasonCode          []CodeableConcept        `bson:"reasonCode,omitempty" json:"reasonCode,omitempty"`
	ReasonReference     []Reference              `bson:"reasonReference,omitempty" json:"reasonReference,omitempty"`
	DeliverFrom         *Reference               `bson:"deliverFrom,omitempty" json:"deliverFrom,omitempty"`
	DeliverTo           *Reference               `bson:"deliverTo,omitempty" json:"deliverTo,omitempty"`
}
type SupplyRequestParameter struct {
	Id                   *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension            []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension    []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Code                 *CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	ValueCodeableConcept *CodeableConcept `bson:"valueCodeableConcept,omitempty" json:"valueCodeableConcept,omitempty"`
	ValueQuantity        *Quantity        `bson:"valueQuantity,omitempty" json:"valueQuantity,omitempty"`
	ValueRange           *Range           `bson:"valueRange,omitempty" json:"valueRange,omitempty"`
	ValueBoolean         *bool            `bson:"valueBoolean,omitempty" json:"valueBoolean,omitempty"`
}
type OtherSupplyRequest SupplyRequest

// MarshalJSON marshals the given SupplyRequest as JSON into a byte slice
func (r SupplyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherSupplyRequest
		ResourceType string `json:"resourceType"`
	}{
		OtherSupplyRequest: OtherSupplyRequest(r),
		ResourceType:       "SupplyRequest",
	})
}

// UnmarshalSupplyRequest unmarshals a SupplyRequest.
func UnmarshalSupplyRequest(b []byte) (SupplyRequest, error) {
	var supplyRequest SupplyRequest
	if err := json.Unmarshal(b, &supplyRequest); err != nil {
		return supplyRequest, err
	}
	return supplyRequest, nil
}
