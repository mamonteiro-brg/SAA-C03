Below are the valid EC2 lifecycle instance states:

    pending – The instance is preparing to enter the running state. An instance enters the pending state when it launches for the first time, or when it is restarted after being in the stopped state.

    running – The instance is running and ready for use.

    stopping – The instance is preparing to be stopped. Take note that you will not billed if it is preparing to stop however, you will still be billed if it is just preparing to hibernate.

    stopped – The instance is shut down and cannot be used. The instance can be restarted at any time.

    shutting-down – The instance is preparing to be terminated.

    terminated – The instance has been permanently deleted and cannot be restarted. Take note that Reserved Instances that applied to terminated instances are still billed until the end of their term according to their payment option.


Elastic Fabric Adapter (EFA) 
    Is a network device that you can attach to your Amazon EC2 instance to accelerate High Performance Computing (HPC) and machine learning applications. EFA enables you to achieve the application performance of an on-premises HPC cluster with the scalability, flexibility, and elasticity provided by the AWS Cloud.


Virtual Interface (VIF)

Elastic Network Interface (ENI)

Elastic Network Adapter (ENA)

High Performance Computing (HPC)


Amazon EC2 provides the following purchasing options to enable you to optimize your costs based on your needs:

    On-Demand Instances – Pay, by the second, for the instances that you launch.

    Savings Plans – Reduce your Amazon EC2 costs by making a commitment to a consistent amount of usage, in USD per hour, for a term of 1 or 3 years.

    Reserved Instances – Reduce your Amazon EC2 costs by making a commitment to a consistent instance configuration, including instance type and Region, for a term of 1 or 3 years.

    Spot Instances – Request unused EC2 instances, which can reduce your Amazon EC2 costs significantly.

    Dedicated Hosts – Pay for a physical host that is fully dedicated to running your instances, and bring your existing per-socket, per-core, or per-VM software licenses to reduce costs.

    Dedicated Instances – Pay, by the hour, for instances that run on single-tenant hardware.

    Capacity Reservations – Reserve capacity for your EC2 instances in a specific Availability Zone for any duration.



    
Placement groups
When you launch a new EC2 instance, the EC2 service attempts to place the instance in such a way that all of your instances are spread out across underlying hardware to minimize correlated failures. You can use placement groups to influence the placement of a group of interdependent instances to meet the needs of your workload. Depending on the type of workload, you can create a placement group using one of the following placement strategies:

    Cluster – packs instances close together inside an Availability Zone. This strategy enables workloads to achieve the low-latency network performance necessary for tightly-coupled node-to-node communication that is typical of high-performance computing (HPC) applications.

    Partition – spreads your instances across logical partitions such that groups of instances in one partition do not share the underlying hardware with groups of instances in different partitions. This strategy is typically used by large distributed and replicated workloads, such as Hadoop, Cassandra, and Kafka.

    Spread – strictly places a small group of instances across distinct underlying hardware to reduce correlated failures.