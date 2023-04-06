package dbUtil

import (
	"bytes"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"os"
)

func UploadImage(path string, name string) string {
	file, err := os.Open(path) //get nute arrau strea,
	if err != nil {
		panic(err)
	}
	oId, err := GetGridfsBucket().UploadFromStream(name, file)

	return oId.Hex()
}

func DownloadImage(objectId string) *bytes.Buffer {
	id, _ := primitive.ObjectIDFromHex(objectId)
	filBuf := bytes.NewBuffer(nil)
	if _, err := GetGridfsBucket().DownloadToStream(id, filBuf); err != nil {
		panic(err)
	}
	return filBuf
}

func UploadBytes(bytes *bytes.Buffer, name string) string {
	oId, err := GetGridfsBucket().UploadFromStream(name, bytes)
	if err != nil {
		panic(err)
	}
	return oId.Hex()
}
