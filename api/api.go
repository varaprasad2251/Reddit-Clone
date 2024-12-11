// api/api.go

package api

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "cop5615-project4/Engine"
    "cop5615-project4/messages"
)

type API struct {
    engine *Engine.Engine
    router *gin.Engine
}

func NewAPI(engine *Engine.Engine) *API {
    api := &API{
        engine: engine,
        router: gin.Default(),
    }
    api.setupRoutes()
    return api
}

func (api *API) setupRoutes() {
    // User-related routes
    api.router.POST("/api/register", api.registerUser)
    api.router.GET("/api/user/:username", api.getUserInfo)
    api.router.POST("/api/user/:username/join", api.joinSubreddit)
    api.router.POST("/api/user/:username/leave", api.leaveSubreddit)

    // Subreddit-related routes
    api.router.GET("/api/subreddit/:name", api.getSubredditInfo)
    api.router.POST("/api/subreddit", api.createSubreddit)

    // Post-related routes
    api.router.POST("/api/submit", api.createPost)
    api.router.GET("/api/posts/:id", api.getPost)
    api.router.POST("/api/posts/:id/upvote", api.upvotePost)
    api.router.POST("/api/posts/:id/downvote", api.downvotePost)

    // Comment-related routes
    api.router.POST("/api/comment", api.createComment)
    api.router.GET("/api/comments/:id", api.getComment)

    // Message-related routes
    api.router.POST("/api/message/compose", api.sendDirectMessage)
    api.router.POST("/api/message/reply-all", api.replyToAllDMs)
    api.router.GET("/api/message/inbox", api.getDirectMessages)

    // Feed-related routes
    api.router.GET("/api/feed", api.getFeed)
}

func (api *API) Run(addr string) error {
    return api.router.Run(addr)
}

func (api *API) registerUser(c *gin.Context) {
    var user struct {
        UserName string `json:"username"`
    }
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    api.engine.RegisterUser(user.UserName)
    c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func (api *API) getUserInfo(c *gin.Context) {
    username := c.Param("username")
    userInfo, exists := api.engine.GetUserData(username)
    if !exists {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }
    c.JSON(http.StatusOK, userInfo)
}

func (api *API) joinSubreddit(c *gin.Context) {
    username := c.Param("username")
    var subreddit struct {
        Name string `json:"name"`
    }
    if err := c.ShouldBindJSON(&subreddit); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    api.engine.SubredditSpecificOp("join", username, subreddit.Name)
    c.JSON(http.StatusOK, gin.H{"message": "User joined subreddit successfully"})
}

func (api *API) leaveSubreddit(c *gin.Context) {
    username := c.Param("username")
    var subreddit struct {
        Name string `json:"name"`
    }
    if err := c.ShouldBindJSON(&subreddit); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    api.engine.SubredditSpecificOp("leave", username, subreddit.Name)
    c.JSON(http.StatusOK, gin.H{"message": "User left subreddit successfully"})
}

func (api *API) getSubredditInfo(c *gin.Context) {
    name := c.Param("name")
    subredditInfo, exists := api.engine.GetSubRedditData(name)
    if !exists {
        c.JSON(http.StatusNotFound, gin.H{"error": "Subreddit not found"})
        return
    }
    c.JSON(http.StatusOK, subredditInfo)
}

func (api *API) createSubreddit(c *gin.Context) {
    var subreddit struct {
        Name string `json:"name"`
    }
    if err := c.ShouldBindJSON(&subreddit); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    api.engine.CreateSubReddit(subreddit.Name)
    c.JSON(http.StatusOK, gin.H{"message": "Subreddit created successfully"})
}

func (api *API) createPost(c *gin.Context) {
    var post struct {
        UserName      string `json:"username"`
        SubredditName string `json:"subreddit"`
        Content       string `json:"content"`
    }
    if err := c.ShouldBindJSON(&post); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    newPost := messages.Post{
        ID:        api.engine.GetSubRedditPostCount(post.SubredditName) + 1,
        Content:   post.Content,
        UserName:  post.UserName,
        Upvotes:   0,
        Downvotes: 0,
        Comments:  []messages.Comment{},
    }

    api.engine.CreatePost(post.UserName, post.SubredditName, newPost)
    c.JSON(http.StatusOK, gin.H{"message": "Post created successfully"})
}

func (api *API) getPost(c *gin.Context) {
    // Implementation for getting a specific post
    // This would require adding a method to your Engine to fetch a post by ID
    c.JSON(http.StatusNotImplemented, gin.H{"error": "Not implemented"})
}

func (api *API) upvotePost(c *gin.Context) {
    var vote struct {
        UserName string `json:"username"`
    }
    if err := c.ShouldBindJSON(&vote); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    api.engine.UpvoteRandomPost(vote.UserName)
    c.JSON(http.StatusOK, gin.H{"message": "Post upvoted successfully"})
}

func (api *API) downvotePost(c *gin.Context) {
    var vote struct {
        UserName string `json:"username"`
    }
    if err := c.ShouldBindJSON(&vote); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    api.engine.DownvoteRandomPost(vote.UserName)
    c.JSON(http.StatusOK, gin.H{"message": "Post downvoted successfully"})
}

func (api *API) createComment(c *gin.Context) {
    var comment struct {
        UserName      string `json:"username"`
        SubredditName string `json:"subreddit"`
        PostID        int    `json:"post_id"`
        CommentID     int    `json:"comment_id"`
        Content       string `json:"content"`
    }
    if err := c.ShouldBindJSON(&comment); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    api.engine.ReplyToComment(comment.UserName, comment.SubredditName, comment.PostID, comment.CommentID, comment.Content)
    c.JSON(http.StatusOK, gin.H{"message": "Comment created successfully"})
}

func (api *API) getComment(c *gin.Context) {
    // Implementation for getting a specific comment
    // This would require adding a method to your Engine to fetch a comment by ID
    c.JSON(http.StatusNotImplemented, gin.H{"error": "Not implemented"})
}

func (api *API) sendDirectMessage(c *gin.Context) {
    var dm struct {
        Sender  string `json:"sender"`
        Content string `json:"content"`
    }
    if err := c.ShouldBindJSON(&dm); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    api.engine.SendDMtoUser(dm.Sender, dm.Content)
    c.JSON(http.StatusOK, gin.H{"message": "Direct message sent successfully"})
}

func (api *API) replyToAllDMs(c *gin.Context) {
    var reply struct {
        Sender  string `json:"sender"`
        Content string `json:"content"`
    }
    if err := c.ShouldBindJSON(&reply); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    api.engine.ReplyToAllDMs(reply.Sender, reply.Content)
    c.JSON(http.StatusOK, gin.H{"message": "Replied to all DMs successfully"})
}

func (api *API) getDirectMessages(c *gin.Context) {
    username := c.Query("username")
    user, exists := api.engine.GetUserData(username)
    if !exists {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }

    c.JSON(http.StatusOK, user.Dm)
}

func (api *API) getFeed(c *gin.Context) {
    username := c.Query("username")
    user, exists := api.engine.GetUserData(username)
    if !exists {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }

    var feed []messages.Post
    for _, subredditName := range user.JointSubReddit {
        subreddit, exists := api.engine.GetSubRedditData(subredditName)
        if exists {
            feed = append(feed, subreddit.ListOfPosts...)
        }
    }

    c.JSON(http.StatusOK, feed)
}
