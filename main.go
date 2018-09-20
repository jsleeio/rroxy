package main

import (
  "flag"
  "net/http"
  "net/url"
  "github.com/vulcand/oxy/forward"
  "github.com/vulcand/oxy/roundrobin"
  "github.com/vulcand/oxy/buffer"
  )

func main() {
  // Forwards incoming requests to whatever location URL points to, adds proper forwarding headers
  listen := flag.String("listen", ":8000", "TCP port to listen on for requests")
  retrypredicate := flag.String("retry-predicate", "IsNetworkError() && Attempts() < 2", "Retry predicate (see: go doc github.com/vulcand/oxy/buffer.Retry)")
  flag.Parse()

  fwd, _ := forward.New()
  lb, _ := roundrobin.New(fwd)

  for _,backend := range flag.Args() {
    url,err := url.Parse(backend)
    if err != nil {
      panic(err)
    }
    err = lb.UpsertServer(url)
    if err != nil {
      panic(err)
    }
  }

  buffer, err := buffer.New(lb, buffer.Retry(*retrypredicate))
  if err != nil {
    panic(err)
  }

  s := &http.Server{
    Addr:    *listen,
    Handler: buffer,
  }
  s.ListenAndServe()
}
