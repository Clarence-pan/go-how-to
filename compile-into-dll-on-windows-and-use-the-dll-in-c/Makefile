SUBDIRS =example-dll-bridge example-go-lib test

.PHONY: subdirs ${SUBDIRS}

default: subdirs
	./test/test.exe && echo Test passed.

subdirs: ${SUBDIRS}

${SUBDIRS}:
	${MAKE} -C $@

test: example-dll-bridge
example-dll-bridge: example-go-lib

clean:
	rm -f example.dll .\example-dll-bridge\example.dll .\example-go-lib\example.a .\example-go-lib\example.h .\test\test.exe

rebuild: clean default
