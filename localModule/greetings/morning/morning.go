package morning

import "fmt"

func GetMorningGreeting(name string) string {
    message := fmt.Sprintf("Good morning, %v", name)
    return message
}
