package engine

type User struct {
	ID                 int
	Username           string
	Karma              int
	SubscribedSubreddits map[int]bool // Subreddit IDs
}

func (e *Engine) RegisterUser(username string) *User {
	e.Lock()
	defer e.Unlock()
	id := len(e.Users) + 1
	user := &User{
		ID:                 id,
		Username:           username,
		Karma:              0,
		SubscribedSubreddits: make(map[int]bool),
	}
	e.Users[id] = user
	return user
}

// func (e *Engine) UpdateKarma(userID, points int) {
// 	e.RLock()
// 	user, exists := e.Users[userID]
// 	e.RUnlock()
// 	if exists {
// 		user.Karma += points
// 	}
// }
