package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// attributes to render a page
type attributes struct {
	Title    string
	NotFound string
}

// s3Bucket to send the client from aws-s3 route
type s3Bucket struct {
	Name        string  `json:"name"`
	IsTruncated bool    `json:"is-truncated"`
	BaseUrl     string  `json:"base-url"`
	List        []Image `json:"list"`
}

// index route
func index(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	attr := attributes{Title: "Time Lapse"}
	renderTemplate(res, "index", &attr)
}

// notFound route
func notFound(res http.ResponseWriter, req *http.Request) {
	attr := attributes{Title: "Time Lapse", NotFound: "404 Bucket Not Found"}
	renderTemplate(res, "index", &attr)
}

// awsS3 route
func awsS3(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	buckets := req.URL.Query()["b"]
	var bucketName string
	if len(buckets) > 0 {
		bucketName = buckets[0]
	}
	resp, err := s3BucketList(bucketName)
	if err != nil {
		sendJSON(res, errorJSON{Error: err.Error()})
		return
	}
	var list []Image
	for _, image := range resp.Contents {
		newImage := Image{
			ETag:         *image.ETag,
			Key:          *image.Key,
			LastModified: *image.LastModified,
			Size:         *image.Size,
			StorageClass: *image.StorageClass,
		}
		list = append(list, newImage)
	}
	bucket := s3Bucket{
		Name:        *resp.Name,
		IsTruncated: *resp.IsTruncated,
		BaseUrl:     "https://s3.amazonaws.com",
		List:        list,
	}
	sendJSON(res, bucket)
}
