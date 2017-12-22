cd /d %~dp0 ^
 && rsrc -manifest test.manifest -o rsrc.syso ^
 && go build -ldflags="-H windowsgui" -o test.exe ^
 && .\test.exe
