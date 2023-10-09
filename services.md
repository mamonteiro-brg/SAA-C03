

Amazon Detective
    Amazon Detective is commonly used to analyze, investigate, and quickly identify the root cause of potential security issues in your AWS workloads, as well as for detecting suspicious activities.

Amazon GuardDuty
    Amazon GuardDuty is a threat detection service that continuously monitors your AWS accounts and workloads for malicious activity and delivers detailed security findings for visibility and remediation.

    Amazon GuardDuty analyzes and processes data from foundational data sources to detect anomalies involving AWS Identity and Access Management (IAM) access keys and Amazon Elastic Compute Cloud (Amazon EC2). You can also activate GuardDuty protection plans to analyze additional data from other AWS services in your AWS environment to protect workloads using Amazon S3, Amazon EKS, Amazon RDS, and AWS Lambda.

    Can generate findings based on suspicious activities such as requests coming from known malicious IP addresses, changing of bucket policies/ACLs to expose an S3 bucket publicly, or suspicious API call patterns that attempt to discover misconfigured bucket permissions.

    To detect possibly malicious behavior, GuardDuty uses a combination of anomaly detection, machine learning, and continuously updated threat intelligence.

Amazon Macie
    Amazon Macie is a data security service that discovers sensitive data by using machine learning and pattern matching, provides visibility into data security risks, and enables automated protection against those risks.

    To help you manage the security posture of your organization's Amazon Simple Storage Service (Amazon S3) data estate, Macie provides you with an inventory of your S3 buckets, and automatically evaluates and monitors the buckets for security and access control. If Macie detects a potential issue with the security or privacy of your data, such as a bucket that becomes publicly accessible, Macie generates a finding for you to review and remediate as necessary.


Amazon Monitron
    Amazon Monitron is simply a service that detects abnormal conditions in industrial equipment such as fans, compressors, motors, etc.


AWS Database Migration Service (DMS)
    Is a cloud service that makes it easy to migrate relational databases, data warehouses, NoSQL databases, and other types of data stores. You can use AWS DMS to migrate your data into the AWS Cloud or between combinations of cloud and on-premises setups.

    At a basic level, AWS DMS is a server in the AWS Cloud that runs replication software. You create a source and target connection to tell AWS DMS where to extract data from and where to load it. Next, you schedule a task that runs on this server to move your data. AWS DMS creates the tables and associated primary keys if they don't exist on the target. You can create the target tables yourself if you prefer. Or you can use AWS Schema Conversion Tool (AWS SCT) to create some or all of the target tables, indexes, views, triggers, and so on.



AWS Backup
    AWS Backup is a fully-managed service that makes it easy to centralize and automate data protection across AWS services, in the cloud, and on premises. Using this service, you can configure backup policies and monitor activity for your AWS resources in one place. It allows you to automate and consolidate backup tasks that were previously performed service-by-service, and removes the need to create custom scripts and manual processes. With a few clicks in the AWS Backup console, you can automate your data protection policies and schedules.

    AWS Backup does not govern backups you take in your AWS environment outside of AWS Backup. Therefore, if you want a centralized, end-to-end solution for business and regulatory compliance requirements, start using AWS Backup today.

AWS Config 
    Is a service that enables you to assess, audit, and evaluate the configurations of your AWS resources. Config continuously monitors and records your AWS resource configurations and allows you to automate the evaluation of recorded configurations against desired configurations. With Config, you can review changes in configurations and relationships between AWS resources, dive into detailed resource configuration histories, and determine your overall compliance against the configurations specified in your internal guidelines. This enables you to simplify compliance auditing, security analysis, change management, and operational troubleshooting.


AWS Trusted Advisor
    Amazon Trusted Advisor is an application that draws upon best practices learned from Amazon Web Services’ aggregated operational history of serving hundreds of thousands of Amazon Web Services customers. Trusted Advisor inspects your  Amazon Web Services environment and makes recommendations for saving money, improving system performance, or closing security gaps.



AWS Organizations 
    Is an account management service that enables you to consolidate multiple AWS accounts into an organization that you create and centrally manage. AWS Organizations includes account management and consolidated billing capabilities that enable you to better meet the budgetary, security, and compliance needs of your business. As an administrator of an organization, you can create accounts in your organization and invite existing accounts to join the organization.

    AWS Organizations offers the following features:
        Centralized management of all of your AWS accounts



Amazon EMR (previously called Amazon Elastic MapReduce)
    Is a managed cluster platform that simplifies running big data frameworks, such as Apache Hadoop and Apache Spark, on AWS to process and analyze vast amounts of data. Using these frameworks and related open-source projects, you can process data for analytics purposes and business intelligence workloads. Amazon EMR also lets you transform and move large amounts of data into and out of other AWS data stores and databases, such as Amazon Simple Storage Service (Amazon S3) and Amazon DynamoDB.


