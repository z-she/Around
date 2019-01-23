# Around: Your Stories Where They Happened

Around is a social network service that organizes stories of its users by geographic location. Users of the service make image-centric posts about their lives and what they see, and others will be able to find the posts where the story happens.

## Stories So Far

This project is created to help me learning about the Google Cloud Platform, Golang, and cross-platform interface development. The backend of the service has been under development; front-end should be started in the near future. There is no release date yet. 

## Technologies Used

The HTTP requests are handled by code written in Golang, post searching is carried out by ElasticSearch, and image storage is provided by Google Cloud Storage. Post data is optionally replicated by BigTable; the collected data will be used by BigQuery for analytics. The data transfer from BigTable to BigQuery will be handled by Dataflow. 
