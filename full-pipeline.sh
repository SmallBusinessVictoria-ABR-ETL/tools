#!/usr/bash

# Change to Tools Dir
cd `dirname $0`

# Fetch VIC Records from ABR
./fetchVicExtract.sh

# Diff, Convert to Parquet and Push to S3
./doAllTheThings.sh

python csvToParquet.py

curl -X POST --data-urlencode "payload={\"channel\": \"#abr-etl\", \"username\": \"ABR_ETL_Tools\", \"text\": \"ABR ETL Completed.\", \"icon_emoji\": \":ghost:\"}" https://hooks.slack.com/services/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

shutdown -t now