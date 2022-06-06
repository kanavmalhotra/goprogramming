# Prefix Match Problem

### Packaging:

As part of packaging, I have used the following directories
* `internal` : To keep the packages private to the module. 
* `config`: To keep the prefixies [Since, we have to make them configurable]. additionally, we can also keep a cfg file instead of the defined constants in the main.go file.

### Sample Data:

Added the sample data to the configs directory by the name of **sample_prefixes.txt**

### Pseudo code:

* Read sample prefixes and pre-process it.
* As part of pre-processing, I have decided to create a Trie(Prefix Tree) and use it as a cache across multiple input strings.
*  Since the list of prefix can be a big file, so read the file in chunks and pass it to channels to construct a trie object.
* Take list of input string that needs to verify the longest prefix
* For each input string, spawn a goroutine
* Each goroutine will print the longest prefix present, if any else print a blank string.




