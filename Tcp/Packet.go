package main

type packet struct {
	header  *header
	message string
}

type header struct {
	syn_number int
	ack_number int
}

func makeHeader(syn int, ack int) *header {
	h := header{}
	h.syn_number = syn
	h.ack_number = ack
	return &h

}

func makePacket(msg string, tcp_header *header) *packet {
	p := packet{}
	p.message = msg
	p.header = tcp_header

	return &p

}
