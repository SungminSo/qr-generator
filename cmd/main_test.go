package main

import (
	"bufio"
	"github.com/SungminSo/qr-generator/server"
	"github.com/SungminSo/qr-generator/pkg"
	"github.com/spf13/viper"
	"os"
	"testing"
)

const BIND_ADDR = "0.0.0.0:3506"

func TestNewServer(t *testing.T) {
	viper.AutomaticEnv()

	bindAddr := viper.GetString("BIND_ADDR")
	if bindAddr == "" {
		bindAddr = BIND_ADDR
	}

	if bindAddr != BIND_ADDR {
		t.Errorf("bindAddr not match. expected: %v, get: %v", BIND_ADDR, bindAddr)
	}

	s := server.NewServer(bindAddr)

	if s.Bind != bindAddr {
		t.Error("server's bind address not match with bindAddr")
	}
}

func TestGenerateQR(t *testing.T) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	testContent := scanner.Text()

	qrCode, qrContent, qrToken, err := pkg.GenerateQR(testContent)
	if err != nil {
		t.Error(err.Error())
	}

	if qrCode == nil {
		t.Error("fail to generate QR Code")
	}

	if len(testContent) != 0 && qrContent == "" {
		t.Error("the content should have been issued because of the content")
	}

	if len(qrToken) != 6 {
		t.Error("the token should have been issued")
	}
}