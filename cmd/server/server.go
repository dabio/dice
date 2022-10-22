package main

import (
	_ "embed"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"text/template"

	"dab.io/dice"
)

//go:embed index.html
var index string

type Roll struct {
	Result   []string
	Count    int
	MaxCount int
}

func main() {
	http.HandleFunc("/", redirectHttps(roll()))

	port := env("PORT", "3000")
	address := env("ADDRESS", "0.0.0.0")
	log.Printf("Listening on port %s:%s", address, port)
	log.Fatal(http.ListenAndServe(address+":"+port, nil))
}

func roll() http.HandlerFunc {
	tpl := template.Must(template.New("index").Parse(index))

	return func(w http.ResponseWriter, r *http.Request) {
		count, err := strconv.Atoi(r.FormValue("count"))
		if err != nil || count == 0 {
			count = dice.DefaultCount
		}
		roll := Roll{
			Result:   dice.Roll(count),
			Count:    count,
			MaxCount: dice.MaxCount,
		}

		w.Header().Set("Content-Type", "text/html; charset=utf8")
		if err := tpl.Execute(w, roll); err != nil {
			log.Fatal(err)
		}
	}
}

func redirectHttps(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		target := url.URL{
			Scheme:   "https",
			Host:     r.Host,
			Path:     r.URL.Path,
			RawQuery: r.URL.RawQuery,
		}

		if r.Header.Get("X-Forwarded-Proto") == "http" {
			http.Redirect(w, r, target.String(), http.StatusTemporaryRedirect)
		} else {
			next(w, r)
		}
	}
}

func env(key, defaultValue string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}

	return defaultValue
}
