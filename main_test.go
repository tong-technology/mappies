package mappies_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/tong-technology/mappies"
	"testing"
)

func Test_Contains(t *testing.T) {
	t.Run("Happy, Find Contains in []string", func(t *testing.T) {
		sampleData := []string{"India", "Canada", "Japan", "Germany", "Italy"}

		isContain := mappies.Contains(sampleData, "Canada")

		assert.Equal(t, true, isContain)
	})

	t.Run("Fail, Failed to Contains in []string", func(t *testing.T) {
		sampleData := []string{"India", "Canada", "Japan", "Germany", "Italy"}

		isContain := mappies.Contains(sampleData, "Canada1234")

		assert.Equal(t, false, isContain)
	})
}

func Test_MergeMap(t *testing.T) {
	t.Run("Happy", func(t *testing.T) {
		map1 := map[string]string{"Title": "Test Title"}
		map2 := map[string]string{"Name": "Test Name"}

		newMap := mappies.MergeMaps(map1, map2)

		assert.NotNil(t, newMap)
	})
}
