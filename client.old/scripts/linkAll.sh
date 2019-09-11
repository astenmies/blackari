#!/bin/bash
# This is myProject/scripts/linkAll.sh
# List all your local libraries in this array
libs=( "library01" "library02" )

appDir=$(pwd) # dir of where script is called
packDir="/absolute/path/to/packages"

for i in "${libs[@]}"
do
    if [ ! -d "$appDir/node_modules/$i" ]; then
        ln -s "$packDir/$i" "$appDir/node_modules"
        echo "$i -> Synced!"
    else
        echo "$i -> Was already synced"
    fi
done