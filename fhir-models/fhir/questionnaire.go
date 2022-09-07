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

// Questionnaire is documented here http://hl7.org/fhir/StructureDefinition/Questionnaire
type Questionnaire struct {
	Id                *string             `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta               `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string             `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string             `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative          `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension         `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension         `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Url               *string             `bson:"url,omitempty" json:"url,omitempty"`
	Identifier        []Identifier        `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Version           *string             `bson:"version,omitempty" json:"version,omitempty"`
	Name              *string             `bson:"name,omitempty" json:"name,omitempty"`
	Title             *string             `bson:"title,omitempty" json:"title,omitempty"`
	DerivedFrom       []string            `bson:"derivedFrom,omitempty" json:"derivedFrom,omitempty"`
	Status            PublicationStatus   `bson:"status" json:"status,omitempty"`
	Experimental      *bool               `bson:"experimental,omitempty" json:"experimental,omitempty"`
	SubjectType       []ResourceType      `bson:"subjectType,omitempty" json:"subjectType,omitempty"`
	Date              *time.Time          `bson:"date,omitempty" json:"date,omitempty"`
	Publisher         *string             `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Contact           []ContactDetail     `bson:"contact,omitempty" json:"contact,omitempty"`
	Description       *string             `bson:"description,omitempty" json:"description,omitempty"`
	UseContext        []UsageContext      `bson:"useContext,omitempty" json:"useContext,omitempty"`
	Jurisdiction      []CodeableConcept   `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	Purpose           *string             `bson:"purpose,omitempty" json:"purpose,omitempty"`
	Copyright         *string             `bson:"copyright,omitempty" json:"copyright,omitempty"`
	ApprovalDate      *time.Time          `bson:"approvalDate,omitempty" json:"approvalDate,omitempty"`
	LastReviewDate    *time.Time          `bson:"lastReviewDate,omitempty" json:"lastReviewDate,omitempty"`
	EffectivePeriod   *Period             `bson:"effectivePeriod,omitempty" json:"effectivePeriod,omitempty"`
	Code              []Coding            `bson:"code,omitempty" json:"code,omitempty"`
	Item              []QuestionnaireItem `bson:"item,omitempty" json:"item,omitempty"`
}
type QuestionnaireItem struct {
	Id                *string                         `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	LinkId            string                          `bson:"linkId" json:"linkId,omitempty"`
	Definition        *string                         `bson:"definition,omitempty" json:"definition,omitempty"`
	Code              []Coding                        `bson:"code,omitempty" json:"code,omitempty"`
	Prefix            *string                         `bson:"prefix,omitempty" json:"prefix,omitempty"`
	Text              *string                         `bson:"text,omitempty" json:"text,omitempty"`
	Type              QuestionnaireItemType           `bson:"type" json:"type,omitempty"`
	EnableWhen        []QuestionnaireItemEnableWhen   `bson:"enableWhen,omitempty" json:"enableWhen,omitempty"`
	EnableBehavior    *EnableWhenBehavior             `bson:"enableBehavior,omitempty" json:"enableBehavior,omitempty"`
	Required          *bool                           `bson:"required,omitempty" json:"required,omitempty"`
	Repeats           *bool                           `bson:"repeats,omitempty" json:"repeats,omitempty"`
	ReadOnly          *bool                           `bson:"readOnly,omitempty" json:"readOnly,omitempty"`
	MaxLength         *int                            `bson:"maxLength,omitempty" json:"maxLength,omitempty"`
	AnswerValueSet    *string                         `bson:"answerValueSet,omitempty" json:"answerValueSet,omitempty"`
	AnswerOption      []QuestionnaireItemAnswerOption `bson:"answerOption,omitempty" json:"answerOption,omitempty"`
	Initial           []QuestionnaireItemInitial      `bson:"initial,omitempty" json:"initial,omitempty"`
	Item              []QuestionnaireItem             `bson:"item,omitempty" json:"item,omitempty"`
}
type QuestionnaireItemEnableWhen struct {
	Id                *string                   `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension               `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension               `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Question          string                    `bson:"question" json:"question,omitempty"`
	Operator          QuestionnaireItemOperator `bson:"operator" json:"operator,omitempty"`
	AnswerBoolean     bool                      `bson:"answerBoolean" json:"answerBoolean,omitempty"`
	AnswerDecimal     float64                   `bson:"answerDecimal" json:"answerDecimal,omitempty"`
	AnswerInteger     int                       `bson:"answerInteger" json:"answerInteger,omitempty"`
	AnswerDate        time.Time                 `bson:"answerDate" json:"answerDate,omitempty"`
	AnswerDateTime    time.Time                 `bson:"answerDateTime" json:"answerDateTime,omitempty"`
	AnswerTime        string                    `bson:"answerTime" json:"answerTime,omitempty"`
	AnswerString      string                    `bson:"answerString" json:"answerString,omitempty"`
	AnswerCoding      Coding                    `bson:"answerCoding" json:"answerCoding,omitempty"`
	AnswerQuantity    Quantity                  `bson:"answerQuantity" json:"answerQuantity,omitempty"`
	AnswerReference   Reference                 `bson:"answerReference" json:"answerReference,omitempty"`
}
type QuestionnaireItemAnswerOption struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	ValueInteger      int         `bson:"valueInteger" json:"valueInteger,omitempty"`
	ValueDate         time.Time   `bson:"valueDate" json:"valueDate,omitempty"`
	ValueTime         string      `bson:"valueTime" json:"valueTime,omitempty"`
	ValueString       string      `bson:"valueString" json:"valueString,omitempty"`
	ValueCoding       Coding      `bson:"valueCoding" json:"valueCoding,omitempty"`
	ValueReference    Reference   `bson:"valueReference" json:"valueReference,omitempty"`
	InitialSelected   *bool       `bson:"initialSelected,omitempty" json:"initialSelected,omitempty"`
}
type QuestionnaireItemInitial struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	ValueBoolean      bool        `bson:"valueBoolean" json:"valueBoolean,omitempty"`
	ValueDecimal      float64     `bson:"valueDecimal" json:"valueDecimal,omitempty"`
	ValueInteger      int         `bson:"valueInteger" json:"valueInteger,omitempty"`
	ValueDate         time.Time   `bson:"valueDate" json:"valueDate,omitempty"`
	ValueDateTime     time.Time   `bson:"valueDateTime" json:"valueDateTime,omitempty"`
	ValueTime         string      `bson:"valueTime" json:"valueTime,omitempty"`
	ValueString       string      `bson:"valueString" json:"valueString,omitempty"`
	ValueUri          string      `bson:"valueUri" json:"valueUri,omitempty"`
	ValueAttachment   Attachment  `bson:"valueAttachment" json:"valueAttachment,omitempty"`
	ValueCoding       Coding      `bson:"valueCoding" json:"valueCoding,omitempty"`
	ValueQuantity     Quantity    `bson:"valueQuantity" json:"valueQuantity,omitempty"`
	ValueReference    Reference   `bson:"valueReference" json:"valueReference,omitempty"`
}
type OtherQuestionnaire Questionnaire

// MarshalJSON marshals the given Questionnaire as JSON into a byte slice
func (r Questionnaire) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherQuestionnaire
		ResourceType string `json:"resourceType"`
	}{
		OtherQuestionnaire: OtherQuestionnaire(r),
		ResourceType:       "Questionnaire",
	})
}

// UnmarshalQuestionnaire unmarshals a Questionnaire.
func UnmarshalQuestionnaire(b []byte) (Questionnaire, error) {
	var questionnaire Questionnaire
	if err := json.Unmarshal(b, &questionnaire); err != nil {
		return questionnaire, err
	}
	return questionnaire, nil
}
