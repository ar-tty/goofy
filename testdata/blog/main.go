package main

import (
	"fmt"

	"github.com/ar-tty/testdata/blog/internal/blog"
	"github.com/ar-tty/testdata/blog/internal/timeline"
	"github.com/ar-tty/testdata/blog/internal/users"
)

func main() {
	fmt.Println(users.User{})
	fmt.Println(timeline.Timeline{})
	fmt.Println(blog.Post{})
}
