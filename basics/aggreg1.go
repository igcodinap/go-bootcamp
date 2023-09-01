// package main

// // arrays, slices, maps, structs

// func main() {
// 	{
// 		// su largo es parte del tipo, es fijo, se asigna al compilar. O(n)
// 		// var nums [5]string
// 		// fmt.Println("---arrays---")

// 		// fmt.Printf("len(nums) : %T\n", len(nums))
// 		// fmt.Printf("nums array: %#v\n", nums)
// 		// fmt.Printf("len(nums) : %d\n", len(nums))

// 		// for i := 0; i < len(nums); i++ {
// 		// 	fmt.Printf("nums[%d]: %s\n", i, nums[i])
// 		// }

// 	}

// 	{
// 		// su largo no es parte del tipo, es dinamico. se asigna en el runtime su tamaño. O(n)
// 		// nums := []string{"uno", "dos", "tres"}
// 		// nums2 := []string{"cuatro", "cinco", "seis"}
// 		// fmt.Println("---slices---")
// 		// fmt.Printf("len(nums) : %v\n", len(nums))

// 		// fmt.Printf("nums slice: %#v\n", nums)

// 		// fmt.Printf("len(nums) : %d\n", len(nums))

// 		// for i := 0; i < len(nums); i++ {
// 		// 	fmt.Printf("nums[%d]: %s\n", i, nums[i])
// 		// }

// 		// range
// 		// for index, value := range nums {
// 		// 	fmt.Printf("nums[%d]: %s\n", index, value)
// 		// 	fmt.Println("nums2: ", nums2[index])
// 		// }

// 		// para que dos slices sean iguales, deben tener el mismo largo y los mismos valores

// 		// won't work: the slice is nil.
// 		// fmt.Printf("nums[0]: %d\n", nums[0])
// 		// fmt.Printf("nums[1]: %d\n", nums[1])
// 	}
// 	{

// 		// estructura llave-valor, largo dinamico e informacion esta indexada. funciona como una hash table. O(1)
// 		// fmt.Println("---maps---")

// 		// jugadores := map[string]int{
// 		// 	"Alexis":  7,
// 		// 	"Claudio": 1,
// 		// 	"Arturo":  8,
// 		// }

// 		// for nombre, dorsal := range jugadores {
// 		// 	fmt.Printf("Name: %s, Dorsal: %d\n", nombre, dorsal)
// 		// }

// 		// fmt.Println("jugadores[\"Alexis\"]", jugadores["Alexis"])
// 		// delete(jugadores, "Alexis")
// 		// jugadores["Edu"] = 11
// 		// clear(jugadores)
// 		// fmt.Println(jugadores)

// 	}

// }
