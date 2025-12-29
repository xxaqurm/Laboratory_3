package list

import (
	"math/rand/v2"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var arraysize int = 10000

func BenchmarkAddStart(t *testing.B) {
	fl := List{nil, nil}
	for range arraysize {
		fl.AddHead("123")
	}
}

func BenchmarkAddEnd(t *testing.B) {
	fl := List{nil, nil}
	t.ResetTimer()
	for i := 0; i < arraysize; i++ {
		fl.AddTail("123")
	}
}

func BenchmarkAddRandom(t *testing.B) {
	fl := List{nil, nil}
	indexes := make([]int, arraysize)
	for i := 0; i < arraysize; i++ {
		indexes[i] = rand.IntN(i + 1)
	}
	t.ResetTimer()
	for i := 0; i < arraysize; i++ {
		fl.Insert("123", indexes[i])
	}
}

func BenchmarkRemoveStart(t *testing.B) {
	fl := List{nil, nil}
	for i := 0; i < arraysize; i++ {
		fl.AddHead("123")
	}
	t.ResetTimer()
	for i := 0; i < arraysize; i++ {
		fl.RemoveHead()
	}
}

func BenchmarkRemoveEnd(t *testing.B) {
	fl := List{nil, nil}
	for i := 0; i < arraysize; i++ {
		fl.AddHead("123")
	}
	t.ResetTimer()
	for i := 0; i < arraysize; i++ {
		fl.RemoveTail()
	}
}

func TestAddGet(t *testing.T) {
	fl := List{nil, nil}
	fl.AddHead("123")
	fl.AddTail("789")
	fl.Insert("456", 1)
	fl.Insert("456", 4)
	fl.Insert("456", 5)
	assert.Equal(t, "123", fl.Head.Key)
	assert.Equal(t, "456", fl.Head.Next.Key)
	assert.Equal(t, "789", fl.Head.Next.Next.Key)
	assert.Nil(t, fl.Head.Next.Next.Next)
	fl.Insert("1", 0)
	assert.Equal(t, "1", fl.Head.Key)
}

func TestRemove(t *testing.T) {
	fl := List{nil, nil}
	fl.AddHead("1")
	fl.AddTail("2")
	fl.AddTail("3")
	fl.AddTail("4")
	assert.Equal(t, true, fl.RemoveHead()) //1
	assert.Equal(t, true, fl.RemoveTail()) //4
	assert.Equal(t, false, fl.Remove(2))   //no
	assert.Equal(t, false, fl.Remove(4))   //no
	assert.Equal(t, true, fl.Remove(1))    //3
	assert.Equal(t, true, fl.RemoveKey("2", 1))
	require.Nil(t, fl.Head)

	fl.AddTail("1")
	assert.Equal(t, true, fl.Remove(0))     //1
	assert.Equal(t, false, fl.Remove(1))    //no
	assert.Equal(t, false, fl.RemoveHead()) //no
	assert.Equal(t, false, fl.RemoveTail()) //no
	require.Nil(t, fl.Head)

	fl.AddTail("1")
	fl.AddTail("2")
	fl.AddTail("1")
	fl.AddTail("3")
	assert.Equal(t, false, fl.RemoveKey("1", 3))
	assert.Equal(t, false, fl.RemoveKey("1", 0))
	assert.Equal(t, true, fl.RemoveKey("1", 2))
	assert.Equal(t, "3", fl.Head.Next.Next.Key)
}

func TestFind(t *testing.T) {
	fl := List{nil, nil}
	fl.AddTail("1")
	fl.AddTail("2")
	fl.AddTail("1")
	assert.Nil(t, fl.Find("3", 1))
	assert.Nil(t, fl.Find("3", 0))
	assert.Equal(t, fl.Head, fl.Find("1", 1))
	assert.NotNil(t, fl.Find("1", 2))
	assert.Nil(t, fl.Find("1", 2).Next)
}

func TestPrint(t *testing.T) {
	fl := List{nil, nil}
	assert.Equal(t, "", fl.PrintFromHead())
	assert.Equal(t, "", fl.PrintFromTail())
	fl.AddTail("1")
	fl.AddTail("2")
	fl.AddTail("3")
	assert.Equal(t, "1 2 3", fl.PrintFromHead())
	assert.Equal(t, "3 2 1", fl.PrintFromTail())
}

func TestIO(t *testing.T) {
	textTestFile := ".texttestfile"
	binTestFile := ".bintestfile"
	//empty
	fl := &List{nil, nil}
	require.Nil(t, fl.WriteToFile(textTestFile))
	require.Nil(t, fl.WriteToFileBinary(binTestFile))
	fl, err := ReadListFromFile(textTestFile)
	require.Nil(t, err)
	require.Nil(t, fl.Head)
	fl, err = ReadListFromFileBinary(binTestFile)
	require.Nil(t, err)
	require.Nil(t, fl.Head)

	//not empty
	fl.AddTail("123")
	fl.AddTail("456")
	require.Nil(t, fl.WriteToFile(textTestFile))
	require.Nil(t, fl.WriteToFileBinary(binTestFile))

	textArr, err := ReadListFromFile(textTestFile)
	require.Nil(t, err)
	binArr, err := ReadListFromFileBinary(binTestFile)
	require.Nil(t, err)
	require.Equal(t, textArr.PrintFromHead(), binArr.PrintFromHead())
	require.Equal(t, textArr.PrintFromHead(), fl.PrintFromHead())

	//wrong file paths
	badPath := "/.txt...."
	assert.Error(t, fl.WriteToFile(badPath))
	assert.Error(t, fl.WriteToFileBinary(badPath))
	_, err = ReadListFromFile(badPath)
	assert.Error(t, err)
	_, err = ReadListFromFileBinary(badPath)
	assert.Error(t, err)
}
