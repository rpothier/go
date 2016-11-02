/*
You have 50 bitcoins to distribute to 10 users:
Matthew, Sarah, Augustus, Heidi, Emilie, Peter, Giana, Adriano, Aaron, Elizabeth
The coins will be distributed based on the vowels contained in each name where:

a: 1 coin e: 1 coin i: 2 coins o: 3 coins u: 4 coins

and a user can’t get more than 10 coins.
Print a map with each user’s name and the amount of coins distributed.
After distributing all the coins, you should have 2 coins left.

The output should look something like that:

map[Matthew:2 Peter:2 Giana:4 Adriano:7 Elizabeth:5 Sarah:2 Augustus:10 Heidi:5 Emilie:6 Aaron:5]
Coins left: 2
*/

package main

import (
	"fmt"
	"strings"
)

var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie",
		"Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

// using string function
func countChars1(user string) int {
	fmt.Println("countChars:", user)
	num := strings.Count(user, "a")
	num += strings.Count(user, "e")
	num += 2 * strings.Count(user, "i")
	num += 3 * strings.Count(user, "o")
	num += 4 * strings.Count(user, "u")
	if num >= 10 {
		num = 10
	}
	return num
}

// using case statement
func countChars2(user string) int {
	var ret int = 0
	for i := 0; i < len(user); i++ {
		switch user[i] {
		case 'a':
			ret++
		case 'e':
			ret++
		case 'i':
			ret = ret + 2
		case 'o':
			ret = ret + 3
		case 'u':
			ret = ret + 4
		}
	}
	if ret > 10 {
		ret = 10
	}
		return ret
}

func main() {

	for i := 0; i < len(users); i++ {
		distribution[users[i]] = countChars2(strings.ToLower(users[i]))
		coins -= distribution[users[i]]
	}

	fmt.Println(distribution)
	fmt.Println("Coins left:", coins)
}
