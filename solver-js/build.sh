#!/bin/sh

npm run build
rm -rf ./dist/pkg
cp -r pkg/ ./dist/
