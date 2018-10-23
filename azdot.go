package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/smtp"
	"strings"
)

func readlines(x string, c chan string) {
	count := 0
	scanner := bufio.NewScanner(strings.NewReader(x))
	for scanner.Scan() {
		c <- scanner.Text()
		count++
	}
	fmt.Printf("Read %d lines\n", count)
	close(c)
}

// Readlines splits a string into lines
// and returns each line via a channel
func Readlines(txt string) chan string {
	ch := make(chan string)
	go readlines(txt, ch)
	return ch
}

func search(txt string, searches []string) bool {
	for l := range Readlines(txt) {
		for _, search := range searches {
			if strings.Contains(l, search) {
				//fmt.Println("found")
				return true
			}
		}
	}
	//fmt.Println("not found")
	return false
}

type Mail struct {
	senderId string
	toIds    []string
	subject  string
	body     string
}

type SmtpServer struct {
	host string
	port string
}

func (s *SmtpServer) ServerName() string {
	return s.host + ":" + s.port
}

func (mail *Mail) BuildMessage() string {
	message := ""
	message += fmt.Sprintf("From: %s\r\n", mail.senderId)
	if len(mail.toIds) > 0 {
		message += fmt.Sprintf("To: %s\r\n", strings.Join(mail.toIds, ";"))
	}

	message += fmt.Sprintf("Subject: %s\r\n", mail.subject)
	message += "\r\n" + mail.body

	return message
}

func send(subject string) {
	from := "joemooney777@gmail.com"
	pass := "Da01rse1!777"
	to := "joe.mooney@gmail.com"
	body := "empty"

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n\n" +
		body

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}
	log.Print("sent email: " + subject)
}

// Try to send email thru cox
// GetPage return
func GetPage(url string) *[]byte {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	x, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	return &x
}

func main() {
	//send("testing123")

	//if len(os.Args) != 2 {
	//	fmt.Fprintf(os.Stderr, "Usage: %s URL\n", os.Args[0])
	//	os.Exit(1)
	//}

	positiveUrls := [...]string{ // we are checking for a string to be there
		"https://azdot.gov/media/blog",
	}

	negativeUrls := [...]string{ // we are hoping for a string NOT to be there
		"https://www.azdot.gov/motor-vehicles/VehicleServices/PlatesandPlacards/energy-efficient-plate-program",
	}

	negativeSearchers := []string{
		"The Energy Efficient Plate Program has reached its maximum limit of 10,000 vehicles",
	}
	positiveSearches := []string{
		"license plate",
	}
	pages := make(map[string]string)

	for _, url := range positiveUrls {
		pageText, found := pages[url]
		if found {
			//fmt.Println("reuse " + url)
		} else {
			//fmt.Println("fetch " + url)
			pageText = string(*GetPage(url))
			pages[url] = pageText
		}
		if search(pageText, positiveSearches) {
			send("Blue Plate: Blog updated")
		} else {
			send("Blue Plate: No updates")
		}
	}
	for _, url := range negativeUrls {
		pageText, found := pages[url]
		if found {
			//fmt.Println("reuse " + url)
		} else {
			//fmt.Println("fetch " + url)
			pageText = string(*GetPage(url))
			pages[url] = pageText
		}
		if !search(pageText, negativeSearchers) {
			fmt.Println("red alert! Text about not being available has changed")
			send("Blue Plate Alert! Text about not being available has changed")
		}
	}
}
