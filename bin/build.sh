#!/usr/bin/env sh
# Gundo Sifhufhi © 2023
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
# shellcheck disable=SC2039
echo  "${BLUE}Cleaning up and tidying...${NC}"
if ! go mod tidy; then
    # shellcheck disable=SC2039
    echo  "${RED}An error occurred while tidying modules.${NC}"
    exit 1
fi

echo  "${BLUE}Verifying modules...${NC}"
if ! go mod verify; then
    echo  "${RED}An error occurred while verifying modules.${NC}"
    exit 1
fi

echo  "${BLUE}Downloading modules...${NC}"
if ! go mod download; then
    echo  "${RED}An error occurred while downloading modules.${NC}"
    exit 1
fi

echo  "${BLUE}Printing module dependency graph...${NC}"
if ! go mod graph; then
    echo  "${RED}An error occurred while printing module dependency graph.${NC}"
    exit 1
fi

echo "${GREEN}Executing build...${NC}"