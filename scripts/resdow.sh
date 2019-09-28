#!/usr/bin/env bash

######### resdow -- Resource Downloader ##########
# check for curl
if ! [ -x "$(command -v curl)" ]; then
  echo 'cannot download resources, curl is not installed' >&2
  exit 1
fi

# ensure dir and expand path
dir_resolve() {
  mkdir -p "$1" 2>/dev/null
  cd "$1" 2>/dev/null || exit $?
  echo "$(pwd -P)" # output full, link-resolved path
}

# resource links
resources=(
  "https://raw.githubusercontent.com/smart-tool/smart/master/data/englishTexts/bible.txt"
  "https://raw.githubusercontent.com/smart-tool/smart/master/data/englishTexts/world192.txt"
  "http://pizzachili.dcc.uchile.cl/texts/nlang/english.50MB.gz"
  "http://pizzachili.dcc.uchile.cl/texts/nlang/english.200MB.gz"
  "http://pizzachili.dcc.uchile.cl/texts/nlang/english.1024MB.gz"
)

cur_dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
cur_dir+="/../corpora"
output=$(dir_resolve "${cur_dir}")
cd "$output" || exit 1

# download resources
for res in "${resources[@]}"; do
  curl -OJL "$res"
done

# extract compressed files
find . -name '*.gz' -execdir gunzip '{}' \;