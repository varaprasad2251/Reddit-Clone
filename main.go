package main

import (
	"github.com/asynkron/protoactor-go/actor"
	"github.com/sirupsen/logrus"
	"cop5615-project4/Engine"
	"cop5615-project4/simulation"
	"cop5615-project4/api"
	"flag"
	"fmt"
	"sync"
	"time"
	"os"
)

func main() {
	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
	log.SetOutput(os.Stderr)

	startTime := time.Now()
	
	actorSystem := actor.NewActorSystem()
	log.Info("Actor system created")
	var wg sync.WaitGroup

	engine := Engine.NewEngine(actorSystem, &wg)
	log.Info("Reddit engine created")
	props := actor.PropsFromProducer(func() actor.Actor {
		return engine
	})
	enginePID := actorSystem.Root.Spawn(props)

	apiHandler := api.NewAPI(engine)
    go func() {
        if err := apiHandler.Run(":8080"); err != nil {
            log.Fatal("Failed to start API server: ", err)
        }
    }()
	
	runSimulation := flag.Bool("simulate", false, "Run the simulation")
	numUsers := flag.Int("users", 10, "Number of users to simulate")
	flag.Parse()
	
	if *runSimulation {
        log.Info(fmt.Sprintf("Starting Simulation with %d users", *numUsers))

        for i := 0; i < *numUsers; i++ {
            wg.Add(1)
            userName := fmt.Sprintf("%d", i+1)
            go simulation.SimulateUser(enginePID, actorSystem, userName, &wg)
        }

        wg.Wait()
        simulationTime := time.Since(startTime)
        engine.Stats.SimulationTime = simulationTime
        log.Info(fmt.Sprintf("Simulation Completed in %v", simulationTime))

        engine.PrintStats()
    } else {
        log.Info("API server is running. Press Ctrl+C to stop.")
        // Keep the main goroutine alive
        select {}
    }
}
