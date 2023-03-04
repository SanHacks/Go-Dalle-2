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
 Sass() {
        echo "${GREEN}Initiating Sass build...${GREEN}"
        npm install -g sass --prefix ../Frontend
        if ! sass; then
            echo "${RED}An error occurred while installing Sass.${NC}"
            exit 1
        fi
        echo "${GREEN}Closing...${GREEN}"
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
        echo "${GREEN}Install Redis...${GREEN}"
        #install redis to the system path and add it to the path
          sudo apt-get install redis-server
        if ! redis-server --version; then
            echo "${RED}An error occurred while installing Redis.${NC}"
            exit 1
        fi
          go get github.com/go-redis/redis
        if ! go get github.com/go-redis/redis; then
            echo "${RED}An error occurred while installing Redis.${NC}"
            exit 1
        fi
        echo "${GREEN}Done with installation...${GREEN}"
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

build() {
  echo "${BLUE}Configuring platform...${NC}"
  echo "${BLUE}Configuring platform...${NC}" >> platform.log

  if go mod tidy; then
      echo "${GREEN}Platform clean and tidy.${NC}"
      echo "${GREEN}Platform clean and tidy.${NC}" >> platform.log
  else
      echo "${RED}An error occurred while tidying modules.${NC}"
      echo "${RED}An error occurred while tidying modules.${NC}" >> platform.log
      exit 1
  fi

  if go mod verify; then
      echo "${GREEN}Platform modules verified.${NC}"
      echo "${GREEN}Platform modules verified.${NC}" >> platform.log
  else
      echo "${RED}An error occurred while verifying modules.${NC}"
      echo "${RED}An error occurred while verifying modules.${NC}" >> platform.log
      exit 1
  fi

  if go mod download; then
      echo "${GREEN}Platform modules downloaded.${NC}"
      echo "${GREEN}Platform modules downloaded.${NC}" >> platform.log
  else
      echo "${RED}An error occurred while downloading modules.${NC}"
      echo "${RED}An error occurred while downloading modules.${NC}" >> platform.log
      exit 1
  fi

  if go mod graph; then
      echo "${GREEN}Platform module dependency graph printed.${NC}"
      echo "${GREEN}Platform module dependency graph printed.${NC}" >> platform.log
  else
      echo "${RED}An error occurred while printing module dependency graph.${NC}"
      echo "${RED}An error occurred while printing module dependency graph.${NC}" >> platform.log
      exit 1
  fi

  echo "${GREEN}🚀🚀🚀 Launching platform  🚀🚀🚀${NC}"
  if go run .; then
      echo "${GREEN}Platform started.${NC}"
      echo "${GREEN}Platform started.${NC}" >> platform.log
      echo "${RED}An error occurred while starting platform.${NC}"
      echo "${RED}An error occurred while starting platform.${NC}" >> platform.log
      exit 1
  fi
  }
  synchronize(){

  echo "${BLUE}Updating local repository...${NC}"
  if ! git checkout main && git pull; then
      echo "An error occurred while updating the local repository."
      exit 1
  fi

  echo "${BLUE}Executing clone repos...${NC}"

  # Clone repos and make changes...

  # Check if there are any changes
  if ! git diff-index --quiet HEAD --; then
    # If there are changes, create a new branch with a name based on the time of day
    BRANCH_NAME="changes-$(date +%Y%m%d-%H%M%S)"
    if ! git checkout -b $BRANCH_NAME; then
      echo "An error occurred while creating a new branch."
      exit 1
    fi
  else
    # If there are no changes, stay on the current branch
    BRANCH_NAME=$(git branch --show-current)
  fi
   if ! git diff-index --quiet HEAD --; then
     BRANCH_NAME="changes-$(date +%Y%m%d-%H%M%S)"
     if ! git checkout -b "$BRANCH_NAME"; then
       echo "An error occurred while creating a new branch."
       exit 1
     fi
   else
     BRANCH_NAME=$(git branch --show-current)
   fi

   if ! git add --all :!*.gitignore; then
       echo "An error occurred while adding changes to commit."
       exit 1
   fi

   if ! git commit -m "Added changes at $(date +'%Y-%m-%d %H:%M:%S')"; then
       echo "An error occurred while committing changes."
       exit 1
   fi

   # Push changes to remote repository
   if ! git push -u origin "$BRANCH_NAME"; then
       echo "An error occurred while pushing changes to remote repository."
       exit 1
   fi

    }

echo "${GREEN} AIGEN CLI "
#choices
echo "1. Build"
echo "2. Deploy to Azure"
echo "3. Test DB"
echo "4. Admin"
echo "5. Dependency Graph"
echo "6 Github Pipeline"
echo "7. Exit"
echo "--help for help"
echo "--version for version"
#read input from user
# shellcheck disable=SC2039
read -p "Enter your choice:" choice


#if [[ $choice == "--help" ]]; then
#  # Print help message
#  echo "Help message"
#
## shellcheck disable=SC2039
#elif [[ $choice == "--version" ]]; then
#  # Print version
#  echo "Version message"
#fi

case $choice in
1) echo "${BLUE}Executing build...${NC}${NC}"
      build
   ;;
2) echo "${BLUE}Synchonization Initiated...${NC}"
      synchronize
    ;;

3) echo "${BLUE}Executing ping test...${NC}"
  read -p "Enter the URL or IP address to ping: " url
  ping -c 5 $url
  ;;

4) echo "${BLUE}Admin Tools...${NC}"
    echo -e "${YELLOW}Please select an action:${NC}"
    echo "1. Create User"
    echo "2. Delete User"
    echo "3. List Users"
    echo "4. Invetory"
    echo "5. AiGen Tools"
    echo "6. Exit Admin Tools"
    read -p "Enter your choice (1-4): " picker
    case $picker in
      1) echo "${BLUE}Creating User...${NC}"
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
      3) echo "${BLUE}Listing Users...${NC}"
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
      4) echo "${BLUE}Invetory...${NC}"
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
         echo "${GREEN}Configuring platform...${GREEN}"
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
            # shellcheck disable=SC2162
            # shellcheck disable=SC2039
            read -p "Enter your choice [ 1 - 18 ] :" choices
            case $choices in
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
                      DevOps
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
                      NewRelic
                    ;;
            17) echo "${PINK}Opening...${NC}${NC}"
                      Bitbucket
                    ;;
            18) echo "${PINK}Opening...${NC}${NC}"
                      Kubernetes
                    ;;
            esac
        ;;
      6) echo "${BLUE}Exiting Admin Tools...${NC}"
        echo "Exiting Admin Tools"
        ;;
      *) echo "${RED}Invalid choice.${NC}"
        ;;
    esac

    ;;
5) echo "${BLUE}Executing dependency graph...${NC}"
    go mod graph
    ;;

6) echo "${BLUE}Executing mirror...${NC}"
    git clone
    ;;
*) echo "${RED}Invalid choice.${NC}"
    ;;

  esac

