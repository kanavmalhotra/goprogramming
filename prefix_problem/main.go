package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"example.com/prefix-problem/internal/trie"
)

const (
	prefixFilePath = "configs/sample_prefixes.txt"
	// prefixFilePath           = "configs/sample_pfx1.txt"
	maxLinesToRead           = 16
	maxDataOnBufferedChannel = 16
)

// PrefixChannel is a custom channel type
type PrefixChannel struct {
	data   []string
	isdone bool
}

// constructTrie function is responsible for constructing
// the trie from the list of lines passed over to this goroutine
func constructTrie(
	obj trie.Trie,
	ch chan PrefixChannel,
	wg *sync.WaitGroup,
) {
	defer close(ch)
	defer wg.Done()

	for {
		msg, more := <-ch
		if more {
			for _, val := range msg.data {
				// Adding to the Trie
				obj.Insert(val)
			}

			// Identifies whether the current message
			// is the last message or not
			if msg.isdone {
				break
			}
		}
	}
}

func main() {
	// Pseudo code:
	// - Read sample prefixes and pre-process it.
	// - As part of pre-processing, I have decided to create a
	//   Trie(Prefix Tree) and use it as a cache across multiple
	//   input strings.
	// - Since the list of prefix can be a big file, so read the
	//   file in chunks and pass it to channels to construct a trie object.
	// - Take list of input string that needs to
	//   verify the longest prefix
	// - For each input string, spawn a goroutine
	// - Each goroutine will print the longest prefix present, if
	//   any else print a blank string.

	absoluteFilePath, err := filepath.Abs(prefixFilePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Absolute path is:", absoluteFilePath)

	// Opening a file
	file, err := os.Open(absoluteFilePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Opening a buffer channel for reading line by line
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	// Creating a buffered Channel for sending the prefixies
	prefixChannel := make(chan PrefixChannel, maxDataOnBufferedChannel)
	wg := new(sync.WaitGroup)

	// Build a trie cache for prefix match
	trieCache := trie.NewPrefixTrie()
	go constructTrie(trieCache, prefixChannel, wg)

	count := 0
	var lines []string
	wg.Add(1)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
		count = count + 1

		// fmt.Println("read string: ", scanner.Text())
		if count == maxLinesToRead {
			prefixChannel <- PrefixChannel{
				data:   lines,
				isdone: false,
			}

			// Send the lines to channel
			// set count to 0 and reset lines
			// fmt.Printf("Setting count: %v\n", lines)
			lines = make([]string, 0)
			count = 0
		}
	}
	// Evaulting the last set of lines
	if count >= 0 {
		prefixChannel <- PrefixChannel{
			data:   lines,
			isdone: true,
		}
		// fmt.Printf("Setting count: %v\n", lines)
		count = 0
	}
	// Waiting for the completion of the building Trie cache
	wg.Wait()

	inputString := []string{"applz", "adpls", "applesdshjfkdsfbas", "abcd", "appppppp"}
	for index := range inputString {
		// leveraging the same wait group
		wg.Add(1)
		go func(i string) {
			defer wg.Done()
			prefix, val := trieCache.SearchLongestPrefix(i)
			fmt.Println("String: ", i, "---> Prefix: ", prefix, "Val: ", val)
		}(inputString[index])
	}
	// Waiting for all the goroutines to complete their execution
	wg.Wait()
}
