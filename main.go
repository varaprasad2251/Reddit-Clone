package main

import (
	"dosp-proj3/redditEngine"
	"dosp-proj3/simulation"
	"fmt"
	"sync"
	"time"
	"github.com/asynkron/protoactor-go/actor"
)

func main() {
	startTime := time.Now()
	// Create a new ActorSystem
	actorSystem := actor.NewActorSystem()
	fmt.Println("ActorSystem created")

	// Initialize WaitGroup
	var wg sync.WaitGroup

	// Initialize RedditEngine with the ActorSystem and WaitGroup
	engine := redditEngine.NewRedditEngine(actorSystem, &wg)
	fmt.Println("RedditEngine initialized")

	// Create props for RedditEngine actor
	props := actor.PropsFromProducer(func() actor.Actor {
		return engine
	})

	// Spawn the RedditEngine actor using the ActorSystem
	enginePID := actorSystem.Root.Spawn(props)
	fmt.Printf("RedditEngine actor spawned with PID: %v\n", enginePID)

	// Simulate multiple users concurrently
	numUsers := 10 // Define how many users to simulate
	for i := 0; i < numUsers; i++ {
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
