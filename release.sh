#!/usr/bin/env sh
set -e

REPO=$(cd $(dirname $0); pwd)
RELEASE=""
PRELEASE=""
LASTTAG=`git describe --match "v[0-9]\.[0-9]\.[0-9]" --abbrev=0 --tags $(git rev-list --tags --max-count=1)`

checkParam () {
  echo "👀 Checking semver format"

  if [ $# -ne 1 ]; then
    echo "❌ This release script requires a single argument corresponding to the semver to be released. See semver.org"
    exit 1
  fi

  GREP="grep"
  if [ -x "$(command -v ggrep)" ]; then
    GREP="ggrep"
  fi

  semver=$(echo "$1" | $GREP -P '^v(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)')

  if [ $? -ne 0 ]; then
    echo "❌ Not valid semver format. See semver.org"
    exit 1
  fi
}
rollback () {
  cd $REPO
  checkParam $LASTTAG
  git tag -d "$LASTTAG"
  git push origin :refs/tags/"$LASTTAG"
}
release () {
  cd $REPO
  
  checkParam $1
  
  if [ "$LASTTAG" = "$1" ]; then
    rollback
  fi

#  echo "🧼  Tidying up go modules"
#  go mod tidy

  echo "🐑 Creating a new commit for the new release"
  git commit --allow-empty -am "chore: version $semver"
  git tag "$1"
  git push
  git push --tags origin

  echo "📦 Done! $semver released."
}

while getopts "br:d" o; do
  case "${o}" in
    r)
      RELEASE=${OPTARG}
      ;;
    b)
      PRELEASE="true"
      ;;
    *)
      usage
      ;;
  esac
done
shift $((OPTIND-1))
if [ "$RELEASE" != "" ]; then
  release $RELEASE
fi

if [ "$PRELEASE" = "true" ]; then
  rollback
fi
