package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"path"
	"regexp"
	"strings"
	"time"

	"github.com/google/generative-ai-go/genai"
	"github.com/immanu10/blog"
	"google.golang.org/api/option"
)


func main(){
	//posts, err := blog.NewPostFromFs(os.DirFS("posts"))
	//if err != nil{
		//panic(err)
	//}

	//for _,post := range posts{

		//pr, err := blog.NewPostRenderer()
		//if err != nil{
			//panic(err)
		//}

	//}

	


	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-1.5-flash")
	model.SetTemperature(1)
	model.SetTopK(40)
	model.SetTopP(0.95)
	model.SetMaxOutputTokens(8192)
	model.ResponseMIMEType = "text/plain"
	model.SystemInstruction = &genai.Content{
		Parts: []genai.Part{genai.Text("You generate .md string/blog post on a technology-related content. Choose any 1 topic randomly from the topics list array given by the user and generate single post which focus on a single, interesting concept within that single topic. Avoid explaining multiple concepts within a single blog post.\nFollow this template structure:\nTitle: A concise title (Do not use any special characters)\nA well-structured explanation of the chosen topic with example or code snippets if required.")},
	  }
	topicsList := []string{"Javascript Interview Questions", "Typescript tips and Trick", "Go topics", "High level system design", "Low level system design", "DSA Interview Questions", "React Interview Questions"}
	randomTopic := topicsList[rand.Intn(len(topicsList))]

	resp, err := model.GenerateContent(ctx, genai.Text("Topic: "+randomTopic)) 
	if err != nil {
		log.Fatal(err)
	}

	newMdFileName, err := writeToFile(resp)
	if err != nil{
		log.Fatalf("Error writing to md file: %v", err)
	}

	pr, err := blog.NewPostRenderer()
	if err != nil{
		panic(err)
	}
	file, err := os.Create(path.Join("static", newMdFileName+".html"))

	if err != nil{
		log.Fatalf("Error creating %s.html file: %v", newMdFileName, err)
	}

	defer file.Close()
	
	post, err := blog.NewPostFromFile(os.DirFS("posts"), newMdFileName+".md")
	if err != nil{
		log.Fatalf("Error reading %s.md file: %v", newMdFileName, err)
	}
	renderErr := pr.Render(file, post)
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

func writeToFile(resp *genai.GenerateContentResponse) (string, error) {
	if len(resp.Candidates) == 0 && len(resp.Candidates[0].Content.Parts) == 0 {
		return "", fmt.Errorf("No content generated") 
	}

	mdtext := resp.Candidates[0].Content.Parts[0].(genai.Text)
	
	scanner := bufio.NewScanner(strings.NewReader(string(mdtext)))
	if !scanner.Scan(){
		return "", fmt.Errorf("Failed to read title")
	}
	title := strings.TrimPrefix(scanner.Text(), "Title: ")
	filename := sanitizeFilename(title)
	file, err := os.Create(path.Join("posts", filename+".md"))

	if err != nil{
		log.Fatalf("Error creating %s file: %v", filename, err)
	}
	defer file.Close()

	file.WriteString(fmt.Sprintf("Title: %s\n", title))

	currentDate := time.Now().Format("02-Jan-2006")
	
	file.WriteString(fmt.Sprintf("Date: %s\n", currentDate))

	for scanner.Scan(){
		file.WriteString(scanner.Text()+"\n")
	}

	if err := scanner.Err(); err != nil{
		return "", err
	}
	return filename, nil

}

func sanitizeFilename(name string) string{
	// remove all special characters, convert to lower case and replace spaces with hyphens
	re := regexp.MustCompile(`[^a-zA-Z0-9\s]`)
	sanitize := re.ReplaceAllString(name, "")

	sanitize = strings.ToLower(sanitize)

	return strings.ReplaceAll(sanitize, " ", "-")

}