#!/bin/bash


EXTRACT_DATE="`date +%y%m%d`"

echo "DATE IS "${EXTRACT_DATE}

mkdir ${EXTRACT_DATE}
unzip ${EXTRACT_DATE}-VIC_ABR_Extract.zip -d ${EXTRACT_DATE}/
#cd ${EXTRACT_DATE}
gzip ./${EXTRACT_DATE}/VIC${EXTRACT_DATE}_ABR_Agency_Data.txt
aws s3 sync ${EXTRACT_DATE} s3://sbv-abr-etl/raw/${EXTRACT_DATE}
