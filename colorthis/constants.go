package colorthis

type Color int 

const  (
	Reset Color = 0
	Red Color = 31
	Green Color = 32
	Yellow Color = 33
	Blue Color = 34
	Mangeta Color = 35
	Cyan Color = 36
	White Color = 37
)

// Base is a function that receive a color and returns the string to print this.
func Base() string{
	return "\x1b[%dm"
}
