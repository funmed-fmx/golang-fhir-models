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

// EpisodeOfCare is documented here http://hl7.org/fhir/StructureDefinition/EpisodeOfCare
type EpisodeOfCare struct {
	Id                   *string                      `bson:"id,omitempty" json:"id,omitempty"`
	Meta                 *Meta                        `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules        *string                      `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language             *string                      `bson:"language,omitempty" json:"language,omitempty"`
	Text                 *Narrative                   `bson:"text,omitempty" json:"text,omitempty"`
	Extension            []Extension                  `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension    []Extension                  `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier           []Identifier                 `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status               EpisodeOfCareStatus          `bson:"status" json:"status"`
	StatusHistory        []EpisodeOfCareStatusHistory `bson:"statusHistory,omitempty" json:"statusHistory,omitempty"`
	Type                 []CodeableConcept            `bson:"type,omitempty" json:"type,omitempty"`
	Diagnosis            []EpisodeOfCareDiagnosis     `bson:"diagnosis,omitempty" json:"diagnosis,omitempty"`
	Patient              Reference                    `bson:"patient" json:"patient"`
	ManagingOrganization *Reference                   `bson:"managingOrganization,omitempty" json:"managingOrganization,omitempty"`
	Period               *Period                      `bson:"period,omitempty" json:"period,omitempty"`
	ReferralRequest      []Reference                  `bson:"referralRequest,omitempty" json:"referralRequest,omitempty"`
	CareManager          *Reference                   `bson:"careManager,omitempty" json:"careManager,omitempty"`
	Team                 []Reference                  `bson:"team,omitempty" json:"team,omitempty"`
	Account              []Reference                  `bson:"account,omitempty" json:"account,omitempty"`
}
type EpisodeOfCareStatusHistory struct {
	Id                *string             `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension         `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension         `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Status            EpisodeOfCareStatus `bson:"status" json:"status"`
	Period            Period              `bson:"period" json:"period"`
}
type EpisodeOfCareDiagnosis struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Condition         Reference        `bson:"condition" json:"condition"`
	Role              *CodeableConcept `bson:"role,omitempty" json:"role,omitempty"`
	Rank              *int             `bson:"rank,omitempty" json:"rank,omitempty"`
}
type OtherEpisodeOfCare EpisodeOfCare

// MarshalJSON marshals the given EpisodeOfCare as JSON into a byte slice
func (r EpisodeOfCare) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherEpisodeOfCare
		ResourceType string `json:"resourceType"`
	}{
		OtherEpisodeOfCare: OtherEpisodeOfCare(r),
		ResourceType:       "EpisodeOfCare",
	})
}

// UnmarshalEpisodeOfCare unmarshals a EpisodeOfCare.
func UnmarshalEpisodeOfCare(b []byte) (EpisodeOfCare, error) {
	var episodeOfCare EpisodeOfCare
	if err := json.Unmarshal(b, &episodeOfCare); err != nil {
		return episodeOfCare, err
	}
	return episodeOfCare, nil
}
