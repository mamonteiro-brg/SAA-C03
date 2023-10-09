AWS KMS
    AWS Key Management Service (AWS KMS) is a managed service that makes it easy for you to create and control the cryptographic keys that are used to protect your data. AWS KMS uses hardware security modules (HSM) to protect and validate your AWS KMS keys under the FIPS 140-2 Cryptographic Module Validation Program. China (Beijing) and China (Ningxia) Regions do not support the FIPS 140-2 Cryptographic Module Validation Program. AWS KMS uses OSCCA certified HSMs to protect KMS keys in China Regions.

    AWS KMS integrates with most other AWS services that encrypt your data. AWS KMS also integrates with AWS CloudTrail to log use of your KMS keys for auditing, regulatory, and compliance needs.


CMKs - Customer Manages Master Key

    Key Rotation

    Customer Master Keys


What is AWS KMS?
    AWS KMS is a safe and resilient service that uses hardware security protocols that are tested or are in the process of being tested to protect our keys. AWS Key Management Service provides a highly available key storage, management, and auditing solution for you to encrypt data within your own applications and control the encryption of stored data across AWS services.


Why AWS Key Management Service?
    Key Management Service is used to encrypt data in AWS. The main purpose of the AWS KMS is to store and manage those encryption keys. Data encryption is vital if you have sensitive data that must not be accessed by unauthorized users. Implement data encryption for both data at rest and data in transit.

    Two main methods to implement encryption at rest are Client-Side Encryption and Server Side Encryption.

        Client-Side Encryption is where you can encrypt the data on the client side and send it all the way to the server or any backend services like S3, EBS, Redshift, etc. In short, we can say in client-side encryption you encrypt your data and manage your own keys.
        
        Server-Side Encryption AWS encrypts the data and manages the keys for you, whereas you let your backend services encrypt the data and manage those keys on your behalf.



AWS KMS Architecture
 
KMS Practical Workflow
    Create a customer managed key (CMK) - Symmetric or Asymmetric key
    Define the administrative User & key User
    Encrypt and decrypt data with the CMK
        aws kms list-keys --region us-east-1
        aws kms encrypt --key-id <from previous command> --plaintext <fileb://file that you want to ecrypt> --region us-east-1
            output : ciphertextBlob
        aws kms decrypt --ciphertextBlob <ciphertextBlob> --region us-east-1


Envelope Encription wirh KMS
    Generate 1 CMK (customer master key) 
        --> aws kms generate-date-key --key-id <> --key-spec SAES_256 
            This will generate the plaintext and ciphertext data key
    Generate the data key. aws returns PT (plain text) and CT (cypher text) version of it
    Use plain text data key to encrypt files in server
    Store cipher text data key along with encrypted file

    
