include .env

init:
	mkdir -p ./$(year)/$$(printf "%02d" "$(day)"); \
	curl --cookie "session=${AOC_COOKIE}" https://adventofcode.com/$(year)/day/$(day)/input > ./$(year)/$$(printf "%02d" "$(day)")/input.txt
	truncate -s -1 ./$(year)/$$(printf "%02d" "$(day)")/input.txt
	cp ./templates/basic/basic.go ./$(year)/$$(printf "%02d" "$(day)")/main.go