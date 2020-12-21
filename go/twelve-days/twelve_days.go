package twelve

type phraseType struct {
	numstr string
	phrase string
}

var phrases = [...]phraseType{
	{"first", "a Partridge in a Pear Tree."},
	{"second", "two Turtle Doves, and "},
	{"third", "three French Hens, "},
	{"fourth", "four Calling Birds, "},
	{"fifth", "five Gold Rings, "},
	{"sixth", "six Geese-a-Laying, "},
	{"seventh", "seven Swans-a-Swimming, "},
	{"eighth", "eight Maids-a-Milking, "},
	{"ninth", "nine Ladies Dancing, "},
	{"tenth", "ten Lords-a-Leaping, "},
	{"eleventh", "eleven Pipers Piping, "},
	{"twelfth", "twelve Drummers Drumming, "},
}

const dayOfChristmas = " day of Christmas my true love gave to me: "

func Song() string {
	var poem = ""

	for i := 1; i < 12; i++ {
		poem = poem + Verse(i) + "\n"
	}
	poem = poem + Verse(12)
	return poem
}

func Verse(in int) string {
	verse := "On the " + phrases[in-1].numstr + dayOfChristmas

	for j := in - 1; j >= 0; j-- {
		verse = verse + phrases[j].phrase
	}
	//fmt.Println(verse)
	return verse
}
