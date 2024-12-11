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

	var result map[string]interface{}
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (c *RedditClient) RegisterUser(username string) (map[string]interface{}, error) {
	return c.makeRequest("POST", "/api/register", map[string]string{"username": username})
}

func (c *RedditClient) CreateSubreddit(name string) (map[string]interface{}, error) {
	return c.makeRequest("POST", "/api/subreddit", map[string]string{"name": name})
}

func (c *RedditClient) JoinSubreddit(username, subredditName string) (map[string]interface{}, error) {
	return c.makeRequest("POST", fmt.Sprintf("/api/user/%s/join", username), map[string]string{"name": subredditName})
}

func (c *RedditClient) LeaveSubreddit(username, subredditName string) (map[string]interface{}, error) {
	return c.makeRequest("POST", fmt.Sprintf("/api/user/%s/leave", username), map[string]string{"name": subredditName})
}

func (c *RedditClient) CreatePost(username, subredditName, content string) (map[string]interface{}, error) {
	return c.makeRequest("POST", "/api/submit", map[string]string{
		"username":  username,
		"subreddit": subredditName,
		"content":   content,
	})
}

func (c *RedditClient) UpvotePost(username string) (map[string]interface{}, error) {
	return c.makeRequest("POST", "/api/posts/1/upvote", map[string]string{"username": username})
}

func (c *RedditClient) DownvotePost(username string) (map[string]interface{}, error) {
	return c.makeRequest("POST", "/api/posts/1/downvote", map[string]string{"username": username})
}

func (c *RedditClient) SendDirectMessage(sender, content string) (map[string]interface{}, error) {
	return c.makeRequest("POST", "/api/message/compose", map[string]string{
		"sender":  sender,
		"content": content,
	})
}

func (c *RedditClient) GetUserInfo(username string) (map[string]interface{}, error) {
	return c.makeRequest("GET", fmt.Sprintf("/api/user/%s", username), nil)
}

func main() {
	client := NewRedditClient()

	// Register users
	fmt.Println("Registering users...")
	printResponse(client.RegisterUser("alice"))
	printResponse(client.RegisterUser("bob"))

	// Create subreddit
	fmt.Println("\nCreating subreddit...")
	printResponse(client.CreateSubreddit("programming"))

	// Join subreddit
	fmt.Println("\nJoining subreddit...")
	printResponse(client.JoinSubreddit("alice", "programming"))

	// Create post
	fmt.Println("\nCreating post...")
	printResponse(client.CreatePost("alice", "programming", "Hello, World!"))

	// Upvote post
	fmt.Println("\nUpvoting post...")
	printResponse(client.UpvotePost("bob"))

	// Downvote post
	fmt.Println("\nDownvoting post...")
	printResponse(client.DownvotePost("alice"))

	// Send direct message
	fmt.Println("\nSending direct message...")
	printResponse(client.SendDirectMessage("alice", "Hey Bob, how are you?"))

	// Get user info
	fmt.Println("\nGetting user info...")
	printResponse(client.GetUserInfo("alice"))

	// Leave subreddit
	fmt.Println("\nLeaving subreddit...")
	printResponse(client.LeaveSubreddit("alice", "programming"))
}

func printResponse(response map[string]interface{}, err error) {
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Response: %v\n", response)
}
