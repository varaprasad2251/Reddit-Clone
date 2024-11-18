package engine


import (
	"sort"
	"sync"
)

// Engine manages all users, subreddits, and posts
type Engine struct {
	Users      map[int]*User
	Subreddits map[int]*Subreddit
	Messages []*Message
	sync.RWMutex
}

func NewEngine() *Engine {
	return &Engine{
		Users:      make(map[int]*User),
		Subreddits: make(map[int]*Subreddit),
		Messages:   []*Message{},
	}
}


func (e *Engine) UpdateKarma(userID, points int) {
	e.RLock()
	user, exists := e.Users[userID]
	e.RUnlock()
	if exists {
		user.Karma += points
	}
}

func (e *Engine) GetFeed(userID int) []*Post {
	e.RLock()
	defer e.RUnlock()
	user, exists := e.Users[userID]
	if !exists {
		return nil
	}

	var feed []*Post
	for subredditID := range user.SubscribedSubreddits {
		if subreddit, exists := e.Subreddits[subredditID]; exists {
			feed = append(feed, subreddit.Posts...)
		}
	}
	sort.Slice(feed, func(i, j int) bool {
		return feed[i].Upvotes > feed[j].Upvotes
	})
	return feed
}