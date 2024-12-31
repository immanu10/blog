package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
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

		//file, err := os.Create(path.Join("static", post.FileName+".html"))

		//if err != nil{
			//log.Fatalf("Error creating %s.md file: %v", post.FileName, err)
		//}

		//defer file.Close()

		//renderErr := pr.Render(file, post)
		//if renderErr != nil{
			//log.Fatalf("Error rendering %s.md file: %v", post.FileName, err)
		//}
	//}

	
	//pr, err := blog.NewPostRenderer()
	//if err != nil{
		//panic(err)
	//}

	//file, err := os.Create("static/index.html")
	//if err != nil{
		//log.Fatalf("Error creating index.html file: %v", err)
	//}
	//defer file.Close()

	//indexRendererr := pr.RenderIndex(file, posts)
	//if indexRendererr != nil{
		//log.Fatalf("Error rendering index.html file: %+v", err)
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
		Parts: []genai.Part{genai.Text("You generate .md string/blog post on a technology-related topics. Do not generate generic or broad content. Each post should focus on a single, interesting concept within a topic provided by the user, chosen from the list of topics user provide. Avoid explaining multiple concepts within a single blog post.\nFollow this template structure:\nTitle: A concise title (Do not use any special characters)\nA well-structured explanation of the chosen topic with example or code snippets if required.")},
	  }
	resp, err := model.GenerateContent(ctx, genai.Text("Topics: Javascript Interview Questions")) 
	if err != nil {
		log.Fatal(err)
	}

	printResponse(resp)
}
func printResponse(resp *genai.GenerateContentResponse) {
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				fmt.Println(part)
			}
		}
	}
	fmt.Println("---")
}
