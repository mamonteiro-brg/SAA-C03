

By default, each VPC comes with 1 route table pre-configured with a “local” route.

The scope of the “local” route is only within the subnet defined for the entire VPC

## Subnet route tables
    Your VPC has an implicit router, and you use route tables to control where network traffic is directed. Each subnet in your VPC must be associated with a route table, which controls the routing for the subnet (subnet route table). You can explicitly associate a subnet with a particular route table. Otherwise, the subnet is implicitly associated with the main route table. A subnet can only be associated with one route table at a time, but you can associate multiple subnets with the same subnet route table.

    Subnet route table—A route table that's associated with a subnet.

Gateway route table—A route table that's associated with an internet gateway or virtual private gateway.


Edge association—A route table that you use to route inbound VPC traffic to an appliance. You associate a route table with the internet gateway or virtual private gateway, and specify the network interface of your appliance as the target for VPC traffic.

Transit gateway route table—A route table that's associated with a transit gateway. For more information, see Transit gateway route tables in Amazon VPC Transit Gateways.


Local gateway route table—A route table that's associated with an Outposts local gateway. For more information, see Local gateways in the AWS Outposts User Guide.

