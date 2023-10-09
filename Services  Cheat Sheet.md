
## Route Table Cheat Sheet

    Each VPC has a main route table that provides local routing throughout each VPC.

    Each subnet, when created using the VPC dashboard, is implicitly associated with the main route table.

    Don’t add additional routes to a main route table. Leaving the main route table in its default state ensures that if the main route table remains associated to a subnet by mistake, the worst that can happen is that local routing is enabled. If additional routes are added to the main route table, the additional routes will be available from each new subnet due to the default association with the main route table.

    The main route table cannot be deleted; however, it can be ignored and will remain unassigned if you do not associate it with any subnets within the VPC.

    Create and assign a custom route table for custom routes required by a subnet.

    Subnet destinations are matched with the most definitive route within the route table that matches the traffic request.

## Security groups

    A security group acts like a firewall at the EC2 instance, protecting all attached network interfaces.

    Security groups support both IPv4 and IPv6 traffic.

    A security group controls both outgoing (egress) and incoming (ingress) traffic.

    For each security group, rules control the inbound traffic that is allowed to reach the associated EC2 instances.

    Separate sets of rules control both the inbound and the outbound traffic.

    Each security group includes an outbound rule that allows all outbound traffic by default. Outbound rules can be modified and, if necessary, deleted.

    Security groups allow traffic based on protocols and port numbers.

    Security groups define allow rules. (It is not possible to create rules that explicitly deny access.)

    Security group rules allow you to direct traffic outbound from one security group inbound to another security group within the same VPC.

    Changes made to a security group take effect immediately.

    Security groups don’t deny traffic explicitly; instead, they deny traffic implicitly by defining only allowed traffic.

    Security groups are stateful; for requests that are allowed in, their response traffic is allowed out, and vice versa.

    For each rule, you define the protocol, the port or port range, and the source inbound and output destination for the traffic.

    The protocols allowed with security groups are TCP, UDP, or ICMP.

## Network ACL

    A NACL is an optional security control for subnets.

    Each VPC is assigned a default NACL that allows all inbound and outbound traffic across all subnets.

    NACLs are stateless in design; inbound and outbound rules are enforced independently.

    Each NACL is a collection of deny or allow rules for both inbound and outbound traffic.

    The default NACL can be modified.

    A NACL has both allow and deny rules.

    A NACL applies to both ingress and egress subnet traffic; it does not apply to traffic within the subnet.

    You can create custom NACLs and associate them to any subnet in a VPC.

    A custom NACL can be associated with more than one subnet.

    A subnet can be associated with only one NACL.

    A NACL is a first line of defense at the subnet level; a security group is a second line of defense at the instance.

## AWS Direct Connect Cheat Sheet

    You can configure an AWS Direct Connect connection with one or more virtual interfaces (VIFs).

    Public VIFs allow access to services such as Amazon S3 buckets and Amazon DynamoDB tables.

    Private VIFs allow access only to VPCs.

    An AWS Direct Connect connection allows connections to all availability zones within the region where the connection has been established.

    You are charged for AWS Direct Connect connections based on data transfer and port hours used.

    AWS Direct Connect dedicated connections are available at 1 Gbps up to 100 Gbps speeds.

    You can order speeds of 50 Mbps up to 200 Mbps through a hosted connection through AWS Direct Connect partners.

    An AWS Direct Connect gateway allows you to connect to multiple VPCs.

    An AWS Direct Connect gateway can connect to virtual private gateways and private virtual interfaces owned by the same AWS account.

    An AWS Direct Connect gateway can be associated with AWS Transit Gateway, extending an organization’s private network.

    An AWS Direct Connect connection can also be used with an IPsec VPN connection for additional security.

## AWS CloudTrail Cheat Sheet

    AWS CloudTrail records all activity on an AWS account, including API calls and authentications.

    Custom AWS CloudWatch trails can deliver events to an S3 bucket or a CloudWatch log group.

    AWS CloudTrail events can be used for auditing AWS account activity.

    AWS CloudTrail reports activity for each AWS account.

    AWS CloudTrail can be integrated with an AWS Organization.

    AWS CloudTrail tracks both data and management events.

    AWS CloudTrail records can be encrypted using S3 server-side encryption.


## CloudFront Cheat Sheet

    Control access to your public-facing content by mandating access via HTTPS endpoints using TLS 1.3.

    Origins include S3 buckets, AWS Elemental MediaStore container, an Application Load Balancer, a Lambda function URL, or a custom origin web server.

    Securing content access by using signed URLs and cookies.

    Use origin access identity (OAI) to restrict direct access to S3 bucket access, making it only accessible from CloudFront.

    Origin failover automatically serves content from the secondary origin when the primary origin is not available.

    Lambda@Edge functions support customizations that take from milliseconds to seconds to execute.

    CloudFront functions are lightweight functions that take less than one millisecond to execute.