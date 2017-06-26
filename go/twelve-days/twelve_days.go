package twelve

const testVersion = 1

func Song() string {
	output := ""
	for i := 1; i <= 12; i++ {
		output += Verse(i)
		output += "\n"
	}
	return output

}

func Verse(line int) string {
	var days = [][]string{
		{"first", "a Partridge in a Pear Tree"},
		{"second", "two Turtle Doves"},
		{"third", "three French Hens"},
		{"fourth", "four Calling Birds"},
		{"fifth", "five Gold Rings"},
		{"sixth", "six Geese-a-Laying"},
		{"seventh", "seven Swans-a-Swimming"},
		{"eighth", "eight Maids-a-Milking"},
		{"ninth", "nine Ladies Dancing"},
		{"tenth", "ten Lords-a-Leaping"},
		{"eleventh", "eleven Pipers Piping"},
		{"twelfth", "twelve Drummers Drumming"},
	}

	output := "On the " + days[line-1][0] + " day of Christmas my true love gave to me"
	for i := line - 1; i >= 0; i-- {
		output += ", "
		if i == 0 && line != 1 {
			output += "and "
		}
		output += days[i][1]
	}
	output += "."
	return output
}
