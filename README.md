# SAA-C03
Preparation Material for SAA-C03


Exam Tips
- the company want to generate the keys then store it in AWS with enabling cross-region replication then you have to configure S3 buckets to use SSE with AWS KMS-Managed Keys (SSE-KMS) with imported key material in both regions
- to secure the data in EBS volume = encrypt the data using KMS
- security group refer to another security group.
–  To distribute the traffic between both environments (Active – Active)= Use Use Route 53 weighted routing policy .
– To route domain traffic to an ELB load balancer, use Amazon Route 53 to create an alias record that points to your load balancer.
- configure the TTL is the only way to control caching in Amazon API Gateway.
- to restrict users based on their location = Use Amazon CloudFront Geo Restriction.
- to give the client one static IP , you have to create ELB and register the EC2 instances in the ELB with Auto Scaling Group
- Use CORS to allow allow communication between applications from different domains.
- orders are processed exactly one + handle large increases in the number of requests = Amazon SQS FIFO
- IPv6 traffic to the Internet = Use an egress-only internet gateway
- to map domain to another domain, use CNAME record.
- to backup EBS volume =  Create periodic snapshots of the Amazon EBS volumes using lifecycle policy
- Server Name Indication (SNI) is the solution to this problem. Browsers that support SNI will immediately communicate the name of the website the visitor wants to connect to, during the initialization of the secured connection, so that the server knows which certificate to send back. This allows browsers/clients and servers supporting SNI to connect multiple certificates for multiple domains to one IP address. So, your website visitor will not notice any differences.
- Application load balancer is used to provide request routing rules at the HTTP level
- Global Accelerator will provide us with the two static IP.
- to build a fast session store for your online applications = Use Amazon ElastiCache for Redis
- to access AWS resources from a private subnet without using internet = Use VPC endpoint
- cost effective and scalable storage +store mapping data in hundreds of data files+ can grow to tens of terabytes = Amazon S3 Standard
- to route a trafic to static pages hosted on Amazon S3 , you must have :-
  a – registered domain name.
  b –  The bucket must have the same name as your domain or subdomain
- API Gateway supports multiple mechanisms for controlling and managing access to your API.
You can use the following mechanisms for authentication and authorization:
     Resource policies let you create resource-based policies to allow or deny access to your APIs and methods from specified source IP addresses or VPC endpoints.
     Standard AWS IAM roles and policies offer flexible and robust access controls that can be applied to an entire API or individual methods. IAM roles and policies can be used for controlling who can 
     create and manage your APIs, as well as who can invoke them.
     IAM tags can be used together with IAM policies to control access.
     Endpoint policies for interface VPC endpoints allow you to attach IAM resource policies to interface VPC endpoints to improve the security of your private APIs.
     Lambda authorizers are Lambda functions that control access to REST API methods using bearer token authentication—as well as information described by headers, paths, query strings, stage variables, or context variables request parameters. Lambda authorizers are used to control who can invoke REST API methods.
     Amazon Cognito user pools let you create customizable authentication and authorization solutions for your REST APIs. Amazon Cognito user pools are used to control who can invoke REST API methods.
- to attach elastic network interface (ENI) with running EC2 instance , you have to use Hot Attach
- 





An AWS KMS customer master key (CMK) is a logical representation of a master key. In addition to the CMK identifiers and other metadata, a CMK contains the key material used to encrypt and decrypt data. When you create a CMK, by default AWS KMS generates the key material for that CMK. But you can create a CMK without key material and then import your own key material into that CMK, a feature often known as “bring your own key” (BYOK).
Imported key material is supported only for symmetric CMKs in AWS KMS key stores. It is not supported on asymmetric CMKs or CMKs in custom key stores.
When you use imported key material, you remain responsible for the key material while allowing AWS KMS to use a copy of it. You might choose to do this for one or more of the following reasons:
• To prove that you generated the key material using a source of entropy that meets your requirements.
• To use key material from your own infrastructure with AWS services, and to use AWS KMS to manage the lifecycle of that key material within AWS.
• To set an expiration time for the key material in AWS and to manually delete it, but to also make it available again in the future. In contrast, scheduling key deletion requires a waiting period of 7 to 30 days, after which you cannot recover the deleted CMK.
• To own the original copy of the key material, and to keep it outside of AWS for additional durability and disaster recovery during the complete lifecycle of the key material.
Reference

If your VPC has a VPC peering connection with another VPC, a security group rule can reference another security group in the peer VPC. This allows instances that are associated with the referenced security group and those that are associated with the referencing security group to communicate with each other.
If the owner of the peer VPC deletes the referenced security group, or if you or the owner of the peer VPC deletes the VPC peering connection, the security group rule is marked as stale. You can delete stale security group rules as you would any other security group rule.

