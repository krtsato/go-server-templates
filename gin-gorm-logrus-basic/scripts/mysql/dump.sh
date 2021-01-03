#!/bin/sh -eu

# example: ./scripts/mysql/dump.sh db-name table-name
# table-name is optional. Empty string is a default value.
table=
suffix=_
readonly DUMP_DIR=/var/dump

if [ -z $1 ]; then
    echo "Enter the DB name.\nYou can also fill in the table name if you want."
    exit 1
fi
suffix="${suffix}$1"

if [ ! -z ${2-} ]; then
  suffix="${suffix}_$2"
  table=$2
fi

mysqldump -u root -ppass_root -h mysql -P 3306 --protocol=tcp \
  --extended-insert \
  --complete-insert \
  --compress \
	--default-character-set=binary \
	--single-transaction \
	--quick \
	$1 $table > ${DUMP_DIR}/$(date "+%Y%m%d")${suffix}.dump

echo "Dumped your database into the container ${DUMP_DIR}."
