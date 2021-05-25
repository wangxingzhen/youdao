package data

func InitData() {
	_, err := GetDB()
	if err != nil {
		panic(err.Error())
	}
}
