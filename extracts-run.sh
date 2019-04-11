#!/usr/bin/env bash



AWS_PROFILE=sbv go run aws-query/query.go extracts.sql.tpl 6931
AWS_PROFILE=sbv go run aws-query/query.go extracts.sql.tpl 6932
AWS_PROFILE=sbv go run aws-query/query.go extracts.sql.tpl 7291



#https://sbv-abr-etl.s3.ap-southeast-2.amazonaws.com/custom-extract/6931/4e1495c1-3052-4052-8441-5ea3aebb0fe1.csv?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAIFEAKDMR4MJKRFLA%2F20180917%2Fap-southeast-2%2Fs3%2Faws4_request&X-Amz-Date=20180917T115236Z&X-Amz-Expires=86400&X-Amz-SignedHeaders=host&X-Amz-Signature=1994dc8b6c8e5ab3bb0b5c55b7eb325a76c409b326b519e83021d6152be20be4
#https://sbv-abr-etl.s3.ap-southeast-2.amazonaws.com/custom-extract/6932/331d8baf-97f3-41ac-af5a-d70b6e6dd44c.csv?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAIFEAKDMR4MJKRFLA%2F20180917%2Fap-southeast-2%2Fs3%2Faws4_request&X-Amz-Date=20180917T115617Z&X-Amz-Expires=86400&X-Amz-SignedHeaders=host&X-Amz-Signature=2440e13320e3a6fb4ae6b2c7c1e8d4e728c1e0c84bbd20403084d41e11dd2b6c
#https://sbv-abr-etl.s3.ap-southeast-2.amazonaws.com/custom-extract/7291/34602e5a-d298-4385-815a-ddecc1d43c49.csv?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAIFEAKDMR4MJKRFLA%2F20180917%2Fap-southeast-2%2Fs3%2Faws4_request&X-Amz-Date=20180917T115547Z&X-Amz-Expires=86400&X-Amz-SignedHeaders=host&X-Amz-Signature=927580b9d9a8c5db0e2a034f0da73e7c8ebf88df93278c95e4b847a47cec3fb6


