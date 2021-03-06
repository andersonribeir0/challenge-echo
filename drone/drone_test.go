package drone

import (
	"strconv"
	"testing"
)

func TestPhotosCommand(t *testing.T) {
	d := Drone{}
	a := NewArea(15, 20)

	d.Command("0112ODEDFFFEDF", a)
	if d.photos != 4 {
		t.Errorf("Expected 5 photos, but %s photos has been taken.", strconv.Itoa(d.photos))
	}

	d2 := Drone{}
	d2.Command("0112NFF", a)
	if d2.photos != 0 {
		t.Errorf("Photos should't have been taken. Photos: %d", d.photos)
	}
}

func TestDroneFinalPosition(t *testing.T) {
	d := Drone{}
	a := NewArea(20, 20)

	d.Command("1010NE", a)
	expected := Drone{
		photos: 1,
		x:      9,
		y:      10,
		z:      "O",
	}
	if expected.x != d.x {
		t.Errorf("Invalid 'x' value: '%d'. Expected: '%d'", d.x, expected.x)
	}
	if expected.y != d.y {
		t.Errorf("Invalid 'y' value: '%d'. Expected: '%d'", d.y, expected.y)
	}
	if expected.z != d.z {
		t.Errorf("Invalid 'z' value: '%s'. Expected: '%s'", d.z, expected.z)
	}

	d2 := Drone{}
	d2.Command("1010NFFDD", a)
	expected2 := Drone{
		photos: 2,
		x:      12,
		y:      12,
		z:      "S",
	}
	if expected2.x != d2.x {
		t.Errorf("Invalid 'x' value: '%d'. Expected: '%d'", d2.x, expected2.x)
	}
	if expected2.y != d2.y {
		t.Errorf("Invalid 'y' value: '%d'. Expected: '%d'", d2.y, expected2.y)
	}
	if expected2.z != d2.z {
		t.Errorf("Invalid 'z' value: '%s'. Expected: '%s'", d2.z, expected2.z)
	}
}

func TestValidateCommand(t *testing.T) {
	d := Drone{}

	err := d.validateCommand("1620NDEF")
	if err != nil {
		t.Error("A valid command was invalidated.", err)
	}

	err = d.validateCommand("11A5LDED")
	if err == nil {
		t.Error("An invalid command was accepted: ", "11A5LDED")
	}
}

func TestDronePosOutOfRange(t *testing.T) {
	d := Drone{}
	a := NewArea(1, 1)

	err := d.Command("0304NE", a)
	if err == nil {
		t.Error("Drone is out of range. Error must be returned by Command function.")
	}

	d2 := Drone{}
	err = d2.Command("0101OEEEE", a)
	if err == nil {
		t.Error("Drone has gone out of range. Error must be returned by Command function.")
	}
}
