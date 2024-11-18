// package main

// import (
// 	"fmt"
// 	"reddit-clone/cmd/simulator"
// 	"reddit-clone/internal/engine"
// 	"sync"
// )

// func main() {
// 	redditEngine := engine.NewEngine()

// 	var wg sync.WaitGroup // Create a WaitGroup for synchronization
// 	simulator.RunSimulation(redditEngine, &wg)

// 	// Wait for all goroutines to finish
// 	wg.Wait()

// 	fmt.Println("Part I simulation completed.")
// }


package main

import (
	"fmt"
	"log"
	"reddit-clone/internal/engine"
	"sync"
)

func simulateUser(userID int, engine *engine.Engine, wg *sync.WaitGroup) {
	defer wg.Done()

	// Create posts in both subreddits
	for subredditID := 1; subredditID <= 2; subredditID++ {
		post, err := engine.CreatePost(subredditID, userID, fmt.Sprintf("Post by User %d in Subreddit %d", userID, subredditID))
		if err != nil {
			log.Printf("Error creating post for User %d in Subreddit %d: %v", userID, subredditID, err)
		} else {
			log.Printf("User %d created post in Subreddit %d: %s", userID, subredditID, post.Content)
		}
	}
}

func main() {
	engine := engine.NewEngine()

	// Create 2 subreddits
	for i := 1; i <= 2; i++ {
		subreddit, err := engine.CreateSubreddit(fmt.Sprintf("Subreddit%d", i))
		if err != nil {
			log.Fatalf("Failed to create Subreddit%d: %v", i, err)
		}
		log.Printf("Created Subreddit: ID=%d, Name=%s", subreddit.ID, subreddit.Name)
	}

	// Create 3 users
	for i := 1; i <= 3; i++ {
		user, err := engine.RegisterUser(fmt.Sprintf("user%d", i))
		if err != nil {
			log.Fatalf("Failed to create User%d: %v", i, err)
		}
		log.Printf("Created User: ID=%d, Username=%s", user.ID, user.Username)
	}

	var wg sync.WaitGroup

	// Simulate user actions
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go simulateUser(i, engine, &wg)
	}

	wg.Wait()
	log.Println("Simulation completed.")
}