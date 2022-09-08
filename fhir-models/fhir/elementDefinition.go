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

// THIS FILE IS GENERATED BY https://github.com/samply/golang-fhir-models
// PLEASE DO NOT EDIT BY HAND

// ElementDefinition is documented here http://hl7.org/fhir/StructureDefinition/ElementDefinition
type ElementDefinition struct {
	Id                              *string                       `bson:"id,omitempty" json:"id,omitempty"`
	Extension                       []Extension                   `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension               []Extension                   `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Path                            string                        `bson:"path,omitempty" json:"path,omitempty"`
	Representation                  []PropertyRepresentation      `bson:"representation,omitempty" json:"representation,omitempty"`
	SliceName                       *string                       `bson:"sliceName,omitempty" json:"sliceName,omitempty"`
	SliceIsConstraining             *bool                         `bson:"sliceIsConstraining,omitempty" json:"sliceIsConstraining,omitempty"`
	Label                           *string                       `bson:"label,omitempty" json:"label,omitempty"`
	Code                            []Coding                      `bson:"code,omitempty" json:"code,omitempty"`
	Slicing                         *ElementDefinitionSlicing     `bson:"slicing,omitempty" json:"slicing,omitempty"`
	Short                           *string                       `bson:"short,omitempty" json:"short,omitempty"`
	Definition                      *string                       `bson:"definition,omitempty" json:"definition,omitempty"`
	Comment                         *string                       `bson:"comment,omitempty" json:"comment,omitempty"`
	Requirements                    *string                       `bson:"requirements,omitempty" json:"requirements,omitempty"`
	Alias                           []string                      `bson:"alias,omitempty" json:"alias,omitempty"`
	Min                             *int                          `bson:"min,omitempty" json:"min,omitempty"`
	Max                             *string                       `bson:"max,omitempty" json:"max,omitempty"`
	Base                            *ElementDefinitionBase        `bson:"base,omitempty" json:"base,omitempty"`
	ContentReference                *string                       `bson:"contentReference,omitempty" json:"contentReference,omitempty"`
	Type                            []ElementDefinitionType       `bson:"type,omitempty" json:"type,omitempty"`
	DefaultValueBase64Binary        *string                       `bson:"defaultValueBase64Binary,omitempty" json:"defaultValueBase64Binary,omitempty"`
	DefaultValueBoolean             *bool                         `bson:"defaultValueBoolean,omitempty" json:"defaultValueBoolean,omitempty"`
	DefaultValueCanonical           *string                       `bson:"defaultValueCanonical,omitempty" json:"defaultValueCanonical,omitempty"`
	DefaultValueCode                *string                       `bson:"defaultValueCode,omitempty" json:"defaultValueCode,omitempty"`
	DefaultValueDate                *string                       `bson:"defaultValueDate,omitempty" json:"defaultValueDate,omitempty"`
	DefaultValueDateTime            *string                       `bson:"defaultValueDateTime,omitempty" json:"defaultValueDateTime,omitempty"`
	DefaultValueDecimal             *float64                      `bson:"defaultValueDecimal,omitempty" json:"defaultValueDecimal,omitempty"`
	DefaultValueId                  *string                       `bson:"defaultValueId,omitempty" json:"defaultValueId,omitempty"`
	DefaultValueInstant             *string                       `bson:"defaultValueInstant,omitempty" json:"defaultValueInstant,omitempty"`
	DefaultValueInteger             *int                          `bson:"defaultValueInteger,omitempty" json:"defaultValueInteger,omitempty"`
	DefaultValueMarkdown            *string                       `bson:"defaultValueMarkdown,omitempty" json:"defaultValueMarkdown,omitempty"`
	DefaultValueOid                 *string                       `bson:"defaultValueOid,omitempty" json:"defaultValueOid,omitempty"`
	DefaultValuePositiveInt         *int                          `bson:"defaultValuePositiveInt,omitempty" json:"defaultValuePositiveInt,omitempty"`
	DefaultValueString              *string                       `bson:"defaultValueString,omitempty" json:"defaultValueString,omitempty"`
	DefaultValueTime                *string                       `bson:"defaultValueTime,omitempty" json:"defaultValueTime,omitempty"`
	DefaultValueUnsignedInt         *int                          `bson:"defaultValueUnsignedInt,omitempty" json:"defaultValueUnsignedInt,omitempty"`
	DefaultValueUri                 *string                       `bson:"defaultValueUri,omitempty" json:"defaultValueUri,omitempty"`
	DefaultValueUrl                 *string                       `bson:"defaultValueUrl,omitempty" json:"defaultValueUrl,omitempty"`
	DefaultValueUuid                *string                       `bson:"defaultValueUuid,omitempty" json:"defaultValueUuid,omitempty"`
	DefaultValueAddress             *Address                      `bson:"defaultValueAddress,omitempty" json:"defaultValueAddress,omitempty"`
	DefaultValueAge                 *Age                          `bson:"defaultValueAge,omitempty" json:"defaultValueAge,omitempty"`
	DefaultValueAnnotation          *Annotation                   `bson:"defaultValueAnnotation,omitempty" json:"defaultValueAnnotation,omitempty"`
	DefaultValueAttachment          *Attachment                   `bson:"defaultValueAttachment,omitempty" json:"defaultValueAttachment,omitempty"`
	DefaultValueCodeableConcept     *CodeableConcept              `bson:"defaultValueCodeableConcept,omitempty" json:"defaultValueCodeableConcept,omitempty"`
	DefaultValueCodeableReference   *CodeableReference            `bson:"defaultValueCodeableReference,omitempty" json:"defaultValueCodeableReference,omitempty"`
	DefaultValueCoding              *Coding                       `bson:"defaultValueCoding,omitempty" json:"defaultValueCoding,omitempty"`
	DefaultValueContactPoint        *ContactPoint                 `bson:"defaultValueContactPoint,omitempty" json:"defaultValueContactPoint,omitempty"`
	DefaultValueCount               *Count                        `bson:"defaultValueCount,omitempty" json:"defaultValueCount,omitempty"`
	DefaultValueDistance            *Distance                     `bson:"defaultValueDistance,omitempty" json:"defaultValueDistance,omitempty"`
	DefaultValueDuration            *Duration                     `bson:"defaultValueDuration,omitempty" json:"defaultValueDuration,omitempty"`
	DefaultValueHumanName           *HumanName                    `bson:"defaultValueHumanName,omitempty" json:"defaultValueHumanName,omitempty"`
	DefaultValueIdentifier          *Identifier                   `bson:"defaultValueIdentifier,omitempty" json:"defaultValueIdentifier,omitempty"`
	DefaultValueMoney               *Money                        `bson:"defaultValueMoney,omitempty" json:"defaultValueMoney,omitempty"`
	DefaultValuePeriod              *Period                       `bson:"defaultValuePeriod,omitempty" json:"defaultValuePeriod,omitempty"`
	DefaultValueQuantity            *Quantity                     `bson:"defaultValueQuantity,omitempty" json:"defaultValueQuantity,omitempty"`
	DefaultValueRange               *Range                        `bson:"defaultValueRange,omitempty" json:"defaultValueRange,omitempty"`
	DefaultValueRatio               *Ratio                        `bson:"defaultValueRatio,omitempty" json:"defaultValueRatio,omitempty"`
	DefaultValueRatioRange          *RatioRange                   `bson:"defaultValueRatioRange,omitempty" json:"defaultValueRatioRange,omitempty"`
	DefaultValueReference           *Reference                    `bson:"defaultValueReference,omitempty" json:"defaultValueReference,omitempty"`
	DefaultValueSampledData         *SampledData                  `bson:"defaultValueSampledData,omitempty" json:"defaultValueSampledData,omitempty"`
	DefaultValueSignature           *Signature                    `bson:"defaultValueSignature,omitempty" json:"defaultValueSignature,omitempty"`
	DefaultValueTiming              *Timing                       `bson:"defaultValueTiming,omitempty" json:"defaultValueTiming,omitempty"`
	DefaultValueContactDetail       *ContactDetail                `bson:"defaultValueContactDetail,omitempty" json:"defaultValueContactDetail,omitempty"`
	DefaultValueContributor         *Contributor                  `bson:"defaultValueContributor,omitempty" json:"defaultValueContributor,omitempty"`
	DefaultValueDataRequirement     *DataRequirement              `bson:"defaultValueDataRequirement,omitempty" json:"defaultValueDataRequirement,omitempty"`
	DefaultValueExpression          *Expression                   `bson:"defaultValueExpression,omitempty" json:"defaultValueExpression,omitempty"`
	DefaultValueParameterDefinition *ParameterDefinition          `bson:"defaultValueParameterDefinition,omitempty" json:"defaultValueParameterDefinition,omitempty"`
	DefaultValueRelatedArtifact     *RelatedArtifact              `bson:"defaultValueRelatedArtifact,omitempty" json:"defaultValueRelatedArtifact,omitempty"`
	DefaultValueTriggerDefinition   *TriggerDefinition            `bson:"defaultValueTriggerDefinition,omitempty" json:"defaultValueTriggerDefinition,omitempty"`
	DefaultValueUsageContext        *UsageContext                 `bson:"defaultValueUsageContext,omitempty" json:"defaultValueUsageContext,omitempty"`
	DefaultValueDosage              *Dosage                       `bson:"defaultValueDosage,omitempty" json:"defaultValueDosage,omitempty"`
	MeaningWhenMissing              *string                       `bson:"meaningWhenMissing,omitempty" json:"meaningWhenMissing,omitempty"`
	OrderMeaning                    *string                       `bson:"orderMeaning,omitempty" json:"orderMeaning,omitempty"`
	FixedBase64Binary               *string                       `bson:"fixedBase64Binary,omitempty" json:"fixedBase64Binary,omitempty"`
	FixedBoolean                    *bool                         `bson:"fixedBoolean,omitempty" json:"fixedBoolean,omitempty"`
	FixedCanonical                  *string                       `bson:"fixedCanonical,omitempty" json:"fixedCanonical,omitempty"`
	FixedCode                       *string                       `bson:"fixedCode,omitempty" json:"fixedCode,omitempty"`
	FixedDate                       *string                       `bson:"fixedDate,omitempty" json:"fixedDate,omitempty"`
	FixedDateTime                   *string                       `bson:"fixedDateTime,omitempty" json:"fixedDateTime,omitempty"`
	FixedDecimal                    *float64                      `bson:"fixedDecimal,omitempty" json:"fixedDecimal,omitempty"`
	FixedId                         *string                       `bson:"fixedId,omitempty" json:"fixedId,omitempty"`
	FixedInstant                    *string                       `bson:"fixedInstant,omitempty" json:"fixedInstant,omitempty"`
	FixedInteger                    *int                          `bson:"fixedInteger,omitempty" json:"fixedInteger,omitempty"`
	FixedMarkdown                   *string                       `bson:"fixedMarkdown,omitempty" json:"fixedMarkdown,omitempty"`
	FixedOid                        *string                       `bson:"fixedOid,omitempty" json:"fixedOid,omitempty"`
	FixedPositiveInt                *int                          `bson:"fixedPositiveInt,omitempty" json:"fixedPositiveInt,omitempty"`
	FixedString                     *string                       `bson:"fixedString,omitempty" json:"fixedString,omitempty"`
	FixedTime                       *string                       `bson:"fixedTime,omitempty" json:"fixedTime,omitempty"`
	FixedUnsignedInt                *int                          `bson:"fixedUnsignedInt,omitempty" json:"fixedUnsignedInt,omitempty"`
	FixedUri                        *string                       `bson:"fixedUri,omitempty" json:"fixedUri,omitempty"`
	FixedUrl                        *string                       `bson:"fixedUrl,omitempty" json:"fixedUrl,omitempty"`
	FixedUuid                       *string                       `bson:"fixedUuid,omitempty" json:"fixedUuid,omitempty"`
	FixedAddress                    *Address                      `bson:"fixedAddress,omitempty" json:"fixedAddress,omitempty"`
	FixedAge                        *Age                          `bson:"fixedAge,omitempty" json:"fixedAge,omitempty"`
	FixedAnnotation                 *Annotation                   `bson:"fixedAnnotation,omitempty" json:"fixedAnnotation,omitempty"`
	FixedAttachment                 *Attachment                   `bson:"fixedAttachment,omitempty" json:"fixedAttachment,omitempty"`
	FixedCodeableConcept            *CodeableConcept              `bson:"fixedCodeableConcept,omitempty" json:"fixedCodeableConcept,omitempty"`
	FixedCodeableReference          *CodeableReference            `bson:"fixedCodeableReference,omitempty" json:"fixedCodeableReference,omitempty"`
	FixedCoding                     *Coding                       `bson:"fixedCoding,omitempty" json:"fixedCoding,omitempty"`
	FixedContactPoint               *ContactPoint                 `bson:"fixedContactPoint,omitempty" json:"fixedContactPoint,omitempty"`
	FixedCount                      *Count                        `bson:"fixedCount,omitempty" json:"fixedCount,omitempty"`
	FixedDistance                   *Distance                     `bson:"fixedDistance,omitempty" json:"fixedDistance,omitempty"`
	FixedDuration                   *Duration                     `bson:"fixedDuration,omitempty" json:"fixedDuration,omitempty"`
	FixedHumanName                  *HumanName                    `bson:"fixedHumanName,omitempty" json:"fixedHumanName,omitempty"`
	FixedIdentifier                 *Identifier                   `bson:"fixedIdentifier,omitempty" json:"fixedIdentifier,omitempty"`
	FixedMoney                      *Money                        `bson:"fixedMoney,omitempty" json:"fixedMoney,omitempty"`
	FixedPeriod                     *Period                       `bson:"fixedPeriod,omitempty" json:"fixedPeriod,omitempty"`
	FixedQuantity                   *Quantity                     `bson:"fixedQuantity,omitempty" json:"fixedQuantity,omitempty"`
	FixedRange                      *Range                        `bson:"fixedRange,omitempty" json:"fixedRange,omitempty"`
	FixedRatio                      *Ratio                        `bson:"fixedRatio,omitempty" json:"fixedRatio,omitempty"`
	FixedRatioRange                 *RatioRange                   `bson:"fixedRatioRange,omitempty" json:"fixedRatioRange,omitempty"`
	FixedReference                  *Reference                    `bson:"fixedReference,omitempty" json:"fixedReference,omitempty"`
	FixedSampledData                *SampledData                  `bson:"fixedSampledData,omitempty" json:"fixedSampledData,omitempty"`
	FixedSignature                  *Signature                    `bson:"fixedSignature,omitempty" json:"fixedSignature,omitempty"`
	FixedTiming                     *Timing                       `bson:"fixedTiming,omitempty" json:"fixedTiming,omitempty"`
	FixedContactDetail              *ContactDetail                `bson:"fixedContactDetail,omitempty" json:"fixedContactDetail,omitempty"`
	FixedContributor                *Contributor                  `bson:"fixedContributor,omitempty" json:"fixedContributor,omitempty"`
	FixedDataRequirement            *DataRequirement              `bson:"fixedDataRequirement,omitempty" json:"fixedDataRequirement,omitempty"`
	FixedExpression                 *Expression                   `bson:"fixedExpression,omitempty" json:"fixedExpression,omitempty"`
	FixedParameterDefinition        *ParameterDefinition          `bson:"fixedParameterDefinition,omitempty" json:"fixedParameterDefinition,omitempty"`
	FixedRelatedArtifact            *RelatedArtifact              `bson:"fixedRelatedArtifact,omitempty" json:"fixedRelatedArtifact,omitempty"`
	FixedTriggerDefinition          *TriggerDefinition            `bson:"fixedTriggerDefinition,omitempty" json:"fixedTriggerDefinition,omitempty"`
	FixedUsageContext               *UsageContext                 `bson:"fixedUsageContext,omitempty" json:"fixedUsageContext,omitempty"`
	FixedDosage                     *Dosage                       `bson:"fixedDosage,omitempty" json:"fixedDosage,omitempty"`
	PatternBase64Binary             *string                       `bson:"patternBase64Binary,omitempty" json:"patternBase64Binary,omitempty"`
	PatternBoolean                  *bool                         `bson:"patternBoolean,omitempty" json:"patternBoolean,omitempty"`
	PatternCanonical                *string                       `bson:"patternCanonical,omitempty" json:"patternCanonical,omitempty"`
	PatternCode                     *string                       `bson:"patternCode,omitempty" json:"patternCode,omitempty"`
	PatternDate                     *string                       `bson:"patternDate,omitempty" json:"patternDate,omitempty"`
	PatternDateTime                 *string                       `bson:"patternDateTime,omitempty" json:"patternDateTime,omitempty"`
	PatternDecimal                  *float64                      `bson:"patternDecimal,omitempty" json:"patternDecimal,omitempty"`
	PatternId                       *string                       `bson:"patternId,omitempty" json:"patternId,omitempty"`
	PatternInstant                  *string                       `bson:"patternInstant,omitempty" json:"patternInstant,omitempty"`
	PatternInteger                  *int                          `bson:"patternInteger,omitempty" json:"patternInteger,omitempty"`
	PatternMarkdown                 *string                       `bson:"patternMarkdown,omitempty" json:"patternMarkdown,omitempty"`
	PatternOid                      *string                       `bson:"patternOid,omitempty" json:"patternOid,omitempty"`
	PatternPositiveInt              *int                          `bson:"patternPositiveInt,omitempty" json:"patternPositiveInt,omitempty"`
	PatternString                   *string                       `bson:"patternString,omitempty" json:"patternString,omitempty"`
	PatternTime                     *string                       `bson:"patternTime,omitempty" json:"patternTime,omitempty"`
	PatternUnsignedInt              *int                          `bson:"patternUnsignedInt,omitempty" json:"patternUnsignedInt,omitempty"`
	PatternUri                      *string                       `bson:"patternUri,omitempty" json:"patternUri,omitempty"`
	PatternUrl                      *string                       `bson:"patternUrl,omitempty" json:"patternUrl,omitempty"`
	PatternUuid                     *string                       `bson:"patternUuid,omitempty" json:"patternUuid,omitempty"`
	PatternAddress                  *Address                      `bson:"patternAddress,omitempty" json:"patternAddress,omitempty"`
	PatternAge                      *Age                          `bson:"patternAge,omitempty" json:"patternAge,omitempty"`
	PatternAnnotation               *Annotation                   `bson:"patternAnnotation,omitempty" json:"patternAnnotation,omitempty"`
	PatternAttachment               *Attachment                   `bson:"patternAttachment,omitempty" json:"patternAttachment,omitempty"`
	PatternCodeableConcept          *CodeableConcept              `bson:"patternCodeableConcept,omitempty" json:"patternCodeableConcept,omitempty"`
	PatternCodeableReference        *CodeableReference            `bson:"patternCodeableReference,omitempty" json:"patternCodeableReference,omitempty"`
	PatternCoding                   *Coding                       `bson:"patternCoding,omitempty" json:"patternCoding,omitempty"`
	PatternContactPoint             *ContactPoint                 `bson:"patternContactPoint,omitempty" json:"patternContactPoint,omitempty"`
	PatternCount                    *Count                        `bson:"patternCount,omitempty" json:"patternCount,omitempty"`
	PatternDistance                 *Distance                     `bson:"patternDistance,omitempty" json:"patternDistance,omitempty"`
	PatternDuration                 *Duration                     `bson:"patternDuration,omitempty" json:"patternDuration,omitempty"`
	PatternHumanName                *HumanName                    `bson:"patternHumanName,omitempty" json:"patternHumanName,omitempty"`
	PatternIdentifier               *Identifier                   `bson:"patternIdentifier,omitempty" json:"patternIdentifier,omitempty"`
	PatternMoney                    *Money                        `bson:"patternMoney,omitempty" json:"patternMoney,omitempty"`
	PatternPeriod                   *Period                       `bson:"patternPeriod,omitempty" json:"patternPeriod,omitempty"`
	PatternQuantity                 *Quantity                     `bson:"patternQuantity,omitempty" json:"patternQuantity,omitempty"`
	PatternRange                    *Range                        `bson:"patternRange,omitempty" json:"patternRange,omitempty"`
	PatternRatio                    *Ratio                        `bson:"patternRatio,omitempty" json:"patternRatio,omitempty"`
	PatternRatioRange               *RatioRange                   `bson:"patternRatioRange,omitempty" json:"patternRatioRange,omitempty"`
	PatternReference                *Reference                    `bson:"patternReference,omitempty" json:"patternReference,omitempty"`
	PatternSampledData              *SampledData                  `bson:"patternSampledData,omitempty" json:"patternSampledData,omitempty"`
	PatternSignature                *Signature                    `bson:"patternSignature,omitempty" json:"patternSignature,omitempty"`
	PatternTiming                   *Timing                       `bson:"patternTiming,omitempty" json:"patternTiming,omitempty"`
	PatternContactDetail            *ContactDetail                `bson:"patternContactDetail,omitempty" json:"patternContactDetail,omitempty"`
	PatternContributor              *Contributor                  `bson:"patternContributor,omitempty" json:"patternContributor,omitempty"`
	PatternDataRequirement          *DataRequirement              `bson:"patternDataRequirement,omitempty" json:"patternDataRequirement,omitempty"`
	PatternExpression               *Expression                   `bson:"patternExpression,omitempty" json:"patternExpression,omitempty"`
	PatternParameterDefinition      *ParameterDefinition          `bson:"patternParameterDefinition,omitempty" json:"patternParameterDefinition,omitempty"`
	PatternRelatedArtifact          *RelatedArtifact              `bson:"patternRelatedArtifact,omitempty" json:"patternRelatedArtifact,omitempty"`
	PatternTriggerDefinition        *TriggerDefinition            `bson:"patternTriggerDefinition,omitempty" json:"patternTriggerDefinition,omitempty"`
	PatternUsageContext             *UsageContext                 `bson:"patternUsageContext,omitempty" json:"patternUsageContext,omitempty"`
	PatternDosage                   *Dosage                       `bson:"patternDosage,omitempty" json:"patternDosage,omitempty"`
	Example                         []ElementDefinitionExample    `bson:"example,omitempty" json:"example,omitempty"`
	MinValueDate                    *string                       `bson:"minValueDate,omitempty" json:"minValueDate,omitempty"`
	MinValueDateTime                *string                       `bson:"minValueDateTime,omitempty" json:"minValueDateTime,omitempty"`
	MinValueInstant                 *string                       `bson:"minValueInstant,omitempty" json:"minValueInstant,omitempty"`
	MinValueTime                    *string                       `bson:"minValueTime,omitempty" json:"minValueTime,omitempty"`
	MinValueDecimal                 *float64                      `bson:"minValueDecimal,omitempty" json:"minValueDecimal,omitempty"`
	MinValueInteger                 *int                          `bson:"minValueInteger,omitempty" json:"minValueInteger,omitempty"`
	MinValuePositiveInt             *int                          `bson:"minValuePositiveInt,omitempty" json:"minValuePositiveInt,omitempty"`
	MinValueUnsignedInt             *int                          `bson:"minValueUnsignedInt,omitempty" json:"minValueUnsignedInt,omitempty"`
	MinValueQuantity                *Quantity                     `bson:"minValueQuantity,omitempty" json:"minValueQuantity,omitempty"`
	MaxValueDate                    *string                       `bson:"maxValueDate,omitempty" json:"maxValueDate,omitempty"`
	MaxValueDateTime                *string                       `bson:"maxValueDateTime,omitempty" json:"maxValueDateTime,omitempty"`
	MaxValueInstant                 *string                       `bson:"maxValueInstant,omitempty" json:"maxValueInstant,omitempty"`
	MaxValueTime                    *string                       `bson:"maxValueTime,omitempty" json:"maxValueTime,omitempty"`
	MaxValueDecimal                 *float64                      `bson:"maxValueDecimal,omitempty" json:"maxValueDecimal,omitempty"`
	MaxValueInteger                 *int                          `bson:"maxValueInteger,omitempty" json:"maxValueInteger,omitempty"`
	MaxValuePositiveInt             *int                          `bson:"maxValuePositiveInt,omitempty" json:"maxValuePositiveInt,omitempty"`
	MaxValueUnsignedInt             *int                          `bson:"maxValueUnsignedInt,omitempty" json:"maxValueUnsignedInt,omitempty"`
	MaxValueQuantity                *Quantity                     `bson:"maxValueQuantity,omitempty" json:"maxValueQuantity,omitempty"`
	MaxLength                       *int                          `bson:"maxLength,omitempty" json:"maxLength,omitempty"`
	Condition                       []string                      `bson:"condition,omitempty" json:"condition,omitempty"`
	Constraint                      []ElementDefinitionConstraint `bson:"constraint,omitempty" json:"constraint,omitempty"`
	MustSupport                     *bool                         `bson:"mustSupport,omitempty" json:"mustSupport,omitempty"`
	IsModifier                      *bool                         `bson:"isModifier,omitempty" json:"isModifier,omitempty"`
	IsModifierReason                *string                       `bson:"isModifierReason,omitempty" json:"isModifierReason,omitempty"`
	IsSummary                       *bool                         `bson:"isSummary,omitempty" json:"isSummary,omitempty"`
	Binding                         *ElementDefinitionBinding     `bson:"binding,omitempty" json:"binding,omitempty"`
	Mapping                         []ElementDefinitionMapping    `bson:"mapping,omitempty" json:"mapping,omitempty"`
}
type ElementDefinitionSlicing struct {
	Id            *string                                 `bson:"id,omitempty" json:"id,omitempty"`
	Extension     []Extension                             `bson:"extension,omitempty" json:"extension,omitempty"`
	Discriminator []ElementDefinitionSlicingDiscriminator `bson:"discriminator,omitempty" json:"discriminator,omitempty"`
	Description   *string                                 `bson:"description,omitempty" json:"description,omitempty"`
	Ordered       *bool                                   `bson:"ordered,omitempty" json:"ordered,omitempty"`
	Rules         SlicingRules                            `bson:"rules,omitempty" json:"rules,omitempty"`
}
type ElementDefinitionSlicingDiscriminator struct {
	Id        *string           `bson:"id,omitempty" json:"id,omitempty"`
	Extension []Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	Type      DiscriminatorType `bson:"type,omitempty" json:"type,omitempty"`
	Path      string            `bson:"path,omitempty" json:"path,omitempty"`
}
type ElementDefinitionBase struct {
	Id        *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	Path      string      `bson:"path,omitempty" json:"path,omitempty"`
	Min       int         `bson:"min,omitempty" json:"min,omitempty"`
	Max       string      `bson:"max,omitempty" json:"max,omitempty"`
}
type ElementDefinitionType struct {
	Id            *string                `bson:"id,omitempty" json:"id,omitempty"`
	Extension     []Extension            `bson:"extension,omitempty" json:"extension,omitempty"`
	Code          string                 `bson:"code,omitempty" json:"code,omitempty"`
	Profile       []string               `bson:"profile,omitempty" json:"profile,omitempty"`
	TargetProfile []string               `bson:"targetProfile,omitempty" json:"targetProfile,omitempty"`
	Aggregation   []AggregationMode      `bson:"aggregation,omitempty" json:"aggregation,omitempty"`
	Versioning    *ReferenceVersionRules `bson:"versioning,omitempty" json:"versioning,omitempty"`
}
type ElementDefinitionExample struct {
	Id                       *string             `bson:"id,omitempty" json:"id,omitempty"`
	Extension                []Extension         `bson:"extension,omitempty" json:"extension,omitempty"`
	Label                    string              `bson:"label,omitempty" json:"label,omitempty"`
	ValueBase64Binary        string              `bson:"valueBase64Binary,omitempty" json:"valueBase64Binary,omitempty"`
	ValueBoolean             bool                `bson:"valueBoolean,omitempty" json:"valueBoolean,omitempty"`
	ValueCanonical           string              `bson:"valueCanonical,omitempty" json:"valueCanonical,omitempty"`
	ValueCode                string              `bson:"valueCode,omitempty" json:"valueCode,omitempty"`
	ValueDate                string              `bson:"valueDate,omitempty" json:"valueDate,omitempty"`
	ValueDateTime            string              `bson:"valueDateTime,omitempty" json:"valueDateTime,omitempty"`
	ValueDecimal             float64             `bson:"valueDecimal,omitempty" json:"valueDecimal,omitempty"`
	ValueId                  string              `bson:"valueId,omitempty" json:"valueId,omitempty"`
	ValueInstant             string              `bson:"valueInstant,omitempty" json:"valueInstant,omitempty"`
	ValueInteger             int                 `bson:"valueInteger,omitempty" json:"valueInteger,omitempty"`
	ValueMarkdown            string              `bson:"valueMarkdown,omitempty" json:"valueMarkdown,omitempty"`
	ValueOid                 string              `bson:"valueOid,omitempty" json:"valueOid,omitempty"`
	ValuePositiveInt         int                 `bson:"valuePositiveInt,omitempty" json:"valuePositiveInt,omitempty"`
	ValueString              string              `bson:"valueString,omitempty" json:"valueString,omitempty"`
	ValueTime                string              `bson:"valueTime,omitempty" json:"valueTime,omitempty"`
	ValueUnsignedInt         int                 `bson:"valueUnsignedInt,omitempty" json:"valueUnsignedInt,omitempty"`
	ValueUri                 string              `bson:"valueUri,omitempty" json:"valueUri,omitempty"`
	ValueUrl                 string              `bson:"valueUrl,omitempty" json:"valueUrl,omitempty"`
	ValueUuid                string              `bson:"valueUuid,omitempty" json:"valueUuid,omitempty"`
	ValueAddress             Address             `bson:"valueAddress,omitempty" json:"valueAddress,omitempty"`
	ValueAge                 Age                 `bson:"valueAge,omitempty" json:"valueAge,omitempty"`
	ValueAnnotation          Annotation          `bson:"valueAnnotation,omitempty" json:"valueAnnotation,omitempty"`
	ValueAttachment          Attachment          `bson:"valueAttachment,omitempty" json:"valueAttachment,omitempty"`
	ValueCodeableConcept     CodeableConcept     `bson:"valueCodeableConcept,omitempty" json:"valueCodeableConcept,omitempty"`
	ValueCodeableReference   CodeableReference   `bson:"valueCodeableReference,omitempty" json:"valueCodeableReference,omitempty"`
	ValueCoding              Coding              `bson:"valueCoding,omitempty" json:"valueCoding,omitempty"`
	ValueContactPoint        ContactPoint        `bson:"valueContactPoint,omitempty" json:"valueContactPoint,omitempty"`
	ValueCount               Count               `bson:"valueCount,omitempty" json:"valueCount,omitempty"`
	ValueDistance            Distance            `bson:"valueDistance,omitempty" json:"valueDistance,omitempty"`
	ValueDuration            Duration            `bson:"valueDuration,omitempty" json:"valueDuration,omitempty"`
	ValueHumanName           HumanName           `bson:"valueHumanName,omitempty" json:"valueHumanName,omitempty"`
	ValueIdentifier          Identifier          `bson:"valueIdentifier,omitempty" json:"valueIdentifier,omitempty"`
	ValueMoney               Money               `bson:"valueMoney,omitempty" json:"valueMoney,omitempty"`
	ValuePeriod              Period              `bson:"valuePeriod,omitempty" json:"valuePeriod,omitempty"`
	ValueQuantity            Quantity            `bson:"valueQuantity,omitempty" json:"valueQuantity,omitempty"`
	ValueRange               Range               `bson:"valueRange,omitempty" json:"valueRange,omitempty"`
	ValueRatio               Ratio               `bson:"valueRatio,omitempty" json:"valueRatio,omitempty"`
	ValueRatioRange          RatioRange          `bson:"valueRatioRange,omitempty" json:"valueRatioRange,omitempty"`
	ValueReference           Reference           `bson:"valueReference,omitempty" json:"valueReference,omitempty"`
	ValueSampledData         SampledData         `bson:"valueSampledData,omitempty" json:"valueSampledData,omitempty"`
	ValueSignature           Signature           `bson:"valueSignature,omitempty" json:"valueSignature,omitempty"`
	ValueTiming              Timing              `bson:"valueTiming,omitempty" json:"valueTiming,omitempty"`
	ValueContactDetail       ContactDetail       `bson:"valueContactDetail,omitempty" json:"valueContactDetail,omitempty"`
	ValueContributor         Contributor         `bson:"valueContributor,omitempty" json:"valueContributor,omitempty"`
	ValueDataRequirement     DataRequirement     `bson:"valueDataRequirement,omitempty" json:"valueDataRequirement,omitempty"`
	ValueExpression          Expression          `bson:"valueExpression,omitempty" json:"valueExpression,omitempty"`
	ValueParameterDefinition ParameterDefinition `bson:"valueParameterDefinition,omitempty" json:"valueParameterDefinition,omitempty"`
	ValueRelatedArtifact     RelatedArtifact     `bson:"valueRelatedArtifact,omitempty" json:"valueRelatedArtifact,omitempty"`
	ValueTriggerDefinition   TriggerDefinition   `bson:"valueTriggerDefinition,omitempty" json:"valueTriggerDefinition,omitempty"`
	ValueUsageContext        UsageContext        `bson:"valueUsageContext,omitempty" json:"valueUsageContext,omitempty"`
	ValueDosage              Dosage              `bson:"valueDosage,omitempty" json:"valueDosage,omitempty"`
}
type ElementDefinitionConstraint struct {
	Id           *string            `bson:"id,omitempty" json:"id,omitempty"`
	Extension    []Extension        `bson:"extension,omitempty" json:"extension,omitempty"`
	Key          string             `bson:"key,omitempty" json:"key,omitempty"`
	Requirements *string            `bson:"requirements,omitempty" json:"requirements,omitempty"`
	Severity     ConstraintSeverity `bson:"severity,omitempty" json:"severity,omitempty"`
	Human        string             `bson:"human,omitempty" json:"human,omitempty"`
	Expression   *string            `bson:"expression,omitempty" json:"expression,omitempty"`
	Xpath        *string            `bson:"xpath,omitempty" json:"xpath,omitempty"`
	Source       *string            `bson:"source,omitempty" json:"source,omitempty"`
}
type ElementDefinitionBinding struct {
	Id          *string         `bson:"id,omitempty" json:"id,omitempty"`
	Extension   []Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	Strength    BindingStrength `bson:"strength,omitempty" json:"strength,omitempty"`
	Description *string         `bson:"description,omitempty" json:"description,omitempty"`
	ValueSet    *string         `bson:"valueSet,omitempty" json:"valueSet,omitempty"`
}
type ElementDefinitionMapping struct {
	Id        *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	Identity  string      `bson:"identity,omitempty" json:"identity,omitempty"`
	Language  *string     `bson:"language,omitempty" json:"language,omitempty"`
	Map       string      `bson:"map,omitempty" json:"map,omitempty"`
	Comment   *string     `bson:"comment,omitempty" json:"comment,omitempty"`
}
