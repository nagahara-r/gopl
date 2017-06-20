# DLに使用したシェルスクリプト
for FIGURE in `seq 1 1000`
do
    echo "${FIGURE}"
    curl https://xkcd.com/${FIGURE}/info.0.json -o ${FIGURE}.json
done
