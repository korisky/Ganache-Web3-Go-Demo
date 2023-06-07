package main

import (
	"context"
	"encoding/json"
	"github.com/davecgh/go-spew/spew"
	"github.com/portto/solana-go-sdk/program/token"
	"io"
	"log"
	"net/http"
	"testing"
)

type TokenList struct {
	Tokens []Token `json:"tokens"`
}

type Token struct {
	ChainID  int    `json:"chainId"`
	Address  string `json:"address"`
	Symbol   string `json:"symbol"`
	Name     string `json:"name"`
	Decimals uint8  `json:"decimals"`
	LogoURI  string `json:"logoURI"`
}

// Test_MintAccountInfo -> get the mint account's info (decimal, etc.)
func Test_MintAccountInfo(t *testing.T) {
	accountInfo, _ := cli.GetAccountInfo(context.Background(), "4jfNnrE97f4CzUt6z9C36NqpFxpv5T9erKU9JsG6Kr5N")
	tokenAccount, _ := token.TokenAccountFromData(accountInfo.Data)
	mintAccountInfo, _ := cli.GetAccountInfo(context.Background(), tokenAccount.Mint.ToBase58())
	mintAccount, _ := token.MintAccountFromData(mintAccountInfo.Data)
	spew.Dump(mintAccount)
}

// Test_GetExternalTokenInfo -> get token names & logo from github
func Test_GetExternalTokenInfo(t *testing.T) {
	resp, _ := http.Get("https://raw.githubusercontent.com/solana-labs/token-list/main/src/tokens/solana.tokenlist.json")
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var tokenList TokenList
	err := json.Unmarshal(body, &tokenList)
	if err != nil {
		log.Fatalf("Error on unmarshal, %v", err)
	}
	spew.Dump(tokenList)
}
