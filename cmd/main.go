package main

import (
	"log"
	"os"
	"path"

	"github.com/immanu10/blog"
	"github.com/immanu10/blog/genblogai"
)


func main(){
	args := os.Args
	if len(args)>1{
		if args[1] == "build-all-posts"{
			readAndRenderAllPost()
			return
		}
	}
	return
	scanner, err := genblogai.GenerateBlogFromAI()
	if err != nil{
		log.Fatalf("Error generating blog from AI: %v", err)
	}

	blogFilename, blogTitle, err := genblogai.GetFilenameFromAiContent(scanner)
	if err != nil{
		log.Fatalf("Error getting filename from AI content: %v", err)
	}

	mdFile, err := os.Create(path.Join("posts", blogFilename+".md"))
	if err != nil{
		log.Fatalf("Error creating %s file: %v", blogFilename, err)
	}
	defer mdFile.Close()

	err = genblogai.WriteAiContentToFile(blogTitle, scanner, mdFile)
	if err != nil{
		log.Fatalf("Error writing AI content to %s file: %v", blogFilename, err)
	}

	pr, err := blog.NewPostRenderer()
	if err != nil{
		panic(err)
	}

	htmlFile, err := os.Create(path.Join("static", blogFilename+".html"))
	if err != nil{
		log.Fatalf("Error creating %s.html file: %v", blogFilename, err)
	}
	defer htmlFile.Close()
	
	post, err := blog.NewPostFromFile(os.DirFS("posts"), blogFilename+".md")
	if err != nil{
		log.Fatalf("Error reading %s.md file: %v", blogFilename, err)
	}

	renderErr := pr.Render(htmlFile, post)
	if renderErr != nil{
		log.Fatalf("Error rendering %s.md file: %v", post.FileName, err)
	}


	posts, err := blog.NewPostFromFs(os.DirFS("posts"))
	if err != nil{
		panic(err)
	}


	file, indexerr := os.Create("static/index.html")
	if indexerr != nil{
		log.Fatalf("Error creating index.html file: %v", err)
	}
	defer file.Close()

	indexRendererr := pr.RenderIndex(file, posts)
	if indexRendererr != nil{
		log.Fatalf("Error rendering index.html file: %+v", err)
	}
}


func readAndRenderAllPost(){

	pr, err := blog.NewPostRenderer()
	if err != nil{
		panic(err)
	}

	posts, err := blog.NewPostFromFs(os.DirFS("posts"))

	for _, post := range posts{
		htmlFile, err := os.Create(path.Join("static", post.FileName+".html"))
		if err != nil{
			log.Fatalf("Error creating %s.html file: %v", post.FileName, err)
		}
		defer htmlFile.Close()

		renderErr := pr.Render(htmlFile, post)
		if renderErr != nil{
			log.Fatalf("Error rendering %s.md file: %v", post.FileName, err)
		}
	}
}