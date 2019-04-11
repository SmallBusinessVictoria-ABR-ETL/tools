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

export EXTRACT_DATE="`date +%Y%m%d`"
export EXTRACT_DATE_SHORT="`date +%Y%m%d` | tail -c 7"
export SBV_BUCKET="sbv-abr-etl"

mkdir ${EXTRACT_DATE}
cd ${EXTRACT_DATE}

unzip ../${EXTRACT_DATE}-VIC_ABR_Extract.zip

aws s3 cp VIC${EXTRACT_DATE_SHORT}_ABR_Agency_Data.txt s3://${SBV_BUCKET}/DATA/Agency_Data/Date/${EXTRACT_DATE}/ABR_Agency_Data.txt
aws s3 cp VIC${EXTRACT_DATE_SHORT}_ABR_ACNC.txt s3://${SBV_BUCKET}/DATA/ACNC/Date/${EXTRACT_DATE}/ABR_ACNC.txt
aws s3 cp VIC${EXTRACT_DATE_SHORT}_ABR_Associates.txt s3://${SBV_BUCKET}/DATA/Associates/Date/${EXTRACT_DATE}/ABR_Associates.txt
aws s3 cp VIC${EXTRACT_DATE_SHORT}_ABR_Businesslocation.txt s3://${SBV_BUCKET}/DATA/Businesslocation/Date/${EXTRACT_DATE}/ABR_Businesslocation.txt
aws s3 cp VIC${EXTRACT_DATE_SHORT}_ABR_Businessname.txt s3://${SBV_BUCKET}/DATA/Businessname/Date/${EXTRACT_DATE}/ABR_Businessname.txt
aws s3 cp VIC${EXTRACT_DATE_SHORT}_ABR_Funds.txt s3://${SBV_BUCKET}/DATA/Funds/Date/${EXTRACT_DATE}/ABR_Funds.txt
aws s3 cp VIC${EXTRACT_DATE_SHORT}_ABR_Othtrdnames.txt s3://${SBV_BUCKET}/DATA/Othtrdnames/Date/${EXTRACT_DATE}/ABR_Othtrdnames.txt
aws s3 cp VIC${EXTRACT_DATE_SHORT}_ABR_Replacedabn.txt s3://${SBV_BUCKET}/DATA/Replacedabn/Date/${EXTRACT_DATE}/ABR_Replacedabn.txt
aws s3 cp VIC${EXTRACT_DATE_SHORT}_ABR_Summary.txt s3://${SBV_BUCKET}/DATA/Summary/Date/${EXTRACT_DATE}/ABR_Summary.txt

# Diff, Convert to Parquet and Push to S3
# bash ./doAllTheThings.sh

# Update related tables + Transport location
# python csvToParquet.py

# Notify users via Slack
# curl -X POST --data-urlencode "payload={\"channel\": \"#abr-etl\", \"username\": \"ABR_ETL_Tools\", \"text\": \"ABR ETL Completed.\", \"icon_emoji\": \":ghost:\"}" https://hooks.slack.com/services/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

# Shutdown
# shutdown -t now