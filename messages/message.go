package messages

import "time"

type RegisterUser struct {
	UserName       string
	JointSubReddit []string
}

type UserDataType struct {
	UserName       string
	JointSubReddit []string
	Dm             []DM
	KarmaPoints    int
	IsConnected bool
}

type DM struct {
	ID int
	UserName string
	Content  string
	Replies []DM
}

type SubReddit struct {
	ListOfPosts []Post
}

type Post struct {
	ID        int
	Content   string
	UserName  string
	Upvotes   int
	Downvotes int
	Comments  []Comment
	CreatedAt time.Time
}

type Comment struct {
	ID        int
	Content   string
	UserName  string
	Upvotes   int
	Downvotes int
	Replies   []Comment
}

type ReplyToComment struct {
	UserName      string
	SubRedditName string
	PostID        int
	CommentID     int
	ReplyContent  string
}

type CJLSubreddite struct{}
type PostInSubreddit struct{}
type CommentInSubReddit struct{}
type VoteInReddit struct{}
type GetFeed struct {
    UserName string
    Limit    int
}

type GetDM struct{}

type CreateSubreddit struct {
	Name string
}

type UserJoinSubReddit struct {
	UserName      string
	SubRedditName string
}

type UserLeaveSubReddit struct {
	UserName      string
	SubRedditName string
}

type CreatePost struct {
	UserName      string
	SubredditName string
	Content       Post
}

type SendDmToUser struct {
	UserName string
	Content  string
}

type ReplyToDm struct {
	UserName string
	Content  string
}

type UpVotePost struct {
    UserName string
    PostID   int
	TargetUser  string
}

type DownVotePost struct {
    UserName string
    PostID   int
	TargetUser  string
}

type UserStat struct {
    Karma     int
    PostCount int
}

type UpdateConnectionStatus struct {
    UserName    string
    IsConnected bool
}

type GetDirectMessages struct {
    UserName string
}