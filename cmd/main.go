package main

import (
	"log"
	"os"
	"path"

	"github.com/immanu10/blog"
)


func main(){
	posts, err := blog.NewPostFromFs(os.DirFS("posts"))
	if err != nil{
		panic(err)
	}

	for _,post := range posts{

		pr, err := blog.NewPostRenderer()
		if err != nil{
			panic(err)
		}

		file, err := os.Create(path.Join("static", post.FileName+".html"))

		if err != nil{
			log.Fatalf("Error creating %s.md file: %v", post.FileName, err)
		}

		defer file.Close()

		renderErr := pr.Render(file, post)
		if renderErr != nil{
			log.Fatalf("Error rendering %s.md file: %v", post.FileName, err)
		}
	}

	
	pr, err := blog.NewPostRenderer()
	if err != nil{
		panic(err)
	}

	file, err := os.Create("static/index.html")
	if err != nil{
		log.Fatalf("Error creating index.html file: %v", err)
	}
	defer file.Close()

	indexRendererr := pr.RenderIndex(file, posts)
	if indexRendererr != nil{
		log.Fatalf("Error rendering index.html file: %v", err)
	}
}