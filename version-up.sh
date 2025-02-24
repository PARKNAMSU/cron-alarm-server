#!/bin/bash

export PROJECT_VERSION=$(echo $(sed -n -e '1,1p' VERSION) | cut -d ' ' -f2)
export READ_ME=$(sed -n -e '2,$p' VERSION) 

echo "previous ${PROJECT_VERSION}"

export MAJOR=$(echo $PROJECT_VERSION | cut -d '.' -f1)
export MINOR=$(echo $PROJECT_VERSION | cut -d '.' -f2)
export PATCH=$(echo $PROJECT_VERSION | cut -d '.' -f3)

export COMMIT_EXPLAIN=""



if [ "$1" == "" ] || [ "$1" == "--patch" ]; then
    PATCH="$(($PATCH + 1))"
elif [ "$1" == "--minor" ]; then
    MINOR="$(($MINOR + 1))"
    PATCH="0"
elif [ "$1" == "--major" ]; then
    MAJOR="$(($MAJOR + 1))"
    PATCH="0"
    MINOR="0"
else
    PATCH="$(($PATCH + 1))"
    COMMIT_EXPLAIN="$1"
fi

while [ $# -ne 0 ]; do
    if [ "$COMMIT_EXPLAIN" == "" ]; then
        COMMIT_EXPLAIN="$2"
    else
        COMMIT_EXPLAIN="$COMMIT_EXPLAIN $2"
    fi
    shift
done

PROJECT_VERSION="VERSION ${MAJOR}.${MINOR}.${PATCH}"

echo "now ${PROJECT_VERSION}"

echo $PROJECT_VERSION > ./VERSION

cat << EOF >> ./VERSION
$READ_ME
EOF

git add .
git commit -m "[${PROJECT_VERSION}]
${COMMIT_EXPLAIN}"