package main

import (
	"fmt"
	"io/fs"
	"net/http"
	"time"
)

var shortlinks = map[string]string{
	"/github":  "https://github.com/aamtt",
	"/twitter": "https://twitter.com/0xA5A5",
	"/x":       "https://x.com/0xA5A5",
	"/discord": "https://discord.gg/R3QtA68Cbf",
}

func root(w http.ResponseWriter, r *http.Request) {
	data := page()

	if wantsPlain(w, r) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Header().Set("Cache-Control", "public, max-age=0, s-maxage=300, stale-while-revalidate=86400")
		plainTemplate.ExecuteTemplate(w, "plain.txt", data)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Cache-Control", cacheFor(data.settled(), "public, max-age=0, s-maxage=300, stale-while-revalidate=86400"))
	pageTemplate.ExecuteTemplate(w, "index.html", data)
}
func shortlink(w http.ResponseWriter, r *http.Request) {
	target, ok := shortlinks[r.URL.Path]
	if !ok {
		notFound(w, r)
		return
	}

	w.Header().Set("Cache-Control", "public, max-age=3600, s-maxage=86400")
	http.Redirect(w, r, target, http.StatusFound)
}

type research struct {
	Title      string
	Status     string
	InProgress bool
}

var papers = []research{
	{Title: "E-Graphs (title pending)", Status: "in progress, no preprint", InProgress: true},
	{Title: "Optimization (title pending)", Status: "in progress, no preprint", InProgress: true},
}

type academicView struct {
	Papers       []research
	CanonicalURL string
	Portrait     string
}

func academic(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "public, max-age=0, s-maxage=3600")

	if wantsPlain(w, r) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		fmt.Fprint(w, "  M. BAKER\n  Optimization & Graph Theory\n\n")
		fmt.Fprint(w, "  I work on optimization and graph theory, and on the\n  intersection between them.\n\n  RESEARCH\n")
		for n, item := range papers {
			fmt.Fprintf(w, "    [%d] %s\n        %s\n", n+1, item.Title, item.Status)
		}
		fmt.Fprintf(w, "\n  m [at] thew [dot] sh / https://%s/\n", canonicalHost)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	academicTemplate.ExecuteTemplate(w, "matthew.html", academicView{
		Papers:       papers,
		CanonicalURL: "https://" + canonicalHost + "/~matthew",
		Portrait:     portraitFrame(),
	})
}
func receipt(w http.ResponseWriter, r *http.Request) {
	body, err := archiveFS.ReadFile("archive/index.html")
	if err != nil {
		notFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Cache-Control", "public, max-age=0, s-maxage=3600, stale-while-revalidate=86400")
	w.Write(body)
}
func exported() http.Handler {
	files, err := fs.Sub(archiveFS, "archive")
	if err != nil {
		panic(err)
	}
	handler := http.FileServerFS(files)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "public, max-age=31536000, s-maxage=31536000, immutable")
		http.StripPrefix("/receipt/", handler).ServeHTTP(w, r)
	})
}

type guestbookView struct {
	view
	Origin    string
	Namespace string
	Signers   []signer
}

func guestbook(w http.ResponseWriter, r *http.Request) {
	data := page()
	data.CanonicalURL = "https://" + canonicalHost + "/guestbook"
	data.Home = true

	list := signers()

	if wantsPlain(w, r) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Header().Set("Cache-Control", "public, max-age=0, s-maxage=60, stale-while-revalidate=86400")
		if len(list) == 0 {
			fmt.Fprintf(w, "    Nobody has signed yet.\n")
			return
		}

		for _, entry := range list {
			fmt.Fprintf(w, "    %s  %s\n", entry.Signed(), entry.Display())
		}
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Cache-Control", cacheFor(data.settled(), "public, max-age=0, s-maxage=60, stale-while-revalidate=86400"))
	guestbookTemplate.ExecuteTemplate(w, "guestbook.html", guestbookView{
		view:      data,
		Origin:    "https://" + canonicalHost,
		Namespace: signNamespace,
		Signers:   list,
	})
}

type errorView struct {
	Home      bool
	Path      string
	Name      string
	Year      int
	DiscordID string
	Github    githubPreview
	Presence  presencePreview
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "public, max-age=0, s-maxage=60")

	if wantsPlain(w, r) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "  Not Found\n\n  There is nothing at %s\n\n  https://%s/\n", trimPath(r.URL.Path), canonicalHost)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusNotFound)

	errorTemplate.ExecuteTemplate(w, "404.html", errorView{
		Path:      trimPath(r.URL.Path),
		Name:      displayName,
		Year:      time.Now().In(zone()).Year(),
		DiscordID: discordID,
		Github:    githubMemo.get(),
		Presence:  presenceMemo.get(),
	})
}
