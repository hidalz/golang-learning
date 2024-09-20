package main

import (
	"fmt"
	"testing"
)

// 1. Structs in Go

// func getMessageText(m messageToSend) string {
// 	return fmt.Sprintf("Sending message: '%s' to: %v", m.message, m.phoneNumber)
// }

// func Test(t *testing.T) {
// 	type testCase struct {
// 		phoneNumber int
// 		message     string
// 		expected    string
// 	}
// 	tests := []testCase{
// 		{148255510981, "Thanks for signing up", "Sending message: 'Thanks for signing up' to: 148255510981"},
// 		{148255510982, "Love to have you aboard!", "Sending message: 'Love to have you aboard!' to: 148255510982"},
// 	}
// 	if withSubmit {
// 		tests = append(tests, []testCase{
// 			{148255510983, "We're so excited to have you", "Sending message: 'We're so excited to have you' to: 148255510983"},
// 			{148255510984, "", "Sending message: '' to: 148255510984"},
// 			{148255510985, "Hello, World!", "Sending message: 'Hello, World!' to: 148255510985"},
// 		}...)
// 	}

// 	passCount := 0
// 	failCount := 0

// 	for _, test := range tests {
// 		output := getMessageText(messageToSend{
// 			phoneNumber: test.phoneNumber,
// 			message:     test.message,
// 		})
// 		if output != test.expected {
// 			failCount++
// 			t.Errorf(`---------------------------------
// Inputs:     (%v, %v)
// Expecting:  %v
// Actual:     %v
// Fail`, test.phoneNumber, test.message, test.expected, output)
// 		} else {
// 			passCount++
// 			fmt.Printf(`---------------------------------
// Inputs:     (%v, %v)
// Expecting:  %v
// Actual:     %v
// Pass
// `, test.phoneNumber, test.message, test.expected, output)
// 		}
// 	}

// 	fmt.Println("---------------------------------")
// 	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
// }

// // 2. Nested structs in Go
// func Test(t *testing.T) {
// 	type testCase struct {
// 		mToSend  messageToSend
// 		expected bool
// 	}
// 	tests := []testCase{
// 		{messageToSend{
// 			message:   "you have an appointment tomorrow",
// 			sender:    user{name: "Brenda Halafax", number: 16545550987},
// 			recipient: user{name: "Sally Sue", number: 19035558973},
// 		}, true},
// 		{messageToSend{
// 			message:   "you have an event tomorrow",
// 			sender:    user{number: 16545550987},
// 			recipient: user{name: "Suzie Sall", number: 19035558973},
// 		}, false},
// 	}
// 	if withSubmit {
// 		tests = append(tests, []testCase{
// 			{messageToSend{
// 				message:   "you have an birthday tomorrow",
// 				sender:    user{name: "Jason Bjorn", number: 16545550987},
// 				recipient: user{name: "Jim Bond"},
// 			}, false},
// 			{messageToSend{
// 				message:   "you have a party tomorrow",
// 				sender:    user{name: "Njorn Halafax"},
// 				recipient: user{name: "Becky Sue", number: 19035558973},
// 			}, false},
// 			{messageToSend{
// 				message:   "you have a birthday tomorrow",
// 				sender:    user{name: "Eli Halafax", number: 16545550987},
// 				recipient: user{number: 19035558973},
// 			}, false},
// 		}...)
// 	}

// 	passCount := 0
// 	failCount := 0

// 	for _, test := range tests {
// 		output := canSendMessage(test.mToSend)
// 		if output != test.expected {
// 			failCount++
// 			t.Errorf(`---------------------------------
// Test Failed. Inputs:
// * message:          %s
// * sender.name:      %s
// * sender.number:    %d
// * recipient.name:   %s
// * recipient.number: %d
// Expected:           %v
// Actual:             %v
// `,
// 				test.mToSend.message,
// 				test.mToSend.sender.name,
// 				test.mToSend.sender.number,
// 				test.mToSend.recipient.name,
// 				test.mToSend.recipient.number,
// 				test.expected,
// 				output)
// 		} else {
// 			passCount++
// 			fmt.Printf(`---------------------------------
// Test Passed. Inputs:
// * message:          %s
// * sender.name:      %s
// * sender.number:    %d
// * recipient.name:   %s
// * recipient.number: %d
// Expected:           %v
// Actual:             %v
// `,
// 				test.mToSend.message,
// 				test.mToSend.sender.name,
// 				test.mToSend.sender.number,
// 				test.mToSend.recipient.name,
// 				test.mToSend.recipient.number,
// 				test.expected,
// 				output)
// 		}
// 	}

