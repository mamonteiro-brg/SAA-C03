


–  “dynamic port routing” = Application load balancer
–  ” static web application” = CloudFront
–  “geographically redundant ,scalable and highly available”  = Amazon Route 53 and Auto Scaling
–  Keyword “Migrate Docker web application” = Use ECS
–  since the current application use only one docker container for web application so we need to use AWS Elastic Beanstalk Docker Single Container for the web application , and an Amazon RDS for MySQL for the database.
–  Cost Optimized Storage+ retain the logs for only 30 days. + analyze logs “Unstructured data”  =  Use S3 with Athena
–  Spanning multiple AZ’s provides high availability (generally with these questions you don’t need to go across region to satisfy HA type questions unless it specifically says a certain number of miles). The Mutli-AZ Amazon RDS makes the DB highly available too. Once these are done, there is no need to have a separate DR environment and thus can be decommissioned.
–  keyword  “master keys and unencrypted data should never be sent to AWS ” then you have to client-side encryption with a client-side master key to send encrypted data.
–  Shared File System + Concurrent access + Directories permission  = Amazon EFS and control permissions by using file-level permissions.
–  S3 Select is used to query data in S3 using SQL to retrieve subset of the data using server side filtering .
–  to provide High availability to RDS = Configure RDS MySQL to Multi-AZ.
–  ELB connection draining can be used to stop sending requests to instances that are de-registering or unhealthy.
–  running the application from inside a VPC – private subnet-. With this ability we’re able to create a NAT (Network Address Translator) Gateway so that all out-bound connections from the application will exit from the NAT which is assigned to a fixed IP address. This fixed IP address can then be whitelisted by our third-parties.
–  Keyword “NFS” + no code changes + highly available + durability = Amazon EFS
–  Key Word “Scratch Data” = Use instance store
–  create a CloudFront origin access identity to restrict access to your Amazon S3 content, and grant the origin access identity the desired permissions.
–  Keyworkd “Web sockets” = Application load balancer
–  You can use a network address translation (NAT) gateway to enable instances in a private subnet to connect to the internet or other AWS services, but prevent the internet from initiating a connection with those instances.
–  The “fanout” scenario is when an Amazon SNS message is sent to a topic and then replicated and pushed to multiple Amazon SQS queues, HTTP endpoints, or email addresses. This allows for parallel asynchronous processing. For example, you could develop an application that sends an Amazon SNS message to a topic whenever an order is placed for a product. Then, the Amazon SQS queues that are subscribed to that topic would receive identical notifications for the new order. The Amazon EC2 server instance attached to one of the queues could handle the processing or fulfillment of the order, while the other server instance could be attached to a data warehouse for analysis of all orders received.
–  Another way to use “fanout” is to replicate data sent to your production environment with your test environment. Expanding upon the previous example, you could subscribe yet another queue to the same topic for new incoming orders. Then, by attaching this new queue to your test environment, you could continue to improve and test your application using data received from your production environment.
– You can use a network address translation (NAT) gateway to enable instances in a private subnet to connect to the internet or other AWS services, but prevent the internet from initiating a connection with those instances.
– Elastic IP must be associated with the NAT gateway.
– you can take advantage of the additional performance and security of CloudFront by putting it in front of your Application Load Balancers (ALB), AWS Elastic Beanstalk, Amazon S3, and other AWS resources delivering HTTP(S) objects for next to no additional cost.
– Instance store is ideal for temporary storage of information that changes frequently, such as buffers, caches, scratch data, and other temporary content, or for data that is replicated across a fleet of instances, such as a load-balanced pool of web servers.
– routing based on metadata “headers”= ELB Application Load Balancers with host-based routing
– Cost Optimized + Use DynamoDB + the maximum request per second is a few hundreds = Use Serverless architecture = Use API Gateway with AWS Lambda.
– infrequently accessed + retrieved immediately + high available+ LOWEST cost =Amazon S3 Standard-IA
– to prevent deleting an object in S3 for a fixed amount of time = Use S3 Object lock
– keyword “concurrently access , strong data consistency and file locking” = Amazon EFS
– S3 is an origin for CloudFront. EBS volumes would need EC2 instances behind an Elastic Load Balancing load balancer to be an origin
– keyword = random distribution on healthy servers = Multivalue answer
–   keyword ” require operating system permission” this means You cannot use RDS for that.
    highly available database = Multi Availability Zones  
– Application load balancer does not support TCP Passthrough
– Network Load balancer support TCP Passthrough    
– keyword “without provisioning or managing infrastructure.” = Serverless Solution
– accessed through HTTPS with custom domain name = CloudFront with S3 as origin
– it’s static, you don’t need API gateway or Lambda .
– to replicate table data between multi Regions = Use Amazon DynamoDB global tables.
– 
– 
– 
– 
– 
– 
– 
– 
– 
– 
– 
– 
– 
– 
– 
– 
– 
– 
– 
– 
– 
– 
– 
– 
– 
– 
– 
– 
– 
– 
– 
– 
– 
– 
– 
– 
– 
– 
– 
– 
– 
– 
– 
– 
– 
– 
– 
– 
– 
– 
– 
– 
– 
– 
– 
