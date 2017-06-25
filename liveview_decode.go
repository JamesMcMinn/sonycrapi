package sonycrapi

import (
	"encoding/binary"
	"io"
	"log"
)

func (lv *Liveview) decodeCommonHeaders() (header LiveviewCommonHeader, err error) {
	startCode, err := lv.Stream.ReadByte()
	if err != nil {
		return
	}

	if startCode != liveviewCommonHeaderStartCode {
		log.Println("Incorrect liveview common header start code. Expected", liveviewCommonHeaderStartCode, "got", startCode)
	}

	header.PayloadType, err = lv.Stream.ReadByte()
	if err != nil {
		return
	}

	// Sequence Number
	seqNumberBytes := make([]byte, 2, 2)
	if _, err = io.ReadFull(lv.Stream, seqNumberBytes); err != nil {
		return
	}
	header.SequenceNumber = binary.BigEndian.Uint16(seqNumberBytes)

	// Timestamp
	timestmapBytes := make([]byte, 4, 4)
	if _, err = io.ReadFull(lv.Stream, timestmapBytes); err != nil {
		return
	}
	header.Timestamp = binary.BigEndian.Uint32(timestmapBytes)

	return
}

func (lv *Liveview) decodePayload() (payload LiveviewPayload, err error) {
	commonHeader, err := lv.decodeCommonHeaders()
	if err != nil {
		return
	}
	payload.LiveviewCommonHeader = commonHeader

	// Check the payload headers have the correct stat code.
	// If they don't, something is very wrong.
	startBytes := make([]byte, 4, 4)
	if _, err = io.ReadFull(lv.Stream, startBytes); err != nil {
		return
	}

	if string(startBytes) != string(payloadHeaderStartCode) {
		log.Printf("Incorrect liveview payload header start code. Expected %v for %v\n", payloadHeaderStartCode, startBytes)
		return
	}

	// PayloadSize
	payloadSizebytes := make([]byte, 3, 3)
	if _, err = io.ReadFull(lv.Stream, payloadSizebytes); err != nil {
		return
	}
	payload.PayloadSize = binary.BigEndian.Uint32(append([]byte{0x0}, payloadSizebytes...))

	// PaddingSize
	payload.PaddingSize, err = lv.Stream.ReadByte()
	if err != nil {
		return
	}

	if commonHeader.PayloadType == liveviewImagePayloadCode {
		if err = lv.decodeImagePayload(&payload); err != nil {
			return
		}
	} else if commonHeader.PayloadType == liveviewInfoPayloadCode {
		if err = lv.decodeInfoPayload(&payload); err != nil {
			return
		}
		return
	}

	return
}

func (lv *Liveview) decodeImagePayload(payload *LiveviewPayload) (err error) {
	imgPlayload := &LiveviewImagePayload{}
	payload.LiveviewImagePayload = imgPlayload

	// Skip the reserved block...
	reserved := make([]byte, 120, 120)
	if _, err = io.ReadFull(lv.Stream, reserved); err != nil {
		return
	}

	// Read the jpeg data
	jpegData := make([]byte, payload.PayloadSize, payload.PayloadSize)
	if _, err = io.ReadFull(lv.Stream, jpegData); err != nil {
		return
	}
	imgPlayload.JPEGData = jpegData

	// Discard the padding data
	lv.Stream.Discard(int(payload.PaddingSize))

	return
}

func (lv *Liveview) decodeInfoPayload(payload *LiveviewPayload) (err error) {
	log.Println("Info Frame. This are going to break.")
	return
}
