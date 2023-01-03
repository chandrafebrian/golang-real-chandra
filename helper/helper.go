// untuk bisa di import nama package nya harus sesuai dengan nama file nya
package helper

import "strings"

// fungsi dengan akses public bisa di export kemana saja karena dengan huruf besar nama nya,
func FungsiValidasi(username string, lastname string, email string, userTicket uint, remainingTickets uint) (bool, bool, bool) {
	// is adalah untuk boolean variable dan memmbuat validasi bernilai true
	isValidateUser := len(username) >= 2 && len(lastname) >= 2
	isValidateEmail := strings.Contains(email, "@")
	isValidTicket := userTicket > 0 && userTicket <= remainingTickets

	return isValidateUser, isValidateEmail, isValidTicket

}
