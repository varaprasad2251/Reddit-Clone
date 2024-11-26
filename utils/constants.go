package utils

type User struct {
	UserName string
	Karma    int
	// Inbox  chan Message
	//Active bool
}

type Post struct {
	ID        int
	Content   string
	AuthorID  int
	Upvotes   int
	Downvotes int
	// Comments  []*Comment
}

type Comment struct {
	ID        int
	Content   string
	AuthorID  int
	Upvotes   int
	Downvotes int
	// Replies   []*Comment
}

type Message struct {
	SenderID int
	Content  string
}

const RegisterUser = "registerUser"
const CjlSubreddite = "cjlSubreddite"
const PostInSubreddit = "postInSubreddit"
const CommentInSubReddit = "commentInSubReddit"
const VoteInReddit = "voteInReddit"
const GetFeed = "getFeed"
const GetDM = "getDM"
