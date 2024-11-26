// reddit_engine.go
package redditEngine

import "time"

import (
	"cop5615-project4/messages"
	"fmt"
	"sync"
	"github.com/asynkron/protoactor-go/actor"
)

type RedditEngine struct {
	userData      map[string]messages.UserDataType
	subRedditData map[string]messages.SubReddit
	Wg            *sync.WaitGroup
	system        *actor.ActorSystem
	Stats    Stats
	subreddits    map[string]Subreddit
}

type Stats struct {
    TotalUsers      int
    TotalPosts      int
    TotalSubreddits int
    TotalMessages   int
    UserStats       map[string]UserStat
    SimulationTime  time.Duration
}

type UserStat struct {
    Karma     int
    PostCount int
}

type Subreddit struct {
    Name  string
    Posts []messages.Post
    // Add other relevant fields
}

// NewRedditEngine initializes the RedditEngine with a WaitGroup.
func NewRedditEngine(system *actor.ActorSystem, wg *sync.WaitGroup) *RedditEngine {
	fmt.Println("NewRedditEngine: Initializing RedditEngine")

	return &RedditEngine{
		userData:      make(map[string]messages.UserDataType),
		subRedditData: make(map[string]messages.SubReddit),
		Wg:            wg,
		system:        system,
		Stats: Stats{
            UserStats: make(map[string]UserStat),
        },
		subreddits: make(map[string]Subreddit),
	}
}

// Receive method to process messages sent to RedditEngine.
func (engine *RedditEngine) Receive(ctx actor.Context) {
	fmt.Printf("RedditEngine: Entered Receive method ; Received Action-> %v\n", ctx.Message())
	switch msg := ctx.Message().(type) {
	case *messages.RegisterUser:
		engine.Stats.TotalUsers++
        engine.Stats.UserStats[msg.UserName] = UserStat{}
		fmt.Println("RedditEngine: Processing user registration")
		engine.RegisterUser(msg.UserName)
		//		engine.Wg.Done()

	case *messages.UserJoinSubReddit:
		if _, exists := engine.subreddits[msg.SubRedditName]; !exists {
            engine.Stats.TotalSubreddits++
        }
		fmt.Println("RedditEngine: User Join SubReddit Operation")
		engine.SubredditSpecificOp("join", msg.UserName, msg.SubRedditName)
		//		engine.Wg.Done()

	case *messages.UserLeaveSubReddit:
		fmt.Println("RedditEngine: User Leave SubReddit Operation")
		engine.SubredditSpecificOp("leave", msg.UserName, msg.SubRedditName)
		//		engine.Wg.Done()

	case *messages.CreatePost:
		engine.Stats.TotalPosts++
        userStat := engine.Stats.UserStats[msg.UserName]
        userStat.PostCount++
        engine.Stats.UserStats[msg.UserName] = userStat
		fmt.Println("RedditEngine: Create Post Operation")
		engine.CreatePost(msg.UserName, msg.SubredditName, msg.Content)

		//		engine.Wg.Done()

	case *messages.ReplyToComment:
		fmt.Println("RedditEngine: Reply To Comment")
		engine.ReplyToComment(msg.UserName, msg.SubRedditName, msg.PostID, msg.CommentID, msg.ReplyContent)
		//		engine.Wg.Done()

	case *messages.SendDmToUser:
		engine.Stats.TotalMessages++
		fmt.Println("RedditEngine: Send DM")
		engine.SendDMtoUser(msg.UserName, msg.Content)
		//		engine.Wg.Done()
	case *messages.ReplyToDm:
		fmt.Println("RedditEngine: Reply To DM")
		engine.ReplyToAllDMs(msg.UserName, msg.Content)
		//		engine.Wg.Done()
	case *messages.UpVotePost:
		fmt.Println("RedditEngine: Reply To DM")
		engine.UpvoteRandomPost(msg.UserName)
		//		engine.Wg.Done()
	case *messages.DownVotePost:
		fmt.Println("RedditEngine: Reply To DM")
		engine.DownvoteRandomPost(msg.UserName)
		//		engine.Wg.Done()

	default:
		fmt.Println("RedditEngine: Unknown message type received")
	}
	fmt.Println("RedditEngine: Exiting Receive method")
}


func (engine *RedditEngine) PrintStats() {
    fmt.Println("\nSimulation Statistics:\n")
    fmt.Printf("Total Users: %d\n", engine.Stats.TotalUsers)
    fmt.Printf("Total Posts: %d\n", engine.Stats.TotalPosts)
    fmt.Printf("Total Subreddits: %d\n", engine.Stats.TotalSubreddits)
    fmt.Printf("Total Messages: %d\n", engine.Stats.TotalMessages)
    fmt.Printf("Simulation Time: %v\n", engine.Stats.SimulationTime)
    
    fmt.Println("\nUser Statistics:")
    for user, stat := range engine.Stats.UserStats {
        fmt.Printf("%s - Karma: %d, Posts: %d\n", user, stat.Karma, stat.PostCount)
    }
}
