
You can protect data in transit by using SSL or by using client-side encryption

Elastic Fabric Adapter is just a network device that you can attach to your Amazon EC2 instance to accelerate High Performance Computing (HPC) and machine learning applications. EFA enables you to achieve the application performance of an on-premises HPC cluster, with the scalability, flexibility, and elasticity provided by AWS

Individual Amazon S3 objects can range in size from a minimum of 0 bytes to a maximum of 5 terabytes. The largest object that can be uploaded in a single PUT is 5 gigabytes.

Redshift is primarily used for OLAP (Online Analytical Processing) applications

AWS Config provides a detailed view of the configuration of AWS resources in your AWS account. This includes how the resources are related to one another and how they were configured in the past so that you can see how the configurations and relationships change over time.

Creating a RAID 0 array allows you to achieve a higher level of performance for a file system than you can provision on a single Amazon EBS volume. Use RAID 0 when I/O performance is of the utmost importance. With RAID 0, I/O is distributed across the volumes in a stripe. If you add a volume, you get the straight addition of throughput and IOPS. However, keep in mind that performance of the stripe is limited to the worst performing volume in the set, and that the loss of a single volume in the set results in a complete data loss for the array.

An entity that can take an action on an AWS resource is called a principal


AWS also offers enhanced networking features to provide high-performance networking capabilities by using
an Elastic Network Adapter or an Intel 82599 Virtual Function (VF) interface. If you have a High-Performance
Computing workload or machine learning applications, you can attach an Elastic Fabric Adapter to your
instance to provide a higher network throughput than your regular TCP transport.

To protect your data against multiple availability zone failures—­or the failure of an entire region—­you can enable cross-­region replication between  source bucket in one region and a destination bucket in another. Note that cross-­region replication requires versioning to be enabled on both buckets. Also, deletes on the source bucket don’t get replicated.

One of the easiest ways to back up an EBS volume is to take a snapshot of it. AWS stores EBS snapshots across multiple availability zones in S3. You can either create a snapshot manually or use the Amazon Data Lifecycle Manager to automatically create a snapshot for you at regular intervals.


Amazon Elastic File System (Amazon EFS) provides simple, scalable file storage for use with your Amazon ECS tasks. With Amazon EFS, storage capacity is elastic, growing and shrinking automatically as you add and remove files. Your applications can have the storage they need when they need it.

IAM roles are not connected to subnets; they are connected directly—in this example to the EC2 instance. Peering connections cannot be created for direct connection between instances and AWS services in different regions.


AWS Resource Access Manager enables you to share resources hosted on subnets to AWS accounts that are part of an AWS Organization. AWS Control Tower is utilized to create a landing zone for new AWS accounts. IAM policies control access to database resources, but do not share database resources. A service control policy assigned to the master account would apply to all AWS accounts in the AWS organization.

The database security group should only allow incoming traffic from the web server security group. The web servers should be placed in the public subnet and the SQL database instances should be placed on a private subnet. Placing both sets of instances on a public subnet provides no security at all. The database security group should never allow incoming traffic from the Internet directly.


GuardDuty performs the collective monitoring and analysis of CloudTrail. 
VPC flow logs of network traffic, DNS records, and S3 API activity.
Amazon VPC flow logs enable you to monitor the network traffic going to and from network interfaces in a VPC.
AWS Inspector monitors and alerts about potential problems with EC2 instances. 
AWS Macie is responsible for monitoring S3 content and user access.




Amazon S3 is an object storage solution and does not provide a file system interface
CloudTrail is used for auditing not performance monitoring.

If your workload is mainly sending GET requests, you should consider using Amazon CloudFront for performance optimization. By integrating CloudFront with Amazon S3, you can distribute content to your users with low latency and a high data transfer rate.

RAID can be used to increase IOPS, however RAID 1 does not. For example:
    – RAID 0 = 0 striping – data is written across multiple disks and increases performance but no redundancy.
    – RAID 1 = 1 mirroring – creates 2 copies of the data but does not increase performance, only redundancy.

Amazon S3 Select is designed to help analyze and process data within an object in Amazon S3 buckets, faster and cheaper. It works by providing the ability to retrieve a subset of data from an object in Amazon S3 using simple SQL expressions

Amazon Redshift Spectrum allows you to directly run SQL queries against exabytes of unstructured data in Amazon S3. No loading or transformation is required.

DynamoDB DAX is an in-memory cache that increases the performance of DynamoDB. However, it costs money and there is no requirement to increase performance.

Security Groups you can only assign permit rules, you cannot assign deny rules.
NACLs you can have permit and deny rules.


