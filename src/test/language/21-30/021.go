package main

// type TagError struct {
// 	code int
// 	msg  string
// }

// func (e TagError) Error() string {
// 	return fmt.Sprint(e.code, "   ", e.msg)
// }

// func f1(i int) (int, error) {
// 	if i == 42 {
// 		return -1, errors.New("can't work with 42")
// 	}
// 	return i + 3, nil
// }

// func f2(i int) (int, error) {
// 	if i == 42 {
// 		return -1, TagError{404, "a error number"}
// 	}
// 	return i + 3, nil
// }

// func main() {
// 	x, err := f1(42)
// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Println("i = ", x)
// 	}

// 	x, err = f2(42)
// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Println("i = ", x)
// 	}

// 	x, err = f2(42)
// 	if ae, ok := err.(TagError); ok {
// 		fmt.Println(ae.code)
// 		fmt.Println(ae.msg)
// 	}
// }
