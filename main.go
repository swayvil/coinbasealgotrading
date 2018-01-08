package main

func main() {	
	wsocketClient := NewWSocketClient()
	wsocketClient.Listen(GetConfigInstance().Init.Crypto + "-" + GetConfigInstance().Init.Currency)
}