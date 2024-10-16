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
	Status            *PublicationStatus  `bson:"status,omitempty" json:"status,omitempty"`
	Experimental      *bool               `bson:"experimental,omitempty" json:"experimental,omitempty"`
	SubjectType       []ResourceType      `bson:"subjectType,omitempty" json:"subjectType,omitempty"`
	Date              *string             `bson:"date,omitempty" json:"date,omitempty"`
	Publisher         *string             `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Contact           []ContactDetail     `bson:"contact,omitempty" json:"contact,omitempty"`
	Description       *string             `bson:"description,omitempty" json:"description,omitempty"`
	UseContext        []UsageContext      `bson:"useContext,omitempty" json:"useContext,omitempty"`
	Jurisdiction      []CodeableConcept   `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	Purpose           *string             `bson:"purpose,omitempty" json:"purpose,omitempty"`
	Copyright         *string             `bson:"copyright,omitempty" json:"copyright,omitempty"`
	ApprovalDate      *string             `bson:"approvalDate,omitempty" json:"approvalDate,omitempty"`
	LastReviewDate    *string             `bson:"lastReviewDate,omitempty" json:"lastReviewDate,omitempty"`
	EffectivePeriod   *Period             `bson:"effectivePeriod,omitempty" json:"effectivePeriod,omitempty"`
	Code              []Coding            `bson:"code,omitempty" json:"code,omitempty"`
	Item              []QuestionnaireItem `bson:"item,omitempty" json:"item,omitempty"`
}
type QuestionnaireItem struct {
	Id                *string                         `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	LinkId            string                          `bson:"linkId,omitempty" json:"linkId,omitempty"`
	Definition        *string                         `bson:"definition,omitempty" json:"definition,omitempty"`
	Code              []Coding                        `bson:"code,omitempty" json:"code,omitempty"`
	Prefix            *string                         `bson:"prefix,omitempty" json:"prefix,omitempty"`
	Text              *string                         `bson:"text,omitempty" json:"text,omitempty"`
	Type              *QuestionnaireItemType          `bson:"type,omitempty" json:"type,omitempty"`
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
	Id                *string                    `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Question          *string                    `bson:"question,omitempty" json:"question,omitempty"`
	Operator          *QuestionnaireItemOperator `bson:"operator,omitempty" json:"operator,omitempty"`
	AnswerBoolean     *bool                      `bson:"answerBoolean,omitempty" json:"answerBoolean,omitempty"`
	AnswerDecimal     *float64                   `bson:"answerDecimal,omitempty" json:"answerDecimal,omitempty"`
	AnswerInteger     *int                       `bson:"answerInteger,omitempty" json:"answerInteger,omitempty"`
	AnswerDate        *string                    `bson:"answerDate,omitempty" json:"answerDate,omitempty"`
	AnswerDateTime    *string                    `bson:"answerDateTime,omitempty" json:"answerDateTime,omitempty"`
	AnswerTime        *string                    `bson:"answerTime,omitempty" json:"answerTime,omitempty"`
	AnswerString      *string                    `bson:"answerString,omitempty" json:"answerString,omitempty"`
	AnswerCoding      *Coding                    `bson:"answerCoding,omitempty" json:"answerCoding,omitempty"`
	AnswerQuantity    *Quantity                  `bson:"answerQuantity,omitempty" json:"answerQuantity,omitempty"`
	AnswerReference   *Reference                 `bson:"answerReference,omitempty" json:"answerReference,omitempty"`
}
type QuestionnaireItemAnswerOption struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	ValueInteger      *int        `bson:"valueInteger,omitempty" json:"valueInteger,omitempty"`
	ValueDate         *string     `bson:"valueDate,omitempty" json:"valueDate,omitempty"`
	ValueTime         *string     `bson:"valueTime,omitempty" json:"valueTime,omitempty"`
	ValueString       *string     `bson:"valueString,omitempty" json:"valueString,omitempty"`
	ValueCoding       *Coding     `bson:"valueCoding,omitempty" json:"valueCoding,omitempty"`
	ValueReference    *Reference  `bson:"valueReference,omitempty" json:"valueReference,omitempty"`
	InitialSelected   *bool       `bson:"initialSelected,omitempty" json:"initialSelected,omitempty"`
}
type QuestionnaireItemInitial struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	ValueBoolean      *bool       `bson:"valueBoolean,omitempty" json:"valueBoolean,omitempty"`
	ValueDecimal      *float64    `bson:"valueDecimal,omitempty" json:"valueDecimal,omitempty"`
	ValueInteger      *int        `bson:"valueInteger,omitempty" json:"valueInteger,omitempty"`
	ValueDate         *string     `bson:"valueDate,omitempty" json:"valueDate,omitempty"`
	ValueDateTime     *string     `bson:"valueDateTime,omitempty" json:"valueDateTime,omitempty"`
	ValueTime         *string     `bson:"valueTime,omitempty" json:"valueTime,omitempty"`
	ValueString       *string     `bson:"valueString,omitempty" json:"valueString,omitempty"`
	ValueUri          *string     `bson:"valueUri,omitempty" json:"valueUri,omitempty"`
	ValueAttachment   *Attachment `bson:"valueAttachment,omitempty" json:"valueAttachment,omitempty"`
	ValueCoding       *Coding     `bson:"valueCoding,omitempty" json:"valueCoding,omitempty"`
	ValueQuantity     *Quantity   `bson:"valueQuantity,omitempty" json:"valueQuantity,omitempty"`
	ValueReference    *Reference  `bson:"valueReference,omitempty" json:"valueReference,omitempty"`
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
