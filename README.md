# cryptoCoins CLI Tool
![Test](https://github.com/andybroger/cryptoCoins/workflows/Test/badge.svg)

This little tool grabs the latest CryptoCoin -> USD Price and print it to the console.

usage:

```shell
go build
./cryptoCoins <CUR>

# example BTC (default)

$ ./cryptoCoins
$ 2020/03/26 09:14:54 - 1 BTC is 6612.98 USD

# example ETH
$ ./cryptoCoins ETH
$ 2020/03/26 09:16:54 - 1 ETH is 134.99 USD
```