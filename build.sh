DEBUG_FLAG=$1

set GOOS=windows
set GOARCH=amd64
CLIENT_EXE=HyperK.exe

if [[ ${DEBUG_FLAG} == "-debug" ]]; then
    echo "[+] Building ${CLIENT_EXE}... (debug)"
    go build -o ${CLIENT_EXE} ./cmd/client
else
    echo "[+] Building ${CLIENT_EXE}..."
    garble -literals -tiny -seed=random build ./cmd/client
    mv client.exe ${CLIENT_EXE}
fi

set GOOS=windows
set GOARCH=amd64
SERVER_EXE=HyperK-server.exe
echo "[+] Building ${SERVER_EXE}..."
go build -o ${SERVER_EXE} ./cmd/server

echo "[+] Done !"
echo "    ${CLIENT_EXE} ($(du -sh ./${CLIENT_EXE} | cut -f1))"
echo "    ${SERVER_EXE} ($(du -sh ./${SERVER_EXE} | cut -f1))"
