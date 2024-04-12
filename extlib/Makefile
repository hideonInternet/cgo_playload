all: libcomm.a libcomm.so

libcomm.a: static.o
	ar rcs $@ $^

libcomm.so: static.o
	gcc -shared -fPIC -o $@ $^

%.o:%.c
	$(CC) $(CXXFLAGS) -Wall -O2 -c $^