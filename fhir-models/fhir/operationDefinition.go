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

// OperationDefinition is documented here http://hl7.org/fhir/StructureDefinition/OperationDefinition
type OperationDefinition struct {
	Id                *string                        `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                          `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                        `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                        `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative                     `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension                    `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                    `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Url               *string                        `bson:"url,omitempty" json:"url,omitempty"`
	Version           *string                        `bson:"version,omitempty" json:"version,omitempty"`
	Name              string                         `bson:"name" json:"name"`
	Title             *string                        `bson:"title,omitempty" json:"title,omitempty"`
	Status            PublicationStatus              `bson:"status" json:"status"`
	Kind              OperationKind                  `bson:"kind" json:"kind"`
	Experimental      *bool                          `bson:"experimental,omitempty" json:"experimental,omitempty"`
	Date              *time.Time                     `bson:"date,omitempty" json:"date,omitempty"`
	Publisher         *string                        `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Contact           []ContactDetail                `bson:"contact,omitempty" json:"contact,omitempty"`
	Description       *string                        `bson:"description,omitempty" json:"description,omitempty"`
	UseContext        []UsageContext                 `bson:"useContext,omitempty" json:"useContext,omitempty"`
	Jurisdiction      []CodeableConcept              `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	Purpose           *string                        `bson:"purpose,omitempty" json:"purpose,omitempty"`
	AffectsState      *bool                          `bson:"affectsState,omitempty" json:"affectsState,omitempty"`
	Code              string                         `bson:"code" json:"code"`
	Comment           *string                        `bson:"comment,omitempty" json:"comment,omitempty"`
	Base              *string                        `bson:"base,omitempty" json:"base,omitempty"`
	Resource          []ResourceType                 `bson:"resource,omitempty" json:"resource,omitempty"`
	System            bool                           `bson:"system" json:"system"`
	Type              bool                           `bson:"type" json:"type"`
	Instance          bool                           `bson:"instance" json:"instance"`
	InputProfile      *string                        `bson:"inputProfile,omitempty" json:"inputProfile,omitempty"`
	OutputProfile     *string                        `bson:"outputProfile,omitempty" json:"outputProfile,omitempty"`
	Parameter         []OperationDefinitionParameter `bson:"parameter,omitempty" json:"parameter,omitempty"`
	Overload          []OperationDefinitionOverload  `bson:"overload,omitempty" json:"overload,omitempty"`
}
type OperationDefinitionParameter struct {
	Id                *string                                      `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                                  `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                                  `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Name              string                                       `bson:"name" json:"name"`
	Use               OperationParameterUse                        `bson:"use" json:"use"`
	Min               int                                          `bson:"min" json:"min"`
	Max               string                                       `bson:"max" json:"max"`
	Documentation     *string                                      `bson:"documentation,omitempty" json:"documentation,omitempty"`
	Type              *string                                      `bson:"type,omitempty" json:"type,omitempty"`
	TargetProfile     []string                                     `bson:"targetProfile,omitempty" json:"targetProfile,omitempty"`
	SearchType        *SearchParamType                             `bson:"searchType,omitempty" json:"searchType,omitempty"`
	Binding           *OperationDefinitionParameterBinding         `bson:"binding,omitempty" json:"binding,omitempty"`
	ReferencedFrom    []OperationDefinitionParameterReferencedFrom `bson:"referencedFrom,omitempty" json:"referencedFrom,omitempty"`
	Part              []OperationDefinitionParameter               `bson:"part,omitempty" json:"part,omitempty"`
}
type OperationDefinitionParameterBinding struct {
	Id                *string         `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Strength          BindingStrength `bson:"strength" json:"strength"`
	ValueSet          string          `bson:"valueSet" json:"valueSet"`
}
type OperationDefinitionParameterReferencedFrom struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Source            string      `bson:"source" json:"source"`
	SourceId          *string     `bson:"sourceId,omitempty" json:"sourceId,omitempty"`
}
type OperationDefinitionOverload struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	ParameterName     []string    `bson:"parameterName,omitempty" json:"parameterName,omitempty"`
	Comment           *string     `bson:"comment,omitempty" json:"comment,omitempty"`
}
type OtherOperationDefinition OperationDefinition

// MarshalJSON marshals the given OperationDefinition as JSON into a byte slice
func (r OperationDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherOperationDefinition
		ResourceType string `json:"resourceType"`
	}{
		OtherOperationDefinition: OtherOperationDefinition(r),
		ResourceType:             "OperationDefinition",
	})
}

// UnmarshalOperationDefinition unmarshals a OperationDefinition.
func UnmarshalOperationDefinition(b []byte) (OperationDefinition, error) {
	var operationDefinition OperationDefinition
	if err := json.Unmarshal(b, &operationDefinition); err != nil {
		return operationDefinition, err
	}
	return operationDefinition, nil
}
