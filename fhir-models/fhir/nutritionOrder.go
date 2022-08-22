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

// NutritionOrder is documented here http://hl7.org/fhir/StructureDefinition/NutritionOrder
type NutritionOrder struct {
	Id                     *string                       `bson:"id,omitempty" json:"id,omitempty"`
	Meta                   *Meta                         `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules          *string                       `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language               *string                       `bson:"language,omitempty" json:"language,omitempty"`
	Text                   *Narrative                    `bson:"text,omitempty" json:"text,omitempty"`
	Extension              []Extension                   `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension      []Extension                   `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier             []Identifier                  `bson:"identifier,omitempty" json:"identifier,omitempty"`
	InstantiatesCanonical  []string                      `bson:"instantiatesCanonical,omitempty" json:"instantiatesCanonical,omitempty"`
	InstantiatesUri        []string                      `bson:"instantiatesUri,omitempty" json:"instantiatesUri,omitempty"`
	Instantiates           []string                      `bson:"instantiates,omitempty" json:"instantiates,omitempty"`
	Status                 RequestStatus                 `bson:"status" json:"status"`
	Intent                 RequestIntent                 `bson:"intent" json:"intent"`
	Patient                Reference                     `bson:"patient" json:"patient"`
	Encounter              *Reference                    `bson:"encounter,omitempty" json:"encounter,omitempty"`
	DateTime               string                        `bson:"dateTime" json:"dateTime"`
	Orderer                *Reference                    `bson:"orderer,omitempty" json:"orderer,omitempty"`
	AllergyIntolerance     []Reference                   `bson:"allergyIntolerance,omitempty" json:"allergyIntolerance,omitempty"`
	FoodPreferenceModifier []CodeableConcept             `bson:"foodPreferenceModifier,omitempty" json:"foodPreferenceModifier,omitempty"`
	ExcludeFoodModifier    []CodeableConcept             `bson:"excludeFoodModifier,omitempty" json:"excludeFoodModifier,omitempty"`
	OralDiet               *NutritionOrderOralDiet       `bson:"oralDiet,omitempty" json:"oralDiet,omitempty"`
	Supplement             []NutritionOrderSupplement    `bson:"supplement,omitempty" json:"supplement,omitempty"`
	EnteralFormula         *NutritionOrderEnteralFormula `bson:"enteralFormula,omitempty" json:"enteralFormula,omitempty"`
	Note                   []Annotation                  `bson:"note,omitempty" json:"note,omitempty"`
}
type NutritionOrderOralDiet struct {
	Id                   *string                          `bson:"id,omitempty" json:"id,omitempty"`
	Extension            []Extension                      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension    []Extension                      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type                 []CodeableConcept                `bson:"type,omitempty" json:"type,omitempty"`
	Schedule             []Timing                         `bson:"schedule,omitempty" json:"schedule,omitempty"`
	Nutrient             []NutritionOrderOralDietNutrient `bson:"nutrient,omitempty" json:"nutrient,omitempty"`
	Texture              []NutritionOrderOralDietTexture  `bson:"texture,omitempty" json:"texture,omitempty"`
	FluidConsistencyType []CodeableConcept                `bson:"fluidConsistencyType,omitempty" json:"fluidConsistencyType,omitempty"`
	Instruction          *string                          `bson:"instruction,omitempty" json:"instruction,omitempty"`
}
type NutritionOrderOralDietNutrient struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Modifier          *CodeableConcept `bson:"modifier,omitempty" json:"modifier,omitempty"`
	Amount            *Quantity        `bson:"amount,omitempty" json:"amount,omitempty"`
}
type NutritionOrderOralDietTexture struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Modifier          *CodeableConcept `bson:"modifier,omitempty" json:"modifier,omitempty"`
	FoodType          *CodeableConcept `bson:"foodType,omitempty" json:"foodType,omitempty"`
}
type NutritionOrderSupplement struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	ProductName       *string          `bson:"productName,omitempty" json:"productName,omitempty"`
	Schedule          []Timing         `bson:"schedule,omitempty" json:"schedule,omitempty"`
	Quantity          *Quantity        `bson:"quantity,omitempty" json:"quantity,omitempty"`
	Instruction       *string          `bson:"instruction,omitempty" json:"instruction,omitempty"`
}
type NutritionOrderEnteralFormula struct {
	Id                        *string                                      `bson:"id,omitempty" json:"id,omitempty"`
	Extension                 []Extension                                  `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension         []Extension                                  `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	BaseFormulaType           *CodeableConcept                             `bson:"baseFormulaType,omitempty" json:"baseFormulaType,omitempty"`
	BaseFormulaProductName    *string                                      `bson:"baseFormulaProductName,omitempty" json:"baseFormulaProductName,omitempty"`
	AdditiveType              *CodeableConcept                             `bson:"additiveType,omitempty" json:"additiveType,omitempty"`
	AdditiveProductName       *string                                      `bson:"additiveProductName,omitempty" json:"additiveProductName,omitempty"`
	CaloricDensity            *Quantity                                    `bson:"caloricDensity,omitempty" json:"caloricDensity,omitempty"`
	RouteofAdministration     *CodeableConcept                             `bson:"routeofAdministration,omitempty" json:"routeofAdministration,omitempty"`
	Administration            []NutritionOrderEnteralFormulaAdministration `bson:"administration,omitempty" json:"administration,omitempty"`
	MaxVolumeToDeliver        *Quantity                                    `bson:"maxVolumeToDeliver,omitempty" json:"maxVolumeToDeliver,omitempty"`
	AdministrationInstruction *string                                      `bson:"administrationInstruction,omitempty" json:"administrationInstruction,omitempty"`
}
type NutritionOrderEnteralFormulaAdministration struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Schedule          *Timing     `bson:"schedule,omitempty" json:"schedule,omitempty"`
	Quantity          *Quantity   `bson:"quantity,omitempty" json:"quantity,omitempty"`
}
type OtherNutritionOrder NutritionOrder

// MarshalJSON marshals the given NutritionOrder as JSON into a byte slice
func (r NutritionOrder) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherNutritionOrder
		ResourceType string `json:"resourceType"`
	}{
		OtherNutritionOrder: OtherNutritionOrder(r),
		ResourceType:        "NutritionOrder",
	})
}

// UnmarshalNutritionOrder unmarshals a NutritionOrder.
func UnmarshalNutritionOrder(b []byte) (NutritionOrder, error) {
	var nutritionOrder NutritionOrder
	if err := json.Unmarshal(b, &nutritionOrder); err != nil {
		return nutritionOrder, err
	}
	return nutritionOrder, nil
}
