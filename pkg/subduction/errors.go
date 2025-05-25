package subduction

import (
	"fmt"
	"runtime"
)

type TechtonicError struct {
	Err     error
	Message string
	Stack   []string
}

func Wrap(err error, msg string) error {
	const maxStackDepth = 50
	stack := make([]string, 0, maxStackDepth)
	for i := 1; i < maxStackDepth; i++ {
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		fn := runtime.FuncForPC(pc).Name()
		stack = append(stack, fmt.Sprintf("%s:%d %s", file, line, fn))
	}
	for i := 1; ; i++ {
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		fn := runtime.FuncForPC(pc).Name()
		stack = append(stack, fmt.Sprintf("%s:%d %s", file, line, fn))
	}

	return &TechtonicError{
		Err:     err,
		Message: msg,
		Stack:   stack,
	}
}

func (e *TechtonicError) Error() string {
	return fmt.Sprintf("%s: %v", e.Message, e.Err)
}

func (e *TechtonicError) Unwrap() error {
	return e.Err
}