AWS Database Migration Service (AWS DMS) 
    Is a cloud service that makes it easy to migrate relational databases, data warehouses, NoSQL databases, and other types of data stores. You can use AWS DMS to migrate your data into the AWS Cloud, between on-premises instances (through an AWS Cloud setup) or between combinations of cloud and on-premises setups. With AWS DMS, you can perform one-time migrations, and you can replicate ongoing changes to keep sources and targets in sync.


AWS Secrets Manager 
    Is an AWS service that makes it easier for you to manage secrets. Secrets can be database credentials, passwords, third-party API keys, and even arbitrary text. You can store and control access to these secrets centrally by using the Secrets Manager console, the Secrets Manager command line interface (CLI), or the Secrets Manager API and SDKs.




Amazon Translate 
    Is a Neural Machine Translation (MT) service for translating text between supported languages.


Amazon Transcribe 
    Is an AWS service that makes it easy for customers to convert speech-to-text. Using Automatic Speech Recognition (ASR) technology, customers can choose to use Amazon Transcribe for a variety of business applications, including transcription of voice-based customer service calls, generation of subtitles on audio/video content, and conduct (text-based) content analysis on audio/video content.

Amazon Comprehend 
    Amazon Comprehend is a natural language processing (NLP) service that uses machine learning to find insights and relationships in text, no machine learning experience is required. Amazon Comprehend uses machine learning to help you uncover the insights and relationships in your unstructured data. The service identifies the language of the text; extracts key phrases, places, people, brands, or events; understands how positive or negative the text is; analyzes text using tokenization and parts of speech; and automatically organizes a collection of text files by topic. You can also use AutoML capabilities in Amazon Comprehend to build a custom set of entities or text classification models that are tailored uniquely to your organization’s needs. To learn more, see Amazon Comprehend.



Amazon DocumentDB (with MongoDB compatibility) 
    Is a fast, scalable, highly available, and fully managed document database service that supports MongoDB workloads. The Amazon DocumentDB Migration Guide outlines three primary approaches for migrating from MongoDB to Amazon DocumentDB: offline, online, and hybrid.



Amazon S3 gateway endpoint
    You can access Amazon S3 from your VPC using gateway VPC endpoints. After you create the gateway endpoint, you can add it as a target in your route table for traffic destined from your VPC to Amazon S3.

    There is no additional charge for using gateway endpoints.



AWS AppSync 
    Enables developers to connect their applications and services to data and events with secure, serverless and high-performing GraphQL and Pub/Sub APIs. You can do the following with AWS AppSync:

        Access data from one or more data sources from a single GraphQL API endpoint.
        Combine multiple source GraphQL APIs into a single, merged GraphQL API.
        Publish real-time data updates to your applications.
        Leverage built-in security, monitoring, logging, and tracing, with optional caching for low latency.
        Only pay for API requests and any real-time messages that are delivered.

AWS Fargate 
    Is a serverless compute engine for containers that works with both Amazon Elastic Container Service (ECS) and Amazon Elastic Kubernetes Service (EKS). Fargate makes it easy for you to focus on building your applications. Fargate removes the need to provision and manage servers, lets you specify and pay for resources per application, and improves security through application isolation by design.

Service control policies (SCPs) 
    Are a type of organization policy that you can use to manage permissions in your organization. SCPs offer central control over the maximum available permissions for all accounts in your organization. SCPs help you to ensure your accounts stay within your organization’s access control guidelines. SCPs are available only in an organization that has all features enabled. SCPs aren't available if your organization has enabled only the consolidated billing features. For instructions on enabling SCPs, see Enabling and disabling policy types.


SMB network file share 



Elastic Fabric Adapter (EFA)
    Elastic Fabric Adapter (EFA) is a network interface for Amazon EC2 instances that enables customers to run applications requiring high levels of inter-node communications at scale on AWS. Its custom-built operating system (OS) bypass hardware interface enhances the performance of inter-instance communications, which is critical to scaling these applications. With EFA, High Performance Computing (HPC) applications using the Message Passing Interface (MPI) and Machine Learning (ML) applications using NVIDIA Collective Communications Library (NCCL) can scale to thousands of CPUs or GPUs. As a result, you get the application performance of on-premises HPC clusters with the on-demand elasticity and flexibility of the AWS cloud.


Amazon EventBridge

Placement groups
    When you launch a new EC2 instance, the EC2 service attempts to place the instance in such a way that all of your instances are spread out across underlying hardware to minimize correlated failures. You can use placement groups to influence the placement of a group of interdependent instances to meet the needs of your workload. Depending on the type of workload, you can create a placement group using one of the following placement strategies:

        Cluster – packs instances close together inside an Availability Zone. This strategy enables workloads to achieve the low-latency network performance necessary for tightly-coupled node-to-node communication that is typical of high-performance computing (HPC) applications.

        Partition – spreads your instances across logical partitions such that groups of instances in one partition do not share the underlying hardware with groups of instances in different partitions. This strategy is typically used by large distributed and replicated workloads, such as Hadoop, Cassandra, and Kafka.

        Spread – strictly places a small group of instances across distinct underlying hardware to reduce correlated failures.


