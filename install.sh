#!/bin/bash
binary_name=btb
goos=$(uname)
version=0.0.1
if [[ "$OSTYPE" == "linux-gnu"* ]]; then
  goos=linux64
elif [[ "$OSTYPE" == "darwin"* ]]; then
  goos=mac64
elif [[ "$OSTYPE" == "msys" ]]; then
  goos=win64
else
  echo "Error: The current os is not supported at this time" 1>&2
  exit 1
fi


file_name=btb-${goos}-${version}.tar.gz

url=https://github.com/crossoverJie/btb/releases/download/v${version}/${file_name}
echo "Download url:${url}"

curl "$url" -OL --progress --retry 2 2>&1

tar -xvf "${file_name}"

if [[ "${goos}" == "win64" ]]; then
  echo "btb download success."
  exit 0
fi
cp $binary_name /usr/local/bin/

rm -rf $binary_name
rm -rf $file_name

echo "btb install success."