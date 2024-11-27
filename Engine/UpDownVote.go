package Engine

import (
	"fmt"
	"math/rand"
	"time"
)

func (engine *Engine) UpvoteRandomPost(userName string) {
	user, userExists := engine.userData[userName]
	if !userExists {
		return
	}

	if len(user.JointSubReddit) == 0 {
		return
	}

	rand.Seed(time.Now().UnixNano())
	selectedSubReddit := user.JointSubReddit[rand.Intn(len(user.JointSubReddit))]

	subreddit, subredditExists := engine.subRedditData[selectedSubReddit]
	if !subredditExists || len(subreddit.ListOfPosts) == 0 {
		return
	}

	selectedPost := &subreddit.ListOfPosts[rand.Intn(len(subreddit.ListOfPosts))]

	selectedPost.Upvotes++
	fmt.Printf("User %s upvoted post %d in subreddit %s. Total upvotes: %d\n",
		userName, selectedPost.ID, selectedSubReddit, selectedPost.Upvotes)

	if postCreator, creatorExists := engine.userData[selectedPost.UserName]; creatorExists {
		postCreator.KarmaPoints++
		engine.userData[selectedPost.UserName] = postCreator
	}
}

func (engine *Engine) DownvoteRandomPost(userName string) {
	user, userExists := engine.userData[userName]
	if !userExists {
		return
	}

	if len(user.JointSubReddit) == 0 {
		return
	}

	rand.Seed(time.Now().UnixNano())
	selectedSubReddit := user.JointSubReddit[rand.Intn(len(user.JointSubReddit))]

	subreddit, subredditExists := engine.subRedditData[selectedSubReddit]
	if !subredditExists || len(subreddit.ListOfPosts) == 0 {
		return
	}

	selectedPost := &subreddit.ListOfPosts[rand.Intn(len(subreddit.ListOfPosts))]

	if selectedPost.Upvotes-1 >= 0 {
		selectedPost.Upvotes--
	}
	fmt.Printf("User %s downvoted post ID %d in subreddit %s. Total upvotes: %d\n",
		userName, selectedPost.ID, selectedSubReddit, selectedPost.Upvotes)

	if postCreator, creatorExists := engine.userData[selectedPost.UserName]; creatorExists {
		if postCreator.KarmaPoints-1 >= 0 {
			postCreator.KarmaPoints--
		}
		engine.userData[selectedPost.UserName] = postCreator
	}
}