AWS Fargate
Amazon Lex
Amazon Polly

Server Name Indication (SNI)
Subject Alternative Name (SAN)


 AWS Control Tower
    AWS Control Tower offers a straightforward way to set up and govern an AWS multi-account environment, following prescriptive best practices. AWS Control Tower orchestrates the capabilities of several other AWS services, including AWS Organizations, AWS Service Catalog, and AWS IAM Identity Center (successor to AWS Single Sign-On), to build a landing zone in less than an hour. Resources are set up and managed on your behalf.


AWS Migration Service (AWS MGN) 
    AWS Application Migration Service (MGN) is a highly automated lift-and-shift (rehost) solution that simplifies, expedites, and reduces the cost of migrating applications to AWS. It allows companies to lift-and-shift a large number of physical, virtual, or cloud servers without compatibility issues, performance disruption, or long cutover windows. MGN replicates source servers into your AWS account. When you’re ready, it automatically converts and launches your servers on AWS so you can quickly benefit from the cost savings, productivity, resilience, and agility of the Cloud. Once your applications are running on AWS, you can leverage AWS services and capabilities to quickly and easily replatform or refactor those applications – which makes lift-and-shift a fast route to modernization.



AWS DataSync
    AWS DataSync moves large amounts of data online between on-premises storage and Amazon S3, Amazon Elastic File System (Amazon Elastic File System) or Amazon FSx. Manual tasks related to data transfers can slow down migrations and burden IT operations. DataSync eliminates or automatically handles many of these tasks, including scripting copy jobs, scheduling and monitoring transfers, validating data, and optimizing network utilization. The DataSync software agent connects to your Network File System (NFS) and Server Message Block (SMB) storage, so you don’t have to modify your applications. DataSync can transfer hundreds of terabytes and millions of files at speeds up to 10 times faster than open-source tools, over the internet or AWS Direct Connect links. You can use DataSync to migrate active data sets or archives to AWS, transfer data to the cloud for timely analysis and processing, or replicate data to AWS for business continuity.

    Is an online data movement and discovery service that simplifies data migration and helps you quickly, easily, and securely transfer your file or object data to, from, and between AWS storage services.

    You can use AWS DataSync to transfer files from an existing file system to Amazon EFS. AWS DataSync is a data transfer service that simplifies, automates, and accelerates moving and replicating data between on-premises storage systems and AWS storage services over the internet or AWS Direct Connect. AWS DataSync can transfer your file data, and also file system metadata such as ownership, timestamps, and access permissions.

AWS Replication Agent
    The AWS Replication Agent performs an initial block-level read of the content of any volume attached to the server and replicates it to the replication server. The Agent then acts as an OS-level read filter to capture writes and synchronizes any block level modifications to the Elastic Disaster Recovery replication server, ensuring near-zero RPO.

Amazon EventBridge
AWS Personal Health Dashboard events
    AWS Personal Health Dashboard provides alerts and guidance for AWS events that might affect your environment. 

AWS Service Health Dashboard events
    AWS Service Health Dashboard shows public events that may affect several customers in particular regions. It doesn’t show events related to specific EC2 instances on individual AWS accounts. You have to check the events on the AWS Personal Health Dashboard instead.

AWS Health 
    Provides ongoing visibility into your resource performance and the availability of your AWS services and accounts. You can use AWS Health events to learn how service and resource changes might affect your applications running on AWS. AWS Health provides relevant and timely information to help you manage events in progress. AWS Health also helps you be aware of and to prepare for planned activities. The service delivers alerts and notifications triggered by changes in the health of AWS resources, so that you get near-instant event visibility and guidance to help accelerate troubleshooting.

AWS CloudHSM

    AWS CloudHSM is a cloud-hosted Hardware Security Module (HSM) service that allows you to perform cryptographic operations and host encryption keys. It gives customers an extra level of protection for data with contractual, strict corporate and regulatory compliance requirements. It let you generate and use your encryption keys on AWS Cloud.

    With CloudHSM, encryption keys are managed using FIPS 140-2 Level 3 validated HSMs. It is managed by Amazon, but the keys are controlled and managed by the customer only.

    AWS CloudHSM service helps you meet corporate, contractual, and regulatory compliance requirements for data security by using dedicated Hardware Security Module (HSM) instances within the AWS cloud. AWS and AWS Marketplace partners offer a variety of solutions for protecting sensitive data within the AWS platform, but for some applications and data subject to contractual or regulatory mandates for managing cryptographic keys, additional protection may be necessary. CloudHSM complements existing data protection solutions and allows you to protect your encryption keys within HSMs that are designed and validated to government standards for secure key management. CloudHSM allows you to securely generate, store, and manage cryptographic keys used for data encryption in a way that keys are accessible only by you.




