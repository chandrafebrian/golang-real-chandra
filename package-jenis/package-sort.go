package main

// package sort adalah package yang berisikan utilitas untuk proses pengaurutan
// agar data kita bisa di urutkan , kita harus mengimplentasikan kontrak di interface sort.interface

// type User struct {
// 	Name string
// 	Age  int
// }

// // untuk alias kontrak
// type Userslice []User

// func (value Userslice) Len() int {
// 	return len(value)
// }

// func (value Userslice) Less(i, j int) bool {
// 	return value[i].Age < value[j].Age
// }

// func (value Userslice) Swap(i, j int) {
// 	value[i], value[j] = value[j], value[i]
// }

// func main() {

// 	users := []User{
// 		{"chandra", 33},
// 		{"budu", 30},
// 		{"febrian", 40},
// 	}
// 	sort.Sort(Userslice(users))

// 	fmt.Println(users)

// }
