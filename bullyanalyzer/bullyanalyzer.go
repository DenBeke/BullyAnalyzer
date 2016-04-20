package main

import (
	"github.com/DenBeke/BullyAnalyzer"
	"github.com/ant0ine/go-json-rest/rest"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {

	lexiconfile := "../profanity_dutch.txt"

	if len(os.Args) > 1 {
		lexiconfile = os.Args[1]
	}

	// BullyAnalyzer stuff
	a, err := bullyanalyzer.New(lexiconfile)
	if err != nil {
		log.Fatal(err)
	}

	// REST api stuff
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(
		rest.Get("/post/#post", func(w rest.ResponseWriter, req *rest.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			post := strings.Join(strings.Split(req.PathParam("post"), "%20"), " ")
			result := a.AnalyzePost(post)
			w.WriteJson(&result)
		}),
		rest.Get("/post/#post/", func(w rest.ResponseWriter, req *rest.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			post := strings.Join(strings.Split(req.PathParam("post"), "%20"), " ")
			result := a.AnalyzePost(post)
			w.WriteJson(&result)
		}),
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":5000", api.MakeHandler()))
}