VPC Flow Logs
    VPC flow logs enable you to capture information about the IP traffic going to and from a VPC. Flow logs can be used to monitor, troubleshoot, and analyze the network traffic in your VPC.

    VPC flow logs can be enabled at the VPC, subnet, or elastic network interface level, capturing traffic flowing in and out of the specified resource. Flow logs record the IP traffic flowing in and out of your VPC, including information about the source and destination IP addresses, ports, protocols, and packet and byte counts.


Amazon Inspector
    Amazon Inspector allows you to test the security levels of instances you have deployed. After you define an assessment target for Amazon Inspector, which is a group of tagged EC2 instances, Amazon Inspector evaluates the state of each instance by using several rule packages.

    Amazon Inspector uses two types of rules: network accessibility tests that don’t require the Inspector agent to be installed, and host assessment rules that require the Inspector agent to be installed 


Amazon SQS
    SQS queues have a maximum message size of 256KB. You can use the extended client library for Java to use pointers to a payload on S3 but the maximum payload size is 2GB.


VPC Endpoints
    Gateway Endpoints
        This Gateway Endpoints only supports Amazon S3 and DynamoDB
    Interface Endpoints
        Amazon S3
        Amazon SQS
        AWS EC2
        ...
    Gateway Load Balancer Endpoints

VPC Endpoints Services



Amazon FSx For Lustre 
    Is a high-performance ﬁle system for fast processing of workloads. 
    Lustre is a popular open-source parallel ﬁle system which stores data across multiple network ﬁle servers to maximize performance and reduce bottlenecks.

    Scratch file systems are meant to be used for more short-term data processing and temporary data storage. The system does not replicate scratch data, which means it can be lost if a file server malfunctions. The advantage of scratch file systems is that they provide excellent throughput—a big burst that can equal up to six times the standard baseline of 200 MBps per TiB (equivalent to just over a TB) of storage capacity.

    Best use cases for scratch file systems include cost-effective storage for workloads that are heavy on processing and only needed for a short period of time.

    Persistent file systems are meant to be used for workloads that need to be stored for a longer period of time. This storage type, if highly available and stored, is replicated automatically in the AWS Availability Zone where the file system is located. The advantage here is that if a server fails, stored data is replaced in just minutes.



Cross-Region Replication of snapshot backups.




DynamoDB Streams 
    Captures a time-ordered sequence of item-level modifications in any DynamoDB table and stores this information in a log for up to 24 hours. Applications can access this log and view the data items as they appeared before and after they were modified, in near-real time.

    A DynamoDB stream is an ordered flow of information about changes to items in a DynamoDB table. When you enable a stream on a table, DynamoDB captures information about every modification to data items in the table.

Enhanced Networking
    ENIs are not the only option for virtual network interfaces. Enhanced networking offers higher network throughput speeds and lower latency than ENIs. Enhanced networking uses single-root input/output virtualization (SR-IOV) to allow an instance direct access to the physical network interface on the host, thus bypassing the hypervisor and resulting in lower CPU utilization and better network performance.

AWS ENA (Elastic Network Adapter) 
    is a custom network interface optimized to deliver high throughput and packet per second (PPS) performance, and consistently low latencies on EC2 instances.

    The Elastic Network Adapter (ENA) is designed to provide Enhanced Networking to your EC2 instances. 

    With ENA, you can expect high throughput and packet per second (PPS) performance, as well as consistently low latencies on Amazon EC2 instances. Using ENA, you can utilize up to 20 Gbps of network bandwidth on certain EC2 instance types – massively improving your networking throughput compared to other EC2 instances, or on premises machines. ENA-based Enhanced Networking is currently supported on X1 instances.


AWS CloudTrail
    To embrace the DevOps principles of collaboration, communication, and transparency, it’s important to understand who is making modifications to your infrastructure. In AWS, this transparency is provided by AWS CloudTrail. All AWS interactions are handled through AWS API calls that are monitored and logged by AWS CloudTrail. All generated log files are stored in an Amazon S3 bucket that you define. Log files are encrypted using Amazon S3 server-side encryption (SSE). All API calls are logged whether they come directly from a user or on behalf of a user by an AWS service. Numerous groups can benefit from CloudTrail logs, including operations teams for support, security teams for governance, and finance teams for billing.



