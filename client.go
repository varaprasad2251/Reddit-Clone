package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const baseURL = "http://localhost:8080"

type RedditClient struct {
	client *http.Client
}

func NewRedditClient() *RedditClient {
	return &RedditClient{
		client: &http.Client{},
	}
}

func (c *RedditClient) makeRequest(method, endpoint string, body interface{}) (map[string]interface{}, error) {
	var buf bytes.Buffer
	if body != nil {
		if err := json.NewEncoder(&buf).Encode(body); err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, baseURL+endpoint, &buf)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result interface{}
    if err := json.Unmarshal(respBody, &result); err != nil {
        return nil, err
    }

    switch v := result.(type) {
    case map[string]interface{}:
        return v, nil
    case []interface{}:
        return map[string]interface{}{"data": v}, nil
    default:
        return nil, fmt.Errorf("unexpected response format")
    }
}

func (c *RedditClient) RegisterUser(username string) (map[string]interface{}, error) {
	return c.makeRequest("POST", "/api/register", map[string]string{"username": username})
}

func (c *RedditClient) GetUserInfo(username string) (map[string]interface{}, error) {
	return c.makeRequest("GET", fmt.Sprintf("/api/user/%s", username), nil)
}

func (c *RedditClient) JoinSubreddit(username, subredditName string) (map[string]interface{}, error) {
	return c.makeRequest("POST", fmt.Sprintf("/api/user/%s/join", username), map[string]string{"name": subredditName})
}

func (c *RedditClient) LeaveSubreddit(username, subredditName string) (map[string]interface{}, error) {
	return c.makeRequest("POST", fmt.Sprintf("/api/user/%s/leave", username), map[string]string{"name": subredditName})
}

func (c *RedditClient) CreateSubreddit(name string) (map[string]interface{}, error) {
	return c.makeRequest("POST", "/api/subreddit", map[string]string{"name": name})
}

func (c *RedditClient) GetSubredditInfo(name string) (map[string]interface{}, error) {
	return c.makeRequest("GET", fmt.Sprintf("/api/subreddit/%s", name), nil)
}

func (c *RedditClient) CreatePost(username, subredditName, content string) (map[string]interface{}, error) {
	return c.makeRequest("POST", "/api/submit", map[string]string{
		"username":  username,
		"subreddit": subredditName,
		"content":   content,
	})
}

func (c *RedditClient) UpvotePost(postID int, username string) (map[string]interface{}, error) {
	return c.makeRequest("POST", fmt.Sprintf("/api/posts/%d/upvote", postID), map[string]string{"username": username})
}

func (c *RedditClient) DownvotePost(postID int, username string) (map[string]interface{}, error) {
	return c.makeRequest("POST", fmt.Sprintf("/api/posts/%d/downvote", postID), map[string]string{"username": username})
}

func (c *RedditClient) CreateComment(username, subredditName string, postID int, content string) (map[string]interface{}, error) {
	return c.makeRequest("POST", "/api/comment", map[string]interface{}{
		"username":  username,
		"subreddit": subredditName,
		"post_id":   postID,
		"content":   content,
	})
}

func (c *RedditClient) SendDirectMessage(sender, recipient, content string) (map[string]interface{}, error) {
	return c.makeRequest("POST", "/api/message/compose", map[string]string{
		"sender":    sender,
		"recipient": recipient,
		"content":   content,
	})
}

func (c *RedditClient) GetDirectMessages(username string) ([]interface{}, error) {
    resp, err := c.makeRequest("GET", fmt.Sprintf("/api/message/inbox?username=%s", username), nil)
    if err != nil {
        return nil, err
    }
    
    messages, ok := resp["data"].([]interface{})
    if !ok {
        return nil, fmt.Errorf("unexpected response format")
    }
    
    return messages, nil
}

func (c *RedditClient) GetFeed(username string) ([]interface{}, error) {
    resp, err := c.makeRequest("GET", fmt.Sprintf("/api/feed?username=%s", username), nil)
    if err != nil {
        return nil, err
    }
    
    feed, ok := resp["data"].([]interface{})
    if !ok {
        return nil, fmt.Errorf("unexpected response format")
    }
    
    return feed, nil
}

func main() {
	client := NewRedditClient()

	// Register users
	fmt.Println("Registering users...")
	printResponse(client.RegisterUser("alice"))
	printResponse(client.RegisterUser("bob"))

	// Get user info
	fmt.Println("\nGetting user info...")
	printResponse(client.GetUserInfo("alice"))

	// Create subreddit
	fmt.Println("\nCreating subreddit...")
	printResponse(client.CreateSubreddit("programming"))

	// Get subreddit info
	fmt.Println("\nGetting subreddit info...")
	printResponse(client.GetSubredditInfo("programming"))

	// Join subreddit
	fmt.Println("\nJoining subreddit...")
	printResponse(client.JoinSubreddit("alice", "programming"))

	// Create post
	fmt.Println("\nCreating post...")
	printResponse(client.CreatePost("alice", "programming", "Hello, World!"))

	// Upvote post (assuming post ID is 1)
	fmt.Println("\nUpvoting post...")
	printResponse(client.UpvotePost(1, "bob"))

	// Downvote post (assuming post ID is 1)
	fmt.Println("\nDownvoting post...")
	printResponse(client.DownvotePost(1, "alice"))

	// Create comment
	fmt.Println("\nCreating comment...")
	printResponse(client.CreateComment("bob", "programming", 1, "Great post!"))

	// Send direct message
	fmt.Println("\nSending direct message...")
	printResponse(client.SendDirectMessage("alice", "bob", "Hey Bob, how are you?"))

	// Get direct messages
	fmt.Println("\nGetting direct messages...")
	printResponse(client.GetDirectMessages("bob"))

	// Get feed
	fmt.Println("\nGetting feed...")
	printResponse(client.GetFeed("alice"))

	// Leave subreddit
	fmt.Println("\nLeaving subreddit...")
	printResponse(client.LeaveSubreddit("alice", "programming"))
}

func printResponse(response interface{}, err error) {
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }

    switch v := response.(type) {
    case map[string]interface{}:
        fmt.Printf("Response: %v\n", v)
    case []interface{}:
        fmt.Println("Response List:")
        for _, item := range v {
            fmt.Printf("%v\n", item)
        }
    default:
        fmt.Println("Unexpected response type")
    }
}
