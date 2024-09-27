package main

import (
	"sync"

	"github.com/janainamai/go-expert-challenges/cmd/configs"
	"github.com/janainamai/go-expert-challenges/cmd/resources"
	restSetup "github.com/janainamai/go-expert-challenges/internal/infra/rest/setup"
)

func main() {
	cfg := configs.LoadConfigs()

	resources := resources.LoadResources(cfg)

	wg := sync.WaitGroup{}
	wg.Add(2)
	go restSetup.InitServer(resources)

	wg.Wait()
}