A gateway endpoint is a gateway that you specify as a target for a route in your route table for traffic destined to a supported AWS service. The following AWS services are supported:
– Amazon S3
– DynamoDB

Explanation
A VPC endpoint enables you to privately connect your VPC to supported AWS services and VPC endpoint services powered by AWS PrivateLink without requiring an internet gateway, NAT device, VPN connection, or AWS Direct Connect connection. Instances in your VPC do not require public IP addresses to communicate with resources in the service. Traffic between your VPC and the other service does not leave the Amazon network.
Endpoints are virtual devices. They are horizontally scaled, redundant, and highly available VPC components. They allow communication between instances in your VPC and services without imposing availability risks or bandwidth constraints on your network traffic.
There are two types of VPC endpoints: interface endpoints and gateway endpoints. Create the type of VPC endpoint required by the supported service.


Explanation
If you host a website on multiple Amazon EC2 instances, you can distribute traffic to your website across the instances by using an Elastic Load Balancing (ELB) load balancer. The ELB service automatically scales the load balancer as traffic to your website changes over time. The load balancer also can monitor the health of its registered instances and route domain traffic only to healthy instances.
To route domain traffic to an ELB load balancer, use Amazon Route 53 to create an alias record that points to your load balancer. An alias record is a Route 53 extension to DNS. It’s similar to a CNAME record, but you can create an alias record both for the root domain, such as example.com, and for subdomains, such as http://www.example.com. (You can create CNAME records only for subdomains.)

Explanation
NLB enables static IP addresses for each Availability Zone. These static addresses don’t change, so they are good for our firewalls’ whitelisting. However, NLB allows only TCP traffic, no HTTPS offloading, and they have none of the nice layer 7 features of ALB.

Explanation
You can use Amazon Data Lifecycle Manager to automate the creation, retention, and deletion of snapshots taken to back up your Amazon EBS volumes. Automating snapshot management helps you to:
· Protect valuable data by enforcing a regular backup schedule.
· Retain backups as required by auditors or internal compliance.
· Reduce storage costs by deleting outdated backups.
Combined with the monitoring features of Amazon CloudWatch Events and AWS CloudTrail, Amazon Data Lifecycle Manager provides a complete backup solution for EBS volumes at no additional cost.

Explanation
With SNI support we’re making it easy to use more than one certificate with the same ALB. The most common reason you might want to use multiple certificates is to handle different domains with the same load balancer. It’s always been possible to use wildcard and subject-alternate-name (SAN) certificates with ALB, but these come with limitations. Wildcard certificates only work for related subdomains that match a simple pattern and while SAN certificates can support many different domains, the same certificate authority has to authenticate each one. That means you have re-authenticate and re-provision your certificate every time you add a new domain.

The load balancer uses a smart certificate selection algorithm with support for SNI. If the hostname provided by a client matches a single certificate in the certificate list, the load balancer selects this certificate. If a hostname provided by a client matches multiple certificates in the certificate list, the load balancer selects the best certificate that the client can support. Certificate selection is based on the following criteria in the following order:
a – Public key algorithm (prefer ECDSA over RSA)
b – Hashing algorithm (prefer SHA over MD5)
c – Key length (prefer the largest)
d – Validity period


AWS Global Accelerator is a service that improves the availability and performance of your applications with local or global users. It provides static IP addresses that act as a fixed entry point to your application endpoints in a single or multiple AWS Regions, such as your Application Load Balancers, Network Load Balancers or Amazon EC2 instances.
AWS Global Accelerator uses the AWS global network to optimize the path from your users to your applications, improving the performance of your traffic by as much as 60%. You can test the performance benefits from your location with a speed comparison tool. AWS Global Accelerator continually monitors the health of your application endpoints and redirects traffic to healthy endpoints in less than 30 seconds.

A VPC endpoint enables you to privately connect your VPC to supported AWS services and VPC endpoint services powered by AWS PrivateLink without requiring an internet gateway, NAT device, VPN connection, or AWS Direct Connect connection. Instances in your VPC do not require public IP addresses to communicate with resources in the service. Traffic between your VPC and the other service does not leave the Amazon network.
Endpoints are virtual devices. They are horizontally scaled, redundant, and highly available VPC components. They allow communication between instances in your VPC and services without imposing availability risks or bandwidth constraints on your network traffic.
There are two types of VPC endpoints: interface endpoints and gateway endpoints. Create the type of VPC endpoint required by the supported service.


Amazon Simple Storage Service (Amazon S3) provides secure, durable, highly scalable cloud storage. You can configure an S3 bucket to host a static website that can include web pages and client-side scripts. (S3 doesn’t support server-side scripting.)
To route domain traffic to an S3 bucket, use Amazon Route 53 to create an alias record that points to your bucket. An alias record is a Route 53 extension to DNS. It’s similar to a CNAME record, except you can create an alias record both for the root domain, such as example.com, and for subdomains, such as http://www.example.com. You can create CNAME records only for subdomains.

