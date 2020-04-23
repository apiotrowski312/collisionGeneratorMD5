package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

const (
	prefix               = "554" // value in HEX
	infoEveryXIterations = 10000000
)

func Hash(text string) string {
	data, _ := hex.DecodeString(text)

	hasher := md5.New()
	hasher.Write(data)
	firstMD5 := hasher.Sum(nil)

	hasher = md5.New()
	hasher.Write(firstMD5)
	return hex.EncodeToString(hasher.Sum(nil))[:14]
}

func main() {
	tortoise := Hash(prefix)
	hare := Hash(prefix + Hash(prefix))
	counter := 0
	fmt.Printf("Looking for loop.\n")
	for tortoise != hare {
		tortoise = Hash(prefix + tortoise)
		hare = Hash(prefix + Hash(prefix+hare))
		counter++
		if counter%infoEveryXIterations == 0 {
			fmt.Printf("Done %v iterations already.\nExample hashed Tortoise %v Hare %v.\n", counter, tortoise, hare)
		}
	}
	tortoise = ""
	counter = 0
	fmt.Printf("Searching for hash.\n")
	for tortoise != hare {
		prevTortoise := prefix + tortoise
		prevHare := prefix + hare
		tortoise = Hash(prevTortoise)
		hare = Hash(prevHare)
		counter++
		if counter%infoEveryXIterations == 0 {
			fmt.Printf("Done %v iterations already.\n", counter)
		}
		if tortoise == hare {
			fmt.Printf("Following hashes make a MD5(MD5(hex))[:14] collision:\n - tortoise: %v - hash: %v \n - hare: %v - hash: %v\n", prevTortoise, tortoise, prevHare, hare)
			break
		}
	}
}
