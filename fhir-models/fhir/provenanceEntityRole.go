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

// ProvenanceEntityRole is documented here http://hl7.org/fhir/ValueSet/provenance-entity-role
type ProvenanceEntityRole int

const (
	ProvenanceEntityRoleDerivation ProvenanceEntityRole = iota
	ProvenanceEntityRoleRevision
	ProvenanceEntityRoleQuotation
	ProvenanceEntityRoleSource
	ProvenanceEntityRoleRemoval
)

func (code ProvenanceEntityRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *ProvenanceEntityRole) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "derivation":
		*code = ProvenanceEntityRoleDerivation
	case "revision":
		*code = ProvenanceEntityRoleRevision
	case "quotation":
		*code = ProvenanceEntityRoleQuotation
	case "source":
		*code = ProvenanceEntityRoleSource
	case "removal":
		*code = ProvenanceEntityRoleRemoval
	default:
		return fmt.Errorf("unknown ProvenanceEntityRole code `%s`", s)
	}
	return nil
}
func (code ProvenanceEntityRole) String() string {
	return code.Code()
}
func (code ProvenanceEntityRole) Code() string {
	switch code {
	case ProvenanceEntityRoleDerivation:
		return "derivation"
	case ProvenanceEntityRoleRevision:
		return "revision"
	case ProvenanceEntityRoleQuotation:
		return "quotation"
	case ProvenanceEntityRoleSource:
		return "source"
	case ProvenanceEntityRoleRemoval:
		return "removal"
	}
	return "<unknown>"
}
func (code ProvenanceEntityRole) Display() string {
	switch code {
	case ProvenanceEntityRoleDerivation:
		return "Derivation"
	case ProvenanceEntityRoleRevision:
		return "Revision"
	case ProvenanceEntityRoleQuotation:
		return "Quotation"
	case ProvenanceEntityRoleSource:
		return "Source"
	case ProvenanceEntityRoleRemoval:
		return "Removal"
	}
	return "<unknown>"
}
func (code ProvenanceEntityRole) Definition() string {
	switch code {
	case ProvenanceEntityRoleDerivation:
		return "A transformation of an entity into another, an update of an entity resulting in a new one, or the construction of a new entity based on a pre-existing entity."
	case ProvenanceEntityRoleRevision:
		return "A derivation for which the resulting entity is a revised version of some original."
	case ProvenanceEntityRoleQuotation:
		return "The repeat of (some or all of) an entity, such as text or image, by someone who might or might not be its original author."
	case ProvenanceEntityRoleSource:
		return "A primary source for a topic refers to something produced by some agent with direct experience and knowledge about the topic, at the time of the topic's study, without benefit from hindsight."
	case ProvenanceEntityRoleRemoval:
		return "A derivation for which the entity is removed from accessibility usually through the use of the Delete operation."
	}
	return "<unknown>"
}
