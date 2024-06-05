package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"strconv"
)

func fibonacci(n int) *big.Int {
	if n == 0 {
		return big.NewInt(0)
	} else if n == 1 || n == 2 {
		return big.NewInt(1)
	} else {
		a := big.NewInt(1)
		b := big.NewInt(1)
		for i := 3; i <= n; i++ {
			a, b = b, new(big.Int).Add(a, b)
		}
		return b
	}
}

//TODO message名をわかりやすくする
func fibHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		writeErrorResponse(w, http.StatusMethodNotAllowed, "bad request")
		return
	}

	//クエリパラメータを取得
	keys, ok := r.URL.Query()["n"]
	if !ok || len(keys) == 0 || len(keys[0]) == 0 {
		writeErrorResponse(w, http.StatusBadRequest, "error: missing n parameter")
		return
	}

	//クエリパラメータを数字に変換
	number, err := strconv.Atoi(keys[0])
	if err != nil {
		writeErrorResponse(w, http.StatusBadRequest, "error: parse to int")
		return
	}

	//フィボナッチを計算
	if number < 1 {
		writeErrorResponse(w, http.StatusBadRequest, "error: inbvalid n parameter")
		return
	}
	res := fibonacci(number)

	//jsonにして返却
	w.Header().Set("Content-Type", "application/json")
	response := map[string]interface{}{
		"result": res,
	}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		writeErrorResponse(w, http.StatusBadRequest, "error: marshal to json")
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

//TODO 最後とそれぞれのメッセージごとに改行を入れる
func writeErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	response := map[string]interface{}{
		"status":  statusCode,
		"message": message,
	}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		log.Fatal("Marshal error:", err)
	}
	w.Write(jsonResponse)
}

func main() {
	http.HandleFunc("/fib", fibHandler)
	fmt.Println("Server is listening on https://localhost:443/fib...")
	err := http.ListenAndServeTLS(":443", "localhost.pem", "localhost-key.pem", nil)
	if err != nil {
		log.Fatal("Listen and Serve:", err)
	}
}
