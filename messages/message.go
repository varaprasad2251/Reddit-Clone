package messages

type RegisterUser struct {
	UserName       string
	JointSubReddit []string
}

type UserDataType struct {
	UserName       string
	JointSubReddit []string
	Dm             []DM
	KarmaPoints    int
}

type DM struct {
	UserName string
	Content  string
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
type GetFeed struct{}
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
}
type DownVotePost struct {
	UserName string
}

// type Stats struct {
//     TotalUsers        int
//     TotalPosts        int
//     TotalSubreddits   int
//     TotalMessages     int
//     UserStats         map[string]UserStat
//     SimulationTime    time.Duration
// }

type UserStat struct {
    Karma     int
    PostCount int
}