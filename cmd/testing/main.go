package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	// Example raw beacon data string with multiple IDs and their respective RSSIs
	rawData := "1101ecafe01861c755aa975a8bc30c720055c701efaf7f6070c855aa9a5a2a0b1b730055ba"

	// Convert the raw hex string to a byte array
	data, err := hex.DecodeString(rawData)
	if err != nil {
		fmt.Println("Failed to decode hex string:", err)
		return
	}

	// Assuming the structure is:
	// - Header/Prefix: 2 bytes (can be ignored for now)
	// - Each Beacon ID: 16 bytes
	// - Each RSSI: 1 byte

	// Start parsing after the first 2 bytes (Header/Prefix)
	startIndex := 2

	// Define the size of each Beacon ID and RSSI
	idSize := 16
	rssiSize := 1

	// Iterate over the data and extract Beacon IDs and their respective RSSI values
	for i := startIndex; i+idSize+rssiSize <= len(data); {
		// Extract Beacon ID
		id := data[i : i+idSize]

		// Extract RSSI (next byte after the ID)
		rssi := int8(data[i+idSize])

		// Print Beacon ID and its respective RSSI value
		fmt.Printf("Beacon ID: %x, RSSI: %d dBm\n", id, rssi)

		// Debug: Print byte offsets for further analysis
		fmt.Printf("Bytes extracted: %x (from index %d to %d)\n", data[i:i+idSize], i, i+idSize)

		// Update `i` to the start of the next Beacon ID, skipping RSSI byte and 1 additional byte for alignment
		i += idSize + rssiSize + 1
	}
}
