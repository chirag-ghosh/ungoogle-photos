package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"strings"

	"github.com/knadh/koanf"
	kjson "github.com/knadh/koanf/parsers/json"
	"github.com/knadh/koanf/providers/file"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	gphotos "github.com/gphotosuploader/google-photos-api-client-go/v2"
)

const (
	PhotoslibraryScope = "https://www.googleapis.com/auth/photoslibrary"
	RedirectURL        = "http://localhost:3001"
)

var env = koanf.New(".")

func main() {

	if err := env.Load(file.Provider("secrets.json"), kjson.Parser()); err != nil {
		panic(err)
	}

	ctx := context.Background()
	oc := oauth2.Config{
		ClientID:     env.String("google.clientID"),
		ClientSecret: env.String("google.clientSecret"),
		Endpoint:     google.Endpoint,
		Scopes:       []string{PhotoslibraryScope},
		RedirectURL:  RedirectURL,
	}

	authURL := oc.AuthCodeURL(env.String("google.stateSecret"), oauth2.AccessTypeOffline)
	fmt.Printf("Visit the URL for the auth dialog: %v\n", authURL)
	fmt.Println()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("state") == env.String("google.stateSecret") {
			runGooglePhotosClient(ctx, oc, r.URL.Query().Get("code"))
		}
		fmt.Fprintf(w, "Check your terminal")
	})

	log.Fatal(http.ListenAndServe(":3001", nil))
}

func runGooglePhotosClient(ctx context.Context, oc oauth2.Config, code string) {

	token, err := oc.Exchange(ctx, code)
	if err != nil {
		log.Fatal(err)
	}

	tc := oc.Client(ctx, token)
	client, err := gphotos.NewClient(tc)
	if err != nil {
		log.Fatal(err)
	}

	mediaFiles, err := client.MediaItems.ListByAlbum(ctx, env.String("google.photosAlbumID"))
	if err != nil {
		log.Fatal(err)
	}

	var existingImageNames = getExistingImageNames()

	for i := 0; i < len(mediaFiles); i++ {
		if isAlreadyUploaded(existingImageNames, mediaFiles[i].Filename) {
			fmt.Println("Ignoring image: " + mediaFiles[i].Filename)
			continue
		}
		metadata, err := json.Marshal(mediaFiles[i].MediaMetadata)
		if err != nil {
			panic(err)
		}
		go uploadImageToImgur(mediaFiles[i].BaseURL+"=w"+mediaFiles[i].MediaMetadata.Width+"-h"+mediaFiles[i].MediaMetadata.Height, mediaFiles[i].Filename, string(metadata))
	}

	// album, err := client.Albums.Create(ctx, "Testing from App")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("created album: ", album)

	// albumById, err := client.Albums.GetById(ctx, "AF1QipM32m_CFh5XzJKTFWT7oCytZBShIXsmNSW3EiPW")
	// if err != nil {
	// 	log.Fatal(err)
	// }
}

func isAlreadyUploaded(existingImageNames []string, imgTitle string) bool {

	for i := 0; i < len(existingImageNames); i++ {
		if existingImageNames[i] == imgTitle {
			return true
		}
	}
	return false
}

func getExistingImageNames() []string {

	req, err := http.NewRequest("GET", "https://api.imgur.com/3/album/"+env.String("imgur.albumHash"), nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Authorization", "Client-ID "+env.String("imgur.clientID"))

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	b, _ := ioutil.ReadAll(res.Body)

	var albumInfoResponse AlbumInfoResponse
	err = json.Unmarshal(b, &albumInfoResponse)
	if err != nil {
		panic(err)
	}

	var fileNames []string
	for i := 0; i < len(albumInfoResponse.AlbumInfo.Images); i++ {
		fileNames = append(fileNames, albumInfoResponse.AlbumInfo.Images[i].Title)
	}

	return fileNames
}

func uploadImageToImgur(imageURL string, title string, metadata string) {

	fmt.Println("Writing image: " + title)

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	fw, err := writer.CreateFormField("title")
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(fw, strings.NewReader(title))
	if err != nil {
		panic(err)
	}

	fw, err = writer.CreateFormField("description")
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(fw, strings.NewReader(metadata))
	if err != nil {
		panic(err)
	}

	fw, err = writer.CreateFormField("type")
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(fw, strings.NewReader("url"))
	if err != nil {
		panic(err)
	}

	fw, err = writer.CreateFormField("album")
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(fw, strings.NewReader(env.String("imgur.albumDeleteHash")))
	if err != nil {
		panic(err)
	}

	fw, err = writer.CreateFormField("image")
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(fw, strings.NewReader(imageURL))
	if err != nil {
		panic(err)
	}

	writer.Close()

	req, err := http.NewRequest("POST", "https://api.imgur.com/3/image", bytes.NewReader(body.Bytes()))
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Authorization", "Client-ID "+env.String("imgur.clientID"))

	client := &http.Client{}
	_, err = client.Do(req)
	if err != nil {
		panic(err)
	}
}
