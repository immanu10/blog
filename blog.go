package blog

import (
	"bufio"
	"bytes"
	"embed"
	"fmt"
	"html/template"
	"io"
	"io/fs"
	"sort"
	"strings"
	"time"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
)

type Post struct{
	Title string
	Body string
	FileName string
	Date string
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
	sortPostsByDate(posts)

	return posts, nil
}

func NewPostFromFile(filesystem fs.FS, filename string) (Post, error){
	return getPost(filesystem, filename)
}

func getPost(filesystem fs.FS, filename string) (Post, error){
	postFile, err := filesystem.Open(filename)
	if err != nil{
		return Post{}, err
	}
	defer postFile.Close()
	return newPost(postFile, filename)
}

const (
	titleSeparator = "Title: "
	dateSeparator = "Date: "
)

func newPost(postfile io.Reader, filename string) (Post,error){

	scanner := bufio.NewScanner(postfile)
	readMetaLine := func(tagName string) string{
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), tagName)
	}

	return Post{
		Title: readMetaLine(titleSeparator),	
		Date: readMetaLine(dateSeparator),
		Body: readBody(scanner),
		FileName: strings.TrimSuffix(filename, ".md"),
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
var (
	//go:embed "templates/*"
	postTemplates embed.FS
)
type PostRenderer struct{
	templ *template.Template
	mdParser *parser.Parser 
}
func NewPostRenderer() (*PostRenderer, error){
	templ, err := template.ParseFS(postTemplates, "templates/*.gohtml")
	if err != nil{
		return nil, err
	}
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs
	parser := parser.NewWithExtensions(extensions)

	return &PostRenderer{templ: templ, mdParser: parser}, nil
}

func (pr *PostRenderer) Render(w io.Writer, post Post) error{
	return pr.templ.ExecuteTemplate(w, "blog.gohtml", newPostVM(post, pr))
}
type postViewModel struct {
	Post
	HTMLBody template.HTML
	SeoTitle string
}

func newPostVM(p Post, r *PostRenderer) postViewModel {
	vm := postViewModel{Post: p,SeoTitle: p.Title}
	vm.HTMLBody = template.HTML(markdown.ToHTML([]byte(p.Body), r.mdParser, nil))
	return vm
}

func (r *PostRenderer) RenderIndex(w io.Writer, posts []Post) error{
	indexViewModel := struct{
		Posts []Post
		SeoTitle string
	}{
		Posts: posts,
		SeoTitle: "Blog",
	}
	return r.templ.ExecuteTemplate(w,"index.gohtml",indexViewModel)

}

func sortPostsByDate(posts []Post) {
	// Define the date layout for parsing
	const dateLayout = "02-Jan-2006"

	sort.Slice(posts, func(i, j int) bool {
		// Parse the dates
		dateI, err1 := time.Parse(dateLayout, posts[i].Date)
		dateJ, err2 := time.Parse(dateLayout, posts[j].Date)

		// Handle parsing errors by assuming invalid dates are older
		if err1 != nil || err2 != nil {
			fmt.Printf("Warning: Invalid date format for %s or %s\n", posts[i].Date, posts[j].Date)
			return err2 == nil // Keep valid dates before invalid ones
		}

		// Compare dates
		return dateI.After(dateJ)
	})
}