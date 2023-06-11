package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var ReadSlice []string
var Filesave string
var CurrentString = ""

func main() {
	fmt.Println("Перед переводом сохранить txt файл из word в кодировке Юникод UTF-8 без разрывов строк.")
	ReadSlice = make([]string, 0)

	ReadFile()
}

func ReadFile() { //This function reads user txt file and writes it to ReadSlice, also function generates file name to save.
	fmt.Print("Please input the name of file to read: ")
	filename, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	filename = strings.Trim(filename, "\r\n")
	f, err := os.Open(filename)
	BiFilename := strings.Split(filename, ".")
	var array []string //[]string
	if err == nil {
		fmt.Println(" " + filename + " -> " + BiFilename[0] + "_bi." + BiFilename[1])
		readdata := bufio.NewScanner(f)

		for readdata.Scan() {
			array = append(array, readdata.Text())
		}

		for _, arrayslice := range array {
			for I := 0; I < len(arrayslice); I++ {
				switch string(arrayslice[I]) {
				case "!", "?", ";": //Single splitters
					{
						CurrentString += string(arrayslice[I])
						ReadSlice = append(ReadSlice, CurrentString)
						CurrentString = ""
					}
				case ".": //Multy same splitters
					{
						if string(arrayslice[I+1]) == "." {
							CurrentString = CurrentString + "..."
							I = I + 2
							ReadSlice = append(ReadSlice, CurrentString)
							CurrentString = ""
						} else {
							CurrentString += string(arrayslice[I])
							ReadSlice = append(ReadSlice, CurrentString)
							CurrentString = ""

						}
					}
				default: //Next simbol if there is no splitter
					CurrentString = CurrentString + string(arrayslice[I])
				}

			}
			ReadSlice = append(ReadSlice, CurrentString) //Make string if there is no splitter, but EOL present
			CurrentString = ""
		}
		for i, ii := range ReadSlice {
			fmt.Println(i, ii)
		}

	} else {
		fmt.Println("File open error! Please check the filename. " + filename)
	}
	fmt.Println("Прочитано строк: ", len(ReadSlice))
	Filesave = (BiFilename[0] + "_bi." + BiFilename[1])

}
