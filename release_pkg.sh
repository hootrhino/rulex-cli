#! /bin/bash
set -e

RESPOSITORY="https://github.com/hootrhino"

# __build
create_pkg() {

    VERSION=$(git describe --tags --always --abbrev=0)
    echo "Create package: ${rulexc-$1-${VERSION}}"
    if [ "$1" == "x64windows" ]; then
        zip -r ../_release/rulexc-$1-${VERSION}.zip \
        ./rulexc.exe
        rm -rf ./rulexc.exe
    else
        zip -r ../_release/rulexc-$1-${VERSION}.zip \
        ./rulexc-*
        rm -rf ./rulexc-*
    fi
}

#
make_zip() {
    if [ -n $1 ]; then
        create_pkg $1
    else
        echo "Should have release target."
        exit 1
    fi

}

#
build_x64windows() {
    make windows
}

build_x64linux() {
    make x64linux
}

build_arm64linux() {
    make arm64
}

build_arm32linux() {
    make arm32
}

build_mips64linux() {
    make mips64
}

build_mips32linux() {
    make mips32
}
#------------------------------------------
cross_compile() {
    ARCHS=("x64windows" "x64linux" "arm64linux" "arm32linux")
    if [ ! -d "./_release/" ]; then
        mkdir -p ./_release/
    else
        rm -rf ./_release/
        mkdir -p ./_release/
    fi
    cd rulexc
    go get
    for arch in ${ARCHS[@]}; do
        echo -e "\033[34m [★] Compile target =>\033[43;34m ["$arch"]. \033[0m"
        if [[ "${arch}" == "x64windows" ]]; then
            # sudo apt install gcc-mingw-w64-x86-64 -y
            build_x64windows $arch
            make_zip $arch
            echo -e "\033[33m [√] Compile target => ["$arch"] Ok. \033[0m"
        fi
        if [[ "${arch}" == "x86linux" ]]; then
            build_x86linux $arch
            make_zip $arch
            echo -e "\033[33m [√] Compile target => ["$arch"] Ok. \033[0m"

        fi
        if [[ "${arch}" == "x64linux" ]]; then
            build_x64linux $arch
            make_zip $arch
            echo -e "\033[33m [√] Compile target => ["$arch"] Ok. \033[0m"

        fi
        if [[ "${arch}" == "arm64linux" ]]; then
            # sudo apt install gcc-arm-linux-gnueabi -y
            build_arm64linux $arch
            make_zip $arch
            echo -e "\033[33m [√] Compile target => ["$arch"] Ok. \033[0m"

        fi
        if [[ "${arch}" == "arm32linux" ]]; then
            # sudo apt install gcc-arm-linux-gnueabi -y
            build_arm32linux $arch
            make_zip $arch
            echo -e "\033[33m [√] Compile target => ["$arch"] Ok. \033[0m"
        fi
    done
    cd ..
}
#
# gen_changelog
#
gen_changelog() {
    PreviewVersion=$(git describe --tags --abbrev=0 $(git rev-list --tags --skip=1 --max-count=1))
    CurrentVersion=$(git describe --tags --abbrev=0)
    echo "----------------------------------------------------------------"
    echo -e "\033[34m [★] Change log between\033[44;34m ["${PreviewVersion}"] <--> [${CurrentVersion}]. \033[0m"
    echo "----------------------------------------------------------------"
    ChangeLog=$(git log ${PreviewVersion}..${CurrentVersion} --pretty=format:"%h %ad | %s")
    printf "${ChangeLog}\n"
    echo ${ChangeLog} > ./ChangeLog.txt
}

#
#-----------------------------------
#
init_env() {
    if [ ! -d "./__build/" ]; then
        mkdir -p ./__build/rulexc/
    else
        rm -rf ./__build/
        mkdir -p ./__build/rulexc/
    fi
}
#
# 检查是否安装了这些软件
#
check_cmd() {
    DEPS=("git" "jq" "gcc" "make")
    for dep in ${DEPS[@]}; do
        echo -e "\033[34m [*] Check dependcy command: $dep. \033[0m"
        if ! [ -x "$(command -v $dep)" ]; then
            echo -e "\033[31m |x| Error: $dep is not installed. \033[0m"
            exit 1
        else
            echo -e "\033[32m [√] $dep has been installed. \033[0m"
        fi
    done

}
#
#-----------------------------------
#
check_cmd
init_env
cp -r $(ls | egrep -v '^__build$') ./__build/rulexc
cd ./__build
cross_compile
gen_changelog
