SOURCE_CODE=$(shell pwd)/go/wallet
SDK_VERSION=1.0.0
SDK_NAME=WalletMobileSDK

.PHONY: all

all: 
	cd go && make all
	@echo "✅ building all"

# Commands
install: all
	flutter pub run ffigen --config ffigen.yaml
	flutter pub get
	(cd example/ios && pod install)
