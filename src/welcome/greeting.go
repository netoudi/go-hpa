package welcome

import (
  "fmt"
  "math"
)

func Greeting(message string) string {
  x := 0.0001

  for i := 0; i <= 1_000_000_000; i++ {
    x += math.Sqrt(x)
  }

  return fmt.Sprintf("<b>%s</b>", message)
}
