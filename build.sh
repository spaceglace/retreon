#!/bin/bash

rm -rf build
mkdir -p build/ui

cd frontend
npm run build
cd ..
cp -r frontend/dist/* build/ui

cd server
env GOOS=windows GOARCH=amd64 go build -o ../build/retro.exe
go build -o ../build/retro
chmod +x ../build/retro