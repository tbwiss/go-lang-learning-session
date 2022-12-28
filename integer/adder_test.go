package integer

import "testing"
import "fmt"

func ExampleAdd() {
	sum := Add(3, 4)
	fmt.Println(sum)
}

func TestAdder(t *testing.T) {
	t.Run("add function", func(t *testing.T) {
		sun := Add(2, 2)
		expected := 4

		if sun != expected {
			t.Errorf("sun '%d' expected '%d'", sun, expected)
		}
	})

}
