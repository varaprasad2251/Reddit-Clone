package redditEngine

import (
	"fmt"
	"math/rand"
	"time"
)

// UpvoteRandomPost allows a user to randomly upvote a post in a subscribed subreddit.
func (engine *RedditEngine) UpvoteRandomPost(userName string) {
	// Step 1: Check if the user exists
	user, userExists := engine.userData[userName]
	if !userExists {
		fmt.Printf("User %s is not registered and cannot upvote posts.\n", userName)
		return
	}

	// Step 2: Check if the user is subscribed to any subreddits
	if len(user.JointSubReddit) == 0 {
		fmt.Printf("User %s is not subscribed to any subreddits and cannot upvote posts.\n", userName)
		return
	}

	// Step 3: Randomly select a subreddit
	rand.Seed(time.Now().UnixNano())
	selectedSubReddit := user.JointSubReddit[rand.Intn(len(user.JointSubReddit))]

	// Step 4: Check if the subreddit exists and has posts
	subreddit, subredditExists := engine.subRedditData[selectedSubReddit]
	if !subredditExists || len(subreddit.ListOfPosts) == 0 {
		fmt.Printf("Subreddit %s does not exist or has no posts.\n", selectedSubReddit)
		return
	}

	// Step 5: Randomly select a post
	selectedPost := &subreddit.ListOfPosts[rand.Intn(len(subreddit.ListOfPosts))]

	// Step 6: Upvote the selected post
	selectedPost.Upvotes++
	fmt.Printf("User %s upvoted post ID %d in subreddit %s. Total upvotes: %d\n",
		userName, selectedPost.ID, selectedSubReddit, selectedPost.Upvotes)

	// Step 7: Optionally update the karma of the post's creator
	if postCreator, creatorExists := engine.userData[selectedPost.UserName]; creatorExists {
		postCreator.KarmaPoints++
		engine.userData[selectedPost.UserName] = postCreator
		fmt.Printf("Updated karma for user %s. New Post Karma: %d\n", selectedPost.UserName, postCreator.KarmaPoints)
	}
}

// UpvoteRandomPost allows a user to randomly upvote a post in a subscribed subreddit.
func (engine *RedditEngine) DownvoteRandomPost(userName string) {
	// Step 1: Check if the user exists
	user, userExists := engine.userData[userName]
	if !userExists {
		fmt.Printf("User %s is not registered and cannot upvote posts.\n", userName)
		return
	}

	// Step 2: Check if the user is subscribed to any subreddits
	if len(user.JointSubReddit) == 0 {
		fmt.Printf("User %s is not subscribed to any subreddits and cannot upvote posts.\n", userName)
		return
	}

	// Step 3: Randomly select a subreddit
	rand.Seed(time.Now().UnixNano())
	selectedSubReddit := user.JointSubReddit[rand.Intn(len(user.JointSubReddit))]

	// Step 4: Check if the subreddit exists and has posts
	subreddit, subredditExists := engine.subRedditData[selectedSubReddit]
	if !subredditExists || len(subreddit.ListOfPosts) == 0 {
		fmt.Printf("Subreddit %s does not exist or has no posts.\n", selectedSubReddit)
		return
	}

	// Step 5: Randomly select a post
	selectedPost := &subreddit.ListOfPosts[rand.Intn(len(subreddit.ListOfPosts))]

	// Step 6: Upvote the selected post
	if selectedPost.Upvotes-1 >= 0 {
		selectedPost.Upvotes--
	}
	fmt.Printf("User %s upvoted post ID %d in subreddit %s. Total upvotes: %d\n",
		userName, selectedPost.ID, selectedSubReddit, selectedPost.Upvotes)

	// Step 7: Optionally update the karma of the post's creator
	if postCreator, creatorExists := engine.userData[selectedPost.UserName]; creatorExists {
		if postCreator.KarmaPoints-1 >= 0 {
			postCreator.KarmaPoints--
		}

		engine.userData[selectedPost.UserName] = postCreator
		fmt.Printf("Updated karma for user %s. New Post Karma: %d\n", selectedPost.UserName, postCreator.KarmaPoints)
	}
}
