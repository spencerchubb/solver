wasm-pack build --target bundler
rm pkg/.gitignore
rm pkg/package.json
rm pkg/README.md
cp -r pkg/ ../solver-js/