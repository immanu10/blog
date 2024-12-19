package blog_test

import (
	"reflect"
	"testing"
	"testing/fstest"

	"github.com/immanu10/blog"
)

func TestCreateBlogPosts(t *testing.T) {
	body := `Title: My first blog post
Tags: blog, golang
---
Hello, this is my first blog post. I hope you enjoy it.`

	fs := fstest.MapFS{
		"hello-world.ms":{Data: []byte(body)},
	}

	posts, err := blog.NewPostFromFs(fs)

	if err != nil{
		t.Fatal(err)
	}
	assertPost(t, posts[0],blog.Post{
		Title: "My first blog post",
		Tags: []string{"blog","golang"},
		Body:"Hello, this is my first blog post. I hope you enjoy it." ,
	})
	if len(posts) != len(fs){
		t.Errorf("got %d posts, want %d length",len(posts), len(fs))
	}

}

func assertPost(t *testing.T, got blog.Post, want blog.Post){
	t.Helper()	
	if !reflect.DeepEqual(got, want){
		t.Errorf("got %+v, want %+v",got, want)
	}
}