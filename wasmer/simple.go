package main

import (
	"fmt"
	"github.com/wasmerio/wasmer-go/wasmer"
	"os"
)

func main() {
	// read wasm file as bytes[]
	wasmBytes, _ := os.ReadFile("simple.wasm")

	// engine for compiling the module
	engine := wasmer.NewEngine()
	store := wasmer.NewStore(engine)
	module, _ := wasmer.NewModule(store, wasmBytes)

	// instantiating module
	importObject := wasmer.NewImportObject()
	instance, _ := wasmer.NewInstance(module, importObject)

	// get exported function from WebAssembly instance
	sumFunc, _ := instance.Exports.GetFunction("sum")

	// calling the function
	result, _ := sumFunc(15, 25)
	fmt.Println(result)
}
