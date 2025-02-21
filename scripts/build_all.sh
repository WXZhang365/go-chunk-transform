GOOS=windows GOARCH=amd64 go build -o ./dist/gctu.exe
zip ./dist/gctu_win_amd64.zip ./dist/gctu.exe
rm ./dist/gctu.exe
GOOS=linux GOARCH=amd64 go build -o ./dist/gctu
zip ./dist/gctu_linux_amd64.zip ./dist/gctu
rm ./dist/gctu
GOOS=linux GOARCH=arm64 go build -o ./dist/gctu
zip ./dist/gctu_linux_arm64.zip ./dist/gctu
rm ./dist/gctu
