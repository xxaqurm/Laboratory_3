package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var arraysize int = 10000

func BenchmarkPush(t *testing.B) {
	q := Queue{nil, nil}
	t.ResetTimer()
	for range arraysize {
		q.Push("123")
	}
}

func BenchmarkPop(t *testing.B) {
	q := Queue{nil, nil}
	for range arraysize {
		q.Push("123")
	}
	t.ResetTimer()
	for range arraysize {
		q.Pop()
	}
}

func TestPushPop(t *testing.T) {
	q := Queue{nil, nil}
	_, err := q.Pop()
	assert.Error(t, err)
	q.Push("1")
	q.Push("2")
	q.Push("3")
	actual, err := q.Pop()
	assert.Nil(t, err)
	assert.Equal(t, "1", actual)
	actual, err = q.Pop()
	assert.Nil(t, err)
	assert.Equal(t, "2", actual)
	actual, err = q.Pop()
	assert.Nil(t, err)
	assert.Equal(t, "3", actual)
}

func TestToString(t *testing.T) {
	s := Queue{nil, nil}
	assert.Equal(t, "", s.Print())
	s.Push("1")
	s.Push("2")
	s.Push("3")
	assert.Equal(t, "1 2 3 ", s.Print())
}

func TestIO(t *testing.T) {
	textTestFile := ".texttestfile"
	binTestFile := ".bintestfile"
	//empty
	q := &Queue{nil, nil}
	require.Nil(t, q.WriteToFile(textTestFile))
	require.Nil(t, q.WriteToFileBinary(binTestFile))
	q, err := ReadQueueFromFile(textTestFile)
	require.Nil(t, err)
	require.Equal(t, "", q.Print())
	q, err = ReadQueueFromFileBinary(binTestFile)
	require.Nil(t, err)
	require.Equal(t, "", q.Print())

	//not empty
	q.Push("123")
	q.Push("456")
	require.Nil(t, q.WriteToFile(textTestFile))
	require.Nil(t, q.WriteToFileBinary(binTestFile))

	textQueue, err := ReadQueueFromFile(textTestFile)
	require.Nil(t, err)
	binQueue, err := ReadQueueFromFileBinary(binTestFile)
	require.Nil(t, err)
	require.Equal(t, textQueue.Print(), binQueue.Print())
	require.Equal(t, textQueue.Print(), q.Print())

	//wrong file paths
	badPath := "/.txt...."
	assert.Error(t, q.WriteToFile(badPath))
	assert.Error(t, q.WriteToFileBinary(badPath))
	_, err = ReadQueueFromFile(badPath)
	assert.Error(t, err)
	_, err = ReadQueueFromFileBinary(badPath)
	assert.Error(t, err)
}
