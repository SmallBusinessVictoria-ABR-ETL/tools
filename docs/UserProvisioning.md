## Groups

* DataLake_DataEngineer
  * Read / Write access to `s3://sbv-abr-erl`
  * Full access to athena
  * Self Manage Passwords

## Provision New Users

```bash
export AWS_PROFILE=sbv
export SBV_USER=jason.test
export SBV_TMP_PASS=sbv_temp_pass_010
export SBV_GROUP=DataLake_DataEngineer

# Create User Account
aws iam create-user --user-name ${SBV_USER}

# Add User to group
aws iam add-user-to-group --user-name ${SBV_USER} --group-name ${SBV_GROUP}

# Allow User to log into AWS Console
aws iam create-login-profile --user-name ${SBV_USER} --password ${SBV_TMP_PASS} --password-reset-required

# Create API Keys
aws iam create-access-key --user-name ${SBV_USER}
```


Email template for new user
```text
Hi <name>, 

Welcome to the Small Business Victoria - Data Lake, to access the console please log in:

https://402555251914.signin.aws.amazon.com/console/

Username: <username>
Password: sbv_temp_pass_010

You will be required to set a password. Please ensure your password contains more then 8 characters and includes Capitals, Lowercase, Numbers and Symbols

Kind Regards,
Small Business Victoria - Data Lake Suppor

```