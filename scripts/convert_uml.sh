#!/bin/sh
UMLFILE=docs/class
plantuml -tsvg "$UMLFILE.plantuml"
inkscape "$UMLFILE.svg" --export-pdf="$UMLFILE.pdf"
