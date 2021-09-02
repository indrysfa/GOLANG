package database

var connection string

//  initialisasi pertama kali
// dibuat ketika GetDatabase() dijalankan
func init() {
	connection = "MySQL"
}

func GetDatabase() string {
	return connection
}
