EBS Volumes:

    – When you create an EBS volume in an Availability Zone, it is automatically replicated within that zone to prevent data loss due to a failure of any single hardware component.

    – After you create a volume, you can attach it to any EC2 instance in the same Availability Zone

    – Amazon EBS Multi-Attach enables you to attach a single Provisioned IOPS SSD (io1) volume to multiple Nitro-based instances that are in the same Availability Zone. However, other EBS types are not supported.

    – An EBS volume is off-instance storage that can persist independently from the life of an instance. You can specify not to terminate the EBS volume when you terminate the EC2 instance during instance creation.

    – EBS volumes support live configuration changes while in production which means that you can modify the volume type, volume size, and IOPS capacity without service interruptions.

    – Amazon EBS encryption uses 256-bit Advanced Encryption Standard algorithms (AES-256)

    – EBS Volumes offer 99.999% SLA.


Amazon EBS volume types
    Solid State Drives (SSD)
        General Purpose SSD
            gp2 
                Max IOPS**/Volume: 16,000
                Max Throughput***/Volume: 250 MB/s
            gp3
                Max IOPS*/Volume: 16,000
                Max Throughput**/Volume: 1000 MB/s
        Provisioned IOPS SSD
            io1 
                Max IOPS*/Volume: 64,000
                Max Throughput**/Volume: 1,000 MB/s
            io2
                Max IOPS*/Volume: 64,000
                Max Throughput**/Volume: 1,000 MB/s
            io2 Block Express - 
                Max IOPS/Volume: 256,000
                Max Throughput*/Volume: 4,000 MB/s

    Hard Disk Drives (HDD)
        Throughput Optimized HDD
        Cold HDD
    Previous Generation
        Magnetic

        
When an EBS volume is encrypted with a custom key you must share the custom key with the PROD account. You also need to modify the permissions on the snapshot to share it with the PROD account. The PROD account must copy the snapshot before they can then create volumes from the snapshot
Note that you cannot share encrypted volumes created using a default CMK key and you cannot change the CMK key that is used to encrypt a volume.



Amazon EBS volume types
    https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-volume-types.html