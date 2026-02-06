# Assignment - 2 ( Designing a Banking System )

## Task Description
* Design a system that replicates a working of a banking system. It will contain multiple banks and branches, a customer can open a savings account or can take a loan at an interest rate of 12%, can view his account, transactions and loan details (such as loan pending, interest to be paid this year etc), a customer can perform various actions such as deposit and withdraw cash, repay loan, take a loan, etc 
* List all the features
* Design a database ER diagram 
* list all the APIs
* Implement all APIs

## Rationale behind ER Diagram
1. Why have separate entities for Bank and Branches?<br>
A. Since branches can have different ISFC codes and addresses while belonging to the main branch

2. Why have separate entities for Customer and Account?<br>
A. Since customers can have multiple accounts being managed across multiple branches of the bank

## Features
1. Customer Management
- Register new customers with personal details
- Update customer information
- View customer profile
- View all accounts under customer's name
- View all loans under customer's name

2. Account Management

3. Loan Management

4. Transaction Operations 

5. Bank and Branch Management