// 	fmt.Printf("---------------------------------\n")
// 	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
// }

// 5. Embedded Structs

// func getSenderLog(s sender) string {
// 	return fmt.Sprintf(`
// ====================================
// Sender name: %v
// Sender number: %v
// Sender rateLimit: %v
// ====================================
// `, s.name, s.number, s.rateLimit)
// }

// func Test(t *testing.T) {
// 	type testCase struct {
// 		rateLimit int
// 		name      string
// 		number    int
// 		expected  string
// 	}
// 	tests := []testCase{
// 		{
// 			10000,
// 			"Deborah",
// 			18055558790,
// 			`
// ====================================
// Sender name: Deborah
// Sender number: 18055558790
// Sender rateLimit: 10000
// ====================================
// `,
// 		},
// 		{
// 			5000,
// 			"Jason",
// 			18055558791,
// 			`
// ====================================
// Sender name: Jason
// Sender number: 18055558791
// Sender rateLimit: 5000
// ====================================
// `,
// 		},
// 	}
// 	if withSubmit {
// 		tests = append(tests, []testCase{
// 			{
// 				1000,
// 				"Jill",
// 				18055558792,
// 				`
// ====================================
// Sender name: Jill
// Sender number: 18055558792
// Sender rateLimit: 1000
// ====================================
// `,
// 			},
// 		}...)
// 	}

// 	passCount := 0
// 	failCount := 0

// 	for _, test := range tests {
// 		output := getSenderLog(sender{
// 			rateLimit: test.rateLimit,
// 			user: user{
// 				name:   test.name,
// 				number: test.number,
// 			},
// 		})
// 		if output != test.expected {
// 			failCount++
// 			t.Errorf(`---------------------------------
// Inputs:     (%v, %v, %v)
// Expecting:  %v
// Actual:     %v
// Fail`, test.rateLimit, test.name, test.number, test.expected, output)
// 		} else {
// 			passCount++
// 			fmt.Printf(`---------------------------------
// Inputs:     (%v, %v, %v)
// Expecting:  %v
// Actual:     %v
// Pass
// `, test.rateLimit, test.name, test.number, test.expected, output)
// 		}
// 	}

// 	fmt.Println("---------------------------------")
// 	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
// }

// // withSubmit is set at compile time depending
// // on which button is used to run the tests
// var withSubmit = true


// 6. Struct methods in Go

func TestGetBasicAuth(t *testing.T) {
	tests := []struct {
		auth     authenticationInfo
		expected string
	}{
		{authenticationInfo{"Google", "12345"}, "Authorization: Basic Google:12345"},
		{authenticationInfo{"Bing", "98765"}, "Authorization: Basic Bing:98765"},
	}
	if withSubmit {
		tests = append(tests, struct {
			auth     authenticationInfo
			expected string
		}{authenticationInfo{"DDG", "76921"}, "Authorization: Basic DDG:76921"})
	}

	passCount := 0
	failCount := 0

	for _, test := range tests {
		output := test.auth.getBasicAuth()
		if output != test.expected {
			failCount++
			t.Errorf(`---------------------------------
Inputs:     %+v
Expecting:  %s
Actual:     %s
Fail`, test.auth, test.expected, output)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:     %+v
Expecting:  %s
Actual:     %s
Pass
`, test.auth, test.expected, output)
		}
	}

	fmt.Println("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}

// withSubmit is set at compile time depending
// on which button is used to run the tests
var withSubmit = true

// 7. Memory layout