package timeline

import "time"

import (
	"github.com/ar-tty/testdata/blog/internal/blog"
	"github.com/ar-tty/testdata/blog/internal/users"
)

type RecordedAt time.Time

type Timeline struct {
	Post    blog.Post  `goofy:"required" json:"post"`
	User    users.User `goofy:"required" json:"user"`
	Created time.Time  `json:"created"`
	Updated RecordedAt `json:"updated"`
}
