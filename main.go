package main

import (
	"fmt"
	"reddit-clone/internal/engine"
	"reddit-clone/cmd/simulator"
)

func main() {
	// Initialize the engine
	redditEngine := engine.NewEngine()

	// Start the simulator
	simulator.RunSimulation(redditEngine)

	fmt.Println("Simulation completed!")
}
