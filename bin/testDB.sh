#!/usr/bin/env sh


# Define colors
RED='\033[0;31m'
BLUE='\033[0;34m'
TEAL='\033[0;36m'
GOLD='\033[0;33m'
PURPLE='\033[0;35m'
ORANGE='\033[0;33m'
PINK='\033[0;35m'
GREEN='\033[0;32m'
YELLOW='\033[0;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color


#mysql u- ndiGundoSan -p @Sifhufhi2024 -h aigen.mysql.database.azure.com

# Redirect MySQL output to a file
# Attempt to connect to MySQL database and select the "aigen" database
mysql -u ndiGundoSan -p@Sifhufhi2024 -h aigen.mysql.database.azure.com -NB -e "use aigen;"

# Check the exit status of the mysql command
if [ $? -eq 0 ]; then
    # If successful, log success message
    echo "${GREEN}✅  Database connection successful. Health check passed.${NC}"
else
    # If unsuccessful, log failure message
    echo "${RED}MySQL connection failed.${NC}"
fi


#echo -e "${RED}Error:${NC} Something went wrong."
#echo -e "${GREEN}Success:${NC} The operation completed successfully."
#echo -e "${YELLOW}Warning:${NC} This action cannot be undone."
#echo -e "${BLUE}Info:${NC} Here's some information you might find useful."
