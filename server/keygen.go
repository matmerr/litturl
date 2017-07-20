package server

import (
	"bytes"
	"hash/fnv"
)

//KeyGenerator is the type which convert's URL's into sequences of easy words
type KeyGenerator struct {
	WordDict     []string
	WordDictSize int
}

//Domain is the domain to be used for shortening "https://u.short.com/"
var Domain string

//MakeKeyGenerator initializes the key generator
func MakeKeyGenerator(minKeyWordLength, maxKeyWordLength int, shortURL string) (KeyGenerator, error) {
	k := new(KeyGenerator)
	Domain = shortURL
	var err error
	k.WordDict, _ = initializeWords(minKeyWordLength, maxKeyWordLength)
	k.WordDictSize = len(k.WordDict)
	return *k, err
}

//GenerateKey accepts a url, and generates a key based on the words in the Oxford-3000
func (k KeyGenerator) GenerateKey(fullRoute string) string {
	modw := len(fullRoute) / 3
	fir := fullRoute[0:modw]
	sec := fullRoute[modw : modw+modw]
	thir := fullRoute[modw+modw : len(fullRoute)]
	key := k.combineKeys(fir, sec, thir)
	return key
}

func (k KeyGenerator) combineKeys(fir, sec, thir string) string {
	firKey, secKey, thirKey := k.genSubKey(fir), k.genSubKey(sec), k.genSubKey(thir)
	var buffer bytes.Buffer
	buffer.WriteString(firKey)
	buffer.WriteString(secKey)
	buffer.WriteString(thirKey)
	return buffer.String()
}

func (k KeyGenerator) genSubKey(sub string) string {
	i64 := hash64(sub)
	j64 := i64 % uint64(k.WordDictSize)
	subkey := k.WordDict[j64]
	return subkey
}

func hash64(s string) uint64 {
	h := fnv.New64()
	h.Write([]byte(s))
	return h.Sum64()
}
