package engine
import "sync"

type Subreddit struct {
	ID      int
	Name    string
	Posts   []*Post
	Members map[int]*User
	sync.Mutex // Add this for thread-safe operations
}

func (e *Engine) CreateSubreddit(name string) *Subreddit {
	e.Lock()
	defer e.Unlock()
	id := len(e.Subreddits) + 1
	subreddit := &Subreddit{
		ID:      id,
		Name:    name,
		Posts:   []*Post{},
		Members: make(map[int]*User),
	}
	e.Subreddits[id] = subreddit
	return subreddit
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
