

AWS Trusted Advisor Service

    AWS Trusted Advisor is an application that draws upon best practices learned from AWS’s aggregated operational history of serving hundreds of thousands of AWS customers. Trusted Advisor inspects your AWS environment and makes recommendations for saving money, improving system performance, or closing security gaps.


The AWS Trusted Advisor Service limit publishes service limits metric to CloudWatch; thus, you can configure an alarm and send a notification to Amazon SNS. You can also create an AWS Lambda function to read data from specific Trusted Advisor checks. A Lambda function invocation can be scheduled using AWS EventBridge (Amazon CloudWatch Events) to automated the process.