package server

import (
	"io"
	"log"
	"net"
)

func StartAndRunTcpServer() {
	lstnr, err := net.Listen("tcp", "0.0.0.0:7700")
	if err != nil {
		panic(err)
	}
	defer lstnr.Close()
	handleConnection(lstnr)

}

func handleConnection(lstnr net.Listener) {
	conc_conn := 0
	for {
		conn, err := lstnr.Accept()
		if err != nil {
			panic(err)
		}
		conc_conn++
		log.Println("client connected", conn.RemoteAddr(), "concurrent Connections", conc_conn)

		for {
			cmd, err := readCommand(conn)
			if err != nil {
				conc_conn--
				conn.Close()
				log.Println("client disconnected", conn.RemoteAddr(), "concurrent Connections", conc_conn)
				if err == io.EOF {
					log.Println("client disconnection command", err)
					break
				}
			}
			log.Println("Command:", cmd)
			err = respond(conn, cmd)

			if err != nil {
				log.Println("Error occured while responding: ", err)
			}
		}
	}

}

func readCommand(conn net.Conn) (string, error) {
	var buff []byte = make([]byte, 512)
	n, err := conn.Read(buff[:])
	if err != nil {
		return "", err

	}
	return string(buff[:n]), nil

}

func respond(conn net.Conn, cmd string) error {
	_, err := conn.Write([]byte(cmd))
	if err != nil {
		return err
	}
	return nil
}
