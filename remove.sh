#!/bin/bash
cd $HOME/development
goenv uninstall --force 1.13.0 
brew uninstall goenv
brew uninstall dep
rm -fr go
