package blogposts_test

import (
	"errors"
	"io/fs"
	blogposts "learning-go-with-tests/reading-files"
	"reflect"
	"testing"
	"testing/fstest"
)

type StubFailiingFS struct {
}

func (s StubFailiingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("always failing")
}

func assertPost(t *testing.T, got blogposts.Post, want blogposts.Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func TestNewBlogPosts(t *testing.T) {
	const (
		firstBody = `Title: Post 1
Description: Description 1
Tags: tdd, go
---
Hello
World`
		secondBody = `Title: Post 2
Description: Description 2
Tags: rust, borrow-checker
---
B
L
M`
	)

	fs := fstest.MapFS{
		"hello world.md":  {Data: []byte(firstBody)},
		"hello-world2.md": {Data: []byte(secondBody)},
	}

	t.Run("blog post no error", func(t *testing.T) {
		posts, err := blogposts.NewPostsFromFS(fs)

		if err != nil {
			t.Fatal(err)
		}

		assertPost(t, posts[0], blogposts.Post{
			Title:       "Post 1",
			Description: "Description 1",
			Tags:        []string{"tdd", "go"},
			Body: `Hello
World`,
		})
	})

	t.Run("blog post fs always error", func(t *testing.T) {
		_, err := blogposts.NewPostsFromFS(StubFailiingFS{})

		if err == nil {
			t.Fatal("Expected error")
		}
	})

}
