BINARY = rmc

build:
	GOOS=linux go build -ldflags "-s -w" -o $(BINARY)

clean:
	rm -f $(BINARY) $(BINARY).gz

package:
	gzip -c $(BINARY) > $(BINARY).gz
