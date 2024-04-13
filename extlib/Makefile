all: libcomm.a libcomm.so

libcomm.a: static.o
	ar rcs $@ $^

libcomm.so: static.o
	gcc -shared -fPIC -o $@ $^

clean:
	rm *.o
	rm *.a
	rm *.so

%.o:%.c
	$(CC) $(CXXFLAGS) -Wall -O2 -c $^