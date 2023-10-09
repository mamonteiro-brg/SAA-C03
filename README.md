# SAA-C03
Preparation Material for SAA-C03


Exam Tip
the company want to generate the keys then store it in AWS with enabling cross-region replication then you have to configure S3 buckets to use SSE with AWS KMS-Managed Keys (SSE-KMS) with imported key material in both regions

An AWS KMS customer master key (CMK) is a logical representation of a master key. In addition to the CMK identifiers and other metadata, a CMK contains the key material used to encrypt and decrypt data. When you create a CMK, by default AWS KMS generates the key material for that CMK. But you can create a CMK without key material and then import your own key material into that CMK, a feature often known as “bring your own key” (BYOK).
Imported key material is supported only for symmetric CMKs in AWS KMS key stores. It is not supported on asymmetric CMKs or CMKs in custom key stores.
When you use imported key material, you remain responsible for the key material while allowing AWS KMS to use a copy of it. You might choose to do this for one or more of the following reasons:
• To prove that you generated the key material using a source of entropy that meets your requirements.
• To use key material from your own infrastructure with AWS services, and to use AWS KMS to manage the lifecycle of that key material within AWS.
• To set an expiration time for the key material in AWS and to manually delete it, but to also make it available again in the future. In contrast, scheduling key deletion requires a waiting period of 7 to 30 days, after which you cannot recover the deleted CMK.
• To own the original copy of the key material, and to keep it outside of AWS for additional durability and disaster recovery during the complete lifecycle of the key material.
Reference

to secure the data in EBS volume = encrypt the data using KMS
