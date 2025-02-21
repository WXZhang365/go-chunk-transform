GOOS=windows GOARCH=amd64 go build -o ./dist/gctu.exe
GOOS=linux GOARCH=amd64 go build -o ./dist/gctu_linux_amd64
GOOS=linux GOARCH=arm64 go build -o ./dist/gctu_linux_arm64
