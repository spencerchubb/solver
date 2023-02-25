#!/bin/sh

cargo build --target wasm32-unknown-unknown
wasm-pack build --target web
rm pkg/.gitignore
rm pkg/package.json
rm pkg/README.md
cp -r pkg/ ../solver-js/src/
