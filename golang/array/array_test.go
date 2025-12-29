package array

import (
	"math/rand/v2"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var arraysize int = 10000

func BenchmarkAddStart(t *testing.B) {
	arr := NewArray()
	for range arraysize {
		arr.Add("123")
	}
}

func BenchmarkAddEnd(t *testing.B) {
	arr := NewArray()
	t.ResetTimer()
	for i := 0; i < arraysize; i++ {
		arr.Insert("123", i)
	}
}

func BenchmarkAddRandom(t *testing.B) {
	arr := NewArray()
	indexes := make([]int, arraysize)
	for i := 0; i < arraysize; i++ {
		indexes[i] = rand.IntN(i + 1)
	}
	t.ResetTimer()
	for i := 0; i < arraysize; i++ {
		arr.Insert("123", indexes[i])
	}
}

func BenchmarkRemoveStart(t *testing.B) {
	arr := NewArray()
	for i := 0; i < arraysize; i++ {
		arr.Add("123")
	}
	t.ResetTimer()
	for i := 0; i < arraysize; i++ {
		arr.Remove(0)
	}
}

func BenchmarkRemoveEnd(t *testing.B) {
	arr := NewArray()
	for i := 0; i < arraysize; i++ {
		arr.Add("123")
	}
	t.ResetTimer()
	for i := 0; i < arraysize; i++ {
		arr.Remove(arr.GetSize() - 1)
	}
}

func BenchmarkRemoveRandom(t *testing.B) {
	arr := NewArray()
	indexes := make([]int, arraysize)
	for i := 0; i < arraysize; i++ {
		arr.Add("123")
		indexes[i] = rand.IntN(arraysize - i)
	}
	t.ResetTimer()
	for i := 0; i < arraysize; i++ {
		arr.Remove(indexes[i])
	}
}

func TestAddGet(t *testing.T) {
	arr := *NewArray()
	arr.Add("123")
	arr.Add("789")
	arr.Insert("456", 1)
	arr.Insert("111", 3)
	arr.Insert("1", 5)

	assert.Equal(t, 4, arr.GetSize())
	actual, err := arr.Get(0)
	assert.Nil(t, err)
	assert.Equal(t, "123", actual)
	actual, err = arr.Get(1)
	assert.Nil(t, err)
	assert.Equal(t, "456", actual)
	actual, err = arr.Get(2)
	assert.Nil(t, err)
	assert.Equal(t, "789", actual)
	actual, err = arr.Get(3)
	assert.Nil(t, err)
	assert.Equal(t, "111", actual)
	_, err = arr.Get(4)
	assert.Error(t, err)
}

func TestChange(t *testing.T) {
	arr := *NewArray()
	arr.Add("123")
	arr.Change(0, "1")
	arr.Change(1, "2")
	actual, err := arr.Get(0)
	assert.Nil(t, err)
	assert.Equal(t, "1", actual)
}

func TestRemove(t *testing.T) {
	arr := *NewArray()
	arr.Add("123")
	arr.Add("456")
	arr.Remove(0)
	assert.Equal(t, 1, arr.GetSize())
	actual, err := arr.Get(0)
	assert.Nil(t, err)
	assert.Equal(t, "456", actual)
	arr.Remove(1)
	assert.Equal(t, 1, arr.GetSize())
}

func TestToString(t *testing.T) {
	arr := *NewArray()
	arr.Add("123")
	arr.Add("456")
	assert.Equal(t, "123 456", arr.ToString())
}

func TestIO(t *testing.T) {
	textTestFile := ".texttestfile"
	binTestFile := ".bintestfile"
	//empty
	arr := NewArray()
	require.Nil(t, arr.WriteToFile(textTestFile))
	require.Nil(t, arr.WriteToFileBinary(binTestFile))
	arr, err := ReadArrayFromFile(textTestFile)
	require.Nil(t, err)
	require.Equal(t, 0, arr.GetSize())
	arr, err = ReadArrayFromFileBinary(binTestFile)
	require.Nil(t, err)
	require.Equal(t, 0, arr.GetSize())

	//not empty
	arr.Add("123")
	arr.Add("456")
	require.Nil(t, arr.WriteToFile(textTestFile))
	require.Nil(t, arr.WriteToFileBinary(binTestFile))

	textArr, err := ReadArrayFromFile(textTestFile)
	require.Nil(t, err)
	binArr, err := ReadArrayFromFileBinary(binTestFile)
	require.Nil(t, err)
	require.Equal(t, textArr.ToString(), binArr.ToString())
	require.Equal(t, textArr.ToString(), arr.ToString())

	//wrong file paths
	badPath := "/.txt...."
	assert.Error(t, arr.WriteToFile(badPath))
	assert.Error(t, arr.WriteToFileBinary(badPath))
	_, err = ReadArrayFromFile(badPath)
	assert.Error(t, err)
	_, err = ReadArrayFromFileBinary(badPath)
	assert.Error(t, err)
}
