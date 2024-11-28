# printgo

`printgo` is a simple Go package that provides enhanced printing capabilities.

## Installation

```bash
go get github.com/yyle88/printgo
```

## Example Usage

```go
package main

import (
	"fmt"
	"github.com/yyle88/printgo"
)

func main() {
	ptx := printgo.NewPTX()

	ptx.Println("Hello, world!")

	res := ptx.String()
	
	fmt.Println(res)
}
```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

## Support

Welcome to contribute to this project by submitting pull requests or reporting issues.

If you find this package helpful, give it a star on GitHub!

**Thank you for your support!**

**Happy Coding with `printgo`!** 🎉
