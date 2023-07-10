package colorthis

type Color int 

const  (
	Reset Color = 0
	Red Color = 0x1F
	Green Color = 0x20
	Yellow Color = 0x21
	Blue Color = 0x22
	Mangeta Color = 0x23
	Cyan Color = 0x24
	White Color = 0x25
)

// Base is a function that receive a color and returns the string to print this.
func Base() string{
	return "\x1b[%dm"
}
