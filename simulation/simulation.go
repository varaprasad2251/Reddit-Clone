package simulation

import (
    "cop5615-project4/messages"
    "fmt"
    "math"
    "math/rand"
    "sync"
    "time"
    "github.com/asynkron/protoactor-go/actor"
)

func zipf(s float64, v float64, imax uint64) uint64 {
    x := rand.Float64()
    return uint64((float64(imax) * math.Pow(x, -s)) / v)
}

func SimulateUser(enginePID *actor.PID, system *actor.ActorSystem, userName string, wg *sync.WaitGroup) {
    defer wg.Done()

    rand.Seed(time.Now().UnixNano())

    system.Root.Send(enginePID, &messages.RegisterUser{UserName: userName})

    go simulateConnectionStatus(enginePID, system, userName)

    actions := []func(){
        func() { joinRandomSubreddit(enginePID, system, userName) },
        func() { createPost(enginePID, system, userName) },
        func() { replyToRandomComment(enginePID, system, userName) },
        func() { sendDirectMessage(enginePID, system, userName) },
        func() { upvoteRandomPost(enginePID, system, userName) },
        func() { downvoteRandomPost(enginePID, system, userName) },
        func() { getFeed(enginePID, system, userName) },
        func() { getDirectMessages(enginePID, system, userName) },
		func() { randomVote(enginePID, system, userName) },
    }

	var numActions int = 150

	for i := 0; i < numActions; i++ { 
		if rand.Float32() < 0.3 {
			randomVote(enginePID, system, userName)
		} else {
			action := actions[rand.Intn(len(actions))]
			action()
		}
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	}
}

func simulateConnectionStatus(enginePID *actor.PID, system *actor.ActorSystem, userName string) {
    for {
        time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
        isConnected := rand.Float32() < 0.8
        system.Root.Send(enginePID, &messages.UpdateConnectionStatus{UserName: userName, IsConnected: isConnected})
    }
}

func joinRandomSubreddit(enginePID *actor.PID, system *actor.ActorSystem, userName string) {
    subredditName := fmt.Sprintf("subreddit_%d", rand.Intn(100))
    system.Root.Send(enginePID, &messages.UserJoinSubReddit{UserName: userName, SubRedditName: subredditName})
}

func createPost(enginePID *actor.PID, system *actor.ActorSystem, userName string) {
    subredditName := fmt.Sprintf("subreddit_%d", rand.Intn(100))
    postContent := messages.Post{
        ID:        rand.Intn(1000000),
        Content:   fmt.Sprintf("This is post #%d by %s", rand.Intn(1000), userName),
        UserName:  userName,
        Upvotes:   0,
        Downvotes: 0,
        Comments:  []messages.Comment{},
        CreatedAt: time.Now(),
    }
    system.Root.Send(enginePID, &messages.CreatePost{UserName: userName, SubredditName: subredditName, Content: postContent})
}

func replyToRandomComment(enginePID *actor.PID, system *actor.ActorSystem, userName string) {
    system.Root.Send(enginePID, &messages.ReplyToComment{
        UserName:      userName,
        SubRedditName: fmt.Sprintf("subreddit_%d", rand.Intn(100)),
        PostID:        rand.Intn(1000),
        CommentID:     rand.Intn(100),
        ReplyContent:  fmt.Sprintf("Reply from %s: %s", userName, generateRandomContent()),
    })
}

func sendDirectMessage(enginePID *actor.PID, system *actor.ActorSystem, userName string) {
    recipientName := fmt.Sprintf("User%d", rand.Intn(1000))
    system.Root.Send(enginePID, &messages.SendDmToUser{
        UserName: recipientName,
        Content:  fmt.Sprintf("DM from %s: %s", userName, generateRandomContent()),
    })
}

func upvoteRandomPost(enginePID *actor.PID, system *actor.ActorSystem, userName string) {
    system.Root.Send(enginePID, &messages.UpVotePost{
        UserName: userName,
        PostID:   rand.Intn(1000),
    })
}

func downvoteRandomPost(enginePID *actor.PID, system *actor.ActorSystem, userName string) {
    system.Root.Send(enginePID, &messages.DownVotePost{
        UserName: userName,
        PostID:   rand.Intn(1000),
    })
}

func getFeed(enginePID *actor.PID, system *actor.ActorSystem, userName string) {
    system.Root.Send(enginePID, &messages.GetFeed{
        UserName: userName,
        Limit:    20,
    })
}

func getDirectMessages(enginePID *actor.PID, system *actor.ActorSystem, userName string) {
    system.Root.Send(enginePID, &messages.GetDirectMessages{
        UserName: userName,
    })
}

func generateRandomContent() string {
    contents := []string{
        "The moon is just a hologram!",
        "Does anyone else hear that faint buzzing?",
        "I can't believe it's not butter!",
        "Why are ducks so underrated?",
        "Time travel is overrated.",
        "Bananas are a government conspiracy!",
        "This reminds me of my pet hamster, Gerald.",
        "Is anyone else craving tacos right now?",
        "Pineapples on pizza? Let's discuss.",
        "I'm 90% sure this is a simulation.",
        "The cake is a lie!",
        "Who let the dogs out?",
        "This post smells like teen spirit.",
        "One does not simply ignore this post.",
        "Aliens are among us. Trust me.",
        "This made my goldfish do a backflip.",
        "I'm typing this with my toes.",
        "The sky just winked at me. Weird.",
        "I'm not crying, you're crying.",
        "Did you know otters hold hands?",
    }
    return contents[rand.Intn(len(contents))]
}

func randomVote(enginePID *actor.PID, system *actor.ActorSystem, userName string) {
    targetUser := fmt.Sprintf("%d", rand.Intn(5)+1)
    postID := rand.Intn(100)
    if rand.Float32() < 0.7 { // 70% chance of upvote
        system.Root.Send(enginePID, &messages.UpVotePost{
            UserName:   userName,
            TargetUser: targetUser,
            PostID:     postID,
        })
    } else {
        system.Root.Send(enginePID, &messages.DownVotePost{
            UserName:   userName,
            TargetUser: targetUser,
            PostID:     postID,
        })
    }
}