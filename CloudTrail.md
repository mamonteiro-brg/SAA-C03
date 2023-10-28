AWS CloudTrail
    AWS CloudTrail records all AWS API calls carried out within each AWS account related to actions across your AWS infrastructure. AWS CloudTrail also logs all account authentications and event history for your AWS account, including actions taken through the AWS Management Console, AWS SDKs, the AWS CLI command prompt, and other AWS service activity. 

    AWS CloudTrail is an AWS service that helps you enable operational and risk auditing, governance, and compliance of your AWS account. Actions taken by a user, role, or an AWS service are recorded as events in CloudTrail. Events include actions taken in the AWS Management Console, AWS Command Line Interface, and AWS SDKs and APIs.

    CloudTrail is active in your AWS account when you create it and doesn't require any manual setup. When activity occurs in your AWS account, that activity is recorded in a CloudTrail event.



There are two types of events that you configure your CloudTrail for:

    – Management Events
    – Data Events

    Management Events provide visibility into management operations that are performed on resources in your AWS account. These are also known as control plane operations. Management events can also include non-API events that occur in your account.

    Data Events, on the other hand, provide visibility into the resource operations performed on or within a resource. These are also known as data plane operations. It allows granular control of data event logging with advanced event selectors. You can currently log data events on different resource types such as Amazon S3 object-level API activity (e.g. GetObject, DeleteObject, and PutObject API operations), AWS Lambda function execution activity (the Invoke API), DynamoDB Item actions, and many more.


Amazon CloudTrail can be used to log activity on the reports. 

Data events provide visibility into the resource operations performed on or within a resource. These are also known as data plane operations. Data events are often high-volume activities.
Example data events include:
    • Amazon S3 object-level API activity (for example, GetObject, DeleteObject, and PutObject API operations).
    • AWS Lambda function execution activity (the Invoke API).

Management events provide visibility into management operations that are performed on resources in your AWS account. These are also known as control plane operations. Example management events include:
    • Configuring security (for example, IAM AttachRolePolicy API operations)
    • Registering devices (for example, Amazon EC2 CreateDefaultVpc API operations).
Therefore, to log data about access to the S3 objects the solutions architect should log read and write data events.

