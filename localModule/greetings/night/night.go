package night

import "fmt"

func GetNightGreeting(name string) string {
    message := fmt.Sprintf("Good night, %v", name)
    return message
}
