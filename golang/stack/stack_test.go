package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var arraysize int = 10000

func BenchmarkPush(t *testing.B) {
	s := Stack{nil}
	for range arraysize {
		s.Push("123")
	}
}

func BenchmarkPop(t *testing.B) {
	s := Stack{nil}
	for range arraysize {
		s.Push("123")
	}
	t.ResetTimer()
	for range arraysize {
		s.Pop()
	}
}

func TestPushPop(t *testing.T) {
	s := Stack{nil}
	_, err := s.Pop()
	assert.Error(t, err)
	s.Push("1")
	s.Push("2")
	s.Push("3")
	actual, err := s.Pop()
	assert.Nil(t, err)
	assert.Equal(t, "3", actual)
	actual, err = s.Pop()
	assert.Nil(t, err)
	assert.Equal(t, "2", actual)
	actual, err = s.Pop()
	assert.Nil(t, err)
	assert.Equal(t, "1", actual)
}

func TestToString(t *testing.T) {
	s := Stack{}
	assert.Equal(t, "", s.Print())
	s.Push("1")
	s.Push("2")
	s.Push("3")
	assert.Equal(t, "3 2 1", s.Print())
}

func TestIO(t *testing.T) {
	textTestFile := ".texttestfile"
	binTestFile := ".bintestfile"
	//empty
	s := &Stack{nil}
	require.Nil(t, s.WriteToFile(textTestFile))
	require.Nil(t, s.WriteToFileBinary(binTestFile))
	s, err := ReadStackFromFile(textTestFile)
	require.Nil(t, err)
	require.Equal(t, "", s.Print())
	s, err = ReadStackFromFileBinary(binTestFile)
	require.Nil(t, err)
	require.Equal(t, "", s.Print())

	//not empty
	s.Push("123")
	s.Push("456")
	require.Nil(t, s.WriteToFile(textTestFile))
	require.Nil(t, s.WriteToFileBinary(binTestFile))

	textStack, err := ReadStackFromFile(textTestFile)
	require.Nil(t, err)
	binStack, err := ReadStackFromFileBinary(binTestFile)
	require.Nil(t, err)
	require.Equal(t, textStack.Print(), binStack.Print())
	require.Equal(t, textStack.Print(), s.Print())

	//wrong file paths
	badPath := "/.txt...."
	assert.Error(t, s.WriteToFile(badPath))
	assert.Error(t, s.WriteToFileBinary(badPath))
	_, err = ReadStackFromFile(badPath)
	assert.Error(t, err)
	_, err = ReadStackFromFileBinary(badPath)
	assert.Error(t, err)
}
