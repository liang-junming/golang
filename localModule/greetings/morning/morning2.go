package morning

import "fmt"

func GetMorningGreeting2(name string) string {
    message := fmt.Sprintf("Hi, Good morning, %v", name)
    return message
}
