package utils

import (
	"strings"
)

func Nyaize(text string) string {
	// ja-JP
	text = strings.ReplaceAll(text, "な", "にゃ")
	text = strings.ReplaceAll(text, "ナ", "ニャ")
	text = strings.ReplaceAll(text, "ﾅ", "ﾆｬ")

	// en-USは敢えてしない

	return text
}
