// package main

// import "error"

// func sum(a int, b int) int {
// 	return a + b
// }

// // function with multiple return values
// func swap(a int, b int) (int, int) {
// 	return b, a
// }

// // function with err return value
// func divide(a int, b int) (int, error) {
// 	if b == 0 {
// 		return 0, errors.New("No se puede dividir por cero")
// 	}
// 	return a / b, nil
// }

// value, err := divide(10, 0)
// if err != nil {
// 	fmt.Println(err)
// 	return
// }
