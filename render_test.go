package blog_test

import (
	"bytes"
	"testing"

	approvals "github.com/approvals/go-approval-tests"
	"github.com/immanu10/blog"
)

func TestRen(t *testing.T){
	postRenderer, err := blog.NewPostRenderer()
	if err != nil{
		t.Fatal(err)
	}
	t.Run("it converts single post into html", func(t *testing.T){
		buf := bytes.Buffer{}
		err := postRenderer.Render(&buf,SamplePost)
		if err != nil{
			t.Fatal(err)
		}
		approvals.VerifyString(t, buf.String())
	})

	t.Run("it converts posts into index page", func(t *testing.T) {
		buf := bytes.Buffer{}
		err := postRenderer.RenderIndex(&buf, []blog.Post{SamplePost})
		if err != nil{
			t.Fatal(err)
		}
		approvals.VerifyString(t, buf.String())
	})
}