// package main
package main

// START1 OMIT
import (
	"fmt"
	"log"
	"time"

	s "go.bug.st/serial.v1"
)

// STOP1 OMIT

func main() {
	// START2 OMIT
	// How do we connect
	mode := &s.Mode{
		BaudRate: 9600,
	}

	// Create a slice with every port available
	portS, err := s.GetPortsList()
	if err != nil {
		log.Fatal(err)
	}
	if len(portS) == 0 {
		log.Fatal("No serial ports found!")
	}
	// STOP2 OMIT

	// START3 OMIT
	// for each port
	for _, port := range portS {
		fmt.Printf("Found port: %v\n", port)

		// Establish connection
		connect, err := s.Open(port, mode)
		if err != nil {
			log.Println(err)
		}
		time.Sleep(1500 * time.Millisecond)

		err = beginComms(connect)
		if err != nil {
			log.Println(err)
		}

		break
	}
	// STOP3 OMIT
}

// START4 OMIT
func beginComms(connect s.Port) (err error) {
	var challenge string
	challenge = "Hi, who are you?"

	// Write to the serial port
	_, err = connect.Write([]byte(challenge))
	if err != nil {
		return err
	}

	buff := make([]byte, 30)
	n, err := connect.Read(buff)
	if err != nil {
		return err
	}
	fmt.Printf("%v", string(buff[:n]))

	return nil
}

// STOP4 OMIT
