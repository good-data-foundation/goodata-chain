#!/bin/bash

ROOTDIR=$(dirname $(realpath $0))
cd $ROOTDIR

trap 'echo "start.sh: SIGTERM received; Let child processes exit by themselves"' SIGTERM

D="$ROOTDIR/datadir"
if [ ! -d "$D/geth" ]; then
    $ROOTDIR/bin/geth --datadir $D init $ROOTDIR/../genesis.json
fi

exec $ROOTDIR/bin/geth --datadir $D --ws --rpc --rpcaddr 0.0.0.0 --wsaddr 0.0.0.0 \
    --rpccorsdomain '*' --wsorigins '*' --mine --nousb --rpcvhosts '*' \
    --networkid 15 --verbosity 3 \
    --miner.etherbase 0x614CD1D8A1CCe3bb5292557F5498298715F2352f
