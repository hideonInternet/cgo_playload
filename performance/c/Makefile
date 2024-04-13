INC=-I../..//extlib
LIB=../../extlib/libcomm.a
NUM=20

c: test_c.o $(LIB)
	gcc -o $@ $^ $(INC)

%.o:%.c
	gcc -Wall -O2 -c $^ $(INC)

clean:
	rm *.o
	rm c

run:
	./c $(NUM)

.PHONY: c