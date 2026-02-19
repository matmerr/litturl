package server

import (
	"os"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	// Tests that exercise config/word-list features need to run from the project root
	if err := os.Chdir(".."); err != nil {
		panic(err)
	}
	os.Exit(m.Run())
}

func TestMakeURLTranslationKey(t *testing.T) {
	kg, err := MakeKeyGenerator(3, 4, "http://short.ly/")
	if err != nil {
		t.Fatalf("MakeKeyGenerator returned error: %v", err)
	}
	url := "https://github.com/matmerr/litturl"
	key := kg.GenerateKey(url)
	if len(key) == 0 {
		t.Error("GenerateKey returned empty string")
	}
}

func TestGenerateKeyDeterministic(t *testing.T) {
	kg, err := MakeKeyGenerator(3, 4, "http://short.ly/")
	if err != nil {
		t.Fatalf("MakeKeyGenerator returned error: %v", err)
	}
	url := "https://example.com/some/path"
	key1 := kg.GenerateKey(url)
	key2 := kg.GenerateKey(url)
	if key1 != key2 {
		t.Errorf("GenerateKey is not deterministic: %q != %q", key1, key2)
	}
}

func TestGenerateKeyDifferentURLs(t *testing.T) {
	kg, err := MakeKeyGenerator(3, 4, "http://short.ly/")
	if err != nil {
		t.Fatalf("MakeKeyGenerator returned error: %v", err)
	}
	key1 := kg.GenerateKey("https://example.com/a")
	key2 := kg.GenerateKey("https://example.com/b")
	if key1 == key2 {
		t.Error("different URLs produced the same key")
	}
}

func TestMakeURLTranslationFields(t *testing.T) {
	kg, err := MakeKeyGenerator(3, 4, "http://short.ly/")
	if err != nil {
		t.Fatalf("MakeKeyGenerator returned error: %v", err)
	}

	// Save and restore global Config state so this test doesn't affect others.
	prevKeyGenerator := Config.keyGenerator
	prevTinyAddress := Config.TinyAddress
	defer func() {
		Config.keyGenerator = prevKeyGenerator
		Config.TinyAddress = prevTinyAddress
	}()

	Config.keyGenerator = kg
	Config.TinyAddress = "http://short.ly/"

	longURL := "https://example.com/some/long/path"
	ut := MakeURLTranslation(longURL)

	if ut.OldURL != longURL {
		t.Errorf("OldURL mismatch: got %q, want %q", ut.OldURL, longURL)
	}
	if !strings.HasPrefix(ut.NewURL, "http://short.ly/") {
		t.Errorf("NewURL %q does not start with tiny address", ut.NewURL)
	}
	if len(ut.Wordkey) == 0 {
		t.Error("Wordkey is empty")
	}
}

func TestHashPassword(t *testing.T) {
	hash := hashPassword("testpassword")
	if len(hash) == 0 {
		t.Error("hashPassword returned empty string")
	}
	if hash == "testpassword" {
		t.Error("hashPassword returned plaintext password")
	}
}
