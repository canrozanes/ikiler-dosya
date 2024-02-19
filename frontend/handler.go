package frontend

import (
	"embed"
	"errors"
	"html/template"
	"io"
	"io/fs"
	"mime"
	"net/http"
	"path"
	"path/filepath"
	"strings"
)

var ErrDir = errors.New("path is dir")

//go:embed dist/*
var dist embed.FS

func tryRead(fs embed.FS, prefix, requestedPath string, w http.ResponseWriter) error {
	f, err := fs.Open(path.Join(prefix, requestedPath))
	if err != nil {
		return err
	}
	defer f.Close()

	stat, _ := f.Stat()
	if stat.IsDir() {
		return ErrDir
	}

	contentType := mime.TypeByExtension(filepath.Ext(requestedPath))
	w.Header().Set("Content-Type", contentType)
	_, err = io.Copy(w, f)
	return err
}

func tryReadHtml(efs embed.FS, prefix, requestedPath string, w http.ResponseWriter, r *http.Request) error {
	if requestedPath != "index.html" {
		return errors.New("path is not index.html")
	}

	indexHtmlBits, err := fs.ReadFile(efs, path.Join(prefix, requestedPath))

	if err != nil {
		return err
	}

	tpl, err := template.New("index.html").Parse(string(indexHtmlBits))
	if err != nil {
		return err
	}

	err = tpl.Execute(w, nil)

	return err
}

func CreateSpaHandler(serverOrigin, auth0Domain string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		setContentSecurityPolicy(w, serverOrigin, auth0Domain)
		// dist/assets
		err := tryRead(dist, "dist", r.URL.Path, w)
		if err == nil {
			return
		}

		// dist/index.html
		err = tryReadHtml(dist, "dist", "index.html", w, r)
		if err != nil {
			panic(err)
		}
	}
}

// Some assets in the frontend are hosted in other domains.
// When the server is built, the browser will block these requests. To fix this, we need to set the Content-Security-Policy
// header to allow these requests.
func setContentSecurityPolicy(w http.ResponseWriter, serverOrigin, auth0Domain string) {
	var allowedDomains = []string{serverOrigin, auth0Domain, "self", "https://cdn.auth0.com", "https://s.gravatar.com", "https://fonts.googleapis.com", "https://fonts.gstatic.com"}
	allowedDomainsStr := strings.Join(allowedDomains, " ")
	w.Header().Set("Content-Security-Policy", "default-src "+allowedDomainsStr+"; style-src "+allowedDomainsStr+";")
}
