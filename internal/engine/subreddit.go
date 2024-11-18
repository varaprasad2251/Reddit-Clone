package engine

import (
    "fmt"
	"sync"
)

type Subreddit struct {
	ID      int
	Name    string
	Posts   []*Post
	Members map[int]*User
	sync.Mutex // Add this for thread-safe operations
}

// func (e *Engine) CreateSubreddit(name string) *Subreddit {
// 	e.Lock()
// 	defer e.Unlock()
// 	id := len(e.Subreddits) + 1
// 	subreddit := &Subreddit{
// 		ID:      id,
// 		Name:    name,
// 		Posts:   []*Post{},
// 		Members: make(map[int]*User),
// 	}
// 	e.Subreddits[id] = subreddit
// 	return subreddit
// }

func (e *Engine) CreateSubreddit(name string) (*Subreddit, error) {
    e.Lock()
    defer e.Unlock()

    // Check if the subreddit name already exists
    for _, existingSubreddit := range e.Subreddits {
        if existingSubreddit.Name == name {
            return nil, fmt.Errorf("subreddit with name '%s' already exists", name)
        }
    }

    // Validate the subreddit name
    if len(name) == 0 {
        return nil, fmt.Errorf("subreddit name cannot be empty")
    }
    if len(name) > 21 {
        return nil, fmt.Errorf("subreddit name cannot exceed 21 characters")
    }

    id := len(e.Subreddits) + 1
    subreddit := &Subreddit{
        ID:      id,
        Name:    name,
        Posts:   []*Post{},
        Members: make(map[int]*User),
    }
    e.Subreddits[id] = subreddit

    return subreddit, nil
}


func (e *Engine) JoinSubreddit(userID, subredditID int) {
	e.RLock()
	user, userExists := e.Users[userID]
	subreddit, subredditExists := e.Subreddits[subredditID]
	e.RUnlock()
	
	if userExists && subredditExists {
		subreddit.Lock()
		subreddit.Members[userID] = user
		user.SubscribedSubreddits[subredditID] = true
		subreddit.Unlock()
	}
}

func (e *Engine) LeaveSubreddit(userID, subredditID int) {
	e.RLock()
	user, userExists := e.Users[userID]
	subreddit, subredditExists := e.Subreddits[subredditID]
	e.RUnlock()

	if userExists && subredditExists {
		subreddit.Lock()
		delete(subreddit.Members, userID)
		delete(user.SubscribedSubreddits, subredditID)
		subreddit.Unlock()
	}
}