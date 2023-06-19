# MinMaxStack

MinMaxStack is a Go library that provides a stack implementation with real-time tracking of the minimum and maximum values. It offers an efficient and convenient way to keep track of the minimum and maximum elements as new values are pushed or popped from the stack.

## Features

- Push and pop values onto the stack
- Retrieve the minimum and maximum values in constant time
- Efficiently track the minimum and maximum values as the stack changes
- Simple and intuitive API for stack manipulation

## Installation

Use `go get` to install the library:

```shell
go get -u github.com/akl773/MinMaxStack
```

## Usage
Here's an example of how to use the MinMaxStack library:

```shell
import "github.com/akl773/MinMaxStack"

func main() {
	stack := minmaxstack.NewMinMaxStack(5)
	stack.Push(10)
	min := stack.GetMin()
	max := stack.GetMax()
}
```

## Contributing
Contributions are welcome! If you find any issues or have suggestions for improvements, please open an issue or submit a pull request on GitHub.
