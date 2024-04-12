all: code_c file_c file_cpp_bad file_cpp_good

code_c:
	make -C code_c code_c

file_c:
	make -C file_c file_c

file_cpp_bad:
	- make -C file_cpp_bad file_cpp_bad

file_cpp_good:
	make -C file_cpp_good file_cpp_good

run:
	make -C code_c run
	make -C file_c run
	make -C file_cpp_good run

clean:
	make -C code_c clean
	make -C file_c clean
	make -C file_cpp_good clean

.PHONY:. code_c file_c file_cpp_bad file_cpp_good