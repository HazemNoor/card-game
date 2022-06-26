package game

import (
	"reflect"
	"testing"
)

func TestNewSuit(t *testing.T) {
	tests := map[string]struct {
		arg     string
		want    *suit
		wantErr bool
	}{
		"Spades":     {"S", &suit{"S", "Spades"}, false},
		"Diamonds":   {"D", &suit{"D", "Diamonds"}, false},
		"small_case": {"d", &suit{"D", "Diamonds"}, false},
		"Clubs":      {"C", &suit{"C", "Clubs"}, false},
		"Hearts":     {"H", &suit{"H", "Hearts"}, false},
		"wrong_suit": {"Z", nil, true},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := NewSuit(tt.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewSuit() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSuit() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_allSuits(t *testing.T) {
	want := map[string]string{
		"S": "Spades",
		"D": "Diamonds",
		"C": "Clubs",
		"H": "Hearts",
	}

	if got := allSuits(); !reflect.DeepEqual(got, want) {
		t.Errorf("allSuits() = %v, want %v", got, want)
	}
}
