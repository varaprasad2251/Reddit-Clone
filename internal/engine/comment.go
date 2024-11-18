package engine

type Comment struct {
	ID        int
	Content   string
	AuthorID  int
	Upvotes   int
	Downvotes int
	Replies   []*Comment
}

func (e *Engine) CommentOnPost(post *Post, authorID int, content string) *Comment {
	comment := &Comment{
		ID:       len(post.Comments) + 1,
		Content:  content,
		AuthorID: authorID,
	}
	post.Comments = append(post.Comments, comment)
	return comment
}

func (post *Post) AddComment(comment *Comment, parentID int) {
	if parentID == 0 {
		post.Comments = append(post.Comments, comment)
	} else {
		// Recursively find parent comment
		for _, c := range post.Comments {
			addReply(c, comment, parentID)
		}
	}
}

func addReply(parent *Comment, reply *Comment, parentID int) bool {
	if parent.ID == parentID {
		parent.Replies = append(parent.Replies, reply)
		return true
	}
	for _, c := range parent.Replies {
		if addReply(c, reply, parentID) {
			return true
		}
	}
	return false
}