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

// Endpoint is documented here http://hl7.org/fhir/StructureDefinition/Endpoint
type Endpoint struct {
	Id                   *string           `bson:"id,omitempty" json:"id,omitempty"`
	Meta                 *Meta             `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules        *string           `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language             *string           `bson:"language,omitempty" json:"language,omitempty"`
	Text                 *Narrative        `bson:"text,omitempty" json:"text,omitempty"`
	Extension            []Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension    []Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier           []Identifier      `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status               EndpointStatus    `bson:"status" json:"status"`
	ConnectionType       Coding            `bson:"connectionType" json:"connectionType"`
	Name                 *string           `bson:"name,omitempty" json:"name,omitempty"`
	ManagingOrganization *Reference        `bson:"managingOrganization,omitempty" json:"managingOrganization,omitempty"`
	Contact              []ContactPoint    `bson:"contact,omitempty" json:"contact,omitempty"`
	Period               *Period           `bson:"period,omitempty" json:"period,omitempty"`
	PayloadType          []CodeableConcept `bson:"payloadType" json:"payloadType"`
	PayloadMimeType      []string          `bson:"payloadMimeType,omitempty" json:"payloadMimeType,omitempty"`
	Address              string            `bson:"address" json:"address"`
	Header               []string          `bson:"header,omitempty" json:"header,omitempty"`
}
type OtherEndpoint Endpoint

// MarshalJSON marshals the given Endpoint as JSON into a byte slice
func (r Endpoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherEndpoint
		ResourceType string `json:"resourceType"`
	}{
		OtherEndpoint: OtherEndpoint(r),
		ResourceType:  "Endpoint",
	})
}

// UnmarshalEndpoint unmarshals a Endpoint.
func UnmarshalEndpoint(b []byte) (Endpoint, error) {
	var endpoint Endpoint
	if err := json.Unmarshal(b, &endpoint); err != nil {
		return endpoint, err
	}
	return endpoint, nil
}
