all: proxyclean


proxyclean: clean
	go build proxyclean.go

clean:
	rm -f proxyclean

