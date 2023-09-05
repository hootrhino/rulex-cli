#!/bin/bash

install(){
    local source_dir="$PWD"
    chmod +x $source_dir/rulexc
    cp "$source_dir/rulexc" "/usr/local"
    if [ $? -eq 0 ]; then
        echo "rulexc installed."
    else
        echo "Failed to install the rulexc."
    fi
    exit 0
}

uninstall(){
    local working_directory="/usr/local"
    rm $working_directory/rulexc
    echo "rulexc has been uninstalled."
}

#
#
#
main(){
    case "$1" in
        "install" |  "uninstall")
            $1
        ;;
        *)
            echo "Invalid command: $1"
            echo "Usage: $0 <install|uninstall>"
            exit 1
        ;;
    esac
    exit 0
}
#===========================================
# main
#===========================================
if [ "$(id -u)" != "0" ]; then
    echo "This script must be run as root"
    exit 1
else
    main $1
fi