VPC Reachability Analyzer
    Reachability Analyzer is a configuration analysis tool that enables you to perform connectivity testing between a source resource and a destination resource in your virtual private clouds (VPCs). When the destination is reachable, Reachability Analyzer produces hop-by-hop details of the virtual network path between the source and the destination. When the destination is not reachable, Reachability Analyzer identifies the blocking component. For example, paths can be blocked by configuration issues in a security group, network ACL, route table, or load balancer.


What is Amazon VPC?

    With Amazon Virtual Private Cloud (Amazon VPC), you can launch AWS resources in a logically isolated virtual network that you've defined. This virtual network closely resembles a traditional network that you'd operate in your own data center, with the benefits of using the scalable infrastructure of AWS.

    A VPC is a virtual network that closely resembles a traditional network that you'd operate in your own data center. After you create a VPC, you can add subnets.


Subnets
    A subnet is a range of IP addresses in your VPC. A subnet must reside in a single Availability Zone. After you add subnets, you can deploy AWS resources in your VPC.


AWS Storage Gateway
    AWS Storage Gateway is a hybrid cloud storage service that gives you on-premises access to virtually unlimited cloud storage. You can use Storage Gateway to simplify storage management and reduce costs for key hybrid cloud storage use cases. These include moving backups to the cloud, using on-premises file shares backed by cloud storage, and providing low-latency access to data in AWS for on-premises applications.

    To support these use cases, the service provides four different types of gateways – Tape Gateway, Amazon S3 File Gateway, Amazon FSx File Gateway, and Volume Gateway – that seamlessly connect on-premises applications to cloud storage, caching data locally for low-latency access.

AWS DataSync 

    AWS DataSync  moves large amounts of data online between on-premises storage and Amazon S3, Amazon Elastic File System (Amazon Elastic File System) or Amazon FSx. Manual tasks related to data transfers can slow down migrations and burden IT operations. DataSync eliminates or automatically handles many of these tasks, including scripting copy jobs, scheduling and monitoring transfers, validating data, and optimizing network utilization. The DataSync software agent connects to your Network File System (NFS) and Server Message Block (SMB) storage, so you don’t have to modify your applications. DataSync can transfer hundreds of terabytes and millions of files at speeds up to 10 times faster than open-source tools, over the internet or AWS Direct Connect links. You can use DataSync to migrate active data sets or archives to AWS, transfer data to the cloud for timely analysis and processing, or replicate data to AWS for business continuity.


What is AWS Resource Access Manager?
    AWS Resource Access Manager (AWS RAM) helps you securely share your resources across AWS accounts, within your organization or organizational units (OUs), and with AWS Identity and Access Management (IAM) roles and users for supported resource types. If you have multiple AWS accounts, you can create a resource once and use AWS RAM to make that resource usable by those other accounts. If your account is managed by AWS Organizations, you can share resources with all the other accounts in the organization or only those accounts contained by one or more specified organizational units (OUs). You can also share with specific AWS accounts by account ID, regardless of whether the account is part of an organization. Some supported resource types also let you share them with specified IAM roles and users.


Amazon WorkSpaces
    Amazon WorkSpaces enables you to provision virtual, cloud-based Microsoft Windows, Amazon Linux, or Ubuntu Linux desktops for your users, known as WorkSpaces. WorkSpaces eliminates the need to procure and deploy hardware or install complex software. You can quickly add or remove users as your needs change. Users can access their virtual desktops from multiple devices or web browsers.



What Is Amazon EventBridge?
    EventBridge is a serverless service that uses events to connect application components together, making it easier for you to build scalable event-driven applications. Event-driven architecture is a style of building loosely-coupled software systems that work together by emitting and responding to events. Event-driven architecture can help you boost agility and build reliable, scalable applications.

    Use EventBridge to route events from sources such as home-grown applications, AWS services, and third-party software to consumer applications across your organization. EventBridge provides simple and consistent ways to ingest, filter, transform, and deliver events so you can build applications quickly.

AWS Systems Manager Run Command

Amazon Neptune
    Amazon Neptune is a fast, reliable, fully managed graph database service that makes it easy to build and run applications that work with highly connected datasets. The core of Neptune is a purpose-built, high-performance graph database engine. This engine is optimized for storing billions of relationships and querying the graph with milliseconds latency. Neptune supports the popular property-graph query languages Apache TinkerPop Gremlin and Neo4j's openCypher, and the W3C's RDF query language, SPARQL. This enables you to build queries that efficiently navigate highly connected datasets. Neptune powers graph use cases such as recommendation engines, fraud detection, knowledge graphs, drug discovery, and network security.

    Neptune is highly available, with read replicas, point-in-time recovery, continuous backup to Amazon S3, and replication across Availability Zones. Neptune provides data security features, with support for encryption at rest and in transit. Neptune is fully managed, so you no longer need to worry about database management tasks like hardware provisioning, software patching, setup, configuration, or backups.

