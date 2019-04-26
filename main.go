package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"unicode"
)

type PageVariables struct {
	Sentiment string
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Listening on port %s", port)
	http.HandleFunc("/", AnalyzePage)
	http.HandleFunc("/analyze", AnalyzePage)
	http.HandleFunc("/results", ResultsPage)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func AnalyzePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, `
  <!DOCTYPE html>
  <html>
  <head>
    <title>Sentimentalyzer</title>
    <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css" integrity="sha384-ggOyR0iXCbMQv3Xipma34MD+dH/1fQ784/j6cY/iJTQUOhcWr7x9JvoRxT2MZw1T" crossorigin="anonymous">
  </head>
  <body>
    <div class="container">
      <p>Welcome to Sentimentalyzer, the world's simplest sentiment analyzer!</p>
      <form action="/results" method="post">
        <div class="form-group">
          <label for="feelings">How are you feeling?</label>
          <textarea class="form-control" rows="5" name="feelings" id="feelings"></textarea>
        </div>
        <button type="submit" class="btn btn-primary">Submit</button>
      </form>
    </div>
  </body>
  </html>
  `)
}

func ResultsPage(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	ResultsPageVars := PageVariables{
		Sentiment: detectSentiment(r.Form.Get("feelings")),
	}

	err = template.Must(template.New("T").Parse(`
  <!DOCTYPE html>
  <html>
  <head>
    <title>Sentimentalyzer</title>
    <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css" integrity="sha384-ggOyR0iXCbMQv3Xipma34MD+dH/1fQ784/j6cY/iJTQUOhcWr7x9JvoRxT2MZw1T" crossorigin="anonymous">
  </head>
  <body>
    <div class="results container">
      <h1>Results</h1>
      <p>You are feeling: <b>{{.Sentiment}}</b></p>
    </div>
  </body>
  </html>
  `)).Execute(w, ResultsPageVars)
	if err != nil {
		log.Print("Error generating page: ", err)
	}
}

func detectSentiment(feelings string) string {
	cu := countUpperCaseCharacters(feelings)
	cTot := countLetters(feelings)
	if float64(cu)/float64(cTot) > 0.5 {
		return "Angry"
	} else {
		return "Content"
	}
}

func countUpperCaseCharacters(str string) int {
	count := 0
	for _, c := range str {
		if unicode.IsUpper(c) {
			count++
		}
	}
	return count
}

func countLetters(str string) int {
	count := 0
	for _, c := range str {
		if unicode.IsLetter(c) {
			count++
		}
	}
	return count
}
