#!/bin/bash

for i in {1..24} ; do
	echo "package main" >> day"$i".go
	echo "package main" >> day"$i"_test.go
	touch puzzle-inputs/day"$i"-example1
	touch puzzle-inputs/day"$i"-input
done