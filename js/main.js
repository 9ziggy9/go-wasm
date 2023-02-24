function init() {
  const go = new Go();
  WebAssembly.instantiateStreaming(fetch("./bin/main.wasm"), go.importObject)
    .then(result => go.run(result.instance));
}

window.onload = init;