AWS Glue
    AWS Glue is a fully managed extract, transform, and load (ETL) service that makes it easy for customers to prepare and load their data for analytics. You can create and run an ETL job with a few clicks in the AWS Management Console. You simply point AWS Glue to your data stored on AWS, and AWS Glue discovers your data and stores the associated metadata (e.g., table definition and schema) in the AWS Glue Data Catalog. Once cataloged, your data is immediately searchable, queryable, and available for ETL. AWS Glue generates the code to execute your data transformations and data loading processes.

What is AWS Batch?
    AWS Batch is a set of batch management capabilities that enables developers, scientists, and engineers to easily and efficiently run hundreds of thousands of batch computing jobs on AWS. AWS Batch dynamically provisions the optimal quantity and type of compute resources (e.g., CPU or memory optimized compute resources) based on the volume and specific resource requirements of the batch jobs submitted. With AWS Batch, there is no need to install and manage batch computing software or server clusters, allowing you to instead focus on analyzing results and solving problems. AWS Batch plans, schedules, and executes your batch computing workloads using Amazon ECS, Amazon EKS, and AWS Fargate with an option to utilize spot instances.

Amazon Textract 
    Is just an AI service used to extract text data from scanned documents in PNG, JPEG, TIFF, PDF formats and is not capable of running sentiment analysis. 

Memcached with Auto Discovery
Redis Global Datastore
AWS Storage File gateway
SMB share

Amazon Pinpoint journey 
    In Amazon Pinpoint, a journey is a customized, multi-step engagement experience. When you create a journey, you start by choosing a segment that defines which customers will participate in the journey. After that, you add the activities that customers pass through on their journeys. Activities can include sending messages or splitting customers into groups based on their attributes or behaviors.

    There are several different types of journey activities, each with its own specific purpose. For example, you can add a Send email activity to your journey. When a customer arrives on this type of activity, they receive an email message. Another type of journey activity is the Multivariate split activity. When customers arrive on this type of activity, they are separated into multiple paths based on their segment membership or their interactions with previous journey activities. You can learn more about journey activities in Take a tour of journeys.

Amazon Lex 
    Enables you to build applications using a speech or text interface powered by the same technology that powers Amazon Alexa. Amazon Lex provides the deep functionality and flexibility of natural language understanding (NLU) and automatic speech recognition (ASR), so you can build highly engaging user experiences with lifelike conversational interactions and create new categories of products.

    Enables any developer to build conversational chatbots quickly. With Amazon Lex, no deep learning expertise is necessary—to create a bot, you just specify the basic conversation flow in the Amazon Lex console. The console provides a graphical user interface that you use to build a production-ready bot for your application.

AWS Proton 
    Allows you to deploy any serverless or container-based application with increased efficiency, consistency, and control. You can define infrastructure standards and effective continuous delivery pipelines for your organization. Proton breaks down the infrastructure into environment and service (“infrastructure as code” templates).

Amazon QuickSight 
    Is a fast, cloud-powered business intelligence service that delivers insights to everyone in your organization. As a fully managed service, Amazon QuickSight lets you easily create and publish interactive dashboards that include machine learning (ML) insights. To learn more, see Amazon QuickSight

AWS Compute Optimizer 
    Recommends optimal AWS resources for your workloads to reduce costs and improve performance by using machine learning to analyze historical utilization metrics. Overprovisioning resources can lead to unnecessary infrastructure costs, and underprovisioning resources can lead to poor application performance. Compute Optimizer generates recommendations for the following resources:
        -Amazon Elastic Compute Cloud (Amazon EC2) instances
        -Amazon EC2 Auto Scaling groups
        -Amazon Elastic Block Store (Amazon EBS) volumes
        -AWS Lambda functions


Amazon Quantum Ledger Database (Amazon QLDB) 
    Is a fully managed ledger database that provides a transparent, immutable, and cryptographically verifiable transaction log owned by a central trusted authority. Amazon QLDB can be used to track every application data change and maintains a complete and verifiable history of changes over time.

    With Amazon QLDB, you can trust that the history of changes to your application data is accurate. QLDB uses an immutable transactional log, known as a journal, for data storage. The journal tracks every change to your committed data and maintains a complete and verifiable history of changes over time.

Amazon Keyspaces (for Apache Cassandra) 
    Is a scalable, highly available, and managed Apache Cassandra–compatible database service. With Amazon Keyspaces, you can run your Cassandra workloads on AWS using the same Cassandra application code and developer tools that you use today. You don’t have to provision, patch, or manage servers, and you don’t have to install, maintain, or operate software. Amazon Keyspaces is serverless, so you pay for only the resources you use and the service can automatically scale tables up and down in response to application traffic. You can build applications that serve thousands of requests per second with virtually unlimited throughput and storage. Data is encrypted by default and Amazon Keyspaces enables you to back up your table data continuously using point-in-time recovery. Amazon Keyspaces gives you the performance, elasticity, and enterprise features you need to operate business-critical Cassandra workloads at scale

