
Amazon Simple Queue Service (SQS) and Amazon Simple Workflow Service (SWF) are the services that you can use for creating a decoupled architecture in AWS. 


The key fact you need to consider is that duplicate messages cannot be introduced into the queue. 
For this reason alone you must use a FIFO queue. The statement about it not being necessary to maintain the order of the messages is meant to confuse you, as that might lead you to think you can use a standard queue, but standard queues don’t guarantee that duplicates are not introduced into the queue.
FIFO (first-in-first-out) queues preserve the exact order in which messages are sent and received – note that this is not required in the question but exactly once processing is. FIFO queues provide exactly-once processing, which means that each message is delivered once and remains available until a consumer processes it and deletes it.


Amazon SNS supports notifications over multiple transport protocols:
    – HTTP/HTTPS – subscribers specify a URL as part of the subscription registration.
    – Email/Email-JSON – messages are sent to registered addresses as email (text-based or JSON-object).
    – SQS – users can specify an SQS standard queue as the endpoint.
    – SMS – messages are sent to registered phone numbers as SMS text messages.


    