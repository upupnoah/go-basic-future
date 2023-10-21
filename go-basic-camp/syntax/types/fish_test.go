package types

import (
	"fmt"
	"testing"
)

func TestFish(t *testing.T) {
	//fish := Fish{}
	//fish.Swim()
	fakeFish := FakeFish{}
	fakeFish.Name = "fake fish"
	fmt.Println(fakeFish.Name)

}
