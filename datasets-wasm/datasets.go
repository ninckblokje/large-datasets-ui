package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"syscall/js"
)

type Data struct {
	Id int `json:"id"`
	DateTime string `json:"dateTime"`
	Value1 int `json:"value1"`
	Value2 int `json:"value2"`
	Value3 int `json:"value3"`
}

var StoredData []Data

func GetDataSet() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		size := args[0].Int()

		handler := js.FuncOf(func(this js.Value, args []js.Value) interface {} {
			resolve := args[0]
			reject := args[1]

			go func() {
				fmt.Printf("Retrieving dataset with size %d\n", size)

				req, err := http.NewRequest("GET", "http://localhost:8080/api/dataset", nil)
				req.Header.Add("js.fetch:mode", "cors")
				if err != nil {
					fmt.Println(err)
					reject.Invoke("test")
					return
				}

				q := req.URL.Query()
				q.Add("size", strconv.Itoa(size))
				req.URL.RawQuery = q.Encode()

				resp, err := http.DefaultClient.Do(req)
				if err != nil {
					fmt.Println(err)
					reject.Invoke("test")
					return
				}

				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					fmt.Println(err)
					reject.Invoke("test")
					return
				}

				err = json.Unmarshal(body, &StoredData)
				if err != nil {
					fmt.Println(err)
					reject.Invoke("test")
					return
				}

				resolve.Invoke(len(StoredData))

				defer resp.Body.Close()
			}()

			return nil
		})

		promiseConstructor := js.Global().Get("Promise")
		return promiseConstructor.New(handler)
	})
}

func PopulatePage() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		pageIndex := args[0].Int()
		pageSize := args[1].Int()

		handler := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			resolve := args[0]

			fmt.Printf("Populating page %d with size %d\n", pageIndex, pageSize)
			start := pageIndex * pageSize
			end := (pageIndex * pageSize) + pageSize

			if end > len(StoredData) {
				end = len(StoredData)
			}
			fmt.Printf("Start %d and end %d\n", start, end)

			responseData := make([]interface{}, end-start)
			for i, data := range StoredData[start:end] {
				responseData[i] = map[string]interface{} {
					"id": data.Id,
					"dateTime": data.DateTime,
					"value1": data.Value1,
					"value2": data.Value2,
					"value3": data.Value3,
				}
			}

			resolve.Invoke(responseData)

			return nil
		})

		promiseConstructor := js.Global().Get("Promise")
		return promiseConstructor.New(handler)
	})
}

func Reset() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		fmt.Println("Resetting WASM state")
		StoredData = nil

		return nil
	})
}

func main() {
	js.Global().Set("GetDataSet", GetDataSet())
	js.Global().Set("PopulatePage", PopulatePage())
	js.Global().Set("Reset", Reset())

	select {}

	fmt.Println("Hello, WebAssembly!")
}
