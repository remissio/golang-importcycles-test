package providerB

import (
	"fmt"
	"importcycle/interfaces"
)

type ProviderB struct{}

func (p *ProviderB) DoStuff() {
	fmt.Println("ProviderB.DoStuff")
}

func (p *ProviderB) DoStuffWithA(a interfaces.A) {
	fmt.Println("ProviderB.DoStuffWithA")
	a.DoStuff()
}
