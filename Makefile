module=code extlib shared static

build:
	@for m in $(module); do\
		make -C $$m;\
	done

run:
	@for m in $(module); do\
		if [ $$m != extlib ]; then \
			make -C $$m run;\
		fi \
	done

clean:
	@for m in $(module); do\
		make -C $$m clean;\
	done