AWS Application Discovery Service 
    Helps you plan your migration to the AWS cloud by collecting usage and configuration data about your on-premises servers. Application Discovery Service is integrated with AWS Migration Hub, which simplifies your migration tracking as it aggregates your migration status information into a single console. You can view the discovered servers, group them into applications, and then track the migration status of each application from the Migration Hub console in your home region.

AWS Migration Hub (Migration Hub) 
    Provides a single place to discover your existing servers, plan migrations, and track the status of each application migration. The Migration Hub provides visibility into your application portfolio and streamlines planning and tracking. You can visualize the connections and the status of the servers and databases that make up each of the applications you are migrating.

    Migration Hub gives you the choice to start migrating right away and group servers while the migration is underway or to first discover servers and then group them into applications.

Use Amazon Data Lifecycle Manager (Amazon DLM) to:
    Automate the creation, retention, and deletion of snapshots taken to back up your Amazon EBS volumes. Automating snapshot management helps you to:
        – Protect valuable data by enforcing a regular backup schedule.
        – Retain backups as required by auditors or internal compliance.
        – Reduce storage costs by deleting outdated backups.

    You can use Amazon Data Lifecycle Manager to automate the creation, retention, and deletion of EBS snapshots and EBS-backed AMIs. When you automate snapshot and AMI management, it helps you to:
        Protect valuable data by enforcing a regular backup schedule.
        Create standardized AMIs that can be refreshed at regular intervals.
        Retain backups as required by auditors or internal compliance.
        Reduce storage costs by deleting outdated backups.
        Create disaster recovery backup policies that back up data to isolated accounts.



What Is Amazon EventBridge?
    EventBridge is a serverless service that uses events to connect application components together, making it easier for you to build scalable event-driven applications. Event-driven architecture is a style of building loosely-coupled software systems that work together by emitting and responding to events. Event-driven architecture can help you boost agility and build reliable, scalable applications.

    Use EventBridge to route events from sources such as home-grown applications, AWS services, and third-party software to consumer applications across your organization. EventBridge provides simple and consistent ways to ingest, filter, transform, and deliver events so you can build applications quickly.



https://aws.amazon.com/faqs/?nc1=f_dr


AWS Systems Manager Application Manager
    Application Manager, a capability of AWS Systems Manager, helps DevOps engineers investigate and remediate issues with their AWS resources in the context of their applications and clusters. Application Manager aggregates operations information from multiple AWS services and Systems Manager capabilities to a single AWS Management Console.

    In Application Manager, an application is a logical group of AWS resources that you want to operate as a unit. This logical group can represent different versions of an application, ownership boundaries for operators, or developer environments, to name a few. Application Manager support for container clusters includes both Amazon Elastic Kubernetes Service (Amazon EKS) and Amazon Elastic Container Service (Amazon ECS) clusters.

    When you choose Get started on the Application Manager home page, Application Manager automatically imports metadata about your resources that were created in other AWS services or Systems Manager capabilities. For applications, Application Manager imports metadata about all of your AWS resources organized into resource groups. Each resource group is listed in the Custom applications category as a unique application. Application Manager also automatically imports metadata about resources that were created by AWS CloudFormation, AWS Launch Wizard, Amazon ECS, and Amazon EKS. Application Manager then displays those resources in predefined categories.



What Is AWS Migration Hub?
    AWS Migration Hub (Migration Hub) provides a single place to discover your existing servers, plan migrations, and track the status of each application migration. The Migration Hub provides visibility into your application portfolio and streamlines planning and tracking. You can visualize the connections and the status of the servers and databases that make up each of the applications you are migrating, regardless of which migration tool you are using.

    Migration Hub gives you the choice to start migrating right away and group servers while migration is underway, or to first discover servers and then group them into applications. Either way, you can migrate each server in an application and track progress from each tool in the AWS Migration Hub.

    Migration Hub supports migration status updates from the following tools:

    AWS Application Migration Service (Application Migration Service)–AWS Application Migration Service is the primary migration service recommended for lift-and-shift migrations to AWS. For more information about Application Migration Service, see AWS Application Migration Service and Application Migration Service Documentation.

    AWS Database Migration Service (AWS DMS)–For more information about AWS DMS, see AWS Database Migration Service and AWS DMS Documentation.


