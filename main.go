package main

import (
	"poly_go_sdk"

	"fmt"
)

func main() {

	//First, create an PolySDK instance with the NewPolySdk method.
	polySdk := poly_go_sdk.NewPolySdk()

	//Next, create an rpc, rest or websocket client.
	polySdk.NewRpcClient().SetAddress("http://localhost:20336")

	//Get current block height
	height, err := polySdk.GetCurrentBlockHeight()
	if err != nil {
		panic(err)
	}
	fmt.Println(height)

}
