# Errors

NO language construct in Go to represent errors

```
type error interface {
    Error() string
}
```

Typical error pattern

```
func Sqrt(f float64) (float64, error) {
    if f < 0 {
        return 0, errors.New("math: square root of negative number")
    }
    // implementation
}
```

```
// errorString is a trivial implementation of error.
type errorString struct {
    s string
}

func (e *errorString) Error() string {
    return e.s
}

// New returns an error that formats as the given text.
func New(text string) error {
    return &errorString{text}
}
```

```
package main

import (
	"errors"
	"fmt"
)

func main() {
	value, err := Sqrt(-2.234)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Sqrt: ", value)

}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("math: square root of negative number")
	}
	
	// implmentation
	return 1.111, nil
}
```

### Custom Error Types

```
package main

import (
	"fmt"
)

type NegativeSqrtError float64

func (f NegativeSqrtError) Error() string {
	return fmt.Sprintf("math: square root of negative number %g", float64(f))
}

func main() {
	value, err := Sqrt(-2.234)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Sqrt: ", value)
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, NegativeSqrtError(f)
	}
	return 1.111, nil
}

```

```
func main() {
	value, err := Sqrt(-2.234)

	if err != nil {
		switch err.(type) {
		case NegativeSqrtError:
			fmt.Println("NegativeSqrtError:", err)
		default:
			fmt.Println("Error:", err)
		}
		return
	}

	fmt.Println("Sqrt: ", value)
}
```

```
package main

import (
	"fmt"
)

type NegativeSqrtError struct {
	value float64
}

func (f NegativeSqrtError) Error() string {
	return fmt.Sprintf("math: square root of negative number %g", f.value)
}

func main() {
	value, err := Sqrt(-2.234)

	if err != nil {
		switch err.(type) {
		case NegativeSqrtError:
			fmt.Println("NegativeSqrtError:", err)
			fmt.Println("NegativeSqrtError Value:", err.(NegativeSqrtError).value)
		default:
			fmt.Println("Error:", err)
		}
		return
	}

	fmt.Println("Sqrt: ", value)
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, NegativeSqrtError{f}
	}
	return 1.111, nil
}
```

- It is idiomatic in Go to postfix the name of a custom error type with the word Error
- Custom errors are part of your public contract for the package

```
  switch e := err.(type) {
     case *json.UnmarshalTypeError:
        log.Printf("UnmarshalTypeError: Value[%s] Type[%v]\n", e.Value, e.Type)
     case *json.InvalidUnmarshalError:
         log.Printf("InvalidUnmarshalError: Type[%v]\n", e.Type)
     default:
         log.Println(err)
  }

```


### Best Practices

- Don't use Sentinel errors/Constant (`io.EOF`)
- Never inspect output of error.Error()
- Avoid error types so you don't make them as part or your public API
- Assert errors for behavior, not type


```
package net

type Error interface {
    error
    Timeout() bool   // Is the error a timeout?
    Temporary() bool // Is the error temporary?
}
```


```
if nerr, ok := err.(net.Error); ok && nerr.Temporary() {
    time.Sleep(1e9)
    continue
}

if err != nil {
    log.Fatal(err)
}
```


## Logging
Wrapping errors using `github.com/pkg/errors`
```
package main

import (
	"fmt"
	"github.com/pkg/errors"
)

func main() {
	val, err := DoThing1()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Success: " + val)
}

func DoThing1() (string, error) {
	value, err := DoThing2()

	if err != nil {
		return "", errors.Wrap(err, "thing2 failed")
	}

	return value, nil
}

func DoThing2() (string, error) {
	value, err := DoThing3()

	if err != nil {
		return "", errors.Wrap(err, "thing3 failed")
	}

	return value, nil
}

func DoThing3() (string, error) {
	return "", errors.New("bad things happened")
}

```

`errors.stackeTracer` interface is not exported
```
type stackTracer interface {
	StackTrace() errors.StackTrace
}

func main() {
	val, err := DoThing1()

	if err != nil {
		fmt.Println(err)

		if _, ok := err.(stackTracer); ok {
			for _, stackTrace := range err.(stackTracer).StackTrace() {
				fmt.Println(stackTrace)
			}

		}
		return
	}

	fmt.Println("Success: " + val)
}
```

## Panic
Unexpected and unrecoverable scenarios. Starts unwinding the stack, running deferred functions along the way.
If that unwinding reaches the top of the goroutine's stack, the program dies.

```
var user = os.Getenv("USER")

func init() {
    if user == "" {
        panic("no value for $USER")
    }
}
```


#### Recover
```
package main

import "fmt"

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("caught ya ", err)
		}
	}()
	createPanic()
}

func createPanic() {
	panic("this function sucks")
}
```
