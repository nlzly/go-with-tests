package blogposts_test

import (
	"errors"
	"io/fs"
	blogposts "learning-go-with-tests/reading-files"
	"testing"
	"testing/fstest"
)

type StubFailiingFS struct {
}

func (s StubFailiingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("always failing")
}

func TestNewBlogPosts(t *testing.T) {
	fs := fstest.MapFS{
		"hello world.md":  {Data: []byte("hi")},
		"hello-world2.md": {Data: []byte("hola")},
	}

	t.Run("blog post no error", func(t *testing.T) {
		posts, err := blogposts.NewPostsFromFS(fs)

		if err != nil {
			t.Fatal(err)
		}

		if len(posts) != len(fs) {
			t.Errorf("got %d posts, wanted %d posts", len(posts), len(fs))
		}
	})

	t.Run("blog post fs always error", func(t *testing.T) {
		_, err := blogposts.NewPostsFromFS(StubFailiingFS{})

		if err == nil {
			t.Fatal("Expected error")
		}
	})

}
