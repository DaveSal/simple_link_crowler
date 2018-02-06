package main

import (
  "flag"
	"fmt"
	"os"
  "crypto/tls"
  "github.com/jackdanger/collectlinks"
  "web_crowler/linkFixer"
  "net/http"
)

var visited = make(map[string]bool)

func main() {
  flag.Parse()
  args := flag.Args()

  if len(args) < 1 {
    fmt.Println("Вы не ввели начальную страницу.")
    os.Exit(1)
  }

  queue := make(chan string)
  go func() { queue <- args[0] }()

  for uri := range queue {
    retrieve(uri, queue)
  }
}

func retrieve(uri string, queue chan string) {
  fmt.Println("Сканирую ", uri)
  visited[uri] = true

  tlsConfig := &tls.Config {
    InsecureSkipVerify: true,
  }

  transport := &http.Transport {
    TLSClientConfig: tlsConfig,
  }

  client := http.Client {Transport: transport}
  response, error := client.Get(uri)

  if error != nil {
    fmt.Println("Ошибка запроса: ", error)
    return
  }
  defer response.Body.Close()

  links := collectlinks.All(response.Body)
  for _, link := range(links) {
    fixedLink := linkFixer.Fix(link, uri)

    if uri != "" {
      if !visited[fixedLink] {
        go func() { queue <- fixedLink }()
      }
    }
  }
}
