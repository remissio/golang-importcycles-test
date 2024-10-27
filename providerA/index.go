package providerA

import (
	"fmt"
	"importcycle/interfaces"
)

type ProviderA struct{}

func (p *ProviderA) DoStuff() {
	fmt.Println("ProviderA.DoStuff")
}

func (p *ProviderA) DoStuffWithB(b interfaces.B) {
	fmt.Println("ProviderA.DoStuffWithA")
	b.DoStuff()
}
