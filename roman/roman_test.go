package roman

import (
	"testing"
)

func Test1ShouldReturnI(t *testing.T) {
	if number := ConvertNumber(1); number != "I" {
		t.Error("Should be I")
	}
}

func Test2ShouldReturnII(t *testing.T) {
	if number := ConvertNumber(2); number != "II" {
		t.Error("Should be II")
	}
}

func Test3ShouldReturnIII(t *testing.T) {
	if number := ConvertNumber(3); number != "III" {
		t.Error("Should be III")
	}
}

func Test4ShouldReturnIV(t *testing.T) {
	if number := ConvertNumber(4); number != "IV" {
		t.Error("Should be IV")
	}
}

func Test5ShouldReturnV(t *testing.T) {
	if number := ConvertNumber(5); number != "V" {
		t.Error("Should be V")
	}
}

func Test6ShouldReturnVI(t *testing.T) {
	if number := ConvertNumber(6); number != "VI" {
		t.Error("Should be VI got", number)
	}
}

func Test7ShouldReturnVII(t *testing.T) {
	if number := ConvertNumber(7); number != "VII" {
		t.Error("Should be VII got", number)
	}
}

func Test8ShouldReturnVIII(t *testing.T) {
	if number := ConvertNumber(8); number != "VIII" {
		t.Error("Should be VIII got", number)
	}
}

func Test9ShouldReturnIX(t *testing.T) {
	if number := ConvertNumber(9); number != "IX" {
		t.Error("Should be IX got", number)
	}
}

func Test10ShouldReturn10(t *testing.T) {
	if number := ConvertNumber(10); number != "X" {
		t.Error("Should be X got", number)
	}
}
