# Functions in Go

What I've learned in this lesson:

* It's very common for declaration of functions to be in camel case.
* Functions are declared with `func`.
* Return type is after the declaration of the function (show case at the bottom) and it is typed it requires a return.
* Import can be "merged" with (). Example `import ("fmt" "os")`
* `os.Exit(0)` Quits the program with success.
* `os.Exit(-1)` Quits the program with exit status, good for error handling.
* I've learned that it's possible to have a multi "return" function.

```
func introduction() int
```