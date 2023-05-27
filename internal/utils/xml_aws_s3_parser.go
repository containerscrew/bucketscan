package utils

import (
	"encoding/xml"
	"fmt"
)

type ListBucketResult struct {
	XMLName     xml.Name `xml:"ListBucketResult"`
	Name        string   `xml:"Name"`
	Prefix      string   `xml:"Prefix"`
	Marker      string   `xml:"Marker"`
	MaxKeys     int      `xml:"MaxKeys"`
	IsTruncated bool     `xml:"IsTruncated"`
	Contents    []struct {
		Key          string `xml:"Key"`
		LastModified string `xml:"LastModified"`
		ETag         string `xml:"ETag"`
		Size         int    `xml:"Size"`
		StorageClass string `xml:"StorageClass"`
	} `xml:"Contents"`
}

func ListBucketContents(response, url string) {
	data := &ListBucketResult{}

	_ = xml.Unmarshal([]byte(response), &data)

	if len(data.Contents) > 0 {
		for _, content := range data.Contents {
			fmt.Printf("		->%s/%s \n", url, content.Key)
		}
	}
}
