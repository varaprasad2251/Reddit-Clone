package simulation

import (
	"dosp-proj3/messages"
	"math/rand"
	"sync"
	"time"
	"github.com/asynkron/protoactor-go/actor"
)

// SimulateUser performs a series of actions as a simulated user.
func SimulateUser(enginePID *actor.PID, system *actor.ActorSystem, userName string, wg *sync.WaitGroup) {
	defer wg.Done()

	// Random seed for generating actions
	rand.Seed(time.Now().UnixNano())

	// Register user
	system.Root.Send(enginePID, &messages.RegisterUser{UserName: userName})

	// Randomly perform actions like joining a subreddit, creating posts, commenting, etc.
	actions := []func(){
		func() {
			// Join a subreddit
			subRedditName := "golang"
			system.Root.Send(enginePID, &messages.UserJoinSubReddit{UserName: userName, SubRedditName: subRedditName})
		},
		func() {
			// Create a post
			subRedditName := "golang"
			postContent := messages.Post{
				ID:        rand.Intn(100),
				Content:   "This is a simulated post about Go!",
				UserName:  userName,
				Upvotes:   0,
				Downvotes: 0,
				Comments:  []messages.Comment{},
			}
			system.Root.Send(enginePID, &messages.CreatePost{UserName: userName, SubredditName: subRedditName, Content: postContent})
		},
		func() {
			// Reply to a comment
			system.Root.Send(enginePID, &messages.ReplyToComment{
				UserName:      userName,
				SubRedditName: "golang",
				PostID:        1, // Example post ID
				CommentID:     1, // Example comment ID
				ReplyContent:  "This is a reply to a comment",
			})
		},
		func() {
			// send DM
			system.Root.Send(enginePID, &messages.SendDmToUser{
				UserName: "user1",
				Content:  "Hi Bro",
			})
		},
		func() {
			// reply to DM
			system.Root.Send(enginePID, &messages.ReplyToDm{
				UserName: "User1",
				Content:  "Hi Bro",
			})
		},
		func() {
			// reply to DM
			system.Root.Send(enginePID, &messages.UpVotePost{
				UserName: "User1",
			})
		},
		func() {
			// reply to DM
			system.Root.Send(enginePID, &messages.DownVotePost{
				UserName: "User1",
			})
		},
	}

	// Execute actions randomly with some delay
	for i := 0; i < 20; i++ { // Each user performs 5 random actions
		action := actions[rand.Intn(len(actions))]
		action()
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond) // Random delay between actions
	}
}
