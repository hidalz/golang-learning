package main

// 1. Structs in Go

// type messageToSend struct {
// 	phoneNumber int
// 	message string
// }

// 2. Nested structs in Go
// type messageToSend struct {
// 	message   string
// 	sender    user
// 	recipient user
// }

// type user struct {
// 	name   string
// 	number int
// }

// func canSendMessage(mToSend messageToSend) bool {
// 	// ?
// 	if mToSend.sender.number == 0 || mToSend.recipient.number == 0 {
// 		return false
// 	}

// 	if mToSend.sender.name == "" || mToSend.recipient.name == "" {
// 		return false
// 	}
// 	return true
// }

// 5. Embedded structs

// type sender struct {
// 	rateLimit int
// 	user
// }

// type user struct {
// 	name   string
// 	number int
// }

// 6. Struct methods in Go
// type authenticationInfo struct {
// 	username string
// 	password string
// }

// // create the method below
// func (a authenticationInfo) getBasicAuth() string {
// 	myString := fmt.Sprintf("Authorization: Basic %v:%v", a.username, a.password)

// 	return myString
// }

// func main() {
//     // Create an instance of the authenticationInfo struct
//     authInfo := authenticationInfo{
// 		username: "myUsername",
//         password: "myPassword",
//     }

//     // Use the getBasicAuth method
//     authHeader := authInfo.getBasicAuth()
//     fmt.Println(authHeader)
// }
// 7. Memory layout
type contact struct {
	sendingLimit int32
	userID       string
	age          int32
}

type perms struct {
	canSend         bool
	canReceive      bool
	permissionLevel int
	canManage       bool
}

// 8. Empty struct

// 9. Empty struct