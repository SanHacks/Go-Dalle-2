#!/bin/zsh

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

# ADMIN FUNCTIONS
# Define functions for each action
createUser() {
  echo -e "${BLUE}Creating user...${NC}"
  # Add your create user commands here
}

deleteUser() {
  echo -e "${BLUE}Deleting user...${NC}"
  # Add your delete user commands here
}

listUsers() {
  echo -e "${BLUE}Listing users...${NC}"
  # Add your list users commands here
}

echo -e "${YELLOW}Please select an action:${NC}"
echo "1. Create User"
echo "2. Delete User"
echo "3. List Users"
echo "4. Invetory"
echo "5. AiGen Tools"
echo "6. Exit Admin Tools"
read -p "Enter your choice (1-4): " choice

# Execute the selected action

# shellcheck disable=SC1073
case $choice in
1) echo "${BLUE}Creating user...${NC}"
   read -p "Enter the username: " username
   read -p "Enter the password: " password
   read -p "Enter the email: " email
   read -p "Enter the role: " role
   echo "Creating user $username with password $password, email $email and role $role"
   # Add your create user commands here
   mysql -u ndiGundoSan -p@Sifhufhi2024 -h aigen.mysql.database.azure.com -NB -e "use aigen; INSERT INTO adminusers (name, password, email, role) VALUES ('$username', '$password', '$email', '$role');"
   ;;
2) echo "${BLUE}Deleting user...${NC}"
    read -p "Enter the username: " username
    echo "Deleting user $username"
    # Add your delete user commands here
    mysql -u ndiGundoSan -p@Sifhufhi2024 -h aigen.mysql.database.azure.com -NB -e "use aigen; DELETE FROM adminusers WHERE name = '$username';"
    ;;
3) echo "${BLUE}Listing users...${NC}"
   if [ -z "$1" ]; then
     echo "Listing all users"
     # Add your list users commands here
     mysql -u ndiGundoSan -p@Sifhufhi2024 -h aigen.mysql.database.azure.com -NB -e "use aigen; SELECT * FROM adminusers;"
   else
     echo "Listing user $1"
     # Add your list user commands here
     mysql -u ndiGundoSan -p@Sifhufhi2024 -h aigen.mysql.database.azure.com -NB -e "use aigen; SELECT * FROM adminusers WHERE name = '$1';"
   fi
  ;;

#dsplay in a table format
4) echo "${BLUE}Inventory...${NC}"
   echo "Listing inventory"
   # Add your list inventory commands here
   mysql -u ndiGundoSan -p@Sifhufhi2024 -h aigen.mysql.database.azure.com -NB -e "use aigen; SELECT * FROM generatedproducts;"
   if [ -z "$1" ]; then
     echo "Listing all inventory"
     # Add your list inventory commands here
     mysql -u ndiGundoSan -p@Sifhufhi2024 -h aigen.mysql.database.azure.com -NB -e "use aigen; SELECT * FROM generatedproducts;"
   else
     echo "Listing generatedproducts $1"
     # Add your list inventory commands here
     mysql -u ndiGundoSan -p@Sifhufhi2024 -h aigen.mysql.database.azure.com -NB -e "use aigen; SELECT * FROM generatedproducts WHERE name = '$1';"
   fi
  ;;

5) echo "${BLUE}AiGen Tools...${NC}"
   sh aigen.sh