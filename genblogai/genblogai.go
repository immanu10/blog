package genblogai

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"math/rand"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/google/generative-ai-go/genai"
	"github.com/immanu10/blog"
	"google.golang.org/api/option"
)

const (
	systemInstruction = "You are a technical blog writer. Pick any single concept from the topic provided by the user and generator a .md string/blog post. Follow this template structure:\nTitle: A very very short concise title.\nA well-structured explanation of the concept with examples or code snippets if required.(Do not use python, use Go or JavaScript)"
)

var topicsList = []string{
	"JavaScript",
	"TypeScript Ecosystem and Best Practices",
	"TypeScript tips and tricks",
	"Go Programming Language",
	"Key concepts and features of Go",
	"High Level System Design",
	"Low Level System Design", 
	"Distributed Systems",
	"API Design and Microservices",
	"Data Structure and Algorithms",
	"Techniques for Solving DSA problems", 
	"Vim",
	"Vim tips and tricks",
	"Vim motion",
	"React Framework",
	"Frontend Development",
	"Web fundamentals",
	"Performance Optimization for Web Applications",
	"CSS and Responsive Design",
}


func GenerateBlogFromAI(prevPosts []blog.Post) (*bufio.Scanner, error){
	
	var titles []string
	for _,post := range prevPosts{
		titles = append(titles, post.Title)
	}  
	prevTitles := strings.Join(titles, ", ")

	ctx := context.Background()
	client, err := genai.NewClient(ctx,option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	if err != nil{
		return nil, err
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-1.5-pro")

	model.SetTemperature(1)
	model.SetTopK(40)
	model.SetTopP(0.95)
	model.SetMaxOutputTokens(8192)
	model.ResponseMIMEType = "text/plain"
	model.SystemInstruction = &genai.Content{
		Parts: []genai.Part{genai.Text(systemInstruction)},
	  }

	randomTopic := topicsList[rand.Intn(len(topicsList))]
	prompt := fmt.Sprintf("Topic is %s. Do not write about the following previously generated blogs (titles): %s",randomTopic, prevTitles)

	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil{
		return nil, err
	}

	if len(resp.Candidates) == 0 && len(resp.Candidates[0].Content.Parts) == 0{
		return nil, fmt.Errorf("no content generated")
	}

	mdText := resp.Candidates[0].Content.Parts[0].(genai.Text)

	scanner := bufio.NewScanner(strings.NewReader(string(mdText)))

	return scanner, nil
}

func GetFilenameFromAiContent(scanner *bufio.Scanner) (string, string, error){
	if !scanner.Scan(){
		return "", "", fmt.Errorf("no title found")
	}
	title := strings.TrimPrefix(scanner.Text(), "Title: ")
	filename := sanitizeTitle(title)
	return filename, title, nil
}
func WriteAiContentToFile(title string, scanner *bufio.Scanner, file io.Writer) error{

	_, err := file.Write([]byte(fmt.Sprintf("Title: %s\n", title)))
	if err != nil{
		return err
	}

	currentDate := time.Now().Format("02-Jan-2006")
	file.Write([]byte(fmt.Sprintf("Date: %s\n", currentDate)))

	for scanner.Scan(){
		file.Write([]byte(scanner.Text()+"\n"))
	}
	return nil
}
func sanitizeTitle(title string) string{
// remove special character using regex
	rx := regexp.MustCompile("[^a-zA-Z0-9]+")
	sanitized := rx.ReplaceAllString(title, "-")
	return strings.ToLower(sanitized)
}
