#!/bin/bash

aws s3 cp s3://sbv-abr-etl/disabled ./disabled

if grep -q disabled ./disabled; then
  echo "ETL Disabled"
  exit
fi

# Change to Tools Dir
cd `dirname $0`

# Fetch VIC Records from ABR
bash ./fetchVicExtract.sh

# Diff, Convert to Parquet and Push to S3
bash ./doAllTheThings.sh

# Update related tables + Transport location
python csvToParquet.py

# Notify users via Slack
curl -X POST --data-urlencode "payload={\"channel\": \"#abr-etl\", \"username\": \"ABR_ETL_Tools\", \"text\": \"ABR ETL Completed.\", \"icon_emoji\": \":ghost:\"}" https://hooks.slack.com/services/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

# Shutdown
shutdown -t now