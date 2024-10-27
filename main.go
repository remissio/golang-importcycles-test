package main

import (
	"importcycle/providerA"
	"importcycle/providerB"
)

func main() {
	a := providerA.ProviderA{}
	b := providerB.ProviderB{}
	a.DoStuffWithB(&b)
	b.DoStuffWithA(&a)
}
