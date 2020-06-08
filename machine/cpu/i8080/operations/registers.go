package operations

const (
	B uint8 = 0x00
	C uint8 = 0x01
	D uint8 = 0x02
	E uint8 = 0x03
	H uint8 = 0x04
	L uint8 = 0x05
	M uint8 = 0x06
	A uint8 = 0x07

	BC uint8 = 0x00
	DE uint8 = 0x01
	HL uint8 = 0x02
	SP uint8 = 0x03
)

var pc uint16 = 0x0000
var sp uint16 = 0x0000
var memory [0x10000]uint8
var registers [0x08]uint16
var flags uint8

var mnemonics = [256]string{
	"NOP", "LXI B,D16", "STAX B", "INX B", "INR B", "DCR B", "MVI B,D8", "RLC", "UNKNOWN",
	"DAD B", "LDAX B", "DCX B",
}
