package game

import (
	"reflect"
	"testing"
)

func TestNewCardFromCode(t *testing.T) {
	tests := map[string]struct {
		arg     string
		want    *card
		wantErr bool
	}{
		"AS":  {"AS", &card{value{"A", "Ace"}, suit{"S", "Spades"}, "AS"}, false},
		"2S":  {"2S", &card{value{"2", "2"}, suit{"S", "Spades"}, "2S"}, false},
		"3S":  {"3S", &card{value{"3", "3"}, suit{"S", "Spades"}, "3S"}, false},
		"4S":  {"4S", &card{value{"4", "4"}, suit{"S", "Spades"}, "4S"}, false},
		"5S":  {"5S", &card{value{"5", "5"}, suit{"S", "Spades"}, "5S"}, false},
		"6S":  {"6S", &card{value{"6", "6"}, suit{"S", "Spades"}, "6S"}, false},
		"7S":  {"7S", &card{value{"7", "7"}, suit{"S", "Spades"}, "7S"}, false},
		"8S":  {"8S", &card{value{"8", "8"}, suit{"S", "Spades"}, "8S"}, false},
		"9S":  {"9S", &card{value{"9", "9"}, suit{"S", "Spades"}, "9S"}, false},
		"10S": {"10S", &card{value{"10", "10"}, suit{"S", "Spades"}, "10S"}, false},
		"JS":  {"JS", &card{value{"J", "Jack"}, suit{"S", "Spades"}, "JS"}, false},
		"QS":  {"QS", &card{value{"Q", "Queen"}, suit{"S", "Spades"}, "QS"}, false},
		"KS":  {"KS", &card{value{"K", "King"}, suit{"S", "Spades"}, "KS"}, false},
		"AD":  {"AD", &card{value{"A", "Ace"}, suit{"D", "Diamonds"}, "AD"}, false},
		"2D":  {"2D", &card{value{"2", "2"}, suit{"D", "Diamonds"}, "2D"}, false},
		"3D":  {"3D", &card{value{"3", "3"}, suit{"D", "Diamonds"}, "3D"}, false},
		"4D":  {"4D", &card{value{"4", "4"}, suit{"D", "Diamonds"}, "4D"}, false},
		"5D":  {"5D", &card{value{"5", "5"}, suit{"D", "Diamonds"}, "5D"}, false},
		"6D":  {"6D", &card{value{"6", "6"}, suit{"D", "Diamonds"}, "6D"}, false},
		"7D":  {"7D", &card{value{"7", "7"}, suit{"D", "Diamonds"}, "7D"}, false},
		"8D":  {"8D", &card{value{"8", "8"}, suit{"D", "Diamonds"}, "8D"}, false},
		"9D":  {"9D", &card{value{"9", "9"}, suit{"D", "Diamonds"}, "9D"}, false},
		"10D": {"10D", &card{value{"10", "10"}, suit{"D", "Diamonds"}, "10D"}, false},
		"JD":  {"JD", &card{value{"J", "Jack"}, suit{"D", "Diamonds"}, "JD"}, false},
		"QD":  {"QD", &card{value{"Q", "Queen"}, suit{"D", "Diamonds"}, "QD"}, false},
		"KD":  {"KD", &card{value{"K", "King"}, suit{"D", "Diamonds"}, "KD"}, false},
		"AC":  {"AC", &card{value{"A", "Ace"}, suit{"C", "Clubs"}, "AC"}, false},
		"2C":  {"2C", &card{value{"2", "2"}, suit{"C", "Clubs"}, "2C"}, false},
		"3C":  {"3C", &card{value{"3", "3"}, suit{"C", "Clubs"}, "3C"}, false},
		"4C":  {"4C", &card{value{"4", "4"}, suit{"C", "Clubs"}, "4C"}, false},
		"5C":  {"5C", &card{value{"5", "5"}, suit{"C", "Clubs"}, "5C"}, false},
		"6C":  {"6C", &card{value{"6", "6"}, suit{"C", "Clubs"}, "6C"}, false},
		"7C":  {"7C", &card{value{"7", "7"}, suit{"C", "Clubs"}, "7C"}, false},
		"8C":  {"8C", &card{value{"8", "8"}, suit{"C", "Clubs"}, "8C"}, false},
		"9C":  {"9C", &card{value{"9", "9"}, suit{"C", "Clubs"}, "9C"}, false},
		"10C": {"10C", &card{value{"10", "10"}, suit{"C", "Clubs"}, "10C"}, false},
		"JC":  {"JC", &card{value{"J", "Jack"}, suit{"C", "Clubs"}, "JC"}, false},
		"QC":  {"QC", &card{value{"Q", "Queen"}, suit{"C", "Clubs"}, "QC"}, false},
		"KC":  {"KC", &card{value{"K", "King"}, suit{"C", "Clubs"}, "KC"}, false},
		"AH":  {"AH", &card{value{"A", "Ace"}, suit{"H", "Hearts"}, "AH"}, false},
		"2H":  {"2H", &card{value{"2", "2"}, suit{"H", "Hearts"}, "2H"}, false},
		"3H":  {"3H", &card{value{"3", "3"}, suit{"H", "Hearts"}, "3H"}, false},
		"4H":  {"4H", &card{value{"4", "4"}, suit{"H", "Hearts"}, "4H"}, false},
		"5H":  {"5H", &card{value{"5", "5"}, suit{"H", "Hearts"}, "5H"}, false},
		"6H":  {"6H", &card{value{"6", "6"}, suit{"H", "Hearts"}, "6H"}, false},
		"7H":  {"7H", &card{value{"7", "7"}, suit{"H", "Hearts"}, "7H"}, false},
		"8H":  {"8H", &card{value{"8", "8"}, suit{"H", "Hearts"}, "8H"}, false},
		"9H":  {"9H", &card{value{"9", "9"}, suit{"H", "Hearts"}, "9H"}, false},
		"10H": {"10H", &card{value{"10", "10"}, suit{"H", "Hearts"}, "10H"}, false},
		"JH":  {"JH", &card{value{"J", "Jack"}, suit{"H", "Hearts"}, "JH"}, false},
		"QH":  {"QH", &card{value{"Q", "Queen"}, suit{"H", "Hearts"}, "QH"}, false},
		"KH":  {"KH", &card{value{"K", "King"}, suit{"H", "Hearts"}, "KH"}, false},
		"HZ":  {"HZ", nil, true},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := NewCardFromCode(tt.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewCardFromCode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCardFromCode() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCard_MarshalJSON(t *testing.T) {
	want := "{\"value\":\"7\",\"suit\":\"DIAMONDS\",\"code\":\"7D\"}"
	c := &card{
		value: value{"7", "7"},
		suit:  suit{"D", "Diamonds"},
		code:  "7D",
	}
	got, err := c.MarshalJSON()
	if err != nil {
		t.Errorf("MarshalJSON() error = %v", err)
		return
	}

	gotString := string(got)
	if gotString != want {
		t.Errorf("MarshalJSON() got = %v, want %v", gotString, want)
	}
}
