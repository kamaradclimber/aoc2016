package main

import (
  "testing"
)

func TestExtractChunksSuffix(t *testing.T) {
  chunks := ExtractChunks("terisuatldépvoiteraaaa", 4)
  if len(chunks) != 1 || chunks[0] != 'a' {
    t.Fatalf(`Expected to find aaaa chunk, found %v instead`, chunks)
  }

  chunks = ExtractChunks("terisuatldépvoiteraaa", 4)
  if len(chunks) != 0 {
    t.Fatalf(`Expected to find no chunk, found %v instead`, chunks)
  }
}

func TestExtractChunksBasic(t *testing.T) {
  chunks := ExtractChunks("aaaa", 4)
  if len(chunks) != 1 || chunks[0] != 'a' {
    t.Fatalf(`Expected to find aaaa chunk, found %v instead`, chunks)
  }

  chunks = ExtractChunks("aabbaa", 4)
  if len(chunks) != 0 {
    t.Fatalf(`Expected to find no chunk, found %v instead`, chunks)
  }
}

func TestExtractChunksMultipleChunks(t *testing.T) {
  chunks := ExtractChunks("aaaabbbb", 4)
  if len(chunks) != 2 || chunks[0] != 'a' || chunks[1] != 'b' {
    t.Fatalf(`Expected to find two chunks, found %v instead`, chunks)
  }

  chunks = ExtractChunks("aabbaaaaabbbbcccdcdcdc", 4)
  if len(chunks) != 2 || chunks[0] != 'a' || chunks[1] != 'b' {
    t.Fatalf(`Expected to find two chunks, found %v instead`, chunks)
  }
}

func TestHashPart1(t *testing.T) {
  hash := Hash("abc", 0, 2016)
  if hash != "a107ff634856bb300138cac6568c0f24" {
    t.Fatalf(`Expected to find the correct hash but found %s instead`, hash)
  }
}
