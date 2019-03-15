package semver_test

import (
	"fmt"
	"log"
	"testing"

	"github.com/dolmen-go/semver"
)

func Example() {
	v := semver.V(1, 2, 3)
	fmt.Println(v)
	fmt.Println(v.Major())

	var v2 semver.Version
	if err := v2.Set("v4.5.6"); err != nil {
		log.Fatal(err)
	}
	fmt.Println(v2)

	// Output:
	// v1.2.3
	// v1
	// v4.5.6
}

func TestVersion(t *testing.T) {

}

func ExampleVersion_String() {
	fmt.Println(semver.V(1, 2, 3))
	// Output:
	// v1.2.3
}

func ExampleVersion_Major() {
	v := semver.V(1, 2, 3)
	fmt.Printf("%s major: string:%q int:%d", v, v.Major(), v.Major())
	// Output:
	// v1.2.3 major: string:"v1" int:1
}

func ExampleVersion_NextMajor() {
	fmt.Println(semver.V(1, 2, 3).NextMajor())
	// Output:
	// v2.0.0
}

func ExampleVersion_NextMinor() {
	fmt.Println(semver.V(1, 2, 3).NextMinor())
	// Output:
	// v1.3.0
}

func ExampleVersion_NextPatch() {
	fmt.Println(semver.V(1, 2, 3).NextPatch())
	// Output:
	// v1.2.4
}
