module=c cgo golang
P=20

build:
	@for m in $(module); do\
		make -C $$m;\
	done

clean:
	@for m in $(module); do\
		make -C $$m clean;\
	done

run:
	@for m in $(module); do\
		make -C $$m run NUM=$(P);\
	done

M=Call
test:
	go test ./... -bench=$M -count=5 -timeout 20m -v -benchmem
