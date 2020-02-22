#!/usr/bin/env bash

CONN=${CONN:-jdbc:mysql://test-mysql-svc:3306/tpcc?useSSL=false&useServerPrepStmts=true&useConfigs=maxPerformance&allowPublicKeyRetrieval=true}
WAREHOUSES=${WAREHOUSES:-2}
LOADWORKERS=${LOADWORKERS:-3}
TERMINALS=${TERMINALS:-3}
USER=${USER:-test}
PASSWORD=${PASSWORD:-test}
RUN=""
if [[ $1 = "" ]]; then
  RUN="all"
else
  RUN=$1
fi

cat <<EOF >./run/props.mysql
db=mysql
driver=com.mysql.cj.jdbc.Driver

//To run specified transactions per terminal- runMins must equal zero
runTxnsPerTerminal=0
//To run for specified minutes- runTxnsPerTerminal must equal zero
runMins=10
//Number of total transactions per minute
limitTxnsPerMin=0

//Set to true to run in 4.x compatible mode. Set to false to use the
//entire configured database evenly.
terminalWarehouseFixed=true

//The following five values must add up to 100
//The default percentages of 45, 43, 4, 4 & 4 match the TPC-C spec
newOrderWeight=45
paymentWeight=43
orderStatusWeight=4
deliveryWeight=4
stockLevelWeight=4

// Directory name to create for collecting detailed result data.
// Comment this out to suppress.
resultDirectory=my_result_%tY-%tm-%td_%tH%tM%tS
//osCollectorScript=./misc/os_collector_linux.py
//osCollectorInterval=1
//osCollectorSSHAddr=user@dbhost
//osCollectorDevices=net_eth0 blk_sda
EOF
printf "conn=$CONN
warehouses=$WAREHOUSES
loadWorkers=$LOADWORKERS
terminals=$TERMINALS
user=$USER
password=$PASSWORD
" >>./run/props.mysql
cd ./run || exit
if [ $RUN = "all" ] || [ $RUN = "createTable" ]; then
  ./runSQL.sh props.mysql sql.mysql/tableCreates.sql && ./runSQL.sh props.mysql sql.mysql/indexCreates.sql
fi
if [ $RUN = "all" ] || [ $RUN = "naiveLoader" ]; then
  ./runLoader.sh props.mysql
fi
if [ $RUN = "all" ] || [ $RUN = "benchmark" ]; then
  ./runBenchmark.sh props.mysql >/test.log
fi
