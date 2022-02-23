package main

import (
	"flag"

	"github.com/kyokomi/emoji/v2"
)

func main() {
	emojiKeyword := flag.String("e", "Hello :world_map:", "emoji name")
	flag.Parse()

	emoji.Print(*emojiKeyword)
}
