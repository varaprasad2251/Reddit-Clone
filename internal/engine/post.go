package engine

import (
    "fmt"
    "log"
)

type Post struct {
	ID        int
	Content   string
	AuthorID  int
	Upvotes   int
	Downvotes int
	Comments  []*Comment
}

func (e *Engine) CreatePost(subredditID, authorID int, content string) (*Post, error) {
    e.RLock()
    subreddit, exists := e.Subreddits[subredditID]
    e.RUnlock()
    if !exists {
        return nil, fmt.Errorf("subreddit with ID %d does not exist", subredditID)
    }

    subreddit.Lock()
    defer subreddit.Unlock()

    postID := len(subreddit.Posts) + 1
    post := &Post{
        ID:       postID,
        Content:  content,
        AuthorID: authorID,
    }
    subreddit.Posts = append(subreddit.Posts, post)

    // Log successful post creation
    log.Printf("Post created: ID=%d, SubredditID=%d, AuthorID=%d", postID, subredditID, authorID)

    return post, nil
}

// func (e *Engine) CreatePost(subredditID, authorID int, content string) *Post {
//     e.RLock()
//     subreddit, exists := e.Subreddits[subredditID]
//     e.RUnlock()
//     if !exists {
//         return nil
//     }

//     post := &Post{
//         ID:       len(subreddit.Posts) + 1,
//         Content:  content,
//         AuthorID: authorID,
//     }
//     subreddit.Lock()
//     subreddit.Posts = append(subreddit.Posts, post)
//     subreddit.Unlock()

//     // Log successful post creation
//     log.Printf("Post created: ID=%d, SubredditID=%d, AuthorID=%d", post.ID, subredditID, authorID)

//     return post
// }

func (post *Post) AddComment(comment *Comment, parentID int) {
	if parentID == 0 {
		post.Comments = append(post.Comments, comment)
	} else {
		// Find parent comment and add reply
		for _, c := range post.Comments {
			addReply(c, comment, parentID)
		}
	}
}

func addReply(parent *Comment, reply *Comment, parentID int) bool {
	if parent.ID == parentID {
		parent.Replies = append(parent.Replies, reply)
		return true
	}
	for _, c := range parent.Replies {
		if addReply(c, reply, parentID) {
			return true
		}
	}
	return false
}