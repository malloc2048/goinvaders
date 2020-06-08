package operations

import (
	"log"
)

func Execute(opcode uint8) int8 {
	log.Fatalf("unimplemented opcode %02X", opcode)
	return 0
}
