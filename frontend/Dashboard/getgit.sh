#!/bin/bash

# Set FTPS endpoint and login credentials
FTP_ENDPOINT="ftps://waws-prod-blu-383.ftp.azurewebsites.windows.net/site/wwwroot"
FTP_USERNAME="AiGen\$AiGen"
FTP_PASSWORD="LcjqYwBvHpuZZFPvytZ4f41SdYtPfod5bghoAfjNkKd6ok7t1A3bZ35mGH3o"

# Set local directory containing files to be uploaded
LOCAL_DIR="."

# Set remote directory on the FTP server
REMOTE_DIR="/site/wwwroot"

# Set FTPS options
FTP_OPTIONS="--explicit-ssl --disable-epsv"

# Change to local directory
cd $LOCAL_DIR

# Upload files to remote directory
ftps -c "set ftp:ssl-allow no; open '$FTP_ENDPOINT'; user '$FTP_USERNAME' '$FTP_PASSWORD'; mirror --reverse $FTP_OPTIONS . $REMOTE_DIR"

# Exit script
exit 0
