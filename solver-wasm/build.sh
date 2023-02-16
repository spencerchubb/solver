wasm-pack build --target bundler
rm pkg/.gitignore
cp -r pkg/ ../solver-js/