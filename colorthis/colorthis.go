package colorthis

import (
	"fmt"
	"io"
	"log"
	"os"
)

// Colorthis is used to paint yours stdin, stderr, stdout
type colorthis struct{
	FileName string 
	Red Color 
	Reset Color 
	Green Color 
	Yellow Color 
	Blue Color 
	Mangeta Color 
	Cyan Color
	White Color
}

// New instace from color
// Params:
// 	fileName: "" if you not use LogFile, else uses a file name: "example.log"
func New(fileName string) *colorthis{
	if fileName != ""{
		return &colorthis{
			FileName: fileName,
			Red: Red,
			Reset: Reset, 
			Green: Green, 
			Yellow: Yellow, 
			Blue: Blue, 
			Mangeta: Mangeta, 
			Cyan: Cyan,
			White: White,
		}
	}
	return &colorthis{
		Red: Red,
		Reset: Reset, 
		Green: Green, 
		Yellow: Yellow, 
		Blue: Blue, 
		Mangeta: Mangeta, 
		Cyan: Cyan,
		White: White,
	}
}

/*Print the text with a color:
  Params:
	w: stdin, stderr, stdout
	color:  Reset, Red, Green, Yellow, Blue, Mangeta, Cyan or White
	text: string

	example: 
	    c := Colorthis.New()
		c.Print(os.Stdout, c.Red, "Colorthis is easy!")
*/ 
func (c colorthis) Print(w io.Writer, color Color, text string){
	format := Base()
	_, err := fmt.Fprintf(w , format , color)
	
	if err != nil{
		log.Println(err)
	}

	fmt.Fprintf(w, "%s\n", text)
}

func (c colorthis) Log(w io.Writer, color Color, text string){
	format := Base()
	_, err := fmt.Fprintf(w , format , color)
	
	if err != nil{
		log.Println(err)
	}

	log.Println(text)
}

func (c colorthis) LogFile(color Color, text string){
	logger, f := create(c.FileName)
	defer f.Close()

	format := Base()
	_, err := fmt.Fprintf(logger.Writer() , format , color)
	
	if err != nil{
		logger.Println(err)
	}
	logger.Println(text)
}

func create(fileName string) (logger *log.Logger, f *os.File){
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
			log.Fatalln(err)
	}

	logger = log.New(f, "", log.LstdFlags)
	return logger, f
}