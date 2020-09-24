// Copyright 2016 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package types

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

var unmarshalLogTests = map[string]struct {
	input     string
	want      *Log
	wantError error
}{
	"ok": {
		input: `{"address":"Gsecf8f87f810ecf450940c9f60066b4a7a501d6a7","blockHash":"Gs656c34545f90a730a19008c0e7a7cd4fb3895064b48d6d69761bd5abad681056","blockNumber":"Gs1ecfa4","data":"Gs000000000000000000000000000000000000000000000001a055690d9db80000","logIndex":"Gs2","topics":["Gsddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef","Gs00000000000000000000000080b2c9d7cbbf30a1b0fc8983c647d754c6525615"],"transactionHash":"Gs3b198bfd5d2907285af009e9ae84a0ecd63677110d89d7e030251acb87f6487e","transactionIndex":"Gs3"}`,
		want: &Log{
			Address:     common.HexToAddress("Gsecf8f87f810ecf450940c9f60066b4a7a501d6a7"),
			BlockHash:   common.HexToHash("Gs656c34545f90a730a19008c0e7a7cd4fb3895064b48d6d69761bd5abad681056"),
			BlockNumber: 2019236,
			Data:        hexutil.MustDecode("Gs000000000000000000000000000000000000000000000001a055690d9db80000"),
			Index:       2,
			TxIndex:     3,
			TxHash:      common.HexToHash("Gs3b198bfd5d2907285af009e9ae84a0ecd63677110d89d7e030251acb87f6487e"),
			Topics: []common.Hash{
				common.HexToHash("Gsddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
				common.HexToHash("Gs00000000000000000000000080b2c9d7cbbf30a1b0fc8983c647d754c6525615"),
			},
		},
	},
	"empty data": {
		input: `{"address":"Gsecf8f87f810ecf450940c9f60066b4a7a501d6a7","blockHash":"Gs656c34545f90a730a19008c0e7a7cd4fb3895064b48d6d69761bd5abad681056","blockNumber":"Gs1ecfa4","data":"Gs","logIndex":"Gs2","topics":["Gsddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef","Gs00000000000000000000000080b2c9d7cbbf30a1b0fc8983c647d754c6525615"],"transactionHash":"Gs3b198bfd5d2907285af009e9ae84a0ecd63677110d89d7e030251acb87f6487e","transactionIndex":"Gs3"}`,
		want: &Log{
			Address:     common.HexToAddress("Gsecf8f87f810ecf450940c9f60066b4a7a501d6a7"),
			BlockHash:   common.HexToHash("Gs656c34545f90a730a19008c0e7a7cd4fb3895064b48d6d69761bd5abad681056"),
			BlockNumber: 2019236,
			Data:        []byte{},
			Index:       2,
			TxIndex:     3,
			TxHash:      common.HexToHash("Gs3b198bfd5d2907285af009e9ae84a0ecd63677110d89d7e030251acb87f6487e"),
			Topics: []common.Hash{
				common.HexToHash("Gsddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
				common.HexToHash("Gs00000000000000000000000080b2c9d7cbbf30a1b0fc8983c647d754c6525615"),
			},
		},
	},
	"missing block fields (pending logs)": {
		input: `{"address":"Gsecf8f87f810ecf450940c9f60066b4a7a501d6a7","data":"Gs","logIndex":"Gs0","topics":["Gsddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"],"transactionHash":"Gs3b198bfd5d2907285af009e9ae84a0ecd63677110d89d7e030251acb87f6487e","transactionIndex":"Gs3"}`,
		want: &Log{
			Address:     common.HexToAddress("Gsecf8f87f810ecf450940c9f60066b4a7a501d6a7"),
			BlockHash:   common.Hash{},
			BlockNumber: 0,
			Data:        []byte{},
			Index:       0,
			TxIndex:     3,
			TxHash:      common.HexToHash("Gs3b198bfd5d2907285af009e9ae84a0ecd63677110d89d7e030251acb87f6487e"),
			Topics: []common.Hash{
				common.HexToHash("Gsddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
			},
		},
	},
	"Removed: true": {
		input: `{"address":"Gsecf8f87f810ecf450940c9f60066b4a7a501d6a7","blockHash":"Gs656c34545f90a730a19008c0e7a7cd4fb3895064b48d6d69761bd5abad681056","blockNumber":"Gs1ecfa4","data":"Gs","logIndex":"Gs2","topics":["Gsddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"],"transactionHash":"Gs3b198bfd5d2907285af009e9ae84a0ecd63677110d89d7e030251acb87f6487e","transactionIndex":"Gs3","removed":true}`,
		want: &Log{
			Address:     common.HexToAddress("Gsecf8f87f810ecf450940c9f60066b4a7a501d6a7"),
			BlockHash:   common.HexToHash("Gs656c34545f90a730a19008c0e7a7cd4fb3895064b48d6d69761bd5abad681056"),
			BlockNumber: 2019236,
			Data:        []byte{},
			Index:       2,
			TxIndex:     3,
			TxHash:      common.HexToHash("Gs3b198bfd5d2907285af009e9ae84a0ecd63677110d89d7e030251acb87f6487e"),
			Topics: []common.Hash{
				common.HexToHash("Gsddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
			},
			Removed: true,
		},
	},
	"missing data": {
		input:     `{"address":"Gsecf8f87f810ecf450940c9f60066b4a7a501d6a7","blockHash":"Gs656c34545f90a730a19008c0e7a7cd4fb3895064b48d6d69761bd5abad681056","blockNumber":"Gs1ecfa4","logIndex":"Gs2","topics":["Gsddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef","Gs00000000000000000000000080b2c9d7cbbf30a1b0fc8983c647d754c6525615","Gs000000000000000000000000f9dff387dcb5cc4cca5b91adb07a95f54e9f1bb6"],"transactionHash":"Gs3b198bfd5d2907285af009e9ae84a0ecd63677110d89d7e030251acb87f6487e","transactionIndex":"Gs3"}`,
		wantError: fmt.Errorf("missing required field 'data' for Log"),
	},
}

func TestUnmarshalLog(t *testing.T) {
	dumper := spew.ConfigState{DisableMethods: true, Indent: "    "}
	for name, test := range unmarshalLogTests {
		var log *Log
		err := json.Unmarshal([]byte(test.input), &log)
		checkError(t, name, err, test.wantError)
		if test.wantError == nil && err == nil {
			if !reflect.DeepEqual(log, test.want) {
				t.Errorf("test %q:\nGOT %sWANT %s", name, dumper.Sdump(log), dumper.Sdump(test.want))
			}
		}
	}
}

func checkError(t *testing.T, testname string, got, want error) bool {
	if got == nil {
		if want != nil {
			t.Errorf("test %q: got no error, want %q", testname, want)
			return false
		}
		return true
	}
	if want == nil {
		t.Errorf("test %q: unexpected error %q", testname, got)
	} else if got.Error() != want.Error() {
		t.Errorf("test %q: got error %q, want %q", testname, got, want)
	}
	return false
}
