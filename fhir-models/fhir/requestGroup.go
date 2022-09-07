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

// RequestGroup is documented here http://hl7.org/fhir/StructureDefinition/RequestGroup
type RequestGroup struct {
	Id                    *string              `bson:"id,omitempty" json:"id,omitempty"`
	Meta                  *Meta                `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules         *string              `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language              *string              `bson:"language,omitempty" json:"language,omitempty"`
	Text                  *Narrative           `bson:"text,omitempty" json:"text,omitempty"`
	Extension             []Extension          `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension     []Extension          `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier            []Identifier         `bson:"identifier,omitempty" json:"identifier,omitempty"`
	InstantiatesCanonical []string             `bson:"instantiatesCanonical,omitempty" json:"instantiatesCanonical,omitempty"`
	InstantiatesUri       []string             `bson:"instantiatesUri,omitempty" json:"instantiatesUri,omitempty"`
	BasedOn               []Reference          `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	Replaces              []Reference          `bson:"replaces,omitempty" json:"replaces,omitempty"`
	GroupIdentifier       *Identifier          `bson:"groupIdentifier,omitempty" json:"groupIdentifier,omitempty"`
	Status                RequestStatus        `bson:"status" json:"status,omitempty"`
	Intent                RequestIntent        `bson:"intent" json:"intent,omitempty"`
	Priority              *RequestPriority     `bson:"priority,omitempty" json:"priority,omitempty"`
	Code                  *CodeableConcept     `bson:"code,omitempty" json:"code,omitempty"`
	Subject               *Reference           `bson:"subject,omitempty" json:"subject,omitempty"`
	Encounter             *Reference           `bson:"encounter,omitempty" json:"encounter,omitempty"`
	AuthoredOn            *time.Time           `bson:"authoredOn,omitempty" json:"authoredOn,omitempty"`
	Author                *Reference           `bson:"author,omitempty" json:"author,omitempty"`
	ReasonCode            []CodeableConcept    `bson:"reasonCode,omitempty" json:"reasonCode,omitempty"`
	ReasonReference       []Reference          `bson:"reasonReference,omitempty" json:"reasonReference,omitempty"`
	Note                  []Annotation         `bson:"note,omitempty" json:"note,omitempty"`
	Action                []RequestGroupAction `bson:"action,omitempty" json:"action,omitempty"`
}
type RequestGroupAction struct {
	Id                  *string                           `bson:"id,omitempty" json:"id,omitempty"`
	Extension           []Extension                       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension   []Extension                       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Prefix              *string                           `bson:"prefix,omitempty" json:"prefix,omitempty"`
	Title               *string                           `bson:"title,omitempty" json:"title,omitempty"`
	Description         *string                           `bson:"description,omitempty" json:"description,omitempty"`
	TextEquivalent      *string                           `bson:"textEquivalent,omitempty" json:"textEquivalent,omitempty"`
	Priority            *RequestPriority                  `bson:"priority,omitempty" json:"priority,omitempty"`
	Code                []CodeableConcept                 `bson:"code,omitempty" json:"code,omitempty"`
	Documentation       []RelatedArtifact                 `bson:"documentation,omitempty" json:"documentation,omitempty"`
	Condition           []RequestGroupActionCondition     `bson:"condition,omitempty" json:"condition,omitempty"`
	RelatedAction       []RequestGroupActionRelatedAction `bson:"relatedAction,omitempty" json:"relatedAction,omitempty"`
	TimingDateTime      *time.Time                        `bson:"timingDateTime,omitempty" json:"timingDateTime,omitempty"`
	TimingAge           *Age                              `bson:"timingAge,omitempty" json:"timingAge,omitempty"`
	TimingPeriod        *Period                           `bson:"timingPeriod,omitempty" json:"timingPeriod,omitempty"`
	TimingDuration      *Duration                         `bson:"timingDuration,omitempty" json:"timingDuration,omitempty"`
	TimingRange         *Range                            `bson:"timingRange,omitempty" json:"timingRange,omitempty"`
	TimingTiming        *Timing                           `bson:"timingTiming,omitempty" json:"timingTiming,omitempty"`
	Participant         []Reference                       `bson:"participant,omitempty" json:"participant,omitempty"`
	Type                *CodeableConcept                  `bson:"type,omitempty" json:"type,omitempty"`
	GroupingBehavior    *ActionGroupingBehavior           `bson:"groupingBehavior,omitempty" json:"groupingBehavior,omitempty"`
	SelectionBehavior   *ActionSelectionBehavior          `bson:"selectionBehavior,omitempty" json:"selectionBehavior,omitempty"`
	RequiredBehavior    *ActionRequiredBehavior           `bson:"requiredBehavior,omitempty" json:"requiredBehavior,omitempty"`
	PrecheckBehavior    *ActionPrecheckBehavior           `bson:"precheckBehavior,omitempty" json:"precheckBehavior,omitempty"`
	CardinalityBehavior *ActionCardinalityBehavior        `bson:"cardinalityBehavior,omitempty" json:"cardinalityBehavior,omitempty"`
	Resource            *Reference                        `bson:"resource,omitempty" json:"resource,omitempty"`
	Action              []RequestGroupAction              `bson:"action,omitempty" json:"action,omitempty"`
}
type RequestGroupActionCondition struct {
	Id                *string             `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension         `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension         `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Kind              ActionConditionKind `bson:"kind" json:"kind,omitempty"`
	Expression        *Expression         `bson:"expression,omitempty" json:"expression,omitempty"`
}
type RequestGroupActionRelatedAction struct {
	Id                *string                `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension            `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension            `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	ActionId          string                 `bson:"actionId" json:"actionId,omitempty"`
	Relationship      ActionRelationshipType `bson:"relationship" json:"relationship,omitempty"`
	OffsetDuration    *Duration              `bson:"offsetDuration,omitempty" json:"offsetDuration,omitempty"`
	OffsetRange       *Range                 `bson:"offsetRange,omitempty" json:"offsetRange,omitempty"`
}
type OtherRequestGroup RequestGroup

// MarshalJSON marshals the given RequestGroup as JSON into a byte slice
func (r RequestGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherRequestGroup
		ResourceType string `json:"resourceType"`
	}{
		OtherRequestGroup: OtherRequestGroup(r),
		ResourceType:      "RequestGroup",
	})
}

// UnmarshalRequestGroup unmarshals a RequestGroup.
func UnmarshalRequestGroup(b []byte) (RequestGroup, error) {
	var requestGroup RequestGroup
	if err := json.Unmarshal(b, &requestGroup); err != nil {
		return requestGroup, err
	}
	return requestGroup, nil
}
