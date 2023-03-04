#!/bin/zsh

# Gundo Sifhufhi © 2023
# Define colors
RED='\033[0;31m'
BLUE='\033[0;34m'
TEAL='\033[0;36m'
GOLD='\033[0;33m'
PURPLE='\033[0;35m'
PINK='\033[0;33m'
PINK='\033[0;35m'
GREEN='\033[0;32m'
YELLOW='\033[0;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

#AIGEN 3rd Party Dependencies Management
#Sass,
 #Javascript,
 #Python,
 #MySQL,
 #PHP,
 #React,
 #Redis,
 #jQuery ,
 # HTML5,
 #Bootstrap
 #Utilities
 #Google Analytics
 #Dev Ops
 #Grunt,
 # Jenkins,
 #Vagrant,
 # New Relic,
 #Bitbucket,
 #Kubernetes
 #Business Tool
 #Google Apps,
 #Slack,
 #JIRA
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

read -p "Enter your choice [ 1 - 18 ] " choice

case $choice in
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