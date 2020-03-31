package qrcode

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type QRCode struct {
	ID 					primitive.ObjectID 		`json:"_id"bson:"_id"`
	Content 			string 					`json:"content"`
	QRToken				string					`json:"qrToken"`
	CreatedAt			time.Time				`json:"createdAt"`
}

func NewQRCode(Content string, QRToken string) QRCode {
	return QRCode{
		ID: primitive.NewObjectID(),
		Content: Content,
		QRToken: QRToken,
		CreatedAt: time.Now(),
	}
}

func (qr QRCode) QRToBson() bson.D {
	return bson.D {
		{"_id", qr.ID},
		{"content", qr.Content},
		{"qrToken",qr.QRToken},
		{"createdAt", qr.CreatedAt},
	}
}