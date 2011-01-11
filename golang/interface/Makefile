include $(GOROOT)/src/Make.inc

GOFMT=gofmt -spaces=true -tabindent=false -tabwidth=4

all:
	$(GC) interfacetest.go 
	$(LD) -o interfacetest.out  interfacetest.$O

format:
	$(GOFMT) -w interfacetest.go

clean:
	rm -rf *.8 *.o *.out *.6

# run :go-proxy
# 	./go-proxy

