package retriever

import (
  "crypto/tls"
  "fmt"
  "github.com/jackdanger/collectlinks"
  "web_crowler/linkFixer"
  "net/http"
)

func Retrieve(uri string, queue chan string, scannedLink map[string]bool) {
  fmt.Println("Сканирую ", uri)
  scannedLink[uri] = true

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
      if !scannedLink[fixedLink] {
        go func() { queue <- fixedLink }()
      }
    }
  }
}
