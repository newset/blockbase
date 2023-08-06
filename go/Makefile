IOS_OUT=./target/ios
SDK_NAME=libWalletSDK
SDK_VERSION=1.0.0
PKG_NAME=./lib.go

clean:
	@echo "❌ cleaning up..."
	rm -rf ../android/libs/**

ios-arm64:
	CGO_ENABLED=1 \
	GOOS=darwin \
	GOARCH=arm64 \
	SDK=iphoneos \
	CC=$(PWD)/clangwrap.sh \
	CGO_CFLAGS="-fembed-bitcode" \
	go build -buildmode=c-archive -tags ios -o $(IOS_OUT)/arm64/$(SDK_NAME).a $(PKG_NAME)

ios-x86_64:
	CGO_ENABLED=1 \
	GOOS=darwin \
	GOARCH=amd64 \
	SDK=iphonesimulator \
	CC=$(PWD)/clangwrap.sh \
	go build -buildmode=c-archive -tags ios -o $(IOS_OUT)/x86_64/$(SDK_NAME).a $(PKG_NAME)

ios: ios-arm64 ios-x86_64
	lipo $(IOS_OUT)/x86_64/$(SDK_NAME).a $(IOS_OUT)/arm64/$(SDK_NAME).a -create -output $(IOS_OUT)/$(SDK_NAME).a
	cp $(IOS_OUT)/arm64/$(SDK_NAME).h ../ios/Classes/bindings.h
	cp $(IOS_OUT)/$(SDK_NAME).a ../ios/Frameworks/$(SDK_NAME).a
	sed '/^extern/!d' < $(IOS_OUT)/x86_64/$(SDK_NAME).h | sed '1d' > ../ios/Classes/bindings.h

ANDROID_OUT=../android/libs
ANDROID_SDK=$(HOME)/Library/Android/sdk
ANDROID_SDK_NAME=walletbase
NDK_BIN=$(ANDROID_SDK)/ndk/22.0.7026061/toolchains/llvm/prebuilt/darwin-x86_64/bin

android-armv7a:
	CGO_ENABLED=1 \
	GOOS=android \
	GOARCH=arm \
	GOARM=7 \
	CC=$(NDK_BIN)/armv7a-linux-androideabi21-clang \
	go build -buildmode=c-shared -o $(ANDROID_OUT)/armeabi-v7a/$(ANDROID_SDK_NAME).so $(PKG_NAME)

android-arm64:
	CGO_ENABLED=1 \
	GOOS=android \
	GOARCH=arm64 \
	CC=$(NDK_BIN)/aarch64-linux-android21-clang \
	go build -buildmode=c-shared -o $(ANDROID_OUT)/arm64-v8a/$(ANDROID_SDK_NAME).so $(PKG_NAME)

android-x86:
	CGO_ENABLED=1 \
	GOOS=android \
	GOARCH=386 \
	CC=$(NDK_BIN)/i686-linux-android21-clang \
	go build -buildmode=c-shared -o $(ANDROID_OUT)/x86/$(ANDROID_SDK_NAME).so $(PKG_NAME)

android-x86_64:
	CGO_ENABLED=1 \
	GOOS=android \
	GOARCH=amd64 \
	CC=$(NDK_BIN)/x86_64-linux-android21-clang \
	go build -buildmode=c-shared -o $(ANDROID_OUT)/x86_64/$(ANDROID_SDK_NAME).so $(PKG_NAME)

android: clean android-armv7a android-arm64 android-x86 android-x86_64

all: ios android
