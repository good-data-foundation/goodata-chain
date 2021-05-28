package so

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"io"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/state"
)

func TestDataStore(t *testing.T) {
	// Create an empty state database
	db := rawdb.NewMemoryDatabase()
	state, _ := state.New(common.Hash{}, state.NewDatabase(db))

	ds := &DataStore{
		Db:   state,
		Addr: common.HexToAddress("1111111111111111111111111111111111111111")}

	for _, l := range []int{1, 31, 32, 44, 63, 64, 65, 128, 513} {
		fmt.Printf("l: %v\n", l)
		key := common.BigToHash(big.NewInt(int64(333333333 + l)))
		buf := bytes.Buffer{}
		io.CopyN(&buf, rand.Reader, 113)
		ds.Write(key, buf.Bytes())
		r := ds.Read(key)

		if !bytes.Equal(buf.Bytes(), r) {
			t.Errorf("Data not consistent for reading and writing: %d", l)
		}
	}
}
