include $(GOROOT)/src/Make.inc

all:
	$(GC) numbers.go
	$(GC) filestuts.go
	$(GC) test.go
	$(LD) -o test.out test.$O

clean:
	rm -rf *.6 *.8 *.o *.out
	


