package blog

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"strings"
)

type Post struct{
	Title string
	Tags []string
	Body string
}
func NewPostFromFs(filesystem fs.FS) ([]Post, error){
	dir, err := fs.ReadDir(filesystem,".")
	if err != nil{
		return nil, err
	}
	var posts []Post
	for _,f := range dir{
		post, err := getPost(filesystem, f.Name())
		if err != nil{
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func getPost(filesystem fs.FS, filename string) (Post, error){
	postFile, err := filesystem.Open(filename)
	if err != nil{
		return Post{}, err
	}
	defer postFile.Close()
	return newPost(postFile)
}

const (
	titleSeparator = "Title: "
	tagsSeparator = "Tags: "
)

func newPost(postfile io.Reader) (Post,error){

	scanner := bufio.NewScanner(postfile)
	readMetaLine := func(tagName string) string{
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), tagName)
	}

	return Post{
		Title: readMetaLine(titleSeparator),	
		Tags: strings.Split(readMetaLine(tagsSeparator),", "),
		Body: readBody(scanner),
	},nil
}

func readBody(scanner *bufio.Scanner) string{
	scanner.Scan()
	buf := bytes.Buffer{}
	for scanner.Scan(){
		fmt.Fprintln(&buf, scanner.Text())
	}
	return strings.TrimSuffix(buf.String(),"\n")
}