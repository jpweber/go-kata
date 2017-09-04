package main

import (
	"bufio"
	"fmt"
	"hash/fnv"
	"log"
	"os"
)

type BloomFilter struct {
	dataset map[int32]bool
}

func (bf *BloomFilter) Add(value int32) {
	bf.dataset[value] = true
}

func (bf *BloomFilter) New() {
	bf.dataset = make(map[int32]bool)
}

func hashData(value string) int32 {
	hash := fnv.New32()
	hash.Write([]byte(value))
	data := hash.Sum32()
	return int32(data)
}

func main() {

	file, err := os.Open("/usr/share/dict/web2")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	bloom := BloomFilter{}
	bloom.New()

	log.Println("Hashing Data")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		bloom.Add(hashData(scanner.Text()))
	}

	fmt.Println("number of elements in map:", len(bloom.dataset))
}
