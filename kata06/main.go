package main

import (
	"bufio"
	"fmt"
	"hash/fnv"
	"log"
	"os"
	"sort"
	"strings"
)

func hashData(value string) int32 {
	sortedString := SortString(value)
	hash := fnv.New32()
	hash.Write([]byte(sortedString))
	data := hash.Sum32()
	return int32(data)
}

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func main() {

	file, err := os.Open("wordlist")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	wordMap := make(map[int]map[int32][]string)
	log.Println("Hashing Data")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		word := scanner.Text()
		wordLen := len(word)
		wordHash := hashData(word)
		wordList := wordMap[wordLen][wordHash]
		wordList = append(wordList, word)
		if wordMap[wordLen] == nil {
			tempHash := make(map[int32][]string)
			tempHash[wordHash] = []string{word}
			wordMap[wordLen] = tempHash
		} else {
			wordMap[wordLen][wordHash] = wordList
		}

	}
	// fmt.Println(wordMap)
	fmt.Println("number of entries in wordmap", len(wordMap))

	file, err = os.Open("wordlist")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner = bufio.NewScanner(file)

	for scanner.Scan() {
		word := scanner.Text()
		wordLen := len(word)
		wordHash := hashData(word)
		if len(wordMap[wordLen][wordHash]) > 1 {
			fmt.Println(wordMap[wordLen][wordHash])
			wordMap[wordLen][wordHash] = nil
		}
	}

}
