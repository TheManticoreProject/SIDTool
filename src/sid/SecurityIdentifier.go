package sid

import (
	"encoding/binary"
	"fmt"
	"strconv"
	"strings"
)

// Represents a Security Identifier (SID) in various formats and provides methods for manipulation and conversion between them.
//
// Attributes:
// 	revisionLevel (int): The revision level of the SID.
// 	subAuthorityCount (int): The number of sub-authorities in the SID.
// 	identifierAuthority (SID_IDENTIFIER_AUTHORITY): The identifier authority value.
// 	reserved (bytes): Reserved bytes, should always be empty.
// 	subAuthorities (list): A list of sub-authorities.
// 	relativeIdentifier (int): The relative identifier.
//
// Methods:
// 	Parse(RawBytes []byte): Parses the raw bytes to populate the SID fields.
// 	ToString() string: Converts the SID to its string representation.
//  Describe(): prints a detailed description of the SID with the specified indentation level.
//
// See: https://learn.microsoft.com/en-us/windows-server/identity/ad-ds/manage/understand-security-identifiers
// https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-dtyp/f992ad60-0fe4-4b87-9fed-beb478836861

type SID struct {
	RevisionLevel       uint8
	SubAuthorityCount   uint8
	IdentifierAuthority uint64
	SubAuthorities      []uint32
	RelativeIdentifier  uint32
	Reserved            []byte
	// Internal
	RawBytes     []byte
	RawBytesSize uint32
}

// IsWellKnownSID checks if the SID is a well-known SID.
func (sid *SID) IsWellKnownSID() bool {
	// Check if the SID is in the map of well-known SIDs
	_, found := WellKnownSIDs[sid.ToString()]
	return found
}

// LookupName returns the name of the well-known SID if it is well-known, otherwise it returns an empty string.
func (sid *SID) LookupName() string {
	// Check if it's a well-known SID and return its name
	if name, found := WellKnownSIDs[sid.ToString()]; found {
		return name
	}
	// If it's not a well-known SID, return an empty string
	return ""
}

// ToBytes converts the SID struct to its binary representation as a byte slice.
func (sid *SID) ToBytes() []byte {
	// Create a byte slice to hold the result
	buffer := make([]byte, 0)

	// Add the RevisionLevel
	buffer = append(buffer, sid.RevisionLevel)

	// Add the SubAuthorityCount
	buffer = append(buffer, sid.SubAuthorityCount)

	// Convert and add the IdentifierAuthority (6 bytes, big-endian)
	identifierBytes := make([]byte, 6)
	binary.BigEndian.PutUint16(identifierBytes[0:2], uint16(sid.IdentifierAuthority>>32))
	binary.BigEndian.PutUint16(identifierBytes[2:4], uint16(sid.IdentifierAuthority>>16))
	binary.BigEndian.PutUint16(identifierBytes[4:6], uint16(sid.IdentifierAuthority))
	buffer = append(buffer, identifierBytes...)

	// Add each sub-authority (4 bytes each, little-endian)
	for _, subAuthority := range sid.SubAuthorities {
		subAuthBytes := make([]byte, 4)
		binary.LittleEndian.PutUint32(subAuthBytes, subAuthority)
		buffer = append(buffer, subAuthBytes...)
	}

	// Add the Relative Identifier (4 bytes, little-endian)
	relativeIdentifierBytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(relativeIdentifierBytes, sid.RelativeIdentifier)
	buffer = append(buffer, relativeIdentifierBytes...)

	return buffer
}

