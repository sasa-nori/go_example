#!/bin/bash
cd $HOME/development
goenv uninstall 1.13.0
brew uninstall goenv
brew uninstall dep
rm -f go