build:
	env GOOS=linux GOARCH=arm GOARM=6 go build -o bin/red-time main.go

# needs to use an argument instead of arcturus
copy: build
	scp bin/red-time $$REDTIMEHOST:~/.

clean:
	rm -rf bin

run: copy
	ssh $$REDTIMEHOST ./red-time