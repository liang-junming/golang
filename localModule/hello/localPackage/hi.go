package hi

import "fmt"

func GetHi(name string) string {
    message := fmt.Sprintf("Hi, %v", name)
    return message
}
