package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path"

	"github.com/btcsuite/btcd/btcutil/hdkeychain"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/tyler-smith/go-bip39"
	"github.com/urfave/cli/v2"
)

var accountCmd = &cli.Command{
	Name:    "account",
	Aliases: []string{"a"},
	Usage:   "Account manager",
	Subcommands: []*cli.Command{
		{
			Name:    "new",
			Aliases: []string{"n"},
			Usage:   "Create a new account",
			Action: func(cCtx *cli.Context) error {
				return create(cCtx)
			},
		},
	},
}

func create(cCtx *cli.Context) error {
	// get password from flag
	pwd := ""
	var err error
	if cCtx.String("password") != "" { // handle root flag `-password`
		pwd = cCtx.String("password")
	}

	// get password from input
	if pwd == "" {
		pwd, err = getInputPassword()
		if err != nil {
			return err
		}
		if pwd == "" {
			return errors.New("Password is empty")
		}
	}

	// gen Mnemonic
	entropy, err := bip39.NewEntropy(128) // 128\192\256
	if err != nil {
		return err
	}
	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		return err
	}

	// seed
	seed := bip39.NewSeed(mnemonic, "")

	// to master key
	masterKey, err := hdkeychain.NewMaster(seed, &chaincfg.MainNetParams)
	if err != nil {
		return err
	}

	// BIP44 path m/44'/60'/0'/0/0
	childKey, err := masterKey.Derive(hdkeychain.HardenedKeyStart + 44)
	if err != nil {
		return err
	}
	childKey, err = childKey.Derive(hdkeychain.HardenedKeyStart + 60)
	if err != nil {
		return err
	}
	childKey, err = childKey.Derive(hdkeychain.HardenedKeyStart)
	if err != nil {
		return err
	}
	childKey, err = childKey.Derive(0)
	if err != nil {
		return err
	}
	childKey, err = childKey.Derive(0)
	if err != nil {
		return err
	}
	/// end path

	// to private key
	privKeyBytes, err := childKey.ECPrivKey()
	if err != nil {
		return err
	}
	privKey, err := crypto.ToECDSA(privKeyBytes.Serialize())
	if err != nil {
		return err
	}

	keystoreDir := cCtx.String("keystore")

	// keystore
	ks := keystore.NewKeyStore(keystoreDir, keystore.StandardScryptN, keystore.StandardScryptP)

	account, err := ks.ImportECDSA(privKey, pwd)
	if err != nil {
		return err
	}

	// rename
	keystoreOutputPath := path.Join(keystoreDir, fmt.Sprintf("%s.keystore", account.Address.Hex()))
	err = os.Rename(account.URL.Path, keystoreOutputPath)
	if err != nil {
		return err
	}
	log.Println("Keystore saved: ", keystoreOutputPath)

	// output mnemonic
	mnemonicOutputPath := path.Join(keystoreDir, fmt.Sprintf("%s.mnemonic", account.Address.Hex()))
	err = os.WriteFile(mnemonicOutputPath, []byte(mnemonic), 0644)
	if err != nil {
		return err
	}
	log.Println("Mnemonic saved: ", mnemonicOutputPath)
	return nil
}
