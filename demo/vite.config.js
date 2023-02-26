import { defineConfig } from 'vite'
import wasm from "vite-plugin-wasm";
import topLevelAwait from "vite-plugin-top-level-await";

export default defineConfig(({ command, mode, ssrBuild }) => {
    return {
        base: mode === 'development' ? '/' : '/solver/',
        minify: false,
        plugins: [
            wasm(),
            topLevelAwait(),
        ],
        worker: {
            format: "es",
            plugins: [
                wasm(),
                topLevelAwait(),
            ],
        },
    };
});