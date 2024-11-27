// user_operations.go
package Engine

import (
	"cop5615-project4/messages"
	"fmt"
)

func (engine *Engine) RegisterUser(userName string) {
	if _, exists := engine.userData[userName]; exists {
		return
	}

	engine.userData[userName] = messages.UserDataType{
		UserName:       userName,
		JointSubReddit: []string{},
	}
	fmt.Printf("User %s registered successfully.\n", userName)
}

func (engine *Engine) SubredditSpecificOp(actionToPerform string, userName string, subRedditName string) {
	userData, exists := engine.userData[userName]
	if !exists {
		return
	}
	if _, subredditExists := engine.subRedditData[subRedditName]; !subredditExists {
		engine.subRedditData[subRedditName] = messages.SubReddit{ListOfPosts: []messages.Post{}}
	}
	switch actionToPerform {
	case "join":
		for _, subreddit := range userData.JointSubReddit {
			if subreddit == subRedditName {
				return
			}
		}
		userData.JointSubReddit = append(userData.JointSubReddit, subRedditName)
		engine.userData[userName] = userData
		fmt.Printf("User %s joined %s subreddit.\n", userName, subRedditName)

	case "leave":
		for i, subreddit := range userData.JointSubReddit {
			if subreddit == subRedditName {
				userData.JointSubReddit = append(userData.JointSubReddit[:i], userData.JointSubReddit[i+1:]...)
				engine.userData[userName] = userData
				fmt.Printf("User %s left %s subreddit.\n", userName, subRedditName)
				return
			}
		}
	default:
		fmt.Printf("Unknown action: %s. Supported actions are 'join' and 'leave'.\n", actionToPerform)
	}
}

func (engine *Engine) CreatePost(userName string, subRedditName string, content messages.Post) {
	user, userExists := engine.userData[userName]
	if !userExists {
		return
	}

	subreddit, subredditExists := engine.subRedditData[subRedditName]
	if !subredditExists {
		return
	}

	subreddit.ListOfPosts = append(subreddit.ListOfPosts, content)
	engine.subRedditData[subRedditName] = subreddit
	fmt.Printf("User %s created a post in subreddit %s: %s\n", userName, subRedditName, content.Content)

	if !isUserInSubreddit(user.JointSubReddit, subRedditName) {
		user.JointSubReddit = append(user.JointSubReddit, subRedditName)
		engine.userData[userName] = user
		fmt.Printf("User %s was added to %s subreddit.\n", userName, subRedditName)
	}
}

func isUserInSubreddit(subreddits []string, subredditName string) bool {
	for _, sub := range subreddits {
		if sub == subredditName {
			return true
		}
	}
	return false
}

func (engine *Engine) ReplyToComment(userName string, subRedditName string, postID int, commentID int, replyContent string) {
	_, userExists := engine.userData[userName]
	if !userExists {
		return
	}

	subreddit, subredditExists := engine.subRedditData[subRedditName]
	if !subredditExists {
		return
	}
	
	var post *messages.Post
	for i, p := range subreddit.ListOfPosts {
		if p.ID == postID {
			post = &subreddit.ListOfPosts[i]
			break
		}
	}

	if post == nil {
		return
	}

	reply := messages.Comment{
		ID:        len(post.Comments) + 1,
		Content:   replyContent,
		UserName:  userName,
		Upvotes:   0,
		Downvotes: 0,
		Replies:   []messages.Comment{},
	}

	success := addReplyToComment(post.Comments, commentID, reply)
	if success {
		fmt.Printf("User %s replied to comment %d on post %d in subreddit %s.\n", userName, commentID, postID, subRedditName)
	} else {
		fmt.Printf("Comment %d not found in post %d in subreddit %s.\n", commentID, postID, subRedditName)
	}
}


func addReplyToComment(comments []messages.Comment, targetID int, reply messages.Comment) bool {
	for i := range comments {
		if comments[i].ID == targetID {
			comments[i].Replies = append(comments[i].Replies, reply)
			return true
		}
		if addReplyToComment(comments[i].Replies, targetID, reply) {
			return true
		}
	}
	return false
}
