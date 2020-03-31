package pkg

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"github.com/SungminSo/qr-generator/models"
	"github.com/SungminSo/qr-generator/models/qrcode"
	"github.com/gin-gonic/gin"
	qr "github.com/skip2/go-qrcode"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"image"
	"image/jpeg"
	"log"
	"net/http"
)

// Generate a 6-digit random number.
func GenerateRandom() string {
	randomNum, err := rand.Prime(rand.Reader, 32)
	if err != nil {
		log.Fatal(err)
	}
	return randomNum.String()[:6]
}

// Add salt string and encrypt them as base64 to increase the length of the QR code content.
// Because the length of the QR code content affects the complexity of the QR code.
func B64Encrypt(qrToken string) string {
	salt := "/qr-g3nerat0r/"
	data := salt + qrToken
	return base64.StdEncoding.EncodeToString([]byte(data))
}

// If there is a content, generate a QR Code with the content,
// If there is no content, generate a QR Code withe a 6-digit authentication number.
// Return QR Code, 6-digit authentication number and error.
func GenerateQR(content string) (qrCode []byte, qrContent string, qrToken string, err error) {
	qrContent = content

	qrToken = GenerateRandom()
	encryptedQRToken := B64Encrypt(qrToken)

	if len(encryptedQRToken) == 0 {
		return nil, "", "", errors.New("fail to create QR Code Token")
	}

	qrCode, err = qr.Encode(encryptedQRToken, qr.Medium, 256)
	if err != nil {
		return nil, "", "", errors.New("fail to create QR Code")
	}

	return qrCode, qrContent, qrToken, nil
}

func SendQR(c *gin.Context) {
	type Req struct {
		Content string 				`json:"content"`
	}

	reqBody := &Req{}
	err := c.ShouldBindJSON(reqBody)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	qrCode, qrContent, qrToken, err := GenerateQR(reqBody.Content)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	_, err = models.QRCodes.InsertOne(context.Background(), qrcode.NewQRCode(reqBody.Content, qrToken).QRToBson())
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	qrCodeImg, _, _ := image.Decode(bytes.NewReader(qrCode))
	buffer := new(bytes.Buffer)
	err = jpeg.Encode(buffer, qrCodeImg, nil)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	qrCodeStr := base64.StdEncoding.EncodeToString(buffer.Bytes())

	c.JSON(http.StatusOK, gin.H{
		"qrCode": qrCodeStr,
		"qrContent": qrContent,
		"qrToken": qrToken,
	})
}

// Verify by comparing entered QR Token with token,created when QR Code was generated, stored in DB.
func VerifyQR(c *gin.Context) {
	QRToken := c.Param("qrToken")

	var findResult qrcode.QRCode
	filter := bson.D{{Key: "qrToken", Value: QRToken}}
	opts := options.FindOne()
	opts.SetSort(bson.D{{"createdAt", -1}})

	err := models.QRCodes.FindOne(context.Background(), filter, opts).Decode(&findResult)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	if findResult.QRToken == QRToken {
		c.String(http.StatusOK, "Success")
	} else {
		c.String(http.StatusBadRequest, "Wrong QR Token")
	}

	return
}
