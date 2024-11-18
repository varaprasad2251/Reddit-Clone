// package simulator

// import (
// 	"fmt"
// 	"math/rand"
// 	"reddit-clone/internal/engine"
// 	"reddit-clone/pkg/utils"
// 	"sync"
// 	"time"
// )

// func simulateUserActions(userID int, engine *engine.Engine, wg *sync.WaitGroup) {
// 	defer wg.Done()

// 	// Simulate connection/disconnection
// 	for i := 0; i < 5; i++ {
// 		engine.CreatePost(1, userID, fmt.Sprintf("Post %d by User %d", i, userID))
// 		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond) // Simulate disconnection
// 	}
// }

// func simulateZipfSubreddits(engine *engine.Engine, userID int, numSubreddits int) {
// 	zipf := utils.GenerateZipfDistribution(numSubreddits, 1.07)
// 	for i := 0; i < numSubreddits; i++ {
// 		if zipf[i] > 0 {
// 			engine.JoinSubreddit(userID, i+1)
// 		}
// 	}
// }

// func measurePerformance(engine *engine.Engine, numUsers int, wg *sync.WaitGroup) {
// 	start := time.Now()

// 	for i := 1; i <= numUsers; i++ {
// 		wg.Add(1)
// 		go simulateUserActions(i, engine, wg)
// 	}

// 	wg.Wait()
// 	fmt.Printf("Simulation completed in %v\n", time.Since(start))
// }

// func RunSimulation(engine *engine.Engine, wg *sync.WaitGroup) {
// 	numUsers := 100
// 	for i := 1; i <= numUsers; i++ {
// 		engine.RegisterUser(fmt.Sprintf("user%d", i))
// 	}

// 	// Simulate Zipf-based subreddit activity
// 	for i := 1; i <= numUsers; i++ {
// 		simulateZipfSubreddits(engine, i, 10)
// 	}

// 	// Measure performance
// 	measurePerformance(engine, numUsers, wg)
// }


package simulator

import (
	"fmt"
	"math/rand"
	"reddit-clone/internal/engine"
	"reddit-clone/pkg/utils"
	"sync"
	"time"
)

func simulateUserActions(userID int, engine *engine.Engine, wg *sync.WaitGroup) {
	defer wg.Done()
	if engine == nil {
        fmt.Println("Error: Engine is nil")
        return
    }

	// Simulate connection/disconnection
	for i := 0; i < 5; i++ {
		post, err := engine.CreatePost(1, userID, fmt.Sprintf("Post %d by User %d", i, userID))
        // if err != nil {
        //     log.Printf("Error creating post: %v", err)
        //     continue
        // }
		if err != nil {
			fmt.Printf("Error creating post for User %d: %v\n", userID, err)
			continue
		}
		fmt.Printf("User %d created post: %s\n", userID, post.Content)

		// comment := &engine.Comment{
		// 	ID:       rand.Intn(1000), // Generate a random ID
		// 	Content:  fmt.Sprintf("Comment %d by User %d", i, userID),
		// 	AuthorID: userID,
		// }
		comment := engine.CreateComment(post.ID, userID, fmt.Sprintf("Comment %d by User %d", i, userID))
		if comment == nil {
            fmt.Println("Error: Failed to create comment")
            continue
        }
		post.AddComment(comment, 0)
		fmt.Printf("User %d commented: %s\n", userID, comment.Content)

		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond) // Simulate disconnection
	}
}

func simulateZipfSubreddits(engine *engine.Engine, userID int, numSubreddits int) {
	zipf := utils.GenerateZipfDistribution(numSubreddits, 1.07)
	for i := 0; i < numSubreddits; i++ {
		if zipf[i] > 0 {
			engine.JoinSubreddit(userID, i+1)
			fmt.Printf("User %d joined subreddit: %d\n", userID, i+1)

			// Simulate leaving a subreddit
			if rand.Float64() > 0.8 { // 20% chance of leaving
				engine.LeaveSubreddit(userID, i+1)
				fmt.Printf("User %d left subreddit: %d\n", userID, i+1)
			}
		}
	}
}

func measurePerformance(engine *engine.Engine, numUsers int, wg *sync.WaitGroup) {
	start := time.Now()

	for i := 1; i <= numUsers; i++ {
		wg.Add(1)
		go simulateUserActions(i, engine, wg)
	}

	wg.Wait()
	fmt.Printf("Simulation completed in %v\n", time.Since(start))
}

func RunSimulation(engine *engine.Engine, wg *sync.WaitGroup) {
	numUsers := 100
	for i := 1; i <= numUsers; i++ {
		user, err := engine.RegisterUser(fmt.Sprintf("user%d", i))
        if err != nil {
            fmt.Printf("Error creating user%d: %v\n", i, err)
            continue
        }
		fmt.Printf("User created: ID=%d, Username=%s\n", user.ID, user.Username)
	}

	// Simulate Zipf-based subreddit activity
	for i := 1; i <= numUsers; i++ {
		simulateZipfSubreddits(engine, i, 10)
	}

	// Measure performance
	measurePerformance(engine, numUsers, wg)
}
