package helpers

import (
	"strings"

	"gorm.io/gorm"
)

func TranslateErrorMessage(err error) map[string]string {
	errorsMap := make(map[string]string)

	// Handle error dari GORM untuk duplicate entry
	if err != nil {
		// Cek jika error mengandung "Duplicate entry" (duplikasi data di database)
		if strings.Contains(err.Error(), "Duplicate entry") {
			if strings.Contains(err.Error(), "username") {
				errorsMap["Username"] = "Username already exists" // Pesan error jika username sudah ada
			}
			if strings.Contains(err.Error(), "email") {
				errorsMap["Email"] = "Email already exists" // Pesan error jika email sudah ada
			}
		} else if err == gorm.ErrRecordNotFound {
			// Jika data yang dicari tidak ditemukan di database
			errorsMap["Error"] = "Record not found"
		}
	}

	// Mengembalikan map yang berisi pesan error
	return errorsMap
}

func IsDupliateEntryError(err error) bool {
	return strings.Contains(err.Error(), "Duplicate entry")
}
