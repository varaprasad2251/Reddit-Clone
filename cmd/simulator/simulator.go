package simulator

import (
	"fmt"
	"reddit-clone/internal/engine"
	"sync"
	"time"
)

func simulateUserActions(engine *engine.Engine, userID int, subredditID int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 5; i++ {
		engine.CreatePost(subredditID, userID, fmt.Sprintf("Post %d by User %d", i, userID))
		time.Sleep(100 * time.Millisecond)
	}
}

func RunSimulation(engine *engine.Engine) {
	// Register users and create subreddits
	user1 := engine.RegisterUser("user1")
	user2 := engine.RegisterUser("user2")
	subreddit := engine.CreateSubreddit("golang")
	engine.JoinSubreddit(user1.ID, subreddit.ID)
	engine.JoinSubreddit(user2.ID, subreddit.ID)

	// Simulate user activity
	var wg sync.WaitGroup
	wg.Add(2)
	go simulateUserActions(engine, user1.ID, subreddit.ID, &wg)
	go simulateUserActions(engine, user2.ID, subreddit.ID, &wg)
	wg.Wait()
}
