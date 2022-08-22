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
	"fmt"
	"strings"
)

// THIS FILE IS GENERATED BY https://github.com/samply/golang-fhir-models
// PLEASE DO NOT EDIT BY HAND

// StructureMapModelMode is documented here http://hl7.org/fhir/ValueSet/map-model-mode
type StructureMapModelMode int

const (
	StructureMapModelModeSource StructureMapModelMode = iota
	StructureMapModelModeQueried
	StructureMapModelModeTarget
	StructureMapModelModeProduced
)

func (code StructureMapModelMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *StructureMapModelMode) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "source":
		*code = StructureMapModelModeSource
	case "queried":
		*code = StructureMapModelModeQueried
	case "target":
		*code = StructureMapModelModeTarget
	case "produced":
		*code = StructureMapModelModeProduced
	default:
		return fmt.Errorf("unknown StructureMapModelMode code `%s`", s)
	}
	return nil
}
func (code StructureMapModelMode) String() string {
	return code.Code()
}
func (code StructureMapModelMode) Code() string {
	switch code {
	case StructureMapModelModeSource:
		return "source"
	case StructureMapModelModeQueried:
		return "queried"
	case StructureMapModelModeTarget:
		return "target"
	case StructureMapModelModeProduced:
		return "produced"
	}
	return "<unknown>"
}
func (code StructureMapModelMode) Display() string {
	switch code {
	case StructureMapModelModeSource:
		return "Source Structure Definition"
	case StructureMapModelModeQueried:
		return "Queried Structure Definition"
	case StructureMapModelModeTarget:
		return "Target Structure Definition"
	case StructureMapModelModeProduced:
		return "Produced Structure Definition"
	}
	return "<unknown>"
}
func (code StructureMapModelMode) Definition() string {
	switch code {
	case StructureMapModelModeSource:
		return "This structure describes an instance passed to the mapping engine that is used a source of data."
	case StructureMapModelModeQueried:
		return "This structure describes an instance that the mapping engine may ask for that is used a source of data."
	case StructureMapModelModeTarget:
		return "This structure describes an instance passed to the mapping engine that is used a target of data."
	case StructureMapModelModeProduced:
		return "This structure describes an instance that the mapping engine may ask to create that is used a target of data."
	}
	return "<unknown>"
}
