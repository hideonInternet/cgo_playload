shared_lib:
	go build
	cp ../extlib/libcomm.so .

run:
	LD_LIBRARY_PATH=./ ./shared
	
clean:
	- rm shared
	- rm libcomm.so