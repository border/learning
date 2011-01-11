include $(GOROOT)/src/Make.inc

GOFMT=gofmt -spaces=true -tabindent=false -tabwidth=4

all:
	$(GC) mongomap.go 
	$(GC) mongotyped.go 
	$(GC) mongofind.go 
	$(LD) -o mongomap.out  mongomap.$O
	$(LD) -o mongotyped.out  mongotyped.$O
	$(LD) -o mongofind.out  mongofind.$O
	

format:
	$(GOFMT) -w mongomap.go
	$(GOFMT) -w mongotyped.go
	$(GOFMT) -w mongofind.go


clean:
	rm -rf *.8 *.o *.out *.6

# run :go-proxy
# 	./go-proxy

