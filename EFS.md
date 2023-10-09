What is Amazon Elastic File System?

    Amazon Elastic File System (Amazon EFS) provides serverless, fully elastic file storage so that you can share file data without provisioning or managing storage capacity and performance. Amazon EFS is built to scale on demand to petabytes without disrupting applications, growing and shrinking automatically as you add and remove files. Because Amazon EFS has a simple web services interface, you can create and configure file systems quickly and easily. The service manages all the file storage infrastructure for you, meaning that you can avoid the complexity of deploying, patching, and maintaining complex file system configurations.

    Amazon EFS supports the Network File System version 4 (NFSv4.1 and NFSv4.0) protocol, so the applications and tools that you use today work seamlessly with Amazon EFS. Multiple compute instances, including Amazon EC2, Amazon ECS, and AWS Lambda, can access an Amazon EFS file system at the same time. Therefore, an EFS file system can provide a common data source for workloads and applications that are running on more than one compute instance or server.


Amazon EFS lifecycle management

    Amazon EFS lifecycle management automatically manages cost-effective file storage for your file systems. When enabled, lifecycle management migrates files that have not been accessed for a set period of time to the EFS Standard–Infrequent Access (Standard-IA) or One Zone–Infrequent Access (One Zone-IA) storage class, depending on your file system. You define that period of time by using the Transition into IA lifecycle policy.

    Amazon EFS lifecycle management uses an internal timer to track when a file was last accessed, and not the POSIX file system attributes that are publicly viewable. Whenever a file in Standard or One Zone storage is accessed, the lifecycle management timer is reset. After lifecycle management moves a file into one of the IA storage classes, the file remains there indefinitely if Amazon EFS Intelligent-Tiering is not enabled.


There is no lifecycle policy available for deleting files on EFS
