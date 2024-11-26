// user_operations.go
package redditEngine

import (
	"dosp-proj3/messages"
	"fmt"
)

// registerUser registers a user if they are not already in userData.
func (engine *RedditEngine) RegisterUser(userName string) {
	// Check if the user already exists in the map
	if _, exists := engine.userData[userName]; exists {
		fmt.Printf("User %s is already registered.\n", userName)
		return
	}

	// If user does not exist, create a new user entry and add it to the map
	engine.userData[userName] = messages.UserDataType{
		UserName:       userName,
		JointSubReddit: []string{},
	}
	fmt.Printf("User %s registered successfully.\n", userName)
}

// subredditSpecificOp performs subreddit join or leave operations for a user.
func (engine *RedditEngine) SubredditSpecificOp(actionToPerform string, userName string, subRedditName string) {
	// Check if the user already exists in the map
	userData, exists := engine.userData[userName]
	if !exists {
		fmt.Printf("User %s is not registered and cannot join or leave subreddits.\n", userName)
		return
	}
	// Ensure the subreddit exists in the map; if not, create it
	if _, subredditExists := engine.subRedditData[subRedditName]; !subredditExists {
		engine.subRedditData[subRedditName] = messages.SubReddit{ListOfPosts: []messages.Post{}}
		fmt.Printf("Subreddit %s did not exist and was created.\n", subRedditName)
	}
	switch actionToPerform {
	case "join":
		// Check if the user is already part of the subreddit
		for _, subreddit := range userData.JointSubReddit {
			if subreddit == subRedditName {
				fmt.Printf("User %s is already a member of %s subreddit.\n", userName, subRedditName)
				return
			}
		}
		// Add subreddit to the user's list if not already joined
		userData.JointSubReddit = append(userData.JointSubReddit, subRedditName)
		engine.userData[userName] = userData
		fmt.Printf("User %s joined %s subreddit.\n", userName, subRedditName)

	case "leave":
		// Find and remove the subreddit from the user's list if joined
		for i, subreddit := range userData.JointSubReddit {
			if subreddit == subRedditName {
				userData.JointSubReddit = append(userData.JointSubReddit[:i], userData.JointSubReddit[i+1:]...)
				engine.userData[userName] = userData
				fmt.Printf("User %s left %s subreddit.\n", userName, subRedditName)
				return
			}
		}
		fmt.Printf("User %s is not a member of %s subreddit.\n", userName, subRedditName)

	default:
		fmt.Printf("Unknown action: %s. Supported actions are 'join' and 'leave'.\n", actionToPerform)
	}
}

// CreatePost creates a new post in a specified subreddit if the user is registered.
func (engine *RedditEngine) CreatePost(userName string, subRedditName string, content messages.Post) {
	// Check if the user is registered
	user, userExists := engine.userData[userName]
	if !userExists {
		fmt.Printf("User %s is not registered and cannot create posts.\n", userName)
		return
	}

	// Check if the subreddit exists in the map
	subreddit, subredditExists := engine.subRedditData[subRedditName]
	if !subredditExists {
		fmt.Printf("Subreddit %s does not exist.\n", subRedditName)
		return
	}

	// Add the post to the subreddit's list of posts
	subreddit.ListOfPosts = append(subreddit.ListOfPosts, content)
	engine.subRedditData[subRedditName] = subreddit
	fmt.Printf("User %s created a post in subreddit %s: %s\n", userName, subRedditName, content.Content)

	//fmt.Printf("Debug: Subreddit 'golang' has %d posts.\n", len(engine.subRedditData["golang"].ListOfPosts))
	// Automatically add the user to the subreddit if not already joined
	if !isUserInSubreddit(user.JointSubReddit, subRedditName) {
		user.JointSubReddit = append(user.JointSubReddit, subRedditName)
		engine.userData[userName] = user
		fmt.Printf("User %s was automatically added to %s subreddit.\n", userName, subRedditName)
	}
}

// Helper function to check if a user is already part of a subreddit
func isUserInSubreddit(subreddits []string, subredditName string) bool {
	for _, sub := range subreddits {
		if sub == subredditName {
			return true
		}
	}
	return false
}

// ReplyToComment allows a user to reply to a specific comment in a post within a subreddit.
func (engine *RedditEngine) ReplyToComment(userName string, subRedditName string, postID int, commentID int, replyContent string) {
	// Step 1: Check if the user is registered
	_, userExists := engine.userData[userName]
	if !userExists {
		fmt.Printf("User %s is not registered and cannot reply to comments.\n", userName)
		return
	}

	// Step 2: Check if the subreddit exists in the map
	subreddit, subredditExists := engine.subRedditData[subRedditName]
	if !subredditExists {
		fmt.Printf("Subreddit %s does not exist.\n", subRedditName)
		return
	}
	// Debug: Print the number of posts in the subreddit

	// Step 3: Locate the post in the subreddit
	var post *messages.Post
	for i, p := range subreddit.ListOfPosts {
		//fmt.Printf("Debug: Inspecting Post[%d] - ID: %d, Content: %s, UserName: %s\n", i, p.ID, p.Content, p.UserName)
		if p.ID == postID {
			post = &subreddit.ListOfPosts[i]
			break
		}
	}

	if post == nil {
		fmt.Printf("Post with ID %d does not exist in subreddit %s.\n", postID, subRedditName)
		return
	}

	// Step 4: Locate the specific comment by ID and add the reply
	reply := messages.Comment{
		ID:        len(post.Comments) + 1, // Generate a new ID for the reply
		Content:   replyContent,
		UserName:  userName,
		Upvotes:   0,
		Downvotes: 0,
		Replies:   []messages.Comment{},
	}

	success := addReplyToComment(post.Comments, commentID, reply)
	if success {
		fmt.Printf("User %s replied to comment %d in post %d in subreddit %s.\n", userName, commentID, postID, subRedditName)
	} else {
		fmt.Printf("Comment with ID %d not found in post %d in subreddit %s.\n", commentID, postID, subRedditName)
	}
}

// Helper function to locate and add a reply to a specific comment recursively
func addReplyToComment(comments []messages.Comment, targetID int, reply messages.Comment) bool {
	for i := range comments {
		if comments[i].ID == targetID {
			// Found the target comment, add the reply
			comments[i].Replies = append(comments[i].Replies, reply)
			return true
		}
		// Recursively search for the target comment in nested replies
		if addReplyToComment(comments[i].Replies, targetID, reply) {
			return true
		}
	}
	return false
}
