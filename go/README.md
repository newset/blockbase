https://www.kandaoni.com/news/6162.html 

https://gaitatzis.medium.com/compile-golang-as-a-mobile-library-243e38590f23 

https://rogchap.com/2020/09/14/running-go-code-on-ios-and-android/ 

## 编译
```
go build -buildmode=c-shared -o bin/arm64-v8a/WalletSDK.a lib.go
go build -buildmode=c-shared -o bin/arm64/WalletSDK.a lib.go
go build -buildmode=c-shared -o bin/x86_64/WalletSDK.a lib.go
go build -buildmode=c-shared -o bin/x86/WalletSDK.a lib.go
```
simulator
```
export GOOS=ios
export GOARCH=amd64
export SDK=iphonesimulator
```

iPhone
```
export GOOS=ios
export GOARCH=arm64
export SDK=iphoneos
```

android
```
set GOARCH=arm64
set GOOS=android
```

## 多链

## 签名
