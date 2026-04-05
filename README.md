# `vibefs-template`

> Minimal Chi-based static app template — drop your frontend in `./public` and serve it under a configurable path. Automatically redirects `/` to your app root.

---

## Project Layout

```plaintext
.
├── README.md
├── go.mod
├── go.sum
├── main.go          # entrypoint, config-aware server
├── public
│   └── index.html   # your frontend app
└── src
    └── config
        ├── conf.go  # defines Config struct
        └── load.go  # env loader
```

---

## How it works

1. `main.go` loads configuration from environment variables (like `APP_PATH`, `APP_PORT`).
2. Serves `./public` folder at the mount point (`APP_PATH`).
3. Automatically redirects `/` → `APP_PATH` (if it’s not `/`).

---

## Quickstart

1. Set environment variables:

```bash
export APP_PATH="/myprefix/"
export APP_PORT="3000"
```

1. Run the server:

```bash
go run main.go
```

1. Visit:

* `http://localhost:3000/` → redirects to `/myprefix/`
* `http://localhost:3000/myprefix/` → serves `./public/index.html`
* `http://localhost:3000/myprefix/js/main.js` → serves other static assets

---

## Config Example (`src/config/conf.go`)

```go
package config

type Config struct {
 APP_PATH string // required, must end with "/"
 APP_PORT string // optional, default "3000"
}
```

---

## Env Loader (`src/config/load.go`)

* Reads `APP_PATH` and `APP_PORT` from environment variables.
* Fallback to defaults if optional values are missing.

---

## main.go Example

```go
package main

import (
 "const-aka/src/config"
 "fmt"
 "log"
 "net/http"

 "github.com/go-chi/chi/v5"
)

func main() {
 var cfg config.Config
 if err := config.Load(&cfg); err != nil {
  log.Fatal(err)
 }

 r := chi.NewRouter()

 // redirect / → APP_PATH
 if cfg.APP_PATH != "/" {
  r.Get("/", func(w http.ResponseWriter, r *http.Request) {
   http.Redirect(w, r, cfg.APP_PATH, http.StatusFound)
  })
 }

 // serve static files under APP_PATH
 fs := http.FileServer(http.Dir("./public"))
 r.Handle(cfg.APP_PATH+"*", http.StripPrefix(cfg.APP_PATH, fs))

 fmt.Printf("Server running at http://localhost:%s%s\n", cfg.APP_PORT, cfg.APP_PATH)
 log.Fatal(http.ListenAndServe(":"+cfg.APP_PORT, r))
}
```

---

## Notes / Tips

* `APP_PATH` **must end with `/`** for proper routing.
* The template is ready for multiple frontends if you copy `r.Handle(...)` blocks for different folders.
* Perfect for **rapid frontend prototyping** without touching boilerplate.
