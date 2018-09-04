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

# download the previous data dump to ./previous and gunzip it
mkdir previous
aws s3 sync s3://sbv-abr-etl/previous/ ./previous/
gunzip previous/full.txt.gz

# rename the current data dump to full.txt and then run the diff program against the previous one
mv VIC${EXTRACT_DATE_SHORT}_ABR_Agency_Data.txt full.txt
diff-abr full.txt ./previous/full.txt new.txt updated.txt


# gzip the 3 new files
gzip full.txt new.txt updated.txt

# copy the new files to the right location on S3
aws s3 cp full.txt.gz s3://sbv-abr-etl/FACT/AGENCY/date=${DATE_PARTITION}/state=full/
aws s3 cp updated.txt.gz s3://sbv-abr-etl/FACT/AGENCY/date=${DATE_PARTITION}/state=updated/
aws s3 cp new.txt.gz s3://sbv-abr-etl/FACT/AGENCY/date=${DATE_PARTITION}/state=new/

# create a date.txt file with the current date and push date.txt and the full.txt.gz through to S3/previous
echo ${EXTRACT_DATE} > date.txt
aws s3 rm s3://sbv-abr-etl/previous/ --recursive
aws s3 sync . s3://sbv-abr-etl/previous --exclude "*" --include "full.txt.gz" --include "date.txt"

