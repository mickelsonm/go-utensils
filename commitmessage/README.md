# Commit Message

A helper which get you a message you could potentially use in your git commit message.

# How do I use?

```go
package main

import(
  "fmt"
  "log"

  "github.com/curt-labs/go-utensils/commitmessage"
)

func main() {
	msg, err := commitmessage.NewMessage()
	if err != nil {
		log.Fatal(err)
	}
	out, jserr := msg.ToJSON()
	if jserr != nil {
		log.Fatal(jserr)
	}
	fmt.Println(string(out))
}

```

# Where do I go from here?

Find more sources and add them! Happy open-sourcing! :thumbsup:
