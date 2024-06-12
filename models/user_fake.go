package models

type User struct {
	Username  string
	FirstName string
	LastName  string
}

// ประกาศตัวแปร userList เป็นตัวแปรระดับแพ็กเกจ
var userList = []User{
	{
		Username:  "b2",
		FirstName: "Ball",
		LastName:  "K",
	},
	{
		Username:  "b2",
		FirstName: "Bas",
		LastName:  "K",
	},
	{
		Username:  "asxxxxxdas",
		FirstName: "Bas",
		LastName:  "K",
	},
}

// ฟังก์ชันเพื่อดึงข้อมูลผู้ใช้ทั้งหมด
func GetAllUsers() []User {
	return userList
}

// ฟังก์ชันเพื่อสร้างผู้ใช้ใหม่
func CreateUser(user User) User {
	userList = append(userList, user)
	return user
}
