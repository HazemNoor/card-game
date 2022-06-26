package game

import (
	"reflect"
	"testing"
)

func TestNewValue(t *testing.T) {
	tests := map[string]struct {
		arg     string
		want    *value
		wantErr bool
	}{
		"small case":   {"a", &value{"A", "Ace"}, false},
		"capital case": {"A", &value{"A", "Ace"}, false},
		"2":            {"2", &value{"2", "2"}, false},
		"3":            {"3", &value{"3", "3"}, false},
		"4":            {"4", &value{"4", "4"}, false},
		"5":            {"5", &value{"5", "5"}, false},
		"6":            {"6", &value{"6", "6"}, false},
		"7":            {"7", &value{"7", "7"}, false},
		"8":            {"8", &value{"8", "8"}, false},
		"9":            {"9", &value{"9", "9"}, false},
		"10":           {"10", &value{"10", "10"}, false},
		"Jack":         {"J", &value{"J", "Jack"}, false},
		"Queen":        {"Q", &value{"Q", "Queen"}, false},
		"King":         {"K", &value{"K", "King"}, false},
		"wrong value":  {"z", nil, true},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := NewValue(tt.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewValue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewValue() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_allValues(t *testing.T) {
	want := map[string]string{
		"A":  "Ace",
		"2":  "2",
		"3":  "3",
		"4":  "4",
		"5":  "5",
		"6":  "6",
		"7":  "7",
		"8":  "8",
		"9":  "9",
		"10": "10",
		"J":  "Jack",
		"Q":  "Queen",
		"K":  "King",
	}

	if got := allValues(); !reflect.DeepEqual(got, want) {
		t.Errorf("allValues() = %v, want %v", got, want)
	}
}
