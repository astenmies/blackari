#! /bin/bash

if [ ! -f $GOBIN/go-bindata ]; then
    echo "go-bindata not found. Downloading it for you..."
    go get -u github.com/jteeuwen/go-bindata/...
    if [ $? -eq 0 ]; then
        echo "Successfully downloaded go-bindata!"
    else
        echo "Failed downloading go-bindata. Please get it at http://github.com/jteeuwen/go-bindata"
        exit
    fi
fi
echo "Generating schema..."
go-bindata -ignore=\.go -pkg=schema -o=schema/bindata.go schema/...

if [ $? -eq 0 ]; then
    echo "Successfully generated schema!"
else
    echo "Failed generating schema."
fi


if [ ! -f $GOBIN/realize ]; then
    
    echo "realize not found. Downloading it for you..."
    go get -u github.com/oxequa/realize/...
    
    if [ $? -eq 0 ]; then
        echo "Successfully downloaded realize!"
    else
        echo "Failed downloading realize. Please get it at http://github.com/oxequa/realize"
        exit
    fi

fi

if [ ! -f _config/global.json ]; then
    echo "Creating _config/global.json"
    cp _config/global.example.json _config/global.json
    echo "Done."
    echo ">>> Please open _config/global.json and update it with your config."
fi

echo "Starting app with realize..."
realize start --run main.go
