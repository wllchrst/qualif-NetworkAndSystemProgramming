package main

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
)

func dotaSTore() {
	context, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	request, err := http.NewRequestWithContext(context, http.MethodGet, "http://localhost:4321/dota/store", nil)

	if err != nil {
		fmt.Println(err)
		return
	}

	resp, err := http.DefaultClient.Do(request)

	if err != nil {
		if netError, ok := err.(net.Error); ok && netError.Timeout() {
			fmt.Println("Request Timed Out")
		} else {
			fmt.Println(err)
		}
	}

	defer func() {
		resp.Body.Close()
	}()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(body))
}

func dotaBestSkin() {
	context, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	request, err := http.NewRequestWithContext(context, http.MethodGet, "http://localhost:4321/dota/skin", nil)

	if err != nil {
		fmt.Println(err)
		return
	}

	resp, err := http.DefaultClient.Do(request)

	if err != nil {
		if netError, ok := err.(net.Error); ok && netError.Timeout() {
			fmt.Println("Request Timed Out")
		} else {
			fmt.Println(err)
		}
	}

	defer func() {
		resp.Body.Close()
	}()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(body))
}

func currentMeta() {
	var skinType string
	var total int

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Choose Skin to Buy [Rare, Ultra Rare, Arcana]: ")

		skinType, _ = reader.ReadString('\n')
		skinType = strings.TrimSpace(skinType)

		if skinType == "Rare" || skinType == "Ultra Rare" || skinType == "Arcana" {
			break
		} else {
			fmt.Println("Please choose valid skin type")
		}
	}

	for {
		fmt.Print("Total [x >= 0; x <= 6969]: ")
		fmt.Scanf("%d\n", &total)

		if total >= 0 && total <= 6969 {
			break
		} else {
			fmt.Println("Input valid quantity")
		}
	}

	requestBody := new(bytes.Buffer)

	w := multipart.NewWriter(requestBody)

	nameField, err := w.CreateFormField("name")

	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = nameField.Write([]byte(skinType))
	if err != nil {
		fmt.Println(err)
		return
	}

	quantityField, err := w.CreateFormField("total")

	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = quantityField.Write([]byte(string(total)))
	if err != nil {
		fmt.Println(err)
		return
	}

	file, err := os.Open("./buy.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fileField, err := w.CreateFormFile("file", file.Name())
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = io.Copy(fileField, file)
	if err != nil {
		fmt.Println(err)
		return
	}

	w.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	request, err := http.NewRequestWithContext(ctx, http.MethodPost, "http://localhost:4321/dota/buy-skin", requestBody)

	if err != nil {
		fmt.Println(err)
		return
	}

	resp, err := http.DefaultClient.Do(request)

	if err != nil {
		if netError, ok := err.(net.Error); ok && netError.Timeout() {
			fmt.Println("Request Timed Out")
		} else {
			fmt.Println(err)
		}
	}

	defer func() {
		resp.Body.Close()
	}()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(body))
}

func dotaPutFucntion() {
	context, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	request, err := http.NewRequestWithContext(context, http.MethodPut, "http://localhost:4321/dota/put-skin", nil)

	if err != nil {
		fmt.Println(err)
		return
	}

	resp, err := http.DefaultClient.Do(request)

	if err != nil {
		if netError, ok := err.(net.Error); ok && netError.Timeout() {
			fmt.Println("Request Timed Out")
		} else {
			fmt.Println(err)
		}
	}

	defer func() {
		resp.Body.Close()
	}()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(body))
}

func main() {
	for {
		fmt.Println("Dota Store")
		fmt.Println("1. Dota 2 Store")
		fmt.Println("2. Best Skin")
		fmt.Println("3. Current Meta")
		fmt.Println("4. Put Method test :0 ")
		fmt.Print(">> ")

		var choice int

		_, err := fmt.Scanf("%d\n", &choice)

		if err != nil {
			fmt.Println(err)
			return
		}

		if choice == 1 {
			dotaSTore()
		} else if choice == 2 {
			dotaBestSkin()
		} else if choice == 3 {
			currentMeta()
		} else if choice == 4 {
			dotaPutFucntion()
		}
	}
}
