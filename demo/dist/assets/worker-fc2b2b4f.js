let r,l=0,u=null;function d(){return(u===null||u.byteLength===0)&&(u=new Uint8Array(r.memory.buffer)),u}const m=new TextEncoder("utf-8"),L=typeof m.encodeInto=="function"?function(e,n){return m.encodeInto(e,n)}:function(e,n){const t=m.encode(e);return n.set(t),{read:e.length,written:t.length}};function w(e,n,t){if(t===void 0){const a=m.encode(e),c=n(a.length);return d().subarray(c,c+a.length).set(a),l=a.length,c}let o=e.length,i=n(o);const f=d();let s=0;for(;s<o;s++){const a=e.charCodeAt(s);if(a>127)break;f[i+s]=a}if(s!==o){s!==0&&(e=e.slice(s)),i=t(i,o,o=s+e.length*3);const a=d().subarray(i+s,i+o),c=L(e,a);s+=c.written}return l=s,i}let y=null;function S(){return(y===null||y.byteLength===0)&&(y=new Uint32Array(r.memory.buffer)),y}function _(e,n){const t=n(e.length*4);return S().set(e,t/4),l=e.length,t}let b=null;function g(){return(b===null||b.byteLength===0)&&(b=new Int32Array(r.memory.buffer)),b}const h=new TextDecoder("utf-8",{ignoreBOM:!0,fatal:!0});h.decode();function T(e,n){return h.decode(d().subarray(e,e+n))}function x(e,n,t,o,i){try{const a=r.__wbindgen_add_to_stack_pointer(-16),c=w(e,r.__wbindgen_malloc,r.__wbindgen_realloc),A=l,U=w(n,r.__wbindgen_malloc,r.__wbindgen_realloc),M=l,W=_(t,r.__wbindgen_malloc),I=l,R=_(o,r.__wbindgen_malloc),v=l;r.scramble(a,c,A,U,M,W,I,R,v,i);var f=g()[a/4+0],s=g()[a/4+1];return T(f,s)}finally{r.__wbindgen_add_to_stack_pointer(16),r.__wbindgen_free(f,s)}}async function E(e,n){if(typeof Response=="function"&&e instanceof Response){if(typeof WebAssembly.instantiateStreaming=="function")try{return await WebAssembly.instantiateStreaming(e,n)}catch(o){if(e.headers.get("Content-Type")!="application/wasm")console.warn("`WebAssembly.instantiateStreaming` failed because your server does not serve wasm with `application/wasm` MIME type. Falling back to `WebAssembly.instantiate` which is slower. Original error:\n",o);else throw o}const t=await e.arrayBuffer();return await WebAssembly.instantiate(t,n)}else{const t=await WebAssembly.instantiate(e,n);return t instanceof WebAssembly.Instance?{instance:t,module:e}:t}}function k(){const e={};return e.wbg={},e}function O(e,n){return r=e.exports,p.__wbindgen_wasm_module=n,b=null,y=null,u=null,r}async function p(e){typeof e>"u"&&(e=new URL("/assets/solver_wasm_bg-a6f6c119.wasm",self.location));const n=k();(typeof e=="string"||typeof Request=="function"&&e instanceof Request||typeof URL=="function"&&e instanceof URL)&&(e=fetch(e));const{instance:t,module:o}=await E(await e,n);return O(t,o)}onmessage=function(e){p(new URL("/assets/solver_wasm_bg-a6f6c119.wasm",self.location)).then(()=>{const{alg:n,moves:t,onlyOrientation:o,disregard:i,maxSolutions:f}=e.data,s=x(n,t,new Uint32Array(o),new Uint32Array(i),f);postMessage(s)})};