You can create AMIs of the EC2 instances and then copy them across Regions. This provides a point-in-time copy of the state of the EC2 instance in the remote Region.
Once you’ve created AMIs of EC2 instances and copied them to the second Region, you can then launch the EC2 instances from the AMIs in that Region.
This is a good DR strategy as you have moved stateful EC2 instances to another Region.


DynamoDB offers consistent single-digit millisecond latency. However, DynamoDB + DAX further increases performance with response times in microseconds for millions of requests per second for read-heavy workloads.
The DAX cache uses cluster nodes running on Amazon EC2 instances and sits in front of the DynamoDB table

Amazon S3 Transfer Acceleration is used for speeding up uploads of data to Amazon S3 by using the CloudFront network. It is not used for downloading data.


AWS Global Accelerator is a service that improves the availability and performance of applications with local or global users. You can configure the ALB as a target and Global Accelerator will automatically route users to the closest point of presence.

RedShift is a columnar data warehouse DB that is ideal for running long complex queries. RedShift can also improve performance for repeat queries by caching the result and returning the cached result when queries are re-run. Dashboard, visualization, and business intelligence (BI) tools that execute repeat queries see a significant boost in performance due to result caching.


Amazon EFS is file-based storage system it is not object-based.
Amazon S3 is an object-based storage system. 


You can use placement groups to influence the placement of a group of interdependent instances to meet the needs of your workload. Depending on the type of workload, you can create a placement group using one of the following placement strategies:

    Cluster – packs instances close together inside an Availability Zone. This strategy enables workloads to achieve the 
              low-latency network performance necessary for tightly-coupled node-to-node communication that is typical of HPC applications.

    Partition – spreads your instances across logical partitions such that groups of instances in one partition do not 
                share the underlying hardware with groups of instances in different partitions. 
                This strategy is typically used by large distributed and replicated workloads, such as Hadoop, Cassandra, and Kafka.

    Spread – strictly places a small group of instances across distinct underlying hardware to reduce correlated failures.

    Cluster placement groups are recommended for applications that benefit from low network latency, high network throughput, or both. They are also recommended when the majority of the network traffic is between the instances in the group. To provide the lowest latency and the highest packet-per-second network performance for your placement group, choose an instance type that supports enhanced networking.

Auto Scaling can perform rebalancing when it finds that the number of instances across AZs is not balanced. 
Auto Scaling rebalances by launching new EC2 instances in the AZs that have fewer instances first, only then will it start 
terminating instances in AZs that had more instances

Auto Scaling can be configured to send an SNS email when:
    – An instance is launched.
    – An instance is terminated.
    – An instance fails to launch.
    – An instance fails to terminate.

Scheduled Reserved Instances (Scheduled Instances) enable you to purchase capacity reservations that recur on a daily, 
weekly, or monthly basis, with a specified start time and duration, for a one-year term. You reserve the capacity in advance, 
so that you know it is available when you need it. You pay for the time that the instances are scheduled, even if you do not use them.
Scheduled Instances are a good choice for workloads that do not run continuously, but do run on a regular schedule.
For example, you can use Scheduled Instances for an application that runs during business hours or for batch processing that runs at the end of the week.



When data is stored in an encrypted state it is referred to as encrypted “at rest” and when it is encrypted as it is being transferred over a network it is referred to as encrypted “in transit”. You can securely upload/download your data to Amazon S3 via SSL endpoints using the HTTPS protocol (In Transit – SSL/TLS).
You have the option of encrypting the data locally before it is uploaded or uploading using SSL/TLS so it is secure in transit and encrypting on the Amazon S3 side using S3 managed keys. The S3 managed keys will be AES-256 (not AES-128) bit keys
Uploading data using CloudFront with an EC2 origin or using an encrypted EBS volume attached to an EC2 instance is not a solution to this problem as your company wants to backup these records onto S3 (not EC2/EBS).


Multi-AZ RDS creates a replica in another AZ and synchronously replicates to it (DR only).
A failover may be triggered in the following circumstances:
    – Loss of primary AZ or primary DB instance failure
    – Loss of network connectivity on primary
    – Compute (EC2) unit failure on primary
    – Storage (EBS) unit failure on primary
    – The primary DB instance is changed
    – Patching of the OS on the primary DB instance
    – Manual failover (reboot with failover selected on primary)
During failover RDS automatically updates configuration (including DNS endpoint) to use the second node.


The NLB and CLB types of Elastic Load Balancer do not support path-based routing or host-based routing so they cannot be used for this use case.

AWS Shield provides protection against distributed denial of service (DDoS) attacks for AWS resources, at the network and transport layers (layer 3 and 4) and the application layer (layer 7).


