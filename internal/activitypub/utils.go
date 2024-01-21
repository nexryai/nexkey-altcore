package activitypub

import (
	"fmt"
	"regexp"
)

// "acct:alice@my-example.com" 的な文字列からユーザー名を取り出す
func GetUserNameFromSubject(subject string) (string, error) {
	re := regexp.MustCompile(`acct:([^@]+)@`)

	matches := re.FindStringSubmatch(subject)

	if len(matches) < 2 {
		return "", fmt.Errorf("invalid string")
	}

	return matches[1], nil
}
