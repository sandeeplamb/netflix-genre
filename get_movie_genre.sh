#!/bin/bash

for i in seq 2
do
    go run netflix-genre.go
    python netflix-genre.py
done