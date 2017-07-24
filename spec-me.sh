#!/usr/bin/env bash

echo "Building data..."
find data/ -name '*.yml' -exec cat {} \; > data.yml

echo "Combining spec..."
multi-file-swagger index.yml > index.json

echo "Mustaching spec..."
mustache data.yml index.json > api-spec.json

echo "Cleaning up temporary files"
rm data.yml index.json
