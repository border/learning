include $(GOROOT)/src/Make.inc

GOFMT=gofmt -spaces=true -tabindent=false -tabwidth=4

all:
	$(GC) webtest.go 
	$(LD) -o webtest.out  webtest.$O

format:
	$(GOFMT) -w webtest.go

clean:
	rm -rf *.8 *.o *.out *.6