func (sid *SID) FromBytes(RawBytes []byte) {
	sid.RawBytesSize = 0

	sid.RevisionLevel = uint8(RawBytes[0])
	sid.RawBytesSize += 1

	sid.SubAuthorityCount = uint8(RawBytes[1])
	sid.RawBytesSize += 1

	sid.IdentifierAuthority = 0
	sid.IdentifierAuthority += uint64(binary.BigEndian.Uint16(RawBytes[2:4])) >> 16
	sid.IdentifierAuthority += uint64(binary.BigEndian.Uint16(RawBytes[4:6])) >> 8
	sid.IdentifierAuthority += uint64(binary.BigEndian.Uint16(RawBytes[6:8]))
	sid.RawBytesSize += 6

	sid.SubAuthorities = make([]uint32, sid.SubAuthorityCount-1)
	for i := 0; i < int(sid.SubAuthorityCount-1); i++ {
		sid.SubAuthorities[i] = binary.LittleEndian.Uint32(RawBytes[sid.RawBytesSize : sid.RawBytesSize+4])
		sid.RawBytesSize += 4
	}

	sid.RelativeIdentifier = binary.LittleEndian.Uint32(RawBytes[sid.RawBytesSize : sid.RawBytesSize+4])
	sid.RawBytesSize += 4

	sid.RawBytes = RawBytes[:sid.RawBytesSize]
}

func (sid *SID) FromString(sidString string) error {
	// Split the SID string into parts using "-" as the delimiter
	parts := strings.Split(sidString, "-")

	// Ensure the SID string starts with "S" and has at least 4 parts: "S", revision, identifier authority, sub-authorities/RID
	if len(parts) < 4 || parts[0] != "S" {
		return fmt.Errorf("invalid SID string format")
	}

	// Parse the revision level (S-<Revision>)
	revision, err := strconv.Atoi(parts[1])
	if err != nil {
		return fmt.Errorf("invalid revision level in SID: %v", err)
	}
	sid.RevisionLevel = uint8(revision)

	// Parse the identifier authority (S-<Revision>-<IdentifierAuthority>)
	identifierAuthority, err := strconv.ParseUint(parts[2], 10, 64)
	if err != nil {
		return fmt.Errorf("invalid identifier authority in SID: %v", err)
	}
	sid.IdentifierAuthority = identifierAuthority

	// The rest are sub-authorities including the Relative Identifier (RID)
	var subAuthorities []uint32
	for i := 3; i < len(parts); i++ {
		subAuthority, err := strconv.ParseUint(parts[i], 10, 32)
		if err != nil {
			return fmt.Errorf("invalid sub-authority in SID: %v", err)
		}
		subAuthorities = append(subAuthorities, uint32(subAuthority))
	}

	// Determine SubAuthorityCount and set SubAuthorities/RID
	sid.SubAuthorityCount = uint8(len(subAuthorities))
	if sid.SubAuthorityCount > 0 {
		sid.SubAuthorities = subAuthorities[:sid.SubAuthorityCount-1]
		sid.RelativeIdentifier = subAuthorities[sid.SubAuthorityCount-1]
	}

	return nil
}

func (sid *SID) ToString() string {
	sidstring := fmt.Sprintf("S-%d-%d", sid.RevisionLevel, sid.IdentifierAuthority)
	for _, subauthority := range sid.SubAuthorities {
		sidstring += fmt.Sprintf("-%d", subauthority)
	}
	sidstring += fmt.Sprintf("-%d", sid.RelativeIdentifier)
	return sidstring
}

func (sid *SID) Describe() {
	fmt.Printf("<SID '%s'>\n", sid.ToString())
	fmt.Printf(" │ RevisionLevel        : 0x%02x\n", sid.RevisionLevel)
	fmt.Printf(" │ IdentifierAuthority  : 0x%02x\n", sid.IdentifierAuthority)

	if sid.SubAuthorityCount != 0 {
		fmt.Printf(" │ SubAuthorities (%03d) :\n", sid.SubAuthorityCount)
		for index, subauthority := range sid.SubAuthorities {
			fmt.Printf(" │ SubAuthority %02d : 0x%08x\n", index, subauthority)
		}
		fmt.Printf(" └─\n")
	} else {
		fmt.Printf(" │ SubAuthorities (0)   : Empty\n")
	}

	fmt.Printf(" │ RelativeIdentifier   : 0x%02x\n", sid.RelativeIdentifier)
	fmt.Printf(" └─\n")
}
