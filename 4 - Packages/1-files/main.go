package main

import (
	"bufio"
	"math/rand"
	"os"
)

func main() {
	file, error := os.Create("file.txt")

	if error != nil {
		panic(error)
	}

	message := randomText()

	res, error := file.Write([]byte(message))
	// res, error := file.WriteString(message)

	if error != nil {
		panic(error)
	}

	println(res, "bytes written")

	file.Close()

	file1, error1 := os.ReadFile("file.txt")

	if error1 != nil {
		panic(error1)
	}

	println("File content: ", string(file1))

	// If my file is too big, I can read it in chunks
	file2, error2 := os.Open("file.txt")

	if error2 != nil {
		panic(error2)
	}

	reader := bufio.NewReader(file2)
	buffer := make([]byte, 15)

	for {
		bytesRead, error3 := reader.Read(buffer)

		if error3 != nil {
			break
		}

		println("Buffer: ", string(buffer[:bytesRead]))
	}

	file2.Close()

	// Remove file

	err := os.Remove("file.txt")

	if err != nil {
		panic(err)
	}
}

func randomText() string {
	textList := []string{
		"Hola, mundo!",
		"Hello, world!",
		"Bonjour, monde!",
		"Salut, monde!",
		"Ciao, mondo!",
		"Hallo, Welt!",
		"Ola, mundo!",
		"Namaste, duniya!",
		"Merhaba, dunya!",
		"Zdravstvuyte, mir!",
		"Konnichiwa, sekai!",
		"Anyoung haseyo, sesang!",
		"Ni hao, shijie!",
		"Yassou, kosmos!",
		"Salam, dunyo!",
		"Shalom, olam!",
	}

	return textList[rand.Intn(len(textList))]

}
