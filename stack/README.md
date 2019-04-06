# Stack

### Use for
* First-In Last-Out operations
* Depth-First Search (DFS) tree

**Time Complexity**

| Operation    | Time Complexity |
|--------------|-----------------|
| Push         |          `O(1)` |
| Pop          |          `O(1)` |
| Top          |          `O(1)` |


### Example code

```go
import (
    "fmt"
    "github.com/agungdwiprasetyo/go-stl/stack"
)

func main() {
    s := stack.New() // init new stack

    s.Push(100) // push value to stack container

    val := s.Top() // get top element from stack

    val, ok := s.Pop() // pop value in stack

    isEmpty := s.IsEmpty() // check stack is empty or not

    fmt.Println(s) // TODO: print pretty stack
}
```