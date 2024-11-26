package main

import (
	"github.com/asynkron/protoactor-go/actor"
	"github.com/sirupsen/logrus"
	"cop5615-project4/Engine"
	"cop5615-project4/simulation"
	"flag"
	"fmt"
	"sync"
	"time"
	"os"
)

func main() {
	log := logrus.New()
	// log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true, // Enable full timestamps
		TimestampFormat: "2006-01-02 15:04:05", // Set custom timestamp format
	})
	// Optional: Log to a file
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

	numUsers := flag.Int("users", 10, "Number of users to simulate")
	flag.Parse()
	
	log.Info(fmt.Sprintf("Starting Simulation with %d users", *numUsers))

	for i := 0; i < *numUsers; i++ {
		wg.Add(1)
		userName := fmt.Sprintf("%d", i+1)
		go simulation.SimulateUser(enginePID, actorSystem, userName, &wg)
	}

	wg.Wait()
	simulationTime := time.Since(startTime)
    engine.Stats.SimulationTime = simulationTime
    log.Info("Simulation Completed")

    engine.PrintStats()
}
