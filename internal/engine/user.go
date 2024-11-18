package engine

import (
    "fmt"
)

type User struct {
	ID                 int
	Username           string
	Karma              int
	SubscribedSubreddits map[int]bool // Subreddit IDs
}

// func (e *Engine) RegisterUser(username string) *User {
// 	e.Lock()
// 	defer e.Unlock()
// 	id := len(e.Users) + 1
// 	user := &User{
// 		ID:                 id,
// 		Username:           username,
// 		Karma:              0,
// 		SubscribedSubreddits: make(map[int]bool),
// 	}
// 	e.Users[id] = user
// 	return user
// }

func (e *Engine) RegisterUser(username string) (*User, error) {
    e.Lock()
    defer e.Unlock()

    // Check if username already exists
    for _, existingUser := range e.Users {
        if existingUser.Username == username {
            return nil, fmt.Errorf("username '%s' already exists", username)
        }
    }

    id := len(e.Users) + 1
    user := &User{
        ID:                   id,
        Username:             username,
        Karma:                0,
        SubscribedSubreddits: make(map[int]bool),
    }

    e.Users[id] = user

    // Log user creation
    // fmt.Println("User created: ID=%d, Username=%s", id, username)

    return user, nil
}

// func (e *Engine) UpdateKarma(userID, points int) {
// 	e.RLock()
// 	user, exists := e.Users[userID]
// 	e.RUnlock()
// 	if exists {
// 		user.Karma += points
// 	}
// }
