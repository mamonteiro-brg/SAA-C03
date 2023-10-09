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
