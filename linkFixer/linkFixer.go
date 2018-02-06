package linkFixer

import (
  "net/url"
)

func Fix(link, baseLink string) (string) {
  uri, error := url.Parse(link)

  if error != nil {
    return ""
  }

  baseUrl, error := url.Parse(baseLink)

  if error != nil {
    return ""
  }

  uri = baseUrl.ResolveReference(uri)
  return uri.String()
}
