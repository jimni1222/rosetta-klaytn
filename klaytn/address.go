// Copyright 2020 Coinbase, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Modifications Copyright © 2022 Klaytn
// Modified and improved for the Klaytn development

package klaytn

import (
	"log"

	"github.com/klaytn/klaytn/common"
)

// ChecksumAddress ensures an Klaytn hex address
// is in Checksum Format. If the address cannot be converted,
// it returns !ok.
func ChecksumAddress(address string) (string, bool) {
	if isHexAddress := common.IsHexAddress(address); !isHexAddress {
		return "", false
	}

	addr := common.HexToAddress(address)

	return addr.Hex(), true
}

// MustChecksum ensures an address can be converted
// into a valid checksum. If it does not, the program
// will exit.
func MustChecksum(address string) string {
	addr, ok := ChecksumAddress(address)
	if !ok {
		log.Fatalf("invalid address %s", address)
	}

	return addr
}
