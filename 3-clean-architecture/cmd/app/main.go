package main

import (
	"sync"

	"github.com/janainamai/go-expert-challenges/3-clean-architecture/cmd/configs"
	"github.com/janainamai/go-expert-challenges/3-clean-architecture/cmd/resources"
	grpcSetup "github.com/janainamai/go-expert-challenges/3-clean-architecture/internal/infra/grpc/setup"
	restSetup "github.com/janainamai/go-expert-challenges/3-clean-architecture/internal/infra/rest/setup"
)

func main() {
	cfg := configs.LoadConfigs()

	resources := resources.LoadResources(cfg)

	wg := sync.WaitGroup{}
	wg.Add(3)

	go restSetup.InitServer(cfg, resources)
	go grpcSetup.InitService(cfg, resources)

	wg.Wait()
}
