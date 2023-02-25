#!/bin/sh

rm -rf dist/
tsc
cp src/worker.js dist/worker.js
cp -r src/pkg/ dist/pkg/
