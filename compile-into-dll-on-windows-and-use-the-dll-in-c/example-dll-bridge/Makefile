..\example.dll: example.dll
	cp -f example.dll ..\example.dll

example.dll: export.c ..\example-go-lib\example.a ..\example-go-lib\example.h
	gcc --shared -pthread -o example.dll export.c ..\example-go-lib\example.a -lWinMM -lntdll -lWS2_32


