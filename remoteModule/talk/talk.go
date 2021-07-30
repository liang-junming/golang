package talk

import "fmt"

func Talk(word string) string {
    message := fmt.Sprintf("talk %v", word)
    return message
}
