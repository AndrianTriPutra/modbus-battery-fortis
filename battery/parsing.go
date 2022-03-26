package battery

import (
	"strconv"
)

func parsing_pack(data []byte) (str_result [31]string, warning, protect, fault [16]string) {
	var (
		final_int16  [37]int16
		final_uint16 [37]uint16

		buff_float float32
		resp_float [37]float32
	)

	data_conv := BytesToUint16s(BIG_ENDIAN, data)

	for i := 0; i < len(data_conv); i++ {

		switch i {
		//int16
		case 0, 12, 31, 32, 33, 34, 35, 36:
			final_int16[i] = int16(data_conv[i])
			buff_float = float32(final_int16[i])

		//uint16
		case 1, 2, 3, 4, 5, 6, 7, 9, 10, 11, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30:
			final_uint16[i] = uint16(data_conv[i])
			buff_float = float32(final_uint16[i])

		}

		resp_float[i] = buff_float
	}

	str_result = insert_pack(resp_float)

	tobit := asBits(uint64(final_uint16[9]))
	for i := 0; i <= 15; i++ {
		warning[i] = strconv.FormatUint(tobit[i], 10)
	}

	tobit = asBits(uint64(final_uint16[10]))
	for i := 0; i <= 15; i++ {
		protect[i] = strconv.FormatUint(tobit[i], 10)
	}

	tobit = asBits(uint64(final_uint16[11]))
	for i := 0; i <= 15; i++ {
		fault[i] = strconv.FormatUint(tobit[i], 10)
	}

	return str_result, warning, protect, fault
}

func asBits(val uint64) []uint64 {
	var bits = []uint64{}
	for i := 0; i <= 15; i++ {
		bits = append([]uint64{val & 0x1}, bits...)
		val = val >> 1
	}
	return bits
}
