package roman

import (
	"testing"
)

func Test1ShouldReturnI(t *testing.T) {
	if n := ConvertToRoman(1); n != "I" {
		t.Error("Should be I", n)
	}
}

func Test2ShouldReturnII(t *testing.T) {
	if n := ConvertToRoman(2); n != "II" {
		t.Error("Should be II", n)
	}
}

func Test3ShouldReturnIII(t *testing.T) {
	if n := ConvertToRoman(3); n != "III" {
		t.Error("Should be III", n)
	}
}

func Test4ShouldReturnIV(t *testing.T) {
	if n := ConvertToRoman(4); n != "IV" {
		t.Error("Should be IV", n)
	}
}

func Test5ShouldReturnV(t *testing.T) {
	if n := ConvertToRoman(5); n != "V" {
		t.Error("Should be V", n)
	}
}

func Test6ShouldReturnVI(t *testing.T) {
	if n := ConvertToRoman(6); n != "VI" {
		t.Error("Should be VI got", n)
	}
}

func Test7ShouldReturnVII(t *testing.T) {
	if n := ConvertToRoman(7); n != "VII" {
		t.Error("Should be VII got", n)
	}
}

func Test8ShouldReturnVIII(t *testing.T) {
	if n := ConvertToRoman(8); n != "VIII" {
		t.Error("Should be VIII got", n)
	}
}

func Test9ShouldReturnIX(t *testing.T) {
	if n := ConvertToRoman(9); n != "IX" {
		t.Error("Should be IX got", n)
	}
}

func Test10ShouldReturnX(t *testing.T) {
	if n := ConvertToRoman(10); n != "X" {
		t.Error("Should be X got", n)
	}
}

func Test11ShouldReturnXI(t *testing.T) {
	if n := ConvertToRoman(11); n != "XI" {
		t.Error("Should be X got", n)
	}
}

func Test50ShouldReturnL(t *testing.T) {
	if n := ConvertToRoman(50); n != "L" {
		t.Error("Should be L got", n)
	}
}

func Test90ShouldReturnXC(t *testing.T) {
	if n := ConvertToRoman(90); n != "XC" {
		t.Error("Should be XC got", n)
	}
}

func Test99ShouldReturnL(t *testing.T) {
	if n := ConvertToRoman(99); n != "XCIX" {
		t.Error("Should be XCIX got", n)
	}
}

func Test100ShouldReturnC(t *testing.T) {
	if n := ConvertToRoman(100); n != "C" {
		t.Error("Should be C got", n)
	}
}

func Test499ShouldReturnCDXCIX(t *testing.T) {
	if n := ConvertToRoman(499); n != "CDXCIX" {
		t.Error("Should be CDXCIX got", n)
	}
}

func Test500ShouldReturnD(t *testing.T) {
	if n := ConvertToRoman(500); n != "D" {
		t.Error("Should be D got", n)
	}
}

func Test999ShouldReturnCMXCIX(t *testing.T) {
	if n := ConvertToRoman(999); n != "CMXCIX" {
		t.Error("Should be CMXCIX got", n)
	}
}

func TestMShouldReturnM(t *testing.T) {
	if n := ConvertToRoman(1000); n != "M" {
		t.Error("Should be M got", n)
	}
}
