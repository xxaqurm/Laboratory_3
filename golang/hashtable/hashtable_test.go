package hashtable

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var arraysize int = 10000

func BenchmarkAdd(t *testing.B) {
	ht := NewHashTable(3)
	t.ResetTimer()
	for i := 0; i < arraysize; i++ {
		ht.Insert(i, 0)
	}
}

func BenchmarkRemove(t *testing.B) {
	ht := NewHashTable(3)
	for i := 0; i < arraysize; i++ {
		ht.Insert(i, 0)
	}
	t.ResetTimer()
	for i := 0; i < arraysize; i++ {
		ht.Remove(i)
	}
}

func BenchmarkSearch(t *testing.B) {
	ht := NewHashTable(3)
	for i := 0; i < arraysize; i++ {
		ht.Insert(i, 0)
	}
	t.ResetTimer()
	for i := 0; i < arraysize; i++ {
		ht.Get(i)
	}
}

func TestAddGet(t *testing.T) {
	ht := NewHashTable(7)
	ht.Insert(-9, 1)
	ht.Insert(-4, 2)
	ht.Insert(9, 3)
	ht.Insert(5, 4)
	ht.Insert(6, 5)
	ht.Insert(7, 6)
	ht.Insert(7, 7)

	assert.Equal(t, true, ht.Contains(-9))
	actual, err := ht.Get(-9)
	assert.Nil(t, err)
	assert.Equal(t, 1, actual)

	assert.Equal(t, true, ht.Contains(-4))
	actual, err = ht.Get(-4)
	assert.Nil(t, err)
	assert.Equal(t, 2, actual)

	assert.Equal(t, true, ht.Contains(9))
	actual, err = ht.Get(9)
	assert.Nil(t, err)
	assert.Equal(t, 3, actual)

	assert.Equal(t, true, ht.Contains(5))
	actual, err = ht.Get(5)
	assert.Nil(t, err)
	assert.Equal(t, 4, actual)

	assert.Equal(t, true, ht.Contains(6))
	actual, err = ht.Get(6)
	assert.Nil(t, err)
	assert.Equal(t, 5, actual)

	assert.Equal(t, true, ht.Contains(7))
	actual, err = ht.Get(7)
	assert.Nil(t, err)
	assert.Equal(t, 7, actual)

	assert.Equal(t, false, ht.Contains(0))
	_, err = ht.Get(0)
	assert.Error(t, err)
}

func TestRemove(t *testing.T) {
	ht := NewHashTable(7)
	ht.Insert(-9, 1)
	ht.Insert(-4, 2)
	ht.Insert(9, 3)
	ht.Insert(5, 4)

	assert.Equal(t, true, ht.Remove(-9))
	assert.Equal(t, false, ht.Remove(-9))
	assert.Equal(t, true, ht.Remove(5))
	assert.Equal(t, false, ht.Remove(5))
	assert.Equal(t, false, ht.Remove(0))

	_, err := ht.Get(-9)
	assert.Error(t, err)
	_, err = ht.Get(5)
	assert.Error(t, err)
	assert.Equal(t, true, ht.Contains(-4))
	assert.Equal(t, true, ht.Contains(9))
}

func TestToString(t *testing.T) {
	ht := NewHashTable(7)
	ht.Insert(-9, 1)
	ht.Insert(-4, 2)
	ht.Insert(9, 3)
	ht.Insert(5, 4)
	assert.Equal(t, "-4 : 2\n-9 : 1\n9 : 3\n5 : 4\n", ht.ToString())
}

func TestIO(t *testing.T) {
	textTestFile := ".texttestfile"
	binTestFile := ".bintestfile"
	//empty
	ht := NewHashTable(0)
	require.Nil(t, ht.WriteToFile(textTestFile))
	require.Nil(t, ht.WriteToFileBinary(binTestFile))
	ht, err := ReadHashTableFromFile(textTestFile)
	require.Nil(t, err)
	ht, err = ReadHashTableFromFileBinary(binTestFile)
	require.Nil(t, err)

	//not empty
	ht.Insert(1, 11)
	ht.Insert(2, 22)
	ht.Insert(3, 33)
	ht.Insert(4, 44)
	ht.Insert(5, 55)
	require.Nil(t, ht.WriteToFile(textTestFile))
	require.Nil(t, ht.WriteToFileBinary(binTestFile))

	textHT, err := ReadHashTableFromFile(textTestFile)
	require.Nil(t, err)
	binHT, err := ReadHashTableFromFileBinary(binTestFile)
	require.Nil(t, err)
	require.Equal(t, textHT.ToString(), binHT.ToString())
	require.Equal(t, textHT.ToString(), ht.ToString())

	//wrong file paths
	badPath := "/.txt...."
	assert.Error(t, ht.WriteToFile(badPath))
	assert.Error(t, ht.WriteToFileBinary(badPath))
	_, err = ReadHashTableFromFile(badPath)
	assert.Error(t, err)
	_, err = ReadHashTableFromFileBinary(badPath)
	assert.Error(t, err)
}
