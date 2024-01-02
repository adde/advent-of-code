include .env

input:
	mkdir -p ./$(year)/$$(printf "%02d" "$(day)"); \
	curl --cookie "session=${AOC_COOKIE}" https://adventofcode.com/$(year)/day/$(day)/input > ./$(year)/$$(printf "%02d" "$(day)")/input.txt