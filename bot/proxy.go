package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"os"

	"github.com/gotd/td/telegram/dcs"
)

func decodeSecret(secretHex string) ([]byte, error) {
	secret, err := hex.DecodeString(secretHex)
	if err != nil {
		return nil, fmt.Errorf("invalid hex-encoded secret: %w", err)
	}
	return secret, nil
}

func createMTProxyResolver(proxyAddr, secretHex string) (dcs.Resolver, error) {
	secret, err := decodeSecret(secretHex)
	if err != nil {
		return nil, err
	}

	resolver, err := dcs.MTProxy(proxyAddr, secret, dcs.MTProxyOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to create MTProxy resolver: %w", err)
	}
	return resolver, nil
}

func getResolver() dcs.Resolver {
	proxyAddr := os.Getenv("PROXY_ADDR")
	secretHex := os.Getenv("PROXY_SECRET")

	resolver, err := createMTProxyResolver(proxyAddr, secretHex)
	if err != nil {
		log.Fatalf("Setup error: %v", err)
	}

	return resolver
}
