include $(GOROOT)/src/Make.inc

GOFMT=gofmt -spaces=true -tabindent=false -tabwidth=4

all:
	$(GC) jsontest.go 
	$(LD) -o jsontest.out  jsontest.$O

format:
	$(GOFMT) -w jsontest.go

clean:
	rm -rf *.8 *.o *.out *.6

# run :go-proxy
# 	./go-proxy

