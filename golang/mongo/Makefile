include $(GOROOT)/src/Make.inc

all:mongomap.go
	$(GC) mongomap.go 
	$(LD) -o mongomap.out  mongomap.$O
	

clean:
	rm -rf *.8 *.o *.out *.6

# run :go-proxy
# 	./go-proxy

