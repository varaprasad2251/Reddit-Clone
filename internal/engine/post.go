package engine

type Post struct {
	ID        int
	Content   string
	AuthorID  int
	Upvotes   int
	Downvotes int
	Comments  []*Comment
}

func (e *Engine) CreatePost(subredditID, authorID int, content string) *Post {
	e.RLock()
	subreddit, exists := e.Subreddits[subredditID]
	e.RUnlock()
	if !exists {
		return nil
	}

	post := &Post{
		ID:       len(subreddit.Posts) + 1,
		Content:  content,
		AuthorID: authorID,
	}
	subreddit.Lock()
	subreddit.Posts = append(subreddit.Posts, post)
	subreddit.Unlock()
	return post
}
