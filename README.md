# Solution for Receipt Processor Challenge
This is the Public Repository containing web service for Receipt Processor Challenge.

Here are few things I want you to consider before you move further ahead.

Web Service used - RESTful web service

Language selected - Go Language

IDE used - Eclipse with GoClipse extension and Golang Development Tooling

## What are the different edge cases I came across?

1. **Invalid Payloads** - As mentioned, API endpoints accept only JSON payloads. rest all are considered invalid.
2. **Transfer protocol** - Because of using RESTful services, I was only able to use http and https protocols to successfully load API endpoints.
3. **id Value** - As mentioned previously, randomized UUID values were used based on the specification given by this problem. UUID is not tightly-coupled to the receipt once called, which means id's for the same receipt keep chaning.
4. **Input validation** - Input values for Retailer names and description are generally a mix of characters including special characters like @, !, # etc. Inputs accept all characters including alphanumeric and non-alphanumeric.
5. **Date and Time validation** - Errors like 13th month, 31st date, 25th hour of the day, 61st minute of the hour etc., are prone to happen as these entities do not exist.
6. **HTTP methods for endpoints** - Process Receipts endpoint accepts only POST whereas Get Points endpoint accepts only GET. Using methods other than what were assigned like using GET for Process Receipts and PUT for Get Points etc., would result in "Method not allowed" error. 

## Files in my project directory
