# Lab 1

## Deploying Systems
I was asked to write code to develop and deploy systems. We have a pool of developers and deployers. 
Each system needs to be developed by a developer and deployed by a deployer to multiple environments.

I created the `Developer` and `Deployer` types, picked a random developer to develop each system, and picked
a random deployer to deploy the system to each environment.

Now, I've been asked to support special systems which an only be handled by specialists, The specialists can develop
and deploy systems. So, I created a `Specialist` type that can develop and deploy systems. My code won't compile anymore.

Please help me fix the code. I heard I need to use interfaces for these kinds of things.

[Link to Code](../cmd/lab1/lab1.go)


# Lab 2

## Errors

Link to [Code](../cmd/lab2/lab2.go) and [Test](../cmd/lab2/lab2_test.go)
The `DepositCheck` function is currently returning errors with string descriptions of what the issue is.
- Validation errors when check number is empty or amount is <=0
- Send to specialist review for suspicious activity if amount > 100000
- Call FBI for suspicious activity if amount > 1000000
 
The `PerformDeposit` function parsing the error message and doing the below
- if "bad amount", send an error back
- if "specialist review", send a fake confirmation number and send the check to specialist review
- if "call FBI", send a fake confirmation number and send the check to FBI review

Step 1. Change the implementation to use typed errors instead of parsing error string
Step 2. Return a generic custom error with `IsValidationError()`, `RequiresSpecialistReview()`, `SendToFBI()` methods