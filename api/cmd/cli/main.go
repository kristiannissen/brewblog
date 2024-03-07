package main

import (
	"log"
	"reflect"

	"github.com/rpdg/vercel_blob"
)

func init() {
}

func main() {
	client := vercel_blob.NewVercelBlobClient()

	log.Println("Hello", client)

	ListFiles()
}

func ListFiles() {
	client := vercel_blob.NewVercelBlobClient()
	files, err := client.List(vercel_blob.ListCommandOptions{})
	if err != nil {
		log.Println("ListFiles(): ", err)
	}
	log.Printf("%T", files.Blobs)

	for _, v := range files.Blobs {
		log.Println(v.URL, v.PathName)

		r := reflect.ValueOf(&v).Elem()
		rt := r.Type()

		for i := 0; i < r.NumField(); i++ {
			f := rt.Field(i)
			log.Println(f.Name)
		}
	}

}
