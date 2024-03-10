workspace {
    name "Social network"
    description "Social network where user can have his/her own wall and communicate with others using PtP chat"

    !identifiers hierarchical
    !docs documentation
    !adrs decisions

    model {
        properties { 
            structurizr.groupSeparator "/"
        }

        user = person "User" {
            description "Social network user willing to chat"
        }

        social_network = softwareSystem "Social network"  {
            description "Main app where users can have their account, wall and exchange messages"

            front_service = container "UI application" {
                description "Tool for the user to interact with the system in browser via web page."
                technology "CSS/HTML"
                tags "web"
            }

            auth_service = container "Authentication" {
                description "Microservice for providing authentification for each user request."
                technology "C++/Python"
                tags "web"
            }
            
            user_service = container "Working with user data" {
                description "Microservice for handling requests for login, registration."
                technology "C++/Python"
                tags "web"
            }
            
            posts_service = container "Working with posts data" {
                description "Microservice for handling requests for making new posts or deleting/updating the old ones."
                technology "C++/Python"
                tags "web"
            }
            
            msgs_service = container "Working with messages data" {
                description "Microservice for handling requests for sending new messages or deleting/updating the old ones.."
                technology "C++/Python"
                tags "web"
            }

            group "Databases" {
                user_database = container "User Database" {
                    description "User data storage (relational database with indexes)"
                    technology "PostgreSQL 16"
                    tags "database"
                }

                cache_database = container "Cache" {
                    description "User data cache for faster usage"
                    technology "Redis 7"
                    tags "database"
                }

                posts_msgs_database = container "Posts and Messages Database" {
                    description "Document-oriented database for storing posts and messages between users"
                    technology "MongoDb 5"
                    tags "database"
                }
            }

            front_service -> auth_service "Redirecting user requests for auth" "TCP 8081"

            auth_service -> user_service "Sending valid requests about user for processing" "TCP 8087"
            auth_service -> posts_service "Sending valid requests about posts for processing" "TCP 8088"
            auth_service -> msgs_service "Sending valid requests about messages for processing" "TCP 8089"
            auth_service -> cache_database "Check user existence" "TCP 6379"

            user_service -> user_database "Read/Update/Delete user data" "TCP 5432"
            user_service -> cache_database "Read/Update/Delete user data" "TCP 6379"

            posts_service -> posts_msgs_database "Read/Update/Delete posts data" "TCP 27017"
            posts_service -> cache_database "Check user existence" "TCP 6379"

            msgs_service -> posts_msgs_database "Read/Update/Delete msgs data" "TCP 27017"
            msgs_service -> cache_database "Check user existence" "TCP 6379"
            
            user -> front_service "Register/Post/Chat" "REST HTTP:8080"
        }

        user -> social_network "Interacts with his account and wall. Sends messages via PtP chat" "REST HTTP:8080"

        deploymentEnvironment "Production" {
            deploymentNode "Front Server" {
                containerInstance social_network.front_service
                instances 1
            }

            deploymentNode "Auth Server" {
                containerInstance social_network.auth_service
                instances 1
                properties {
                    "CPU" "2"
                    "RAM" "64Gb"
                    "HDD" "2Tb"
                }
                tags "web"
            }

            deploymentNode "User Server" {
                containerInstance social_network.user_service
                instances 1
                properties {
                    "CPU" "2"
                    "RAM" "64Gb"
                    "HDD" "2Tb"
                }
                tags "web"
            }

            deploymentNode "Posts Server" {
                containerInstance social_network.posts_service
                instances 1
                properties {
                    "CPU" "2"
                    "RAM" "64Gb"
                    "HDD" "2Tb"
                }
                tags "web"
            }

            deploymentNode "Messages Server" {
                containerInstance social_network.msgs_service
                instances 1
                properties {
                    "CPU" "2"
                    "RAM" "64Gb"
                    "HDD" "2Tb"
                }
                tags "web"
            }
 
            deploymentNode "Databases" {
                deploymentNode "User database Server" {
                    containerInstance social_network.user_database
                    instances 2
                    tags "database"
                }

                deploymentNode "Posts N Messages database Server" {
                    containerInstance social_network.posts_msgs_database
                    instances 2
                    tags "database"
                }

                deploymentNode "Cache Server" {
                    containerInstance social_network.cache_database
                    instances 2
                    tags "database"
                }
            }
        }
    }

    views {
        themes default

        properties { 
            structurizr.tooltips true
        }

        !script groovy {
            workspace.views.createDefaultViews()
            workspace.views.views.findAll { it instanceof
            com.structurizr.view.ModelView }.each { it.enableAutomaticLayout() }
        }

        dynamic social_network "UC01" "New user registration" {
            autoLayout
            user -> social_network.front_service "Create new {user} with data (POST /user)"
            social_network.front_service -> social_network.auth_service "Sending parsed from front data"
            social_network.auth_service -> social_network.user_service "If auth checks are correct sending request type and data"
            social_network.user_service -> social_network.cache_database "Checking {user} has not existed previously"
            social_network.user_service -> social_network.user_database "Saving new user data"
            social_network.user_service -> social_network.cache_database "Adding new user data to cache"
        }

        dynamic social_network "UC02" "User search by login" {
            autoLayout
            user -> social_network.front_service "Get user with given {login} (GET /user)"
            social_network.front_service -> social_network.auth_service "Sending parsed from front data"
            social_network.auth_service -> social_network.user_service "If auth checks are correct sending request type and data"
            social_network.user_service -> social_network.cache_database "Searching for user with given {login}"
        }

        dynamic social_network "UC03" "User search by mask" {
            autoLayout
            user -> social_network.front_service "Get user with given {mask} (GET /user)"
            social_network.front_service -> social_network.auth_service "Sending parsed from front data"
            social_network.auth_service -> social_network.user_service "If auth checks are correct sending request type and data"
            social_network.user_service -> social_network.cache_database "Searching for user with given {mask}"
        }

        dynamic social_network "UC04" "Adding new post" {
            autoLayout
            user -> social_network.front_service "Adding new post for user with {id} (POST /user/posts)"
            social_network.front_service -> social_network.auth_service "Sending parsed from front data"
            social_network.auth_service -> social_network.cache_database "Checking user-source with {id} existence (Operation enabled only for registrated users)"
            social_network.auth_service -> social_network.posts_service "If auth checks are correct sending request type and data"
            social_network.posts_service -> social_network.posts_msgs_database "Adding new post for the given user with {id}"
        }

        dynamic social_network "UC05" "Getting user posts" {
            autoLayout
            user -> social_network.front_service "Getting user's posts by {id} (GET /user/posts)"
            social_network.front_service -> social_network.auth_service "Sending parsed from front data"
            social_network.auth_service -> social_network.posts_service "If auth checks are correct sending request type and data"
            social_network.posts_service -> social_network.cache_database "Check user with {id} existence"
            social_network.posts_service -> social_network.posts_msgs_database "Getting posts for user with given {id} if posts found"
        }

        dynamic social_network "UC06" "Sending message to the user" {
            autoLayout
            user -> social_network.front_service "Sending msg from user with {id1} to user with {id2}(POST /user/msgs)"
            social_network.front_service -> social_network.auth_service "Sending parsed from front data"
            social_network.auth_service -> social_network.cache_database "Checking user-source {id1} existence"
            social_network.auth_service -> social_network.posts_service "If auth checks are correct sending request type and data"
            social_network.posts_service -> social_network.cache_database "Check user (destination {id2}) existence"
            social_network.posts_service -> social_network.posts_msgs_database "Adding new message for user-source {id1} and user-destination {id2}"
        }

        dynamic social_network "UC07" "Getting messages for the user" {
            autoLayout
            user -> social_network.front_service "Getting msg for user with {id} (GET /user/msgs)"
            social_network.front_service -> social_network.auth_service "Sending parsed from front data"
            social_network.auth_service -> social_network.cache_database "Checking user existence with {id} (operation is possible only for the owner of the account)"
            social_network.auth_service -> social_network.posts_service "If auth checks are correct sending request type and data"
            social_network.posts_service -> social_network.posts_msgs_database "Getting messages where given user with {id} is the destination"
        }

        styles {
            element "database" {
                color #ffffff
                shape cylinder
            }
            
            element "web" {
                color #ffffff
                shape RoundedBox
            }
        }
    }
}