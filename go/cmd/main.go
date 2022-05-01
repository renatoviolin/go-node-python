package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/renatoviolin/go-node/adapters/chi"
	"github.com/renatoviolin/go-node/adapters/fiber"
	"github.com/renatoviolin/go-node/adapters/gin"
	"github.com/renatoviolin/go-node/ports"
	repository "github.com/renatoviolin/go-node/repository/mongodb"
	"github.com/renatoviolin/go-node/usecase"
)

func main() {
	var servers []ports.PortHttp

	repository, err := repository.NewMongoRepository("mongodb://127.0.0.1:27017/", "benchmark", 60)
	if err != nil {
		log.Fatal(err)
	}

	useCase := usecase.NewFindAllUseCase(repository)

	servers = append(servers, chi.NewHandler(useCase))
	servers = append(servers, gin.NewHandler(useCase))
	servers = append(servers, fiber.NewHandler(useCase))

	var wg sync.WaitGroup

	initial_port := 7000
	for i, s := range servers {
		wg.Add(i + 1)
		go func(server ports.PortHttp, port int) {
			server.SetupAndRun(fmt.Sprintf(":%d", port))
			wg.Done()
		}(s, initial_port)
	}
	initial_port += 1
	wg.Wait()
}
