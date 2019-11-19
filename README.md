# Gofake
Gofake is a minimal random value generator. It can be used to generate random
- integers
- names
- emails
- words
- sentences
- paragraphs

## Usage
```
package main

import (
	"fmt"

	"github.com/twharmon/gofake"
)

func main() {
	fmt.Println(gofake.Int())
	fmt.Println(gofake.Int64())
	fmt.Println(gofake.Name())
	fmt.Println(gofake.Email())
	fmt.Println(gofake.Word())
	fmt.Println(gofake.Sentence())
	fmt.Println(gofake.Paragraph())
}

```