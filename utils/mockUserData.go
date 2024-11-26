package utils

func GetTestUserData() []User {
	return []User{
		{UserName: "Shubham", Karma: 0},
		{UserName: "Ananya", Karma: 10},
		{UserName: "Ravi", Karma: 5},
		{UserName: "Priya", Karma: 15},
		{UserName: "Shubham", Karma: 20}, // Duplicate of "Shubham"
		{UserName: "Ravi", Karma: 8},     // Duplicate of "Ravi"
		{UserName: "Akash", Karma: 12},
		{UserName: "Neha", Karma: 7},
		{UserName: "Priya", Karma: 15}, // Duplicate of "Priya"
		{UserName: "Vivek", Karma: 3},
		{UserName: "Shubham", Karma: 0}, // Another duplicate of "Shubham"
		{UserName: "Maya", Karma: 10},
		{UserName: "Suman", Karma: 25},
		{UserName: "Akash", Karma: 5}, // Duplicate of "Akash"
		{UserName: "Karan", Karma: 4},
	}
}
