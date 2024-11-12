#!/bin/bash

# this script takes 2 arguments: the latex file to compile, and the the place to compile it to
# no special care has been taken for error handling, this script is designed to be run correctly by the server

set -e

inputfile="${1%.*}"
inputdir=$(dirname "$inputfile")
outputdir=$2

# compile cv to pdf privately, then move to public once completed
latexmk -output-directory="$inputdir" "$inputfile.tex" && latexmk -c -output-directory="$inputdir" "$inputfile.tex"
mv "$inputfile.pdf" "$outputdir"

# compile cv to html privately, then move to public once completed
pandoc -f latex -t html -o "$inputfile.html" "$inputfile.tex"
mv "$inputfile.html" "$outputdir"


