package Engine

import (
	"github.com/asynkron/protoactor-go/actor"
	"cop5615-project4/messages"
	"fmt"
	"sync"
	"sort"
    "time"
)

type Engine struct {
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
}


func NewEngine(system *actor.ActorSystem, wg *sync.WaitGroup) *Engine {
	fmt.Println("Initializing New Engine")

	return &Engine{
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


func (engine *Engine) Receive(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case *messages.RegisterUser:
		engine.Stats.TotalUsers++
        engine.Stats.UserStats[msg.UserName] = UserStat{}
		// fmt.Println("Engine: Processing user registration")
		engine.RegisterUser(msg.UserName)
		//		engine.Wg.Done()

	case *messages.UserJoinSubReddit:
		if _, exists := engine.subreddits[msg.SubRedditName]; !exists {
            engine.Stats.TotalSubreddits++
        }
		// fmt.Println("Engine: User Join SubReddit Operation")
		engine.SubredditSpecificOp("join", msg.UserName, msg.SubRedditName)
		//		engine.Wg.Done()

	case *messages.UserLeaveSubReddit:
		// fmt.Println("Engine: User Leave SubReddit Operation")
		engine.SubredditSpecificOp("leave", msg.UserName, msg.SubRedditName)
		//		engine.Wg.Done()

	case *messages.CreatePost:
		engine.Stats.TotalPosts++
        userStat := engine.Stats.UserStats[msg.UserName]
        userStat.PostCount++
        engine.Stats.UserStats[msg.UserName] = userStat
		// fmt.Println("Engine: Create Post Operation")
		engine.CreatePost(msg.UserName, msg.SubredditName, msg.Content)

		//		engine.Wg.Done()

	case *messages.ReplyToComment:
		// fmt.Println("Engine: Reply To Comment")
		engine.ReplyToComment(msg.UserName, msg.SubRedditName, msg.PostID, msg.CommentID, msg.ReplyContent)
		//		engine.Wg.Done()

	case *messages.SendDmToUser:
		engine.Stats.TotalMessages++
		// fmt.Println("Engine: Send DM")
		engine.SendDMtoUser(msg.UserName, msg.Content)
		//		engine.Wg.Done()
	case *messages.ReplyToDm:
		// fmt.Println("Engine: Reply To DM")
		engine.ReplyToAllDMs(msg.UserName, msg.Content)
		//		engine.Wg.Done()
	case *messages.UpVotePost:
		// fmt.Println("Engine: Processing Upvote")
		engine.UpvoteRandomPost(msg.UserName)
	
	case *messages.DownVotePost:
		// fmt.Println("Engine: Processing Downvote")
		engine.DownvoteRandomPost(msg.UserName)

	default:
		// fmt.Println("Engine: Unknown message type received")
	}
	// fmt.Println("Engine: Exiting Receive method")
}


func (engine *Engine) GetFeed(userName string, limit int) []messages.Post {
    var feed []messages.Post
    for _, subreddit := range engine.subreddits {
        feed = append(feed, subreddit.Posts...)
    }
    
    sort.Slice(feed, func(i, j int) bool {
        return feed[i].CreatedAt.After(feed[j].CreatedAt)
    })
    
    if len(feed) > limit {
        return feed[:limit]
    }
    return feed
}


func (engine *Engine) GetDirectMessages(userName string) []messages.DM {
    return engine.userData[userName].Dm
}

func (engine *Engine) ReplyToDirectMessage(userName string, messageID int, content string) {
    for i, dm := range engine.userData[userName].Dm {
        if dm.ID == messageID {
            replyDM := messages.DM{
                UserName: userName,
                Content:  content,
            }
            engine.userData[userName].Dm[i].Replies = append(dm.Replies, replyDM)
            break
        }
    }
}


func (engine *Engine) PrintStats() {
    fmt.Println("\nSimulation Statistics:")
    fmt.Printf("Total Users: %d\n", engine.Stats.TotalUsers)
    fmt.Printf("Total Posts: %d\n", engine.Stats.TotalPosts)
    fmt.Printf("Total Subreddits: %d\n", engine.Stats.TotalSubreddits)
    fmt.Printf("Total Messages: %d\n", engine.Stats.TotalMessages)
    fmt.Printf("Total Simulation Time: %v\n", engine.Stats.SimulationTime)
    
    fmt.Println("\nUser Statistics:")
    for user, stat := range engine.Stats.UserStats {
        fmt.Printf("User %s - Posts: %d\n", user, stat.PostCount)
    }
}
