package main

import (
	"cop5615-project4/redditEngine"
	"cop5615-project4/simulation"
	"flag"
	"fmt"
	"sync"
	"time"
	"github.com/asynkron/protoactor-go/actor"
)

func main() {
	startTime := time.Now()
	actorSystem := actor.NewActorSystem()
	fmt.Println("ActorSystem created")
	var wg sync.WaitGroup


	engine := redditEngine.NewRedditEngine(actorSystem, &wg)
	fmt.Println("Reddit Engine created")
	props := actor.PropsFromProducer(func() actor.Actor {
		return engine
	})
	enginePID := actorSystem.Root.Spawn(props)

	// numUsers := 10
	numUsers := flag.Int("users", 10, "number of users to simulate")
	flag.Parse()
	for i := 0; i < *numUsers; i++ {
		wg.Add(1)
		userName := fmt.Sprintf("User%d", i+1)
		go simulation.SimulateUser(enginePID, actorSystem, userName, &wg)
	}

	// Wait until all users have finished their actions
	wg.Wait()
	fmt.Println("Finished simulating all users.")
	simulationTime := time.Since(startTime)
    engine.Stats.SimulationTime = simulationTime
    
    engine.PrintStats()
}
