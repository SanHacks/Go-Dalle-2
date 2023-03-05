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
case $picker in
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
      Sass() {
        echo "${GREEN}Opening...${GREEN}"
        # Add your build commands here
        echo "${GREEN}Opening...${GREEN}"
        # Add your build commands here
        }
      Javascript() {
        echo "${GREEN}Opening...${GREEN}"
        # Add your build commands here
        echo "${GREEN}Opening...${GREEN}"
        # Add your build commands here
        }
      Python() {
        echo "${GREEN}Opening...${GREEN}"
        # Add your build commands here
        echo "${GREEN}Opening...${GREEN}"
        # Add your build commands here
        }
      MySQL() {
        echo "${GREEN}Opening...${GREEN}"
        # Add your build commands here
        echo "${GREEN}Opening...${GREEN}"
        # Add your build commands here
        }
      PHP() {
        echo "${GREEN}Opening...${GREEN}"
        # Add your build commands here
        echo "${GREEN}Opening...${GREEN}"
        # Add your build commands here
        }
      React() {
        echo "${GREEN}Opening...${GREEN}"
        # Add your build commands here
        echo "${GREEN}Opening...${GREEN}"
        # Add your build commands here
        }
      Redis() {
        echo "${GREEN}Opening...${GREEN}"
        # Add your build commands here
        echo "${GREEN}Opening...${GREEN}"
        # Add your build commands here
        }
      jQuery() {
        echo "${GREEN}Opening...${GREEN}"
        # Add your build commands here
        echo "${GREEN}Opening...${GREEN}"
        # Add your build commands here
        }
      HTML5() {
        echo "${GREEN}Opening...${GREEN}"
        # Add your build commands here
        echo "${GREEN}Opening...${GREEN}"
        # Add your build commands here
        }
      Bootstrap() {
        echo "${GREEN}Opening...${GREEN}"
        # Add your build commands here
        echo "${GREEN}Opening...${GREEN}"
        # Add your build commands here
        }
      Utilities() {
        echo "${GREEN}Opening...${GREEN}"
        # Add your build commands here
        echo "${GREEN}Opening...${GREEN}"
        # Add your build commands here
        }
      DevOps() {
        echo "${GREEN}Opening...${GREEN}"
        # Add your build commands here
        echo "${GREEN}Opening...${GREEN}"
        # Add your build commands here
        }
      Grunt() {
        echo "${GREEN}Opening...${GREEN}"
        # Add your build commands here
        echo "${GREEN}Opening...${GREEN}"
        # Add your build commands here
        }
      Jenkins() {
        echo "${GREEN}Opening...${GREEN}"
        # Add your build commands here
        echo "${GREEN}Opening...${GREEN}"
        # Add your build commands here
        }
      Vagrant() {
        echo "${GREEN}Opening...${GREEN}"
        # Add your build commands here
        echo "${GREEN}Opening...${GREEN}"
        # Add your build commands here
        }
      NewRelic() {
        echo "${GREEN}Opening...${GREEN}"
        # Add your build commands here
        echo "${GREEN}Opening...${GREEN}"
        # Add your build commands here
        }
      Bitbucket() {
        echo "${GREEN}Opening...${GREEN}"
        # Add your build commands here
        echo "${GREEN}Opening...${GREEN}"
        # Add your build commands here
        }
      Kubernetes() {
        echo "${GREEN}Opening...${GREEN}"
        # Add your build commands here
        echo "${GREEN}Opening...${GREEN}"
        # Add your build commands here
        }

       echo "${GREEN}Configuring platform...${GREEN}"
       #choices
      echo "1. Sass"
      echo "2. Javascript"
      echo "3. Python"
      echo "4. MySQL"
      echo "5. PHP"
      echo "6. React"
      echo "7. Redis"
      echo "8. jQuery"
      echo "9. HTML5"
      echo "10. Bootstrap"
      echo "11. Utilities"
      echo "12. Dev Ops"
      echo "13. Grunt"
      echo "14. Jenkins"
      echo "15. Vagrant"
      echo "16. New Relic"
      echo "17. Bitbucket"
      echo "18. Kubernetes"

      read -p "Enter your choice [ 1 - 18 ] " choices

      case choices in
      1) echo "${PINK}Opening...${NC}${NC}"
            Sass
         ;;
      2) echo "${PINK}Opening...${NC}${NC}"
            Javascript
          ;;
      3) echo "${PINK}Opening...${NC}${NC}"
            Python
          ;;
      4) echo "${PINK}Opening...${NC}${NC}"
            MySQL
          ;;
      5) echo "${PINK}Opening...${NC}${NC}"
            PHP
          ;;
      6) echo "${PINK}Opening...${NC}${NC}"
            React
          ;;
      7) echo "${PINK}Opening...${NC}${NC}"
            Redis
          ;;
      8) echo "${PINK}Opening...${NC}${NC}"
            jQuery
          ;;
      9) echo "${PINK}Opening...${NC}${NC}"
            HTML5
          ;;
      10) echo "${PINK}Opening...${NC}${NC}"
            Bootstrap
          ;;
      11) echo "${PINK}Opening...${NC}${NC}"
            Utilities
          ;;
      12) echo "${PINK}Opening...${NC}${NC}"
            Dev Ops
          ;;
      13) echo "${PINK}Opening...${NC}${NC}"
            Grunt
          ;;
      14) echo "${PINK}Opening...${NC}${NC}"
            Jenkins
          ;;
      15) echo "${PINK}Opening...${NC}${NC}"
            Vagrant
          ;;
      16) echo "${PINK}Opening...${NC}${NC}"
            New Relic
          ;;
      17) echo "${PINK}Opening...${NC}${NC}"
            Bitbucket
          ;;
      18) echo "${PINK}Opening...${NC}${NC}"
            Kubernetes
          ;;
      *) echo "${RED}Invalid choice...${NC}"
          ;;
        esac