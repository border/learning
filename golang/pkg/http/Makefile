include $(GOROOT)/src/Make.inc

GOFMT=gofmt -spaces=true -tabindent=false -tabwidth=4

all: httpGet httpGet2

httpGet:
	$(GC) httpGet.go 
	$(LD) -o httpGet.out  httpGet.$O

httpGet2:
	$(GC) httpGet2.go 
	$(LD) -o httpGet2.out  httpGet2.$O


format:
	$(GOFMT) -w httpGet.go httpGet2.go

clean:
	rm -rf *.8 *.o *.out *.6


