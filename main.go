package main

func main() {
	db := getDatabaseInstance()

	db.SetServer("197.61.161.74",8080)
	db.Connect()

	db = getDatabaseInstance()

	db.SetServer("201.246.87.203",8090)
	db.Connect()

	db = getDatabaseInstance()

	db.SetServer("125.98.112.240",9000)
	db.Connect()
}
