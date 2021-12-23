package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func read(reader io.Reader) <-chan string {
	out := make(chan string)

	go func() {
		defer close(out)

		scanner := bufio.NewScanner(reader)

		for scanner.Scan() {
			out <- strings.TrimSpace(scanner.Text())
		}
	}()

	return out
}

func min(lhs int, rhs int) int {
	if lhs < rhs {
		return lhs
	}

	return rhs
}

func max(lhs int, rhs int) int {
	if lhs < rhs {
		return rhs
	}

	return lhs
}

func byte2int(bits []byte) int {
	value := 0

	for _, bit := range bits {
		value = (value << 1) | int(bit)
	}

	return value
}

func int2byte(value int, sizeBits int) []byte {
	bits := make([]byte, sizeBits)

	for i := sizeBits - 1; i >= 0 && value > 0; i-- {
		bits[i] = byte(value & 0x1)
		value >>= 1
	}

	return bits
}

func hex2byte(hexStr string) []byte {
	decoded := make([]byte, 0, len(hexStr)*4)

	for _, each := range strings.ToLower(hexStr) {
		var bits []byte

		switch {
		case each >= '0' && each <= '9':
			bits = int2byte(int(each-'0'), 4)
		case each >= 'a' && each <= 'f':
			bits = int2byte(int(each-'a')+10, 4)
		default:
			panic("invalid hex")
		}

		decoded = append(decoded, bits...)
	}

	return decoded
}

const (
	TypeInit = iota - 1
	TypeSum
	TypeProduct
	TypeMin
	TypeMax
	TypeValue
	TypeGreaterThan
	TypeLessThan
	TypeEqual
)

type Packet struct {
	Version int
	Type    int
	Value   int

	Size int
	Sub  []*Packet
}

func (p *Packet) Eval() int {
	value := 0

	switch p.Type {
	case TypeValue:
		value = p.Value
	case TypeSum:
		for _, each := range p.Sub {
			value += each.Eval()
		}
	case TypeProduct:
		for i := 0; i < len(p.Sub); i++ {
			if i > 0 {
				value *= p.Sub[i].Eval()
			} else {
				value = p.Sub[i].Eval()
			}
		}
	case TypeMin:
		for i := 0; i < len(p.Sub); i++ {
			if i > 0 {
				value = min(value, p.Sub[i].Eval())
			} else {
				value = p.Sub[i].Eval()
			}
		}
	case TypeMax:
		for _, each := range p.Sub {
			value = max(value, each.Eval())
		}
	case TypeGreaterThan:
		if len(p.Sub) > 1 && p.Sub[0].Eval() > p.Sub[1].Eval() {
			value = 1
		}
	case TypeLessThan:
		if len(p.Sub) > 1 && p.Sub[0].Eval() < p.Sub[1].Eval() {
			value = 1
		}
	case TypeEqual:
		if len(p.Sub) > 1 && p.Sub[0].Eval() == p.Sub[1].Eval() {
			value = 1
		}
	default:
		panic("invalid type")
	}

	return value
}

func decode(bits []byte) *Packet {
	if len(bits) == 0 {
		return nil
	} else if byte2int(bits) == 0 {
		return nil
	}

	packet := &Packet{
		Size:    6,
		Version: byte2int(bits[:3]),
		Type:    byte2int(bits[3:6]),
	}

	switch packet.Type {
	case TypeValue:
		decoded := []byte{}

		for i := 6; i < len(bits); i += 5 {
			decoded = append(decoded, bits[i+1:i+5]...)

			if bits[i] == 0 {
				packet.Size = i + 5
				break
			}
		}

		packet.Value = byte2int(decoded)
	default:
		lengthType, bits := bits[6], bits[7:]
		packet.Size++

		if lengthType == 0 {
			length, bits := byte2int(bits[:15]), bits[15:]
			payload, bits := bits[:length], bits[length:]
			packet.Size += 15 + length

			for {
				sub := decode(payload)

				if sub == nil {
					break
				}

				payload = payload[sub.Size:]
				packet.Sub = append(packet.Sub, sub)
			}
		} else {
			count, bits := byte2int(bits[:11]), bits[11:]
			packet.Size += 11

			for ; count > 0; count-- {
				sub := decode(bits)

				if sub == nil {
					break
				}

				bits = bits[sub.Size:]
				packet.Sub = append(packet.Sub, sub)
				packet.Size += sub.Size
			}
		}
	}

	return packet
}

func version(packet *Packet) int {
	out := packet.Version

	for _, each := range packet.Sub {
		out += version(each)
	}

	return out
}

func solution(token string) (int, int) {
	bits := hex2byte(token)
	packet := decode(bits)

	if packet == nil {
		return 0, 0
	}

	return version(packet), decode(bits).Eval()
}

func main() {
	token := ""

	for input := range read(os.Stdin) {
		token = input
	}

	fmt.Println(solution(token))
}
