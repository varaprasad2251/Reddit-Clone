package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
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

func simulateUser(username string, wg *sync.WaitGroup) {
	defer wg.Done()

	client := NewRedditClient()

	// Register user
	fmt.Printf("[%s] Registering user...\n", username)
	printResponse(client.RegisterUser(username))

	// Create and join subreddit
	subredditName := username + "_subreddit"
	fmt.Printf("[%s] Creating subreddit %s...\n", username, subredditName)
	printResponse(client.CreateSubreddit(subredditName))

	fmt.Printf("[%s] Joining subreddit %s...\n", username, subredditName)
	printResponse(client.JoinSubreddit(username, subredditName))

	// Create post
	fmt.Printf("[%s] Creating post in %s...\n", username, subredditName)
	printResponse(client.CreatePost(username, subredditName, "Hello from "+username))

	// Upvote and downvote posts
	fmt.Printf("[%s] Upvoting a post...\n", username)
	printResponse(client.UpvotePost(username))

	fmt.Printf("[%s] Downvoting a post...\n", username)
	printResponse(client.DownvotePost(username))

	// Send direct message
	fmt.Printf("[%s] Sending a direct message...\n", username)
	printResponse(client.SendDirectMessage(username, "Hello from "+username))

	// Get user info
	fmt.Printf("[%s] Getting user info...\n", username)
	printResponse(client.GetUserInfo(username))

	// Leave subreddit
	fmt.Printf("[%s] Leaving subreddit %s...\n", username, subredditName)
	printResponse(client.LeaveSubreddit(username, subredditName))
}

func main() {
	numUsers := 5
	var wg sync.WaitGroup

	for i := 1; i <= numUsers; i++ {
		wg.Add(1)
		go simulateUser(fmt.Sprintf("user%d", i), &wg)
		time.Sleep(100 * time.Millisecond) // Small delay to stagger user actions
	}

	wg.Wait()
	fmt.Println("All users have completed their actions.")
}

func printResponse(response map[string]interface{}, err error) {
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Response: %v\n", response)
}
