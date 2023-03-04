#!/bin/bash

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

# Define functions for each action
build() {
  echo -e "${BLUE}Executing build...${NC}"
  # Add your build commands here
}

clone_repos() {
  echo -e "${BLUE}Executing clone repos...${NC}"
  # Add your clone repos commands here
}

ping_test() {
  echo -e "${BLUE}Executing ping test...${NC}"
  read -p "Enter the URL or IP address to ping: " url
  ping -c 5 $url
}

# Prompt the user to select an action
echo -e "${YELLOW}Please select an action:${NC}"
echo "1. Build"
echo "2. Clone Repos"
echo "3. Ping Test"
read -p "Enter your choice (1-3): " choice

# Execute the selected action
case $choice in
  1) build;;
  2) clone_repos;;
  3) ping_test;;
  *) echo -e "${RED}Invalid choice.${NC}";;
esac