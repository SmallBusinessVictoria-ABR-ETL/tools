from pyspark import SparkContext
from pyspark.sql import SQLContext
from pyspark.sql.types import *
from pyspark.sql.functions import *
from decimal import Decimal
import sys
import os.path
import subprocess



if __name__ == "__main__":

    def execCmd(bashCommand):
        process = subprocess.Popen(bashCommand, shell=True, stdout=subprocess.PIPE)
        output, error = process.communicate()

        if error:
            print("ERROR: %s" % error)


    if len(sys.argv) != 3:
        print "Two argument required"
        exit()
    else:
        directory = sys.argv[1]
        date = sys.argv[2]

        sc = SparkContext(appName="csvToParquet")
        sqlContext = SQLContext(sc)




        #
        # ACNC
        #

        filename = "./%s/VIC%s_ABR_ACNC.txt" % (directory, date)

        if os.path.isfile(filename):
            execCmd("rm -rf ACNC")

            schema = StructType([
                StructField("PID", LongType(), True),
                StructField("Ent_Typ_Cd", StringType(), True),
                StructField("Sts_Typ_Cd", StringType(), True),
                StructField("Regn_Dt", StringType(), True)
            ])

            def parse(line):
                items = line.split("\t")
                return (long(items[0]), items[1], items[2], items[3])

            rdd = sc.textFile(filename, use_unicode=False).map(parse)
            df = sqlContext.createDataFrame(rdd, schema)
            df.write.parquet('./ACNC')

            execCmd("aws s3 rm s3://sbv-abr-etl/DIMENSION/ACNC/ --recursive")
            execCmd("aws s3 sync ./ACNC/ s3://sbv-abr-etl/DIMENSION/ACNC/ --exclude '*' --include '*.parquet'")
            execCmd("aws s3 sync ./ACNC/ s3://sbv-abr-etl/DIMENSION_old/%s/ACNC/ --exclude '*' --include '*.parquet'" % date)

            print("%s parsed and uploaded to S3" % filename)
        else:
            print("File does not exist: %s" % filename)


        #
        # Associates
        #

        filename = "./%s/VIC%s_ABR_Associates.txt" % (directory, date)

        if os.path.isfile(filename):
            execCmd("rm -rf ASSOCIATES")

            schema = StructType([
                StructField("PID", LongType(), True),
                StructField("Rltnshp_Cd", StringType(), True),
                StructField("Assoc_Org_Nm", StringType(), True),
                StructField("Assoc_Titl_Cd", StringType(), True),
                StructField("Assoc_Gvn_Nm", StringType(), True),
                StructField("Assoc_Othr_Gvn_Nms", StringType(), True),
                StructField("Assoc_Fmly_Nm", StringType(), True),
                StructField("Assoc_Nm_Sufx_Cd", StringType(), True)
            ])

            def parse(line):
                items = line.split("\t")
                return (long(items[0]), items[1], items[2], items[3], items[4], items[5], items[6], items[7])

            rdd = sc.textFile(filename, use_unicode=False).map(parse)
            df = sqlContext.createDataFrame(rdd, schema)
            df.write.parquet('./ASSOCIATES')

            execCmd("aws s3 rm s3://sbv-abr-etl/DIMENSION/ASSOCIATES/ --recursive")
            execCmd("aws s3 sync ./ASSOCIATES/ s3://sbv-abr-etl/DIMENSION/ASSOCIATES/ --exclude '*' --include '*.parquet'")
            execCmd("aws s3 sync ./ASSOCIATES/ s3://sbv-abr-etl/DIMENSION_old/%s/ASSOCIATES/ --exclude '*' --include '*.parquet'" % date)

            print("%s parsed and uploaded to S3" % filename)
        else:
            print("File does not exist: %s" % filename)



        #
        # Business name
        #

        filename = "./%s/VIC%s_ABR_Businessname.txt" % (directory, date)

        if os.path.isfile(filename):
            execCmd("rm -rf BUSINESSNAME")

            schema = StructType([
                StructField("PID", LongType(), True),
                StructField("Bus_Nm", StringType(), True)
            ])

            def parse(line):
                items = line.split("\t")
                return (long(items[0]), items[1])

            rdd = sc.textFile(filename, use_unicode=False).map(parse)
            df = sqlContext.createDataFrame(rdd, schema)
            df.write.parquet('./BUSINESSNAME')

            execCmd("aws s3 rm s3://sbv-abr-etl/DIMENSION/BUSINESSNAME/ --recursive")
            execCmd("aws s3 sync ./BUSINESSNAME/ s3://sbv-abr-etl/DIMENSION/BUSINESSNAME/ --exclude '*' --include '*.parquet'")
            execCmd("aws s3 sync ./BUSINESSNAME/ s3://sbv-abr-etl/DIMENSION_old/%s/BUSINESSNAME/ --exclude '*' --include '*.parquet'" % date)

            print("%s parsed and uploaded to S3" % filename)
        else:
            print("File does not exist: %s" % filename)




        #
        # Funds
        #

        filename = "./%s/VIC%s_ABR_Funds.txt" % (directory, date)

        if os.path.isfile(filename):
            execCmd("rm -rf FUNDS")

            schema = StructType([
                StructField("PID", LongType(), True),
                StructField("Fnd_Nm", StringType(), True),
                StructField("Role_Typ_Cd", StringType(), True),
                StructField("DGR_Regn_Dt", StringType(), True),
                StructField("DGR_Cancn_Dt", StringType(), True)
            ])

            def parse(line):
                items = line.split("\t")
                return (long(items[0]), items[1], items[2], items[3], items[4])

            rdd = sc.textFile(filename, use_unicode=False).map(parse)
            df = sqlContext.createDataFrame(rdd, schema)
            df.write.parquet('./FUNDS')

            execCmd("aws s3 rm s3://sbv-abr-etl/DIMENSION/FUNDS/ --recursive")
            execCmd("aws s3 sync ./FUNDS/ s3://sbv-abr-etl/DIMENSION/FUNDS/ --exclude '*' --include '*.parquet'")
            execCmd("aws s3 sync ./FUNDS/ s3://sbv-abr-etl/DIMENSION_old/%s/FUNDS/ --exclude '*' --include '*.parquet'" % date)

            print("%s parsed and uploaded to S3" % filename)
        else:
            print("File does not exist: %s" % filename)





        #
        # Other trading names
        #

        filename = "./%s/VIC%s_ABR_Othtrdnames.txt" % (directory, date)

        if os.path.isfile(filename):
            execCmd("rm -rf OTHTRDNAMES")

            schema = StructType([
                StructField("PID", LongType(), True),
                StructField("Trd_Nm", StringType(), True),
            ])

            def parse(line):
                items = line.split("\t")
                return (long(items[0]), items[1])

            rdd = sc.textFile(filename, use_unicode=False).map(parse)
            df = sqlContext.createDataFrame(rdd, schema)
            df.write.parquet('./OTHTRDNAMES')

            execCmd("aws s3 rm s3://sbv-abr-etl/DIMENSION/OTHTRDNAMES/ --recursive")
            execCmd("aws s3 sync ./OTHTRDNAMES/ s3://sbv-abr-etl/DIMENSION/OTHTRDNAMES/ --exclude '*' --include '*.parquet'")
            execCmd("aws s3 sync ./OTHTRDNAMES/ s3://sbv-abr-etl/DIMENSION_old/%s/OTHTRDNAMES/ --exclude '*' --include '*.parquet'" % date)

            print("%s parsed and uploaded to S3" % filename)
        else:
            print("File does not exist: %s" % filename)





        #
        # Replaced ABN
        #

        filename = "./%s/VIC%s_ABR_Replacedabn.txt" % (directory, date)

        if os.path.isfile(filename):
            execCmd("rm -rf REPLACEDABN")

            schema = StructType([
                StructField("PID", LongType(), True),
                StructField("Replc_Abn", StringType(), True),
                StructField("Abn_Chng_Dt", StringType(), True)
            ])

            def parse(line):
                items = line.split("\t")
                return (long(items[0]), items[1], items[2])

            rdd = sc.textFile(filename, use_unicode=False).map(parse)
            df = sqlContext.createDataFrame(rdd, schema)
            df.write.parquet('./REPLACEDABN')

            execCmd("aws s3 rm s3://sbv-abr-etl/DIMENSION/REPLACEDABN/ --recursive")
            execCmd("aws s3 sync ./REPLACEDABN/ s3://sbv-abr-etl/DIMENSION/REPLACEDABN/ --exclude '*' --include '*.parquet'")
            execCmd("aws s3 sync ./REPLACEDABN/ s3://sbv-abr-etl/DIMENSION_old/%s/REPLACEDABN/ --exclude '*' --include '*.parquet'" % date)

            print("%s parsed and uploaded to S3" % filename)
        else:
            print("File does not exist: %s" % filename)





        #
        # Location (joined with other data)
        #

        filename = "./%s/VIC%s_ABR_Businesslocation.txt" % (directory, date)

        if os.path.isfile(filename):
            execCmd("rm -rf BUSINESSLOCATION")

            schema = StructType([
                StructField("PID", LongType(), True),
                StructField("Locn_Typ_Cd", StringType(), True),
                StructField("Locn_Strt_Dt", StringType(), True),
                StructField("Bus_Locn_Addr_Ln_1", StringType(), True),
                StructField("Bus_Locn_Addr_Ln_2", StringType(), True),
                StructField("Bus_Locn_Sbrb", StringType(), True),
                StructField("Bus_Locn_Stt", StringType(), True),
                StructField("Bus_Locn_Pc", StringType(), True),
                StructField("Bus_Locn_Cntry_Cd", StringType(), True),
                StructField("Bus_Locn_DPID", IntegerType(), True),
                StructField("Bus_Locn_Ltd", DecimalType(10, 7), True),
                StructField("Bus_Locn_Lngtd", DecimalType(10, 7), True),
                StructField("Bus_Locn_Msh_Blk", StringType(), True),
                StructField("Bus_Locn_GNAF_PID", StringType(), True),
                StructField("Bus_Locn_Posnl_Rlblty", StringType(), True),
                StructField("Bus_Locn_Ph_Area_Cd", StringType(), True),
                StructField("Bus_Locn_Ph_Num", StringType(), True),
                StructField("Bus_Locn_Ph_Area_Cd_Mbl", StringType(), True),
                StructField("Bus_Locn_Ph_Num_Mbl", StringType(), True),
                StructField("Bus_Locn_Eml", StringType(), True),
                StructField("Bus_Locn_Indy_Clsn", StringType(), True),
                StructField("Bus_Locn_Indy_Clsn_Descn", StringType(), True)
            ])

            def parse(line):
                items = line.split("\t")
                if items[9] == '':
                    items[9] = 0

                if items[10] == '':
                    items[10] = 0

                if items[11] == '':
                    items[11] = 0

                return (long(items[0]), items[1], items[2], items[3], items[4], items[5], items[6], items[7], items[8], long(items[9]), Decimal(items[10]), Decimal(items[11]), items[12], items[13], items[14], items[15], items[16], items[17], items[18], items[19], items[20], items[21])

            sparkBusinessLocation = sc.textFile(filename, use_unicode=False).map(parse)
            dfBusinessLocation = sqlContext.createDataFrame(sparkBusinessLocation, schema)


            dfMeshblocks = sqlContext.read.format('csv').options(header='true', inferSchema='true').load('./meshblocks_to_lga.csv')
            dfVGBO = sqlContext.read.format('csv').options(header='true', inferSchema='true').load('./lga_vgbo_likecouncils.csv')


            dfMeshblocks = dfMeshblocks.withColumn('LGA_NAME_2016', regexp_replace('LGA_NAME_2016', '\((.*?)\)', ''))
            dfMeshblocks = dfMeshblocks.withColumn('LGA_NAME_2016', rtrim(dfMeshblocks.LGA_NAME_2016))


            dfLocationAll = dfBusinessLocation.join(dfMeshblocks, dfBusinessLocation.Bus_Locn_Msh_Blk == dfMeshblocks.MB_CODE_2016, "left").drop("MB_CODE_2016")

            dfLocationAll = dfLocationAll.join(dfVGBO, dfLocationAll.LGA_NAME_2016 == dfVGBO.Local_Government_Area, "left").drop("LGA_NAME_2016")

            dfLocationAll.printSchema()
            dfLocationAll.show()


            dfLocationAll.write.parquet('./BUSINESSLOCATION')


            execCmd("aws s3 rm s3://sbv-abr-etl/DIMENSION/BUSINESSLOCATION/ --recursive")
            execCmd("aws s3 sync --exclude '*' --include '*.parquet' ./BUSINESSLOCATION/ s3://sbv-abr-etl/DIMENSION/BUSINESSLOCATION/")
            execCmd("aws s3 sync --exclude '*' --include '*.parquet' ./BUSINESSLOCATION/ s3://sbv-abr-etl/DIMENSION_old/%s/BUSINESSLOCATION/" % date)

            print("%s parsed and uploaded to S3" % filename)
        else:
            print("File does not exist: %s" % filename)