package hi

import "fmt"

func SayHi(firstName string, lastName string) string {
	return fmt.Sprintf("Hello %s %s", firstName, lastName)
}
