# WebAssembly in Go - Benchmark

Example code to benchmark WebAssembly code in Go and JavaScript


## Run

To run (you have to have Go installed on your system):

```bash
go run server/main.go
```

Open the browser with the address `localhost:8000`.

## Compile 


### Go

To compile, just run the command:

```
GOOS=js GOARCH=wasm go build -o assets/go.wasm wasm/go/main.go
```

A file named `go.wasm` will be created in `assets` folder.

### TinyGo

You need the TinyGo package installed on your system. 
You can see the procedures ["here"](https://tinygo.org/getting-started/).

After install, you can the command to compile the code:

```bash
tinygo build -o ./assets/tiny_go.wasm -target wasm ./wasm/tinygo/main.go
```

A file named `tiny_go.wasm` will be created in `assets` folder.

**Note**: I changed the `wasm_exec.js` from TinyGo project so that both Go and TinyGo will run alongside.

### C

You need the Emscripten package installed on your system to compile it.
You can follow the instructions ["here"](https://emscripten.org/docs/getting_started/downloads.html).

```bash
source ~/emsdk/emsdk_env.sh --build=Release
emcc wasm/c/main.c -o assets/wasm_exec_c.js -Os
```

Two files, named `wasm_exec.js` and `wasm_exec_c.wasm` will be create in `assets` folder.

## Article

This code was done for the article ["WebAssembly in Go vs JavaScript: A Benchmark"](https://medium.com/vacatronics/webassembly-in-go-vs-javascript-a-benchmark-6deb28f24e9d?sk=6a12153dcc84f42459e288135aea4afa)
