package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

const (
	baseURL      = "http://localhost:8080"
	numUsers     = 5
	numActions   = 50
	maxSubreddit = 10
	maxPostID    = 1000
	maxCommentID = 100
)

type RedditClient struct {
	client  *http.Client
	baseURL string
}

func NewRedditClient() *RedditClient {
	return &RedditClient{
		client:  &http.Client{},
		baseURL: baseURL,
	}
}

func (c *RedditClient) makeRequest(method, endpoint string, body interface{}) (map[string]interface{}, error) {
	var buf bytes.Buffer
	if body != nil {
		if err := json.NewEncoder(&buf).Encode(body); err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, c.baseURL+endpoint, &buf)
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

func (c *RedditClient) CreatePost(username, subredditName, content string) (map[string]interface{}, error) {
	return c.makeRequest("POST", "/api/submit", map[string]string{
		"username":  username,
		"subreddit": subredditName,
		"content":   content,
	})
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

func (c *RedditClient) UpvotePost(postID int, username string) (map[string]interface{}, error) {
	return c.makeRequest("POST", fmt.Sprintf("/api/posts/%d/upvote", postID), map[string]string{"username": username})
}

func (c *RedditClient) DownvotePost(postID int, username string) (map[string]interface{}, error) {
	return c.makeRequest("POST", fmt.Sprintf("/api/posts/%d/downvote", postID), map[string]string{"username": username})
}

func (c *RedditClient) GetFeed(username string) (map[string]interface{}, error) {
	return c.makeRequest("GET", fmt.Sprintf("/api/feed?username=%s", username), nil)
}

func (c *RedditClient) GetDirectMessages(username string) (map[string]interface{}, error) {
	return c.makeRequest("GET", fmt.Sprintf("/api/message/inbox?username=%s", username), nil)
}

func (c *RedditClient) GetUserInfo(username string) (map[string]interface{}, error) {
    return c.makeRequest("GET", fmt.Sprintf("/api/user/%s", username), nil)
}

func simulateUser(client *RedditClient, userName string, wg *sync.WaitGroup) {
	defer wg.Done()

	rand.Seed(time.Now().UnixNano())

	_, err := client.RegisterUser(userName)
	if err != nil {
		fmt.Printf("Error registering user %s: %v\n", userName, err)
		return
	}

	actions := []func(){
		func() { joinRandomSubreddit(client, userName) },
		func() { createPost(client, userName) },
		func() { replyToRandomComment(client, userName) },
		func() { sendDirectMessage(client, userName) },
		func() { upvoteRandomPost(client, userName) },
		func() { downvoteRandomPost(client, userName) },
		func() { getFeed(client, userName) },
		func() { getDirectMessages(client, userName) },
		func() { randomVote(client, userName) },
	}

	for i := 0; i < numActions; i++ {
		if rand.Float32() < 0.3 {
			randomVote(client, userName)
		} else {
			action := actions[rand.Intn(len(actions))]
			action()
		}
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	}
}

func joinRandomSubreddit(client *RedditClient, userName string) {
	subredditName := fmt.Sprintf("subreddit_%d", rand.Intn(maxSubreddit))
	_, err := client.JoinSubreddit(userName, subredditName)
	if err != nil {
		fmt.Printf("Error joining subreddit for user %s: %v\n", userName, err)
	}
}

func (c *RedditClient) LeaveSubreddit(username, subredditName string) (map[string]interface{}, error) {
    return c.makeRequest("POST", fmt.Sprintf("/api/user/%s/leave", username), map[string]string{"name": subredditName})
}

func (c *RedditClient) GetSubredditInfo(name string) (map[string]interface{}, error) {
    return c.makeRequest("GET", fmt.Sprintf("/api/subreddit/%s", name), nil)
}

func createPost(client *RedditClient, userName string) {
	subredditName := fmt.Sprintf("subreddit_%d", rand.Intn(maxSubreddit))
	content := fmt.Sprintf("This is post #%d by %s", rand.Intn(1000), userName)
	_, err := client.CreatePost(userName, subredditName, content)
	if err != nil {
		fmt.Printf("Error creating post for user %s: %v\n", userName, err)
	}
}

func replyToRandomComment(client *RedditClient, userName string) {
	subredditName := fmt.Sprintf("subreddit_%d", rand.Intn(maxSubreddit))
	postID := rand.Intn(maxPostID)
	content := fmt.Sprintf("Reply from %s: %s", userName, generateRandomContent())
	_, err := client.CreateComment(userName, subredditName, postID, content)
	if err != nil {
		fmt.Printf("Error replying to comment for user %s: %v\n", userName, err)
	}
}

func sendDirectMessage(client *RedditClient, userName string) {
	recipientName := fmt.Sprintf("User%d", rand.Intn(numUsers))
	content := fmt.Sprintf("DM from %s: %s", userName, generateRandomContent())
	_, err := client.SendDirectMessage(userName, recipientName, content)
	if err != nil {
		fmt.Printf("Error sending DM for user %s: %v\n", userName, err)
	}
}

func upvoteRandomPost(client *RedditClient, userName string) {
	postID := rand.Intn(maxPostID)
	_, err := client.UpvotePost(postID, userName)
	if err != nil {
		fmt.Printf("Error upvoting post for user %s: %v\n", userName, err)
	}
}

func downvoteRandomPost(client *RedditClient, userName string) {
	postID := rand.Intn(maxPostID)
	_, err := client.DownvotePost(postID, userName)
	if err != nil {
		fmt.Printf("Error downvoting post for user %s: %v\n", userName, err)
	}
}

func getFeed(client *RedditClient, userName string) {
	_, err := client.GetFeed(userName)
	if err != nil {
		fmt.Printf("Error getting feed for user %s: %v\n", userName, err)
	}
}

func getDirectMessages(client *RedditClient, userName string) {
	_, err := client.GetDirectMessages(userName)
	if err != nil {
		fmt.Printf("Error getting DMs for user %s: %v\n", userName, err)
	}
}

func randomVote(client *RedditClient, userName string) {
	postID := rand.Intn(maxPostID)
	if rand.Float32() < 0.7 {
		_, err := client.UpvotePost(postID, userName)
		if err != nil {
			fmt.Printf("Error upvoting post for user %s: %v\n", userName, err)
		}
	} else {
		_, err := client.DownvotePost(postID, userName)
		if err != nil {
			fmt.Printf("Error downvoting post for user %s: %v\n", userName, err)
		}
	}
}

func generateRandomContent() string {
	contents := []string{
		"The moon is just a hologram!",
		"Does anyone else hear that faint buzzing?",
		"I can't believe it's not butter!",
		"Why are ducks so underrated?",
		"Time travel is overrated.",
		"Bananas are a government conspiracy!",
		"This reminds me of my pet hamster, Gerald.",
		"Is anyone else craving tacos right now?",
		"Pineapples on pizza? Let's discuss.",
		"I'm 90% sure this is a simulation.",
		"The cake is a lie!",
		"Who let the dogs out?",
		"This post smells like teen spirit.",
		"One does not simply ignore this post.",
		"Aliens are among us. Trust me.",
		"This made my goldfish do a backflip.",
		"I'm typing this with my toes.",
		"The sky just winked at me. Weird.",
		"I'm not crying, you're crying.",
		"Did you know otters hold hands?",
	}
	return contents[rand.Intn(len(contents))]
}


func main() {
    client1 := NewRedditClient()
    client2 := NewRedditClient()

    fmt.Println("Client 1: Registering user Alice")
    printResponse(client1.RegisterUser("Alice"))

    fmt.Println("\nClient 2: Registering user Bob")
    printResponse(client2.RegisterUser("Bob"))

    fmt.Println("\nClient 1: Creating subreddit 'programming'")
    printResponse(client1.CreateSubreddit("programming"))

    fmt.Println("\nClient 2: Creating subreddit 'technology'")
    printResponse(client2.CreateSubreddit("technology"))

    fmt.Println("\nClient 2: Joining subreddit 'programming'")
    printResponse(client2.JoinSubreddit("Bob", "programming"))

    fmt.Println("\nClient 1: Joining subreddit 'technology'")
    printResponse(client1.JoinSubreddit("Alice", "technology"))

    fmt.Println("\nClient 1: Creating post in 'programming'")
    printResponse(client1.CreatePost("Alice", "programming", "Hello, World!"))

    fmt.Println("\nClient 2: Creating post in 'technology'")
    printResponse(client2.CreatePost("Bob", "technology", "AI is the future!"))

    fmt.Println("\nClient 2: Upvoting post in 'programming'")
    printResponse(client2.UpvotePost(1, "Bob"))

    fmt.Println("\nClient 1: Upvoting post in 'technology'")
    printResponse(client1.UpvotePost(2, "Alice"))

    fmt.Println("\nClient 1: Creating another post in 'programming'")
    printResponse(client1.CreatePost("Alice", "programming", "Go vs Rust: A comparison"))

    fmt.Println("\nClient 2: Downvoting post in 'programming'")
    printResponse(client2.DownvotePost(3, "Bob"))

    fmt.Println("\nClient 2: Sending direct message to Alice")
    printResponse(client2.SendDirectMessage("Bob", "Alice", "Hi Alice!"))

    fmt.Println("\nClient 1: Sending direct message to Bob")
    printResponse(client1.SendDirectMessage("Alice", "Bob", "How are you Bob?"))

    fmt.Println("\nClient 1: Sending another direct message to Bob")
    printResponse(client1.SendDirectMessage("Alice", "Bob", "How is your day?"))

    fmt.Println("\nClient 1: Leaving subreddit 'programming'")
    printResponse(client1.LeaveSubreddit("Alice", "programming"))

    fmt.Println("\nClient 2: Getting subreddit info for 'programming'")
    printResponse(client2.GetSubredditInfo("programming"))

    fmt.Println("\nClient 1: Getting subreddit info for 'technology'")
    printResponse(client1.GetSubredditInfo("technology"))
}


func printResponse(response interface{}, err error) {
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }
    fmt.Printf("Response: %v\n", response)
}