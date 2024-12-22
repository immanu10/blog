package blog_test

import (
	"reflect"
	"testing"
	"testing/fstest"

	"github.com/immanu10/blog"
)
var SamplePost = blog.Post{
		Title: "My first blog post",
		Body:"Hello, this is my first blog post. I hope you enjoy it." ,
		FileName: "hello-world",
		Date: "22 Dec 2024",
}

func TestCreateBlogPosts(t *testing.T) {
	body := `Title: My first blog post
Date: 22 Dec 2024
---
Hello, this is my first blog post. I hope you enjoy it.`

	fs := fstest.MapFS{
		"hello-world.md":{Data: []byte(body)},
	}

	posts, err := blog.NewPostFromFs(fs)

	if err != nil{
		t.Fatal(err)
	}
	assertPost(t, posts[0],SamplePost)
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