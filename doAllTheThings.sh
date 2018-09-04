#!/bin/bash

# run this file after running fetchVicExtract.sh
#
# for testing, if you have already run this script today then you should delete the directory that it created (e.g. 20180903)
# in S3 you should also delete the /raw/<date>/ and /FACT/AGENCY/date=<date>/ directories

EXTRACT_DATE="`date +%Y%m%d`"
EXTRACT_DATE_SHORT="`date +%y%m%d`"
DATE_PARTITION="`date +%Y-%m-%d`"

# creates a local directory and extracts the data dump files
mkdir ${EXTRACT_DATE}
unzip ${EXTRACT_DATE}-VIC_ABR_Extract.zip -d ${EXTRACT_DATE}/

cd ${EXTRACT_DATE}

# upload all files excluding the main Agency file to /raw/<date> on S3
aws s3 sync . s3://sbv-abr-etl/raw/${EXTRACT_DATE} --exclude "*" --include "*.txt" --exclude "VIC${EXTRACT_DATE}_ABR_Agency_Data.txt"

# Compare previous with current Agency file generate combine file + Changing Dimensions partitions
diff-abr ../previous-combine.txt VIC${EXTRACT_DATE_SHORT}_ABR_Agency_Data.txt combined.txt ${DATE_PARTITION}

# Replace agency data with new file -- TODO: convert to parquet
gzip combined.txt
aws s3 cp combined.txt.gz s3://sbv-abr-etl/FACT/AGENCY/combined.txt.gz

# Compress files
gzip OrgNameChange.txt
gzip NameChange.txt
gzip TradingNameChange.txt
gzip SONAddressChange.txt
gzip BusAddressChange.txt
gzip EmailChange.txt
gzip IndustryChange.txt

# Add new change partitions
aws s3 cp OrgNameChange.txt.gz s3://sbv-abr-etl/DIMENSION/OrgNameChange/date=${DATE_PARTITION}/OrgNameChange.txt.gz
aws s3 cp NameChange.txt.gz s3://sbv-abr-etl/DIMENSION/NameChange/date=${DATE_PARTITION}/NameChange.txt.gz
aws s3 cp TradingNameChange.txt.gz s3://sbv-abr-etl/DIMENSION/TradingNameChange/date=${DATE_PARTITION}/TradingNameChange.txt.gz
aws s3 cp SONAddressChange.txt.gz s3://sbv-abr-etl/DIMENSION/SONAddressChange/date=${DATE_PARTITION}/SONAddressChange.txt.gz
aws s3 cp BusAddressChange.txt.gz s3://sbv-abr-etl/DIMENSION/BusAddressChange/date=${DATE_PARTITION}/BusAddressChange.txt.gz
aws s3 cp EmailChange.txt.gz s3://sbv-abr-etl/DIMENSION/EmailChange/date=${DATE_PARTITION}/EmailChange.txt.gz
aws s3 cp IndustryChange.txt.gz s3://sbv-abr-etl/DIMENSION/IndustryChange/date=${DATE_PARTITION}/IndustryChange.txt.gz

# Copy current combined into parent folder for next run
cp combined.txt ../previous-combine.txt