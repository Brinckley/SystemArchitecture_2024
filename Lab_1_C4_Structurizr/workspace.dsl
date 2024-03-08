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

            api_service = container "API application" {
                description "Tool for user to interact with the system in browser via web page. Working with requests from user. Authentication and validation. Providing functionality."
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

           
            api_service -> user_database "Read/Update/Delete user data" "TCP 5432"
            api_service -> cache_database "Read/Update/Delete user data" "TCP 6379"
            api_service -> posts_msgs_database "Read/Update/Delete posts and msgs data" "TCP 27017"
            
            user -> api_service "Register/Post/Chat" "REST HTTP:8080"
        }

        user -> social_network "Interacts with his account and wall. Sends messages via PtP chat" "REST HTTP:8080"

        deploymentEnvironment "Production" {
            deploymentNode "API Service Server" {
                containerInstance social_network.api_service
                instances 1
                properties {
                    "CPU" "2"
                    "RAM" "64Gb"
                    "HDD" "2Tb"
                }
                tags "web"
            }

            deploymentNode "Databases" {
                deploymentNode "User database Server 1" {
                    containerInstance social_network.user_database
                    instances 3
                    tags "database"
                }

                deploymentNode "Posts N Messages database Server 2" {
                    containerInstance social_network.posts_msgs_database
                    instances 3
                    tags "database"
                }

                deploymentNode "Cache Server 3" {
                    containerInstance social_network.cache_database
                    instances 1
                    tags "database"
                }
            }
        }
    }

    views {
        themes default

        container social_network {
            include *
            autoLayout
        }

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
            user -> social_network.api_service "Create new user with data (POST /user)"
            social_network.api_service -> social_network.user_database "Saving new user data"
            social_network.api_service -> social_network.cache_database "Adding user data to cache"
        }

        dynamic social_network "UC02" "Addig new post" {
            autoLayout
            user -> social_network.api_service "Add new post (POST /user/posts)"
            social_network.api_service -> social_network.posts_msgs_database "Saving new post to the DB"
        }

        dynamic social_network "UC03" "Search for user by login" {
            autoLayout
            user -> social_network.api_service "Find user with {login} (GET /users)"
            social_network.api_service -> social_network.cache_database "Getting user by {login}"
            social_network.api_service -> social_network.user_database "If not found in cache getting user by {login} or none if not found in cache"
        }

        dynamic social_network "UC04" "Search for user by mask" {
            autoLayout
            user -> social_network.api_service "Find user with {mask} (GET /users)"
            social_network.api_service -> social_network.cache_database "Getting user by {mask}"
            social_network.api_service -> social_network.user_database "If not found in cache getting user by {mask} or none"
        }

        dynamic social_network "UC05" "Getting user's posts" {
            autoLayout
            user -> social_network.api_service "Get user posts (GET /user/posts)"
            social_network.api_service -> social_network.posts_msgs_database "Getting posts for user from the DB"
        }

        dynamic social_network "UC06" "Sending message" {
            autoLayout
            user -> social_network.api_service "Send the message (POST /user/msgs)"
            social_network.api_service -> social_network.posts_msgs_database "Adding new message to the DB with source and destination"
        }

        dynamic social_network "UC07" "Getting messages" {
            autoLayout
            user -> social_network.api_service "Get messages (GET /user/msgs)"
            social_network.api_service -> social_network.posts_msgs_database "Getting messages from the DB with source and destination"
        }

        styles {
            element "database" {
                color #ffffff
                shape cylinder
            }
            
            element "web" {
                color #1168bd
                shape RoundedBox
            }
        }
    }
}