AWS database schema conversion options
    AWS offers two schema conversion solutions to make heterogeneous database migrations predictable, fast, secure, and simple. Customers have the choice to: 1) log in to the AWS Database Migration Service (AWS DMS) console to initiate the AWS DMS Schema Conversion (DMS SC) workflow for a fully managed experience or 2) download the AWS Schema Conversion Tool (AWS SCT) software to their local drive.
    Both options will automatically assess and convert the source database schema and a majority of the database code objects, including views, stored procedures, and functions, to a format compatible with the target database. Any objects that cannot be automatically converted are clearly marked as action items with prescriptive instructions on how to convert, so that they can be manually converted to complete the migration.

    AWS SCT can also scan your application source codes for embedded SQL statements and convert them as part of a database-schema-conversion project. During this process, AWS SCT performs cloud-native code optimization by converting legacy Oracle and SQL Server functions to their equivalent AWS service, helping to modernize the applications at the same time of database migration. Once schema conversion is complete, it can help migrate data from a range of data warehouses to Amazon Redshift using built-in data migration agents.



AWS Trusted Advisor
    AWS Trusted Advisor provides recommendations that help you follow AWS best practices. Trusted Advisor evaluates your account by using checks. These checks identify ways to optimize your AWS infrastructure, improve security and performance, reduce costs, and monitor service quotas. You can then follow the recommendations to optimize your services and resources.
    AWS Basic Support and AWS Developer Support customers can access core security checks and checks for service quotas. AWS Business Support and AWS Enterprise Support customers can access all checks, including cost optimization, security, fault tolerance, performance, and service quotas. For a complete list of checks and descriptions, see the Trusted Advisor Best Practices.

    AWS Trusted Advisor Priority helps you focus on the most important recommendations to optimize your cloud deployments, improve resilience, and address security gaps. Available to AWS Enterprise Support customers, Trusted Advisor Priority provides prioritized and context-driven recommendations that come from your AWS account team as well as machine-generated checks from AWS services.




Amazon RedShift
    Amazon Redshift Database Developer Guide. Amazon Redshift is a fully managed, petabyte-scale data warehouse service in the cloud. Amazon Redshift Serverless lets you access and analyze data without the usual configurations of a provisioned data warehouse. Resources are automatically provisioned and data warehouse capacity is intelligently scaled to deliver fast performance for even the most demanding and unpredictable workloads. You don't incur charges when the data warehouse is idle, so you only pay for what you use. Regardless of the size of the dataset, you can load data and start querying right away in the Amazon Redshift query editor v2 or in your favorite business intelligence (BI) tool. Enjoy the best price performance and familiar SQL features in an easy-to-use, zero administration environment.

    Amazon Redshift is an enterprise-level, petabyte scale, fully managed data warehousing service. It uses columnar storage to improve the performance of complex queries.

AWS VPN CloudHub    

    Building on the AWS managed VPN options described previously, you can securely communicate from one site to another using the AWS VPN CloudHub. The AWS VPN CloudHub operates on a simple hub-and-spoke model that you can use with or without a VPC. Use this approach if you have multiple branch offices and existing internet connections and would like to implement a convenient, potentially low-cost hub-and-spoke model for primary or backup connectivity between these remote offices.

    The following figure shows the AWS VPN CloudHub architecture, with dashed lines indicating network traffic between remote sites being routed over their AWS VPN connections.

Amazon ElastiCache 
   Amazon ElastiCache is an in-memory database. With ElastiCache Memcached there is no data replication or high availability. 

   Therefore, the Redis engine must be used which does support both data replication and clustering.


Amazon Glacier 
You can specify one of the following when initiating a job to retrieve an archive based on your access time and cost requirements.
    Expedited — Expedited retrievals allow you to quickly access your data when occasional urgent requests for a subset of archives are required.
                For all but the largest archives (250 MB+), data accessed using Expedited retrievals are typically made available within 1–5 minutes. 
                Provisioned Capacity ensures that retrieval capacity for Expedited retrievals is available when you need it.
    Standard — Standard retrievals allow you to access any of your archives within several hours. 
               Standard retrievals typically complete within 3–5 hours. This is the default option for retrieval requests that do not specify the retrieval option.
    Bulk — Bulk retrievals are S3 Glacier’s lowest-cost retrieval option, which you can use to retrieve large amounts, even petabytes, of data inexpensively in a day.
           Bulk retrievals typically complete within 5–12 hours.

Amazon Simple Workflow Service (SWF) 
    Is a web service that makes it easy to coordinate work across distributed application components. SWF enables applications for a range of use cases, including media processing, web application back-ends, business process workflows, and analytics pipelines, to be designed as a coordination of tasks.

Amazon Aurora Global Database 
    Is designed for globally distributed applications, allowing a single Amazon Aurora database to span multiple AWS regions. It replicates your data with no impact on database performance, enables fast local reads with low latency in each region, and provides disaster recovery from region-wide outages.
    Aurora Global Database uses storage-based replication with typical latency of less than 1 second, using dedicated infrastructure that leaves your database fully available to serve application workloads. In the unlikely event of a regional degradation or outage, one of the secondary regions can be promoted to full read/write capabilities in less than 1 minute.

AWS OpsWork