– You can attach a network interface to an instance when it’s running (hot attach), when it’s stopped (warm attach), or when the instance is being launched (cold attach).
– You can detach secondary network interfaces when the instance is running or stopped. However, you can’t detach the primary network interface.
– You can move a network interface from one instance to another, if the instances are in the same Availability Zone and VPC but in different subnets.
– When launching an instance using the CLI, API, or an SDK, you can specify the primary network interface and additional network interfaces.
– Launching an Amazon Linux or Windows Server instance with multiple network interfaces automatically configures interfaces, private IPv4 addresses, and route tables on the operating system of the instance.
– A warm or hot attach of an additional network interface may require you to manually bring up the – second interface, configure the private IPv4 address, and modify the route table accordingly. Instances running Amazon Linux or Windows Server automatically recognize the warm or hot attach and configure themselves.
– Attaching another network interface to an instance (for example, a NIC teaming configuration) cannot be used as a method to increase or double the network bandwidth to or from the dual-homed instance.
– If you attach two or more network interfaces from the same subnet to an instance, you may encounter networking issues such as asymmetric routing. If possible, use a secondary private IPv4 address on the primary network interface instead.
- if the deleted object is still existing in Amazon S3 bucket because Amazon S3 DELETE requests are eventually consistent.
- encryption keys which will be generated and managed on premises = Use SSE-C
- cost-effective and fault-tolerant to handle application scalability  = Application load balancer + Auto Scaling group
- FSx for lustre is the optimum file system to handle ML workloads.
- to provide high scalability + distribute the traffic = Use ELB Application Load Balancer
- to provide one-way Internet access = Use NAT Gateway.
- Protecting the infrastructure is the main responsibility at Amazon Side.
- By default, AWS has a limit of 20 instances per region. This includes all instances set up on your AWS account.
- To increase EC2 limits, request a higher limit by providing information about the new limit and regions where it should be applied.
- separate the web tier from the application tier “Decouple” + asynchronously processing = Amazon SQS
- Spread placement groups strictly places a small group of instances across distinct underlying hardware to reduce correlated 
- the job run in monthly basis + without interruption = Use Scheduled Reserved Instances
- Data size is 100 TB = Amazon Snowball Edge.
- to store, access, and automatically rotate API keys = use use AWS Systems Manager Parameter Store
- to enhance the performance of Kinesis Data Streams, you have to increase number of shards using UpdateShardCount
- can be accessed and shared across multiple VPC +  store up to 3300 keys+ integrated with AWS CloudTrail + support MFA = AWS CloudHSM
- to audit log of any changes made to AWS resources in their account = Use AWS CloudTrail
- to encrypt data at rest in Amazon S3 = use client side encryption before saving in S3 or use SSE-C
- to enhance database performance  while reading static data = Use Amazon ElastiCache
- Fully Managed + DynamoDB + low-latency reads = Amazon DynamoDB Accelerator (DAX)
- The default termination policy terminates instance in the AZ with most number of instances
- 




The default termination policy is designed to help ensure that your instances span Availability Zones evenly for high availability. The default policy is kept generic and flexible to cover a range of scenarios.
The default termination policy behavior is as follows:
  1 – Determine which Availability Zones have the most instances, and at least one instance that is not protected from scale in.
  2 – Determine which instances to terminate so as to align the remaining instances to the allocation strategy for the On-Demand or Spot Instance that is terminating. This only applies to an Auto Scaling group that specifies allocation strategies.
For example, after your instances launch, you change the priority order of your preferred instance types. When a scale-in event occurs, Amazon EC2 Auto Scaling tries to gradually shift the On-Demand Instances away from instance types that are lower priority.
  3 – Determine whether any of the instances use the oldest launch template or configuration:
          a – [For Auto Scaling groups that use a launch template] ,Determine whether any of the instances use the oldest launch template unless there are instances that use a launch configuration. Amazon EC2 Auto Scaling terminates instances that use a launch configuration before instances that use a launch template.
          b – [For Auto Scaling groups that use a launch configuration] Determine whether any of the instances use the oldest launch configuration.
   4 – After applying all of the above criteria, if there are multiple unprotected instances to terminate, determine which instances are closest to the next billing hour. If there are multiple unprotected instances closest to the next billing hour, terminate one of these instances at random.
Note that terminating the instance closest to the next billing hour helps you maximize the use of your instances that have an hourly charge. Alternatively, if your Auto Scaling group uses Amazon Linux or Ubuntu, your EC2 usage is billed in one-second increments.

