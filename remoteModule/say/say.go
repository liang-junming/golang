package say

import "fmt"

func Say(word string) string {
    message := fmt.Sprintf("say %v", word)
    return message
}
