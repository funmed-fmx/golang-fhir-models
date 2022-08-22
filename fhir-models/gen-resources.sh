#!/usr/bin/env bash

wget -O definitions.zip https://www.hl7.org/fhir/definitions.json.zip
unzip definitions.zip definitions.json/profiles-resources.json definitions.json/profiles-types.json definitions.json/valuesets.json -d fhir
mv definitions.json/profiles-resources.json fhir/
mv definitions.json/profiles-types.json fhir/
mv definitions.json/valuesets.json fhir/
rm -rf definitions.zip definitions.json
rm definitions.zip

go generate ./fhir
