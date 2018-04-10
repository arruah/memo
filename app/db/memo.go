package db

import (
	"git.jasonc.me/main/bitcoin/bitcoin/script"
	"git.jasonc.me/main/bitcoin/bitcoin/wallet"
	"github.com/btcsuite/btcutil"
	"github.com/cpacia/btcd/chaincfg/chainhash"
	"github.com/jchavannes/jgo/jerr"
	"html"
	"time"
)

type MemoTest struct {
	Id        uint   `gorm:"primary_key"`
	TxHash    []byte `gorm:"unique;size:50"`
	PkHash   []byte
	PkScript  []byte
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (m MemoTest) Save() error {
	result := save(&m)
	if result.Error != nil {
		return jerr.Get("error saving memo test", result.Error)
	}
	return nil
}

func (m MemoTest) GetTransactionHashString() string {
	hash, err := chainhash.NewHash(m.TxHash)
	if err != nil {
		jerr.Get("error getting chainhash from memo test", err).Print()
		return ""
	}
	return hash.String()
}

func (m MemoTest) GetAddressString() string {
	pkHash, err := btcutil.NewAddressPubKeyHash(m.PkHash, &wallet.MainNetParamsOld)
	if err != nil {
		jerr.Get("error getting pubkeyhash from memo test", err).Print()
		return ""
	}
	return pkHash.EncodeAddress()
}

func (m MemoTest) GetScriptString() string {
	return html.EscapeString(script.GetScriptString(m.PkScript))
}

func GetMemoTests() ([]*MemoTest, error) {
	var memoTests []*MemoTest
	err := find(&memoTests)
	if err != nil {
		return nil, jerr.Get("error getting memo tests", err)
	}
	return memoTests, nil
}
