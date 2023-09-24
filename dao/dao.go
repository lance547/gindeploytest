package dao

// date access object用于操作数据库
var Datebase = make(map[string]string)

func Adddate(username, password string) {
	Datebase[username] = password
}

func Selectusername(username string) bool {
	if Datebase[username] == "" {
		return false
	}
	return true
}

func Selectpasswordfromusername(username string) string {
	return Datebase[